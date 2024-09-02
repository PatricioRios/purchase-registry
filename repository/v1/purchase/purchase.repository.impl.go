package purchase

import (
	"errors"

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

// DeletePurchase elimina una compra por su ID y userId, y elimina los artículos relacionados
func (repository *RepositoryCompraImpl) DeletePurchase(id int, userId int) error {
	// Primero elimina los artículos relacionados
	if err := repository.MySQLDB.Where("compra_id = ?", id).Delete(&models.Article{}).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrArticleNotFound
		}
		return ErrPurchaseDeleteFail
	}

	// Luego elimina la compra
	if err := repository.MySQLDB.Where("id = ? AND user_id = ?", id, userId).Delete(&models.Purchase{}).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrPurchaseNotFound
		}
		return ErrPurchaseDeleteFail
	}
	return nil
}

// GetAllPurchases retrieves all compras from the database
func (repository *RepositoryCompraImpl) GetAllPurchases(userId int) ([]models.Purchase, error) {
	var compras []models.Purchase
	err := repository.MySQLDB.Where("user_id = ?", userId).Preload("Category").Find(&compras).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrPurchaseNotFound
		}
		return nil, err
	}
	return compras, nil
}

// GetPurchaseById retrieves a compra by ID from the database
func (repository *RepositoryCompraImpl) GetPurchaseById(id int, userId int) (models.Purchase, error) {
	var compra models.Purchase
	tx := repository.MySQLDB.
		Preload("Category").
		Preload("Articulos").
		Where("id = ? AND user_id = ?", id, userId).
		First(&compra)
	err := tx.Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return compra, ErrPurchaseNotFound
		}
		return compra, err
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
	tx := repository.MySQLDB.Session(&gorm.Session{FullSaveAssociations: true}).Preload("Category").Where("user_id = ?", compra.UserID).Updates(&compra)
	if tx.Error != nil {
		return compra, ErrPurchaseUpdateFail
	}
	return compra, nil
}

func (repository *RepositoryCompraImpl) GetPurchaseBySoloId(id int) (models.Purchase, error) {
	var compra models.Purchase
	tx := repository.MySQLDB.Where("id = ?", id).First(&compra)
	err := tx.Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return compra, ErrPurchaseNotFound
		}
		return compra, err
	}
	return compra, nil
}
