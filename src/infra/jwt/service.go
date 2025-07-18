package jwt

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"os"
	jwtdomain "shop/src/domain/jwt"
	"time"
)

type service struct {
	s jwtdomain.Service
}

func NewService(s jwtdomain.Service) jwtdomain.Service {
	return &service{s: s}
}

func (s *service) GenerateRefreshToken(userID uuid.UUID) (string, error) {
	ttl := os.Getenv("REFRESH_TOKEN_TTL")
	duration, err := time.ParseDuration(ttl)
	if err != nil {
		return "", err
	}
	secret := []byte(os.Getenv("REFRESH_TOKEN_SECRET"))
	claims := jwt.MapClaims{
		"sub":  userID.String(), // как str(user_id)
		"type": "refresh",       // "access"
		"exp":  time.Now().Add(duration).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(secret)
}

func (s *service) GenerateAccessToken(userID uuid.UUID, username string) (string, error) {
	ttl := os.Getenv("ACCESS_TOKEN_TTL")
	duration, err := time.ParseDuration(ttl)
	if err != nil {
		return "", err
	}
	secret := []byte(os.Getenv("ACCESS_TOKEN_SECRET"))
	claims := jwt.MapClaims{
		"sub":      userID.String(), // как str(user_id)
		"type":     "access",        // "access"
		"username": username,        // username
		"exp":      time.Now().Add(duration).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(secret)
}

func (s *service) DecodeRefreshToken(token string) (jwtdomain.JWT, error) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("REFRESH_TOKEN_SECRET")), nil
	})
	if err != nil {
		return jwtdomain.JWT{}, err
	}
	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok || !parsedToken.Valid {
		return jwtdomain.JWT{}, err
	}
	userID, err := uuid.Parse(claims["sub"].(string))
	if err != nil {
		return jwtdomain.JWT{}, err
	}
	return jwtdomain.JWT{
		UserID: userID,
	}, nil
}

func (s *service) DecodeAccessToken(token string) (jwtdomain.JWT, error) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("REFRESH_TOKEN_SECRET")), nil
	})
	if err != nil {
		return jwtdomain.JWT{}, err
	}
	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok || !parsedToken.Valid {
		return jwtdomain.JWT{}, err
	}
	userID, err := uuid.Parse(claims["sub"].(string))
	if err != nil {
		return jwtdomain.JWT{}, err
	}

	username, ok := claims["username"].(string)
	if !ok {
		return jwtdomain.JWT{}, err
	}

	return jwtdomain.JWT{
		UserID:   userID,
		Username: username,
	}, nil
}
