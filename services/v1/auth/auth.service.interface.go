package auth

import (
	"errors"

	"github.com/PatricioRios/Compras/models"
)

type AuthService interface {
	LogInWithEmail(user models.User) (string, string, error)
	LogInWithUserName(user models.User) (string, string, error)
	ValidateToken(tokenString string) error
	ValidateRefreshToken(tokenString string) error
	RefreshToken(refreshTokenString string) (string, error)
	GetUserIdInToken(tokenString string) (int, error)
}

var (
	ErrInvalidEmailFormat = errors.New("FormatoDeEmailInvalido")
	ErrPasswordRequired   = errors.New("PasswordRequired")

	ErrUserNameRequired = errors.New("UserNameRequired")
	ErrUserNotFound     = errors.New("UserNotFound")

	ErrInvalidPassword       = errors.New("WrongPassword")
	ErrTokenInvalid          = errors.New("InvalidToken")
	ErrTokenClaimsParsing    = errors.New("ErrorParsingTokenClaims")
	ErrTokenGenerationFailed = errors.New("TokenGenerationFailed")
)
