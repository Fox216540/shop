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
		return "", uuid.Nil, NewInvalidGenerateRefreshToken(err)
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
		return "", uuid.Nil, NewInvalidGenerateRefreshToken(err)
	}
	return tokenString, jti, nil
}

func (s *service) GenerateAccessToken(userID uuid.UUID) (string, error) {
	config := settings.Config
	ttl := config.AccessTokenTTL
	secret := []byte(config.AccessTokenSecret)
	duration, err := s.toDuration(ttl)
	if err != nil {
		return "", NewInvalidGenerateAccessToken(err)
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

func (s *service) decodeToken(tokenStr string, secret string, newBadRequestError, newNoValidError func(error) error) (jwtdomain.JWTUser, error) {
	parsedToken, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return jwtdomain.JWTUser{}, newBadRequestError(err)
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok || !parsedToken.Valid {
		return jwtdomain.JWTUser{}, newNoValidError(fmt.Errorf("invalid token claims"))
	}

	sub, ok := claims["sub"].(string)
	if !ok {
		return jwtdomain.JWTUser{}, newBadRequestError(fmt.Errorf("missing sub claim"))
	}
	userID, err := uuid.Parse(sub)
	if err != nil {
		return jwtdomain.JWTUser{}, newBadRequestError(err)
	}

	jtiStr, ok := claims["jti"].(string)
	if !ok {
		return jwtdomain.JWTUser{}, newBadRequestError(fmt.Errorf("missing jti claim"))
	}
	jti, err := uuid.Parse(jtiStr)
	if err != nil {
		return jwtdomain.JWTUser{}, newBadRequestError(err)
	}

	return jwtdomain.JWTUser{
		UserID: userID,
		JTI:    jti,
	}, nil
}

func (s *service) DecodeRefreshToken(token string) (jwtdomain.JWTUser, error) {
	return s.decodeToken(token, settings.Config.RefreshTokenSecret, jwtdomain.NewBadRefreshTokenError, jwtdomain.NewNoValidRefreshTokenError)
}

func (s *service) DecodeAccessToken(token string) (jwtdomain.JWTUser, error) {
	return s.decodeToken(token, settings.Config.AccessTokenSecret, jwtdomain.NewBadAccessTokenError, jwtdomain.NewNoValidAccessTokenError)
}
