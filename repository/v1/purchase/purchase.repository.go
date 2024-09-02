package purchase

import (
	"errors"

	"github.com/PatricioRios/Compras/models"
)

type RepositoryCompra interface {
	GetAllPurchases(userId int) ([]models.Purchase, error)
	GetPurchaseById(id int, userId int) (models.Purchase, error)
	CreatePurchase(compra models.Purchase) (models.Purchase, error)
	UpdatePurchase(compra models.Purchase) (models.Purchase, error)
	DeletePurchase(id int, userId int) error
	GetPurchaseBySoloId(id int) (models.Purchase, error)
}

var (
	ErrPurchaseNotFound   = errors.New("purchase not found")
	ErrArticleNotFound    = errors.New("article not found")
	ErrPurchaseDeleteFail = errors.New("failed to delete purchase")
	ErrPurchaseCreateFail = errors.New("failed to create purchase")
	ErrPurchaseUpdateFail = errors.New("failed to update purchase")
)
