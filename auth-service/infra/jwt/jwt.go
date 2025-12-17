package jwt

import (
	"fmt"
	jwtdomain "github.com/Fox216540/shop/auth-service/domain/jwt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"time"
)

type service struct {
	ttlRefresh    time.Duration
	ttlAccess     time.Duration
	secretRefresh string
	secretAccess  string
}

func NewService(ttlAccess, ttlRefresh time.Duration, secretAccess, secretRefresh string) jwtdomain.Repository {
	return &service{
		ttlAccess:     ttlAccess,
		ttlRefresh:    ttlRefresh,
		secretAccess:  secretAccess,
		secretRefresh: secretRefresh,
	}
}

func (s *service) generateToken(userID uuid.UUID, ttl time.Duration, secret string, tokenType string) (string, uuid.UUID, error) {
	jti := uuid.New()
	claims := jwt.MapClaims{
		"sub":  userID.String(),
		"type": tokenType,
		"exp":  time.Now().Add(ttl).Unix(),
		"jti":  jti.String(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", uuid.Nil, err
	}

	return tokenString, jti, nil
}

func (s *service) GenerateRefreshToken(userID uuid.UUID) (string, uuid.UUID, error) {
	token, jti, err := s.generateToken(userID, s.ttlRefresh, s.secretRefresh, "refresh")
	if err != nil {
		return "", uuid.Nil, NewInvalidGenerateRefreshToken(err)
	}
	return token, jti, nil
}

func (s *service) GenerateAccessToken(userID uuid.UUID) (string, error) {
	token, _, err := s.generateToken(userID, s.ttlAccess, s.secretAccess, "access")
	if err != nil {
		return "", NewInvalidGenerateAccessToken(err)
	}
	return token, nil
}

func (s *service) decodeToken(tokenStr, secret string, newBadRequestError, newNoValidError func(error) error) (jwtdomain.JWTUser, error) {
	parsedToken, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
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

	return jwtdomain.JWTUser{UserID: userID, JTI: jti}, nil
}

func (s *service) DecodeRefreshToken(token string) (jwtdomain.JWTUser, error) {
	return s.decodeToken(token, s.secretRefresh, jwtdomain.NewBadRefreshTokenError, jwtdomain.NewNoValidRefreshTokenError)
}

func (s *service) DecodeAccessToken(token string) (jwtdomain.JWTUser, error) {
	return s.decodeToken(token, s.secretAccess, jwtdomain.NewBadAccessTokenError, jwtdomain.NewNoValidAccessTokenError)
}
