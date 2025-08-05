package jwt

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"log"
	"os"
	jwtdomain "shop/src/domain/jwt"
	"strconv"
	"time"
)

type service struct {
}

func NewService() jwtdomain.Service {
	return &service{}
}

func (s *service) toDuration(ttl string) (time.Duration, error) {
	seconds, err := strconv.ParseInt(ttl, 10, 64)
	if err != nil {
		return 0, err
	}
	return time.Duration(seconds) * time.Second, nil
}

func (s *service) GenerateRefreshToken(userID uuid.UUID) (string, uuid.UUID, error) {
	log.Println(userID)
	ttl := os.Getenv("REFRESH_TOKEN_TTL")
	log.Println(ttl)
	secret := []byte(os.Getenv("REFRESH_TOKEN_SECRET"))
	log.Println(secret)

	duration, err := s.toDuration(ttl)
	if err != nil {
		return "", uuid.Nil, fmt.Errorf("invalid REFRESH_TOKEN_TTL: %w", err)
	}
	jti := uuid.New()
	claims := jwt.MapClaims{
		"sub":  userID.String(), // как str(user_id)
		"type": "refresh",       // "access"
		"exp":  time.Now().Add(duration).Unix(),
		"jti":  jti,
	}
	fmt.Println(claims)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	fmt.Println(token)
	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", uuid.Nil, err
	}
	return tokenString, jti, nil
}

func (s *service) GenerateAccessToken(userID uuid.UUID, username string) (string, error) {
	ttl := os.Getenv("ACCESS_TOKEN_TTL")
	secret := []byte(os.Getenv("ACCESS_TOKEN_SECRET"))
	duration, err := s.toDuration(ttl)
	if err != nil {
		return "", fmt.Errorf("invalid REFRESH_TOKEN_TTL: %w", err)
	}
	jti := uuid.New()
	claims := jwt.MapClaims{
		"sub":      userID.String(), // как str(user_id)
		"type":     "access",        // "access"
		"username": username,        // username
		"exp":      time.Now().Add(duration).Unix(),
		"jti":      jti,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secret)
}

func (s *service) DecodeRefreshToken(token string) (jwtdomain.JWTUser, error) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("REFRESH_TOKEN_SECRET")), nil
	})
	if err != nil {
		return jwtdomain.JWTUser{}, err
	}
	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok || !parsedToken.Valid {
		return jwtdomain.JWTUser{}, err
	}
	userID, err := uuid.Parse(claims["sub"].(string))
	if err != nil {
		return jwtdomain.JWTUser{}, err
	}
	jti, err := uuid.Parse(claims["jti"].(string))
	if err != nil {
		return jwtdomain.JWTUser{}, err
	}
	return jwtdomain.JWTUser{
		UserID: userID,
		JTI:    jti,
	}, nil
}

func (s *service) DecodeAccessToken(token string) (jwtdomain.JWTUser, error) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("REFRESH_TOKEN_SECRET")), nil
	})
	if err != nil {
		return jwtdomain.JWTUser{}, err
	}
	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok || !parsedToken.Valid {
		return jwtdomain.JWTUser{}, err
	}
	userID, err := uuid.Parse(claims["sub"].(string))
	if err != nil {
		return jwtdomain.JWTUser{}, err
	}

	username, ok := claims["username"].(string)
	if !ok {
		return jwtdomain.JWTUser{}, err
	}

	jti, err := uuid.Parse(claims["jti"].(string))
	if err != nil {
		return jwtdomain.JWTUser{}, err
	}

	return jwtdomain.JWTUser{
		UserID:   userID,
		Username: username,
		JTI:      jti,
	}, nil
}
