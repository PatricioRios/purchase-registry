package auth

import (
	"fmt"
	"time"

	"github.com/PatricioRios/Compras/models"
	"github.com/PatricioRios/Compras/utils"
	"github.com/golang-jwt/jwt/v4"
)

type AuthRepositoryImpl struct {
	secretKey          string
	secretKeyRefresh   string
	jwtDuration        int
	jwtRefreshDuration int
}

func NewAuthRepositoryImpl(env utils.Env) AuthRepository {
	return &AuthRepositoryImpl{
		secretKey:          env.JWT.Secret,
		secretKeyRefresh:   env.JWT.SecretRefresh,
		jwtDuration:        env.JWT.Duration,
		jwtRefreshDuration: env.JWT.RefreshDuration,
	}
}

func (r AuthRepositoryImpl) GenerateJWT(user models.User) (string, string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":        time.Now().Add(time.Duration(r.jwtDuration) * time.Hour).Unix(),
		"authorized": true,
		"user":       user.UserName,
		"id":         user.Id,
	})

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":        time.Now().Add(time.Duration(r.jwtRefreshDuration*720) * time.Hour).Unix(),
		"authorized": true,
		"user":       user.UserName,
		"id":         user.Id,
	})

	refreshTokenString, err := refreshToken.SignedString([]byte(r.secretKeyRefresh))
	if err != nil {
		fmt.Println(err)
		return "", "", ErrTokenGenerationFailed
	}

	tokenString, err := token.SignedString([]byte(r.secretKey))
	if err != nil {
		fmt.Println(err)
		return "", "", ErrTokenGenerationFailed
	}

	return tokenString, refreshTokenString, nil
}

func (r AuthRepositoryImpl) VerifyToken(tokenString string) error {
	token, err := r.GetToken(tokenString, r.secretKey)
	if err != nil {
		return utils.NewBadRequest(err.Error())
	}

	if !token.Valid {
		return ErrInvalidToken
	}

	return nil
}

func (r AuthRepositoryImpl) VerifyRefreshToken(tokenString string) error {
	token, err := r.GetToken(tokenString, r.secretKeyRefresh)
	if err != nil {
		return utils.NewBadRequest(err.Error())
	}

	if !token.Valid {
		return ErrInvalidToken
	}

	return nil
}

func (r AuthRepositoryImpl) RefreshToken(refreshTokenString string) (string, error) {
	token, err := r.GetToken(refreshTokenString, r.secretKeyRefresh)
	if err != nil {
		return "", utils.NewBadRequest(err.Error())
	}

	if !token.Valid {
		return "", ErrInvalidToken
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", ErrTokenParsingFailed
	}

	userId, ok := claims["id"].(float64)
	if !ok {
		return "", ErrTokenParsingFailed
	}

	var user models.User
	user.UserName, ok = claims["user"].(string)
	if !ok {
		return "", ErrTokenParsingFailed
	}

	user.Id = int(userId)
	newToken, _, err := r.GenerateJWT(user)
	if err != nil {
		return "", ErrTokenGenerationFailed
	}

	return newToken, nil
}

func (r AuthRepositoryImpl) GetToken(tokenString string, key string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrUnexpectedSigningMethod
		}
		return []byte(key), nil
	})

	return token, err
}

func (r AuthRepositoryImpl) GetUserId(token *jwt.Token) (int, error) {
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, ErrTokenParsingFailed
	}

	userId, ok := claims["id"].(float64)
	if !ok {
		return 0, ErrTokenParsingFailed
	}

	return int(userId), nil
}
