package user

import (
	"github.com/PatricioRios/Compras/models"
	"gorm.io/gorm"
)

/*
	GetAllUsers() ([]models.User, error)
	GetUserById(id int) (models.User, error)
	CreateUser(user models.User) (models.User, error)
	UpdateUser(user models.User) (models.User, error)
	DeleteUser(id int) error
*/

type UserRepositoryImpl struct {
	MySQLDB *gorm.DB
}

func NewUserRepositoryImpl(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{MySQLDB: db}
}

func (r *UserRepositoryImpl) GetAllUsers() ([]models.User, error) {
	var users []models.User
	tx := r.MySQLDB.Find(&users)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return users, nil
}

func (r *UserRepositoryImpl) GetUserById(id int) (models.User, error) {
	var user models.User
	tx := r.MySQLDB.First(&user, id)
	if tx.Error != nil {
		return user, tx.Error
	}
	return user, nil
}

func (r *UserRepositoryImpl) CreateUser(user models.User) (models.User, error) {
	tx := r.MySQLDB.Create(&user)
	if tx.Error != nil {
		return user, tx.Error
	}
	return user, nil
}

func (r *UserRepositoryImpl) UpdateUser(user models.User) (models.User, error) {
	tx := r.MySQLDB.Save(&user)
	if tx.Error != nil {
		return user, tx.Error
	}
	return user, nil
}

func (r *UserRepositoryImpl) DeleteUser(id int) error {
	var user models.User
	tx := r.MySQLDB.Delete(&user, id)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (r *UserRepositoryImpl) GetUserByEmail(email string) (models.User, error) {
	var user models.User
	tx := r.MySQLDB.Where("email = ?", email).First(&user)
	if tx.Error != nil {
		return user, tx.Error
	}
	return user, nil
}

func (r *UserRepositoryImpl) GetUserByUserName(userName string) (models.User, error) {

	var user models.User
	tx := r.MySQLDB.Where("user_name = ?", userName).First(&user)
	if tx.Error != nil {
		return user, tx.Error
	}
	return user, nil
}
