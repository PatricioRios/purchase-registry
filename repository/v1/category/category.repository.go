package category

import (
	"errors"

	"github.com/PatricioRios/Compras/models"
)

type RepositoryCategory interface {
	GetAllCategories(userId int) ([]models.CategoryPurchase, error)
	GetCategoryById(id int, userId int) (models.CategoryPurchase, error)
	CreateCategory(category models.CategoryPurchase) (models.CategoryPurchase, error)
	UpdateCategory(category models.CategoryPurchase) (models.CategoryPurchase, error)
	DeleteCategory(id int, userId int) error
	GetCategoryByName(name string, userId int) (models.CategoryPurchase, error)
}

var ErrRecordNotFound = errors.New("RecordNotFound")
