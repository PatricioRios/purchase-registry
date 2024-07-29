package auth

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/PatricioRios/Compras/models"
	"github.com/PatricioRios/Compras/repository/v1/user"
	"github.com/PatricioRios/Compras/utils"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthService struct {
	secretKey        string
	secretKeyRefresh string
	userRepository   user.UserRepository
}

func NewAuthService(
	userRepository user.UserRepository,
	env utils.Env,
) AuthService {
	fmt.Println("AUTH SERVICE PROVIDED")
	return AuthService{
		secretKey:        env.JWT.Secret,
		secretKeyRefresh: env.JWT.SecretRefresh,
		userRepository:   userRepository,
	}
}

func (s *AuthService) LogInWithEmail(user models.User) (string, string, utils.SrvcError) {

	if !strings.Contains(user.Email, "@") || !strings.Contains(user.Email, ".") {
		return "", "", utils.NewError(http.StatusBadRequest, "El email no es correcto")
	}
	if user.Password == "" {
		return "", "", utils.NewError(http.StatusBadRequest, "La contraseña es obligatoria")
	}
	user.Email = strings.ToLower(user.Email)
	userRegistry, err := s.userRepository.GetUserByEmail(user.Email)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return "", "", utils.NewError(http.StatusUnauthorized, "Usuario no encontrado")
		}
		return "", "", utils.NewError(http.StatusInternalServerError, err.Error())
	}
	err = bcrypt.CompareHashAndPassword([]byte(userRegistry.Password), []byte(user.Password))

	if err != nil {
		return "", "", utils.NewError(http.StatusUnauthorized, "Contraseña incorrecta")
	}
	token, refreshToken, err := s.generateJWT(userRegistry)

	if err != nil {
		return "", "", utils.NewError(http.StatusInternalServerError, err.Error())
	}

	return token, refreshToken, nil

}

func (s *AuthService) LogInWithUserName(user models.User) (string, string, utils.SrvcError) {

	if user.Password == "" {
		return "", "", utils.NewError(http.StatusBadRequest, "La contraseña es obligatoria")
	}
	if user.UserName == "" {
		return "", "", utils.NewError(http.StatusBadRequest, "El nombre de usuario es obligatorio")
	}

	userRegistry, err := s.userRepository.GetUserByUserName(user.UserName)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return "", "", utils.NewError(http.StatusUnauthorized, "Usuario no encontrado")
		}
		return "", "", utils.NewError(http.StatusInternalServerError, err.Error())
	}
	err = bcrypt.CompareHashAndPassword([]byte(userRegistry.Password), []byte(user.Password))

	if err != nil {
		return "", "", utils.NewError(http.StatusUnauthorized, "Contraseña incorrecta")
	}
	token, refreshToken, err := s.generateJWT(userRegistry)

	if err != nil {
		return "", "", utils.NewError(http.StatusInternalServerError, err.Error())
	}

	return token, refreshToken, nil

}

func (s *AuthService) generateJWT(user models.User) (string, string, error) {
	//make sure you use HS256 for signingMethods
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		//exp means expiration time. set it 1 minutes to check if it works.
		"exp": time.Now().Add(1 * time.Hour).Unix(),
		//set authorized true
		"authorized": true,
		//as a user we set username from the input you may want to set something else..
		"user": user.UserName,
		"id":   user.Id,
	})
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		//exp means expiration time. set it 1 minutes to check if it works.
		"exp": time.Now().Add(1 * time.Hour).Unix(),
		//set authorized true
		"authorized": true,
		//as a user we set username from the input you may want to set something else..
		"user": user.UserName,
		"id":   user.Id,
	})
	RefreshTokenString, err := refreshToken.SignedString([]byte(s.secretKeyRefresh))
	if err != nil {
		fmt.Println(err)
		return "", "", err
	}
	tokenString, err := token.SignedString([]byte(s.secretKey))
	if err != nil {
		fmt.Println(err)
		return "", "", err
	}

	return tokenString, RefreshTokenString, nil
}

func (s *AuthService) ValidateToken(tokenString string) utils.SrvcError {
	// Parse the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, utils.NewError(http.StatusBadRequest, "unexpected signing method")
		}
		return []byte(s.secretKey), nil
	})

	if err != nil {
		return utils.NewBadRequest(err.Error())
	}

	if !token.Valid {
		return utils.NewError(http.StatusLocked, "Token invalido")
	}

	return nil
}

func (s *AuthService) ValidateRefreshToken(tokenString string) utils.SrvcError {
	// Parse the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, utils.NewError(http.StatusBadRequest, "unexpected signing method")
		}
		return []byte(s.secretKeyRefresh), nil
	})

	if err != nil {
		return utils.NewBadRequest(err.Error())
	}

	if !token.Valid {
		return utils.NewError(http.StatusLocked, "Token invalido")
	}

	return nil
}

func (s *AuthService) RefreshToken(refreshTokenString string) (string, utils.SrvcError) {
	// Parse the token
	token, err := jwt.Parse(refreshTokenString, func(token *jwt.Token) (interface{}, error) {
		// Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, utils.NewError(http.StatusBadRequest, "unexpected signing method")
		}
		return []byte(s.secretKeyRefresh), nil
	})

	if err != nil {
		return "", utils.NewBadRequest(err.Error())
	}

	if !token.Valid {
		return "", utils.NewError(http.StatusLocked, "Token invalido")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", utils.NewError(http.StatusInternalServerError, "Error parsing token claims")
	}

	userId := claims["id"].(float64)

	var user models.User

	user.UserName = claims["user"].(string)
	fmt.Println(user.UserName)
	user.Id = int(userId)

	newToken, _, err := s.generateJWT(user)
	if err != nil {
		return "", utils.NewError(http.StatusInternalServerError, err.Error())
	}

	return newToken, nil
}
