package category

//capa de logica de negocio
import (
	"github.com/PatricioRios/Compras/models"
	categoryPkg "github.com/PatricioRios/Compras/repository/v1/category"
)

// CategoryService maneja operaciones de categorías usando el repositorio de categorías.
func mapRepositoryError(err error) error {
	switch err {
	case categoryPkg.ErrRecordNotFound:
		return ErrNotFound
	default:
		return ErrInternalError
	}
}

// CategoryService maneja operaciones de categorías usando el repositorio de categorías.
type CategoryService struct {
	categoryRepository categoryPkg.RepositoryCategory
}

// NewCategoryService crea una nueva instancia de CategoryService.
func NewCategoryService(categoryRepository categoryPkg.RepositoryCategory) Service {
	return &CategoryService{
		categoryRepository: categoryRepository,
	}
}

// GetAllCategories obtiene todas las categorías.
func (s *CategoryService) GetAllCategories(userId int) (*[]models.CategoryPurchase, error) {
	categorias, err := s.categoryRepository.GetAllCategories(userId)
	if err != nil {
		return nil, mapRepositoryError(err)
	}
	return &categorias, nil
}

// GetCategoryById obtiene una categoría por su ID.
func (s *CategoryService) GetCategoryById(id int, userId int) (models.CategoryPurchase, error) {
	if id < 1 {
		return models.CategoryPurchase{}, ErrInvalidId
	}
	categoria, err := s.categoryRepository.GetCategoryById(id, userId)
	if err != nil {
		return models.CategoryPurchase{}, mapRepositoryError(err)
	}
	return categoria, nil
}

// CreateCategory crea una nueva categoría.
func (s *CategoryService) CreateCategory(category models.CategoryPurchase) (models.CategoryPurchase, error) {
	if category.Name == "" {
		return category, ErrInvalidName
	}

	category, err := s.categoryRepository.CreateCategory(category)
	if err != nil {
		return category, mapRepositoryError(err)
	}
	return category, nil
}

// UpdateCategory actualiza una categoría existente.
func (s *CategoryService) UpdateCategory(category models.CategoryPurchase) (models.CategoryPurchase, error) {
	if category.Id < 1 {
		return models.CategoryPurchase{}, ErrInvalidId
	}

	categoriaExistente, err := s.categoryRepository.GetCategoryById(category.Id, category.UserID)
	if err != nil {
		return categoriaExistente, mapRepositoryError(err)
	}

	categoriaExistente.Name = category.Name

	categoriaActualizada, err := s.categoryRepository.UpdateCategory(categoriaExistente)
	if err != nil {
		return categoriaActualizada, mapRepositoryError(err)
	}
	return categoriaActualizada, nil
}

// DeleteCategory elimina una categoría por su ID.
func (s *CategoryService) DeleteCategory(id int, userId int) error {
	if id < 1 {
		return ErrInvalidId
	}
	err := s.categoryRepository.DeleteCategory(id, userId)
	if err != nil {
		return mapRepositoryError(err)
	}
	return nil
}
