package category

import "github.com/PatricioRios/Compras/models"

type RepositoryCategory interface {
	GetAllCategories() ([]models.CategoryPurchase, error)
	GetCategoryById(id int) (models.CategoryPurchase, error)
	CreateCategory(category models.CategoryPurchase) (models.CategoryPurchase, error)
	UpdateCategory(category models.CategoryPurchase) (models.CategoryPurchase, error)
	DeleteCategory(id int) error
	GetCategoryByName(name string) (models.CategoryPurchase, error)
}
