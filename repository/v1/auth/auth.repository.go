package auth

import (
	"errors"

	"github.com/PatricioRios/Compras/models"
	"github.com/golang-jwt/jwt/v4"
)

type AuthRepository interface {
	GenerateJWT(user models.User) (string, string, error)
	VerifyToken(tokenString string) error
	VerifyRefreshToken(tokenString string) error
	GetToken(token string, key string) (*jwt.Token, error)
	GetUserId(token *jwt.Token) (int, error)
	RefreshToken(tokenString string) (string, error)
}

var (
	ErrInvalidToken            = errors.New("Invalid token")
	ErrUnexpectedSigningMethod = errors.New("Unexpected signing method")
	ErrTokenParsingFailed      = errors.New("Error parsing token claims")
	ErrTokenGenerationFailed   = errors.New("Failed to generate token")
)
