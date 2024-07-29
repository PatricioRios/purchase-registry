package purchase

import (
	"github.com/PatricioRios/Compras/models"
	"gorm.io/gorm"
)

// RepositoryCompraImpl is the implementation of the RepositoryCompra interface using GORM.
type RepositoryCompraImpl struct {
	MySQLDB *gorm.DB
}

// NewRepositoryCompraImpl creates a new instance of RepositoryCompraImpl
func NewRepositoryCompraImpl(db *gorm.DB) RepositoryCompra {
	return &RepositoryCompraImpl{MySQLDB: db}
}

// DeletePurchase deletes a compra by ID
func (repository *RepositoryCompraImpl) DeletePurchase(id int) error {
	var compra models.Purchase
	tx := repository.MySQLDB.Delete(&compra, id)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

// GetAllPurchases retrieves all compras from the database
func (repository *RepositoryCompraImpl) GetAllPurchases() ([]models.Purchase, error) {
	var compras []models.Purchase
	tx := repository.MySQLDB.
		//Preload("Articulos").
		//Select("id", "title", "description", "import", "categoria_id", "created_at", "updated_at"). // Selecciona solo los campos necesarios
		Find(&compras)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return compras, nil
}

// GetPurchaseById retrieves a compra by ID from the database
func (repository *RepositoryCompraImpl) GetPurchaseById(id int) (models.Purchase, error) {
	var compra models.Purchase
	tx := repository.MySQLDB.Preload("Category").Preload("Articulos").First(&compra, id)
	if tx.Error != nil {
		return compra, tx.Error
	}
	return compra, nil
}

// CreatePurchase creates a new compra in the database
func (repository *RepositoryCompraImpl) CreatePurchase(compra models.Purchase) (models.Purchase, error) {
	tx := repository.MySQLDB.Create(&compra)
	if tx.Error != nil {
		return compra, tx.Error
	}
	return compra, nil
}

// UpdatePurchase updates an existing compra in the database
func (repository *RepositoryCompraImpl) UpdatePurchase(compra models.Purchase) (models.Purchase, error) {
	tx := repository.MySQLDB.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&compra)
	if tx.Error != nil {
		return compra, tx.Error
	}
	return compra, nil
}
