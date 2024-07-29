package category

import (
	"github.com/PatricioRios/Compras/models"
	"gorm.io/gorm"
)

// RepositoryCategoryImpl es la implementación de la interfaz RepositoryCategory usando GORM.
type RepositoryCategoryImpl struct {
	MySQLDB *gorm.DB
}

// NewRepositoryCategoryImpl crea una nueva instancia de RepositoryCategoryImpl.
func NewRepositoryCategoryImpl(db *gorm.DB) RepositoryCategory {
	return &RepositoryCategoryImpl{MySQLDB: db}
}

// GetAllCategories obtiene todas las categorías desde la base de datos.
func (repo *RepositoryCategoryImpl) GetAllCategories() ([]models.CategoryPurchase, error) {
	var categorias []models.CategoryPurchase
	tx := repo.MySQLDB.Find(&categorias)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return categorias, nil
}

// GetCategoryById obtiene una categoría por su ID desde la base de datos.
func (repo *RepositoryCategoryImpl) GetCategoryById(id int) (models.CategoryPurchase, error) {
	var categoria models.CategoryPurchase
	tx := repo.MySQLDB.Preload("Compras").First(&categoria, id)
	if tx.Error != nil {
		return categoria, tx.Error
	}
	return categoria, nil
}

// CreateCategory crea una nueva categoría en la base de datos.
func (repo *RepositoryCategoryImpl) CreateCategory(category models.CategoryPurchase) (models.CategoryPurchase, error) {
	tx := repo.MySQLDB.Create(&category)
	if tx.Error != nil {
		return category, tx.Error
	}
	return category, nil
}

// UpdateCategory actualiza una categoría existente en la base de datos.
func (repo *RepositoryCategoryImpl) UpdateCategory(category models.CategoryPurchase) (models.CategoryPurchase, error) {
	tx := repo.MySQLDB.Save(&category)
	if tx.Error != nil {
		return category, tx.Error
	}
	return category, nil
}

// DeleteCategory elimina una categoría de la base de datos por su ID.
func (repo *RepositoryCategoryImpl) DeleteCategory(id int) error {
	tx := repo.MySQLDB.Delete(&models.CategoryPurchase{}, id)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (repo *RepositoryCategoryImpl) GetCategoryByName(name string) (models.CategoryPurchase, error) {
	var article models.CategoryPurchase
	tx := repo.MySQLDB.Where("name = ?", name).First(&article)
	if tx.Error != nil {
		return article, tx.Error
	}
	return article, nil
}
