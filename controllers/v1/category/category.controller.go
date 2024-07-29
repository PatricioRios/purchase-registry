package category

import (
	"net/http"
	"strconv"

	"github.com/PatricioRios/Compras/models"
	"github.com/PatricioRios/Compras/services/v1/category"
	"github.com/PatricioRios/Compras/utils"
	"github.com/gin-gonic/gin"
)

// CategoryController maneja las solicitudes HTTP relacionadas con las categorías.
type CategoryController struct {
	logger  utils.Logger
	service *category.CategoryService
}

func NewCategoryController(logger utils.Logger, service *category.CategoryService) CategoryController {
	return CategoryController{
		logger:  logger,
		service: service,
	}
}

// @Summary Get all categories
// @Description Get all categories
// @Tags Categories
// @Accept json
// @Produce json
// @Success 200 {array} models.CategoryPurchase
// @Failure 500 {object} utils.ResponseError
// @Router /v1/category [get]
func (ctrl *CategoryController) GetAllCategories(c *gin.Context) {
	categorias, err := ctrl.service.GetAllCategories()
	if err != nil {
		c.JSON(err.Code(), utils.ResponseError{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, categorias)
}

// @Summary Get category by ID
// @Description Get a category by its ID
// @Tags Categories
// @Accept json
// @Produce json
// @Param id path int true "Category ID"
// @Success 200 {object} models.CategoryPurchase
// @Failure 400 {object} utils.ResponseError "Invalid ID"
// @Failure 404 {object} utils.ResponseError "Category not found"
// @Failure 500 {object} utils.ResponseError
// @Router /v1/category/{id} [get]
func (ctrl *CategoryController) GetCategoryById(c *gin.Context) {
	param := c.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ResponseError{Message: "ID inválido"})
		return
	}

	categoria, srvcError := ctrl.service.GetCategoryById(id)
	if srvcError != nil {
		c.JSON(srvcError.Code(), utils.ResponseError{Message: srvcError.Error()})
		return
	}
	c.JSON(http.StatusOK, categoria)
}

// @Summary Create a new category
// @Description Create a new category
// @Tags Categories
// @Accept json
// @Produce json
// @Param category body models.CategoryPurchase true "Category to create"
// @Success 201 {object} models.CategoryPurchase
// @Failure 400 {object} utils.ResponseError "Invalid input"
// @Failure 500 {object} utils.ResponseError
// @Router /v1/category [post]
func (ctrl *CategoryController) CreateCategory(c *gin.Context) {
	var categoria models.CategoryPurchase
	if err := c.BindJSON(&categoria); err != nil {
		c.JSON(http.StatusBadRequest, utils.ResponseError{Message: err.Error()})
		return
	}

	categoria, err := ctrl.service.CreateCategory(categoria)
	if err != nil {
		c.JSON(err.Code(), utils.ResponseError{Message: err.Error()})
		return
	}
	c.JSON(http.StatusCreated, categoria)
}

// @Summary Update an existing category
// @Description Update an existing category
// @Tags Categories
// @Accept json
// @Produce json
// @Param category body models.CategoryPurchase true "Category to update"
// @Success 200 {object} models.CategoryPurchase
// @Failure 400 {object} utils.ResponseError "Invalid input"
// @Failure 500 {object} utils.ResponseError
// @Router /v1/category [put]
func (ctrl *CategoryController) UpdateCategory(c *gin.Context) {
	var categoria models.CategoryPurchase
	if err := c.BindJSON(&categoria); err != nil {
		c.JSON(http.StatusBadRequest, utils.ResponseError{Message: err.Error()})
		return
	}

	categoriaActualizada, err := ctrl.service.UpdateCategory(categoria)
	if err != nil {
		c.JSON(err.Code(), utils.ResponseError{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, categoriaActualizada)
}

// @Summary Delete a category by ID
// @Description Delete a category by its ID
// @Tags Categories
// @Accept json
// @Produce json
// @Param id path int true "Category ID"
// @Success 200 {object} gin.H "Category deleted message"
// @Failure 400 {object} utils.ResponseError "Invalid ID"
// @Failure 500 {object} utils.ResponseError
// @Router /v1/category/{id} [delete]
func (ctrl *CategoryController) DeleteCategory(c *gin.Context) {
	id := c.Param("id")
	categoriaId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ResponseError{Message: "ID inválido"})
		return
	}

	errSrvc := ctrl.service.DeleteCategory(categoriaId)
	if errSrvc != nil {
		c.JSON(errSrvc.Code(), utils.ResponseError{Message: errSrvc.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Categoría eliminada"})
}
