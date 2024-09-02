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
func (repo *RepositoryCategoryImpl) GetAllCategories(userId int) ([]models.CategoryPurchase, error) {
	var categorias []models.CategoryPurchase
	if err := repo.MySQLDB.Where("user_id = ?", userId).Find(&categorias).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, ErrRecordNotFound
		}
		return nil, err
	}
	return categorias, nil
}

// GetCategoryById obtiene una categoría por su ID desde la base de datos y verifica que pertenece al userId.
func (repo *RepositoryCategoryImpl) GetCategoryById(id int, userId int) (models.CategoryPurchase, error) {
	var categoria models.CategoryPurchase
	tx := repo.MySQLDB.Preload("Compras").Where("id = ? AND user_id = ?", id, userId).First(&categoria)
	err := tx.Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return categoria, err
		}
		return categoria, err
	}
	return categoria, nil
}

// CreateCategory crea una nueva categoría en la base de datos.
func (repo *RepositoryCategoryImpl) CreateCategory(category models.CategoryPurchase) (models.CategoryPurchase, error) {
	category.Id = 0
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

// DeleteCategory elimina una categoría de la base de datos por su ID y el ID del usuario.
func (repo *RepositoryCategoryImpl) DeleteCategory(id int, userId int) error {
	tx := repo.MySQLDB.Where("id = ? AND user_id = ?", id, userId).Delete(&models.CategoryPurchase{})
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (repo *RepositoryCategoryImpl) GetCategoryByName(name string, userId int) (models.CategoryPurchase, error) {
	var category models.CategoryPurchase
	tx := repo.MySQLDB.Where("name = ? AND user_id = ?", name, userId).First(&category)
	err := tx.Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return category, ErrRecordNotFound
		}
		return category, tx.Error

	}
	return category, nil
}
