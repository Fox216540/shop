package jwt

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"shop/src/core/settings"
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
	config := settings.Config
	ttl := config.RefreshTokenTTL
	secret := []byte(config.RefreshTokenSecret)
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
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", uuid.Nil, err
	}
	return tokenString, jti, nil
}

func (s *service) GenerateAccessToken(userID uuid.UUID) (string, error) {
	config := settings.Config
	ttl := config.AccessTokenTTL
	secret := []byte(config.AccessTokenSecret)
	duration, err := s.toDuration(ttl)
	if err != nil {
		return "", fmt.Errorf("invalid ACCESS_TOKEN_TTL: %w", err)
	}
	jti := uuid.New()
	claims := jwt.MapClaims{
		"sub":  userID.String(), // как str(user_id)
		"type": "access",        // "access"
		"exp":  time.Now().Add(duration).Unix(),
		"jti":  jti,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secret)
}

func (s *service) DecodeRefreshToken(token string) (jwtdomain.JWTUser, error) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(settings.Config.RefreshTokenSecret), nil
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
		return []byte(settings.Config.AccessTokenSecret), nil
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
