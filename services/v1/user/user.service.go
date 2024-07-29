package user

import (
	"net/http"
	"strings"

	"github.com/PatricioRios/Compras/models"
	"github.com/PatricioRios/Compras/repository/v1/user"
	"github.com/PatricioRios/Compras/utils"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userRepository user.UserRepository
}

func NewUserService(userRepository user.UserRepository) UserService {
	return UserService{userRepository: userRepository}
}

func (s *UserService) GetAllUsers() ([]models.User, utils.SrvcError) {

	users, err := s.userRepository.GetAllUsers()
	if err != nil {
		return nil, utils.NewError(http.StatusInternalServerError, err.Error())
	}
	return users, nil
}
func (s *UserService) GetUserById(id int) (models.User, utils.SrvcError) {
	if id < 1 {
		return models.User{}, utils.NewError(http.StatusBadRequest, "El ID no puede ser menor a 1")
	}
	user, err := s.userRepository.GetUserById(id)

	if err != nil {
		return user, utils.NewError(http.StatusInternalServerError, err.Error())
	}

	return user, nil
}

func (s *UserService) CreateUser(user models.User) (models.User, utils.SrvcError) {
	if user.Name == "" {
		return user, utils.NewBadRequest("El nombre es obligatorio")
	}
	if user.LastName == "" {
		return user, utils.NewBadRequest("El apellido es obligatorio")
	}
	if user.UserName == "" {
		return user, utils.NewBadRequest("El nombre de usuario es obligatorio")
	}
	if user.Password == "" {
		return user, utils.NewBadRequest("La contraseña es obligatoria")
	}
	if user.Email == "" {
		return user, utils.NewBadRequest("El email es obligatorio")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 5)
	if err != nil {
		return models.User{}, utils.NewError(http.StatusInternalServerError, err.Error())
	}
	user.Password = string(hashedPassword)

	user.UserName = strings.ToLower(user.UserName)
	user.Email = strings.ToLower(user.Email)
	user.Name = strings.ToLower(user.Name)
	user.LastName = strings.ToLower(user.LastName)

	user, err = s.userRepository.CreateUser(user)

	if err != nil {
		return models.User{}, utils.NewError(http.StatusInternalServerError, err.Error())
	}
	user.Password = ""
	return user, nil
}
