package infrastructure

import (
	"errors"
	"time"

	"task-manager-clean-arch/Domain"

	"github.com/golang-jwt/jwt/v5"
)

type JWTServiceImpl struct {
	secretKey []byte
}

func NewJWTService(secretKey string) domain.JWTService {
	return &JWTServiceImpl{
		secretKey: []byte(secretKey),
	}
}

func (j *JWTServiceImpl) GenerateToken(userID, username, role string) (string, error) {
	claims := domain.JWTClaims{
		UserID:   userID,
		Username: username,
		Role:     role,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":  claims.UserID,
		"username": claims.Username,
		"role":     claims.Role,
		"exp":      time.Now().Add(24 * time.Hour).Unix(),
		"iat":      time.Now().Unix(),
	})

	// Sign the token
	tokenString, err := token.SignedString(j.secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (j *JWTServiceImpl) ValidateToken(tokenString string) (*domain.JWTClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return j.secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("invalid token claims")
	}

	userID, ok := claims["user_id"].(string)
	if !ok {
		return nil, errors.New("invalid user_id claim")
	}

	username, ok := claims["username"].(string)
	if !ok {
		return nil, errors.New("invalid username claim")
	}

	role, ok := claims["role"].(string)
	if !ok {
		return nil, errors.New("invalid role claim")
	}

	// Check if token is expired
	exp, ok := claims["exp"].(float64)
	if !ok {
		return nil, errors.New("invalid exp claim")
	}

	if time.Now().Unix() > int64(exp) {
		return nil, errors.New("token expired")
	}

	return &domain.JWTClaims{
		UserID:   userID,
		Username: username,
		Role:     role,
	}, nil
} 