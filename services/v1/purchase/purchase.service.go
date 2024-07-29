package purchase

import (
	"fmt"
	"net/http"

	"github.com/PatricioRios/Compras/models"
	"github.com/PatricioRios/Compras/repository/v1/category"
	"github.com/PatricioRios/Compras/repository/v1/purchase"
	"github.com/PatricioRios/Compras/utils"
	"gorm.io/gorm"
)

type CompraService struct {
	purchaseRepository purchase.RepositoryCompra
	categoryRepository category.RepositoryCategory
}

func NewCompraService(repository purchase.RepositoryCompra, categoryRepository category.RepositoryCategory) *CompraService {
	return &CompraService{
		purchaseRepository: repository,
		categoryRepository: categoryRepository,
	}
}

func (s *CompraService) GetAllPurchases() (*[]models.Purchase, utils.SrvcError) {
	compras, err := s.purchaseRepository.GetAllPurchases()
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, utils.NewError(http.StatusNoContent, "Compras no encontradas")
		}
		return nil, utils.NewError(http.StatusInternalServerError, err.Error())
	}
	return &compras, nil
}

func (s *CompraService) GetPurchaseById(id int) (models.Purchase, utils.SrvcError) {
	if id < 1 {
		return models.Purchase{}, utils.NewError(http.StatusBadRequest, "La id no puede ser menor a 1")
	}
	compra, err := s.purchaseRepository.GetPurchaseById(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return models.Purchase{}, utils.NewError(http.StatusNotFound, "Compra no encontrada")
		}
		return compra, utils.NewError(http.StatusInternalServerError, err.Error())
	}
	return compra, nil
}

func (s *CompraService) UpdateCompra(compra models.CompraUpdate) (models.Purchase, utils.SrvcError) {
	var compraFinded models.Purchase

	if compra.Id == nil || *compra.Id < 1 {
		return models.Purchase{}, utils.NewError(http.StatusBadRequest, "Id Invalida")
	}

	if compra.Import != nil && *compra.Import < 0 {
		return compraFinded, utils.NewError(http.StatusBadRequest, "El importe no puede ser negativo")
	}

	compraFinded, err := s.purchaseRepository.GetPurchaseById(*compra.Id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return compraFinded, utils.NewError(http.StatusNoContent, "Compra no encontrada")
		}
		return compraFinded, utils.NewError(http.StatusInternalServerError, err.Error())
	}

	//The first paramether for actualize is CategoriaId
	if compra.Category != nil || compra.CategoriaID != nil {
		fmt.Println("ENTRO")
		if compra.CategoriaID != nil {
			fmt.Println("hay una id")

			newCatgory, err := s.categoryRepository.GetCategoryById(*compra.CategoriaID)
			if err != nil {
				if err == gorm.ErrRecordNotFound {
					return compraFinded, utils.NewError(http.StatusNotFound, "Categoría no encontrada")
				} else {
					return compraFinded, utils.NewError(http.StatusInternalServerError, err.Error())
				}
			}
			compraFinded.CategoriaID = newCatgory.Id
			compraFinded.Category = newCatgory
		} else {
			fmt.Println("hay un nombre")

			newCategory, err := s.categoryRepository.GetCategoryByName(compra.Category.Name)
			if err != nil {
				if err == gorm.ErrRecordNotFound {
					var category models.CategoryPurchase
					category.Name = compra.Category.Name
					newCategory, err = s.categoryRepository.CreateCategory(category)
					if err != nil {
						return compraFinded, utils.NewError(http.StatusInternalServerError, err.Error())
					}
				} else {
					return compraFinded, utils.NewError(http.StatusInternalServerError, err.Error())
				}
			}
			if newCategory.Name != "" {
				fmt.Println(newCategory.Name)
				compraFinded.CategoriaID = newCategory.Id
			}
			compraFinded.Category = newCategory
		}
	}
	utils.SetIfNotNil(&compraFinded.Title, compra.Title)
	utils.SetIfNotNil(&compraFinded.Description, compra.Description)
	utils.SetIfNotNil(&compraFinded.Import, compra.Import)

	compraFinded, err = s.purchaseRepository.UpdatePurchase(compraFinded)
	if err != nil {
		return compraFinded, utils.SrvcErrorImpl{
			CodeError: http.StatusInternalServerError,
			Message:   err.Error(),
		}
	}
	return compraFinded, nil
}

func (s *CompraService) CreatePurchase(compra models.Purchase) (models.Purchase, utils.SrvcError) {
	if compra.Title == "" {
		return compra, utils.NewBadRequest("El título es obligatorio")
	}
	if compra.Description == "" {

		return compra, utils.NewBadRequest("La descripción es obligatoria")
	}
	if compra.Import < 0 {
		return compra, utils.NewBadRequest("El importe no puede ser negativo")
	}
	if compra.CategoriaID < 1 {
		return compra, utils.NewBadRequest("La categoría es obligatoria")
	}
	importe := 0.0
	for i, artículo := range compra.Articulos {
		if artículo.Name == "" {
			return compra, utils.NewBadRequest("El nombre del artículo es obligatorio")
		}
		if artículo.Price < 0 {
			return compra, utils.NewBadRequest("El precio del artículo" + artículo.Name + " no puede ser negativo")
		}
		importe += artículo.Price
		compra.Articulos[i].Id = 0
		compra.Articulos[i].CompraID = 0
	}
	compra.Id = 0
	compra.Import = importe
	compra, err := s.purchaseRepository.CreatePurchase(compra)
	if err != nil {
		return compra, utils.NewError(http.StatusInternalServerError, err.Error())
	}
	return compra, nil
}

func (s *CompraService) DeletePurchase(id int) utils.SrvcError {
	if id < 1 {
		return utils.NewError(http.StatusBadRequest, "La id no puede ser menor a 1")
	}
	err := s.purchaseRepository.DeletePurchase(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return utils.NewError(http.StatusNotFound, "Compra no encontrada")
		}
		return utils.NewError(http.StatusInternalServerError, err.Error())
	}
	return nil
}
