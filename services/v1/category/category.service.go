package category

import (
	"net/http"

	"github.com/PatricioRios/Compras/models"
	"github.com/PatricioRios/Compras/repository/v1/category"
	"github.com/PatricioRios/Compras/utils"
	"gorm.io/gorm"
)

// CategoryService maneja operaciones de categorías usando el repositorio de categorías.
type CategoryService struct {
	categoryRepository category.RepositoryCategory
}

// NewCategoryService crea una nueva instancia de CategoryService.
func NewCategoryService(categoryRepository category.RepositoryCategory) *CategoryService {
	return &CategoryService{
		categoryRepository: categoryRepository,
	}
}

// GetAllCategories obtiene todas las categorías.
func (s *CategoryService) GetAllCategories() (*[]models.CategoryPurchase, utils.SrvcError) {
	categorias, err := s.categoryRepository.GetAllCategories()
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, utils.NewError(http.StatusNotFound, "Categorías no encontradas")
		}
		return nil, utils.NewError(http.StatusInternalServerError, err.Error())
	}
	return &categorias, nil
}

// GetCategoryById obtiene una categoría por su ID.
func (s *CategoryService) GetCategoryById(id int) (models.CategoryPurchase, utils.SrvcError) {
	if id < 1 {
		return models.CategoryPurchase{}, utils.NewError(http.StatusBadRequest, "El ID no puede ser menor a 1")
	}
	categoria, err := s.categoryRepository.GetCategoryById(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return models.CategoryPurchase{}, utils.NewError(http.StatusNotFound, "Categoría no encontrada")
		}
		return categoria, utils.NewError(http.StatusInternalServerError, err.Error())
	}
	return categoria, nil
}

// CreateCategory crea una nueva categoría.
func (s *CategoryService) CreateCategory(category models.CategoryPurchase) (models.CategoryPurchase, utils.SrvcError) {
	if category.Name == "" {
		return category, utils.NewBadRequest("El nombre de la categoría es obligatorio")
	}

	category, err := s.categoryRepository.CreateCategory(category)
	if err != nil {
		return category, utils.NewError(http.StatusInternalServerError, err.Error())
	}
	return category, nil
}

// UpdateCategory actualiza una categoría existente.
func (s *CategoryService) UpdateCategory(category models.CategoryPurchase) (models.CategoryPurchase, utils.SrvcError) {
	if category.Id < 1 {
		return models.CategoryPurchase{}, utils.NewError(http.StatusBadRequest, "ID inválido")
	}

	// Obtener la categoría existente para verificar su existencia antes de actualizar
	categoriaExistente, err := s.categoryRepository.GetCategoryById(category.Id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return categoriaExistente, utils.NewError(http.StatusNotFound, "Categoría no encontrada")
		}
		return categoriaExistente, utils.NewError(http.StatusInternalServerError, err.Error())
	}

	// Actualizar los campos de la categoría
	categoriaExistente.Name = category.Name

	categoriaActualizada, err := s.categoryRepository.UpdateCategory(categoriaExistente)
	if err != nil {
		return categoriaActualizada, utils.NewError(http.StatusInternalServerError, err.Error())
	}
	return categoriaActualizada, nil
}

// DeleteCategory elimina una categoría por su ID.
func (s *CategoryService) DeleteCategory(id int) utils.SrvcError {
	if id < 1 {
		return utils.NewError(http.StatusBadRequest, "El ID no puede ser menor a 1")
	}
	err := s.categoryRepository.DeleteCategory(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return utils.NewError(http.StatusNotFound, "Categoría no encontrada")
		}
		return utils.NewError(http.StatusInternalServerError, err.Error())
	}
	return nil
}
