package user

import "github.com/PatricioRios/Compras/models"

type UserRepository interface {
	GetAllUsers() ([]models.User, error)
	GetUserById(id int) (models.User, error)
	CreateUser(user models.User) (models.User, error)
	UpdateUser(user models.User) (models.User, error)
	DeleteUser(id int) error
	GetUserByEmail(email string) (models.User, error)
	GetUserByUserName(userName string) (models.User, error)
}
