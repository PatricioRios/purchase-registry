package auth

import (
	"strings"

	"github.com/PatricioRios/Compras/models"
	"github.com/PatricioRios/Compras/repository/v1/auth"
	"github.com/PatricioRios/Compras/repository/v1/user"
	"github.com/PatricioRios/Compras/utils"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthServiceImpl struct {
	secretKey        string
	secretKeyRefresh string
	userRepository   user.UserAuthRepo
	authRepository   auth.AuthRepository
}

func NewAuthService(
	userRepository user.UserAuthRepo,
	authRepository auth.AuthRepository,
	env utils.Env,
) AuthService {
	return &AuthServiceImpl{
		secretKey:        env.JWT.Secret,
		secretKeyRefresh: env.JWT.SecretRefresh,
		userRepository:   userRepository,
		authRepository:   authRepository,
	}
}

func (s *AuthServiceImpl) LogInWithEmail(user models.User) (string, string, error) {
	if !strings.Contains(user.Email, "@") || !strings.Contains(user.Email, ".") {
		return "", "", ErrInvalidEmailFormat
	}

	if user.Password == "" {
		return "", "", ErrPasswordRequired
	}

	user.Email = strings.ToLower(user.Email)
	userRegistry, err := s.userRepository.GetUserByEmail(user.Email)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return "", "", ErrUserNotFound
		}
		return "", "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(userRegistry.Password), []byte(user.Password))
	if err != nil {
		return "", "", ErrInvalidPassword
	}

	token, refreshToken, err := s.authRepository.GenerateJWT(userRegistry)
	if err != nil {
		return "", "", ErrTokenGenerationFailed
	}

	return token, refreshToken, nil
}

func (s *AuthServiceImpl) LogInWithUserName(user models.User) (string, string, error) {
	if user.Password == "" {
		return "", "", ErrPasswordRequired
	}

	if user.UserName == "" {
		return "", "", ErrUserNameRequired
	}

	userRegistry, err := s.userRepository.GetUserByUserName(user.UserName)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return "", "", ErrUserNotFound
		}
		return "", "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(userRegistry.Password), []byte(user.Password))
	if err != nil {
		return "", "", ErrInvalidPassword
	}

	token, refreshToken, err := s.authRepository.GenerateJWT(userRegistry)
	if err != nil {
		return "", "", ErrTokenGenerationFailed
	}

	return token, refreshToken, nil
}

func (s *AuthServiceImpl) ValidateToken(tokenString string) error {
	token, err := s.authRepository.GetToken(tokenString, s.secretKey)
	if err != nil {
		return err
	}

	if !token.Valid {
		return ErrTokenInvalid
	}

	return nil
}

func (s *AuthServiceImpl) ValidateRefreshToken(tokenString string) error {
	token, err := s.authRepository.GetToken(tokenString, s.secretKeyRefresh)
	if err != nil {
		return err
	}

	if !token.Valid {
		return ErrTokenInvalid
	}

	return nil
}

func (s *AuthServiceImpl) RefreshToken(refreshTokenString string) (string, error) {
	token, err := s.authRepository.GetToken(refreshTokenString, s.secretKeyRefresh)
	if err != nil {
		return "", err
	}

	if !token.Valid {
		return "", ErrTokenInvalid
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", ErrTokenClaimsParsing
	}

	userId, ok := claims["id"].(float64)
	if !ok {
		return "", ErrTokenClaimsParsing
	}

	var user models.User
	user.Id = int(userId)

	newToken, _, err := s.authRepository.GenerateJWT(user)
	if err != nil {
		return "", ErrTokenGenerationFailed
	}

	return newToken, nil
}

func (s *AuthServiceImpl) GetUserIdInToken(tokenString string) (int, error) {
	token, err := s.authRepository.GetToken(tokenString, s.secretKey)
	if err != nil {
		return 0, err
	}

	userId, err := s.authRepository.GetUserId(token)
	if err != nil {
		return 0, err
	}

	return userId, nil
}
