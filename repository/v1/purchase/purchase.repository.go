package purchase

import "github.com/PatricioRios/Compras/models"

type RepositoryCompra interface {
	GetAllPurchases() ([]models.Purchase, error)
	GetPurchaseById(id int) (models.Purchase, error)
	CreatePurchase(compra models.Purchase) (models.Purchase, error)
	UpdatePurchase(compra models.Purchase) (models.Purchase, error)
	DeletePurchase(id int) error
}
