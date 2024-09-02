package purchase

import (
	"errors"

	"github.com/PatricioRios/Compras/models"
)

type Service interface {
	GetAllPurchases(userId int) (*[]models.Purchase, error)
	GetPurchaseById(id int, userId int) (models.Purchase, error)
	UpdateCompra(compra models.CompraUpdate) (models.Purchase, error)
	CreatePurchase(compra models.Purchase) (models.Purchase, error)
	DeletePurchase(id int, userId int) error
}

var (
	ErrPurchaseBadTitle       error = errors.New("PurchaseBadTitle")
	ErrPurchaseBadDescription error = errors.New("PurchaseBadDescription")
	ErrPurchaseNegativeImport error = errors.New("PurchaseNegativeImport")
	ErrPurchaseNotFound       error = errors.New("PurchaseNotFound")
	ErrPurchaseInvalidId      error = errors.New("InvalidPurchaseId")

	ErrCategoryInvalid    error = errors.New("PurchaseInvalidCategory")
	ErrCategoryInvalidId  error = errors.New("PurchaseInvalidId")
	ErrCategoryNotFound   error = errors.New("CategoryNotFound")
	ErrCategoryNotExist   error = errors.New("CategoryNotExist")
	ErrCategoryBadRequest error = errors.New("UseIdOrNameCategoryNotNameAndId")

	ErrArticleNameEmpty     error = errors.New("NameOfArticleIsEmpty")
	ErrArticlePriceNegative error = errors.New("PriceOfArticleIsNegative")

	ErrInternalError error = errors.New("InternalError")
)
