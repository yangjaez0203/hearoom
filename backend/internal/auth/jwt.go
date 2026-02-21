package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/yangjaez0203/hearoom/backend/internal/models"
)

func GenerateToken(secret string, expiry time.Duration, user *models.User) (string, error) {
	claims := models.UserClaims{
		Username:  user.Username,
		Anonymous: user.Anonymous,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   user.ID,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expiry)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

func ValidateToken(secret string, tokenString string) (*models.User, error) {
	token, err := jwt.ParseWithClaims(tokenString, &models.UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*models.UserClaims)
	if !ok || !token.Valid {
		return nil, jwt.ErrTokenInvalidClaims
	}

	return &models.User{
		ID:        claims.Subject,
		Username:  claims.Username,
		Anonymous: claims.Anonymous,
	}, nil
}

func NewAnonymousUser() *models.User {
	return &models.User{
		ID:        uuid.New().String(),
		Username:  GenerateName(),
		Anonymous: true,
	}
}
