package purchase

import (
	"github.com/PatricioRios/Compras/models"
	categoryRepoPgk "github.com/PatricioRios/Compras/repository/v1/category"
	"github.com/PatricioRios/Compras/repository/v1/purchase"
	"github.com/PatricioRios/Compras/utils"
)

// CompraService implementa la interfaz Service y maneja la lógica de negocio.
type CompraService struct {
	purchaseRepository purchase.RepositoryCompra
	categoryRepository categoryRepoPgk.RepositoryCategory
}

// NewCompraService crea una nueva instancia de CompraService.
func NewCompraService(repository purchase.RepositoryCompra, categoryRepository categoryRepoPgk.RepositoryCategory) Service {
	return &CompraService{
		purchaseRepository: repository,
		categoryRepository: categoryRepository,
	}
}

// GetAllPurchases obtiene todas las compras de un usuario.
func (s *CompraService) GetAllPurchases(userId int) (*[]models.Purchase, error) {
	compras, err := s.purchaseRepository.GetAllPurchases(userId)
	if err != nil {
		if err == categoryRepoPgk.ErrRecordNotFound {
			return nil, ErrPurchaseNotFound
		}
		return nil, ErrInternalError
	}
	return &compras, nil
}

// GetPurchaseById obtiene una compra por su ID.
func (s *CompraService) GetPurchaseById(id int, userId int) (models.Purchase, error) {
	if id < 1 {
		return models.Purchase{}, ErrPurchaseInvalidId
	}
	compra, err := s.purchaseRepository.GetPurchaseById(id, userId)
	if err != nil {
		if err == purchase.ErrPurchaseNotFound {
			return models.Purchase{}, ErrPurchaseNotFound
		}
		return compra, ErrInternalError
	}
	return compra, nil
}

// UpdateCompra actualiza una compra existente.
func (s *CompraService) UpdateCompra(compra models.CompraUpdate) (models.Purchase, error) {
	if compra.Id == nil || *compra.Id < 1 {
		return models.Purchase{}, ErrPurchaseInvalidId
	}

	compraFinded, err := s.purchaseRepository.GetPurchaseById(*compra.Id, compra.UserID)
	if err != nil {
		if err == purchase.ErrPurchaseNotFound {
			return models.Purchase{}, ErrPurchaseNotFound
		}
		return models.Purchase{}, ErrInternalError
	}

	// Update category if needed
	if compra.Category != nil || compra.CategoriaID != nil {
		if compra.CategoriaID != nil || compra.Category.Name == "" {
			newCategory, err := s.categoryRepository.GetCategoryById(*compra.CategoriaID, compra.UserID)
			if err != nil {
				if err == categoryRepoPgk.ErrRecordNotFound {
					return compraFinded, ErrCategoryNotExist
				}
				return compraFinded, ErrInternalError
			}
			compraFinded.CategoriaID = newCategory.Id
			compraFinded.Category = newCategory
		} else if compra.Category != nil && compra.Category.Name != "" {
			compraFinded.Category, err = s.categoryRepository.GetCategoryByName(compra.Category.Name, compra.UserID)
			if err != nil {
				if err == categoryRepoPgk.ErrRecordNotFound {
					compra.Category.UserID = compra.UserID
					compraFinded.Category, err = s.categoryRepository.CreateCategory(*compra.Category)
					if err != nil {
						return compraFinded, ErrInternalError
					}
					compraFinded.CategoriaID = compraFinded.Category.Id
				} else {
					return compraFinded, ErrInternalError
				}
			} else {
				compraFinded.CategoriaID = compraFinded.Category.Id
			}
		} else {
			return compraFinded, ErrCategoryBadRequest
		}
	}

	utils.SetIfNotNil(&compraFinded.Title, compra.Title)
	utils.SetIfNotNil(&compraFinded.Description, compra.Description)
	utils.SetIfNotNil(&compraFinded.Import, compra.Import)

	compraFinded, err = s.purchaseRepository.UpdatePurchase(compraFinded)
	if err != nil {
		return compraFinded, ErrInternalError
	}
	return compraFinded, nil
}

// CreatePurchase crea una nueva compra.
func (s *CompraService) CreatePurchase(compra models.Purchase) (models.Purchase, error) {
	if compra.Title == "" {
		return compra, ErrPurchaseBadTitle
	}
	if compra.Description == "" {
		return compra, ErrPurchaseBadDescription
	}
	if compra.Import < 0 {
		return compra, ErrPurchaseNegativeImport
	}
	if compra.CategoriaID < 1 {
		if compra.Category.Name != "" {
			category, err := s.categoryRepository.GetCategoryByName(compra.Category.Name, compra.UserID)
			if err != nil {
				if err == categoryRepoPgk.ErrRecordNotFound {
					compra.Category.UserID = compra.UserID
					category, err = s.categoryRepository.CreateCategory(compra.Category)
					if err != nil {
						return models.Purchase{}, ErrInternalError
					}
					compra.Category = category
					compra.CategoriaID = category.Id
				} else {
					return models.Purchase{}, ErrInternalError
				}
			} else {
				compra.Category = category
				compra.CategoriaID = category.Id
			}
		} else {
			return compra, ErrCategoryInvalid
		}
	}
	if len(compra.Articulos) > 0 {
		importe := 0.0
		for i, articulo := range compra.Articulos {
			if articulo.Name == "" {
				return compra, ErrArticleNameEmpty
			}
			if articulo.Price < 0 {
				return compra, ErrArticlePriceNegative
			}
			importe += articulo.Price
			compra.Articulos[i].Id = 0
			compra.Articulos[i].CompraID = 0
		}
		compra.Import = importe
	}
	compra.Id = 0
	compra, err := s.purchaseRepository.CreatePurchase(compra)
	if err != nil {
		return compra, ErrInternalError
	}
	return compra, nil
}

// DeletePurchase elimina una compra por ID.
func (s *CompraService) DeletePurchase(id int, userId int) error {
	if id < 1 {
		return ErrPurchaseInvalidId
	}
	err := s.purchaseRepository.DeletePurchase(id, userId)
	if err != nil {
		if err == purchase.ErrPurchaseNotFound {
			return ErrCategoryNotFound
		}
		return ErrInternalError
	}
	return nil
}
