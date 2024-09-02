package category

import (
	"errors"

	"github.com/PatricioRios/Compras/models"
)

var (
	ErrBadName       error = errors.New("CategoryBadName")
	ErrNotFound      error = errors.New("CategoryNotFound")
	ErrInternalError error = errors.New("InternalError")
	ErrInvalidId     error = errors.New("CategoryInvalidId")
	ErrInvalidName   error = errors.New("CategoryBadName")
)

type Service interface {
	GetAllCategories(userId int) (*[]models.CategoryPurchase, error)
	GetCategoryById(id int, userId int) (models.CategoryPurchase, error)
	CreateCategory(category models.CategoryPurchase) (models.CategoryPurchase, error)
	UpdateCategory(category models.CategoryPurchase) (models.CategoryPurchase, error)
	DeleteCategory(id int, userId int) error
}
