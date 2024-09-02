package purchase

import (
	"net/http"
	"strconv"

	"github.com/PatricioRios/Compras/models"
	compra "github.com/PatricioRios/Compras/services/v1/purchase"
	"github.com/PatricioRios/Compras/utils"
	"github.com/gin-gonic/gin"
)

type PurchaseController struct {
	logger  utils.Logger
	service compra.Service
}

func NewPurchaseController(logger utils.Logger, service compra.Service) PurchaseController {
	return PurchaseController{
		logger:  logger,
		service: service,
	}
}

// mapErrorToHTTPStatus maps service errors to HTTP status codes.
func mapErrorToHTTPStatus(err error) int {
	switch err {
	case compra.ErrPurchaseBadTitle, compra.ErrPurchaseBadDescription, compra.ErrPurchaseNegativeImport, compra.ErrCategoryInvalid, compra.ErrCategoryInvalidId:
		return http.StatusBadRequest
	case compra.ErrCategoryNotFound:
		return http.StatusNotFound
	case compra.ErrInternalError:
		return http.StatusInternalServerError
	default:
		return http.StatusInternalServerError
	}
}

// GetAllPurchases handles the retrieval of all purchases for a user.
func (ctrl *PurchaseController) GetAllPurchases(c *gin.Context) {
	userId, exist := c.Get("user_id")
	if !exist {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User ID is required"})
		return
	}
	compras, err := ctrl.service.GetAllPurchases(userId.(int))
	if err != nil {
		c.JSON(mapErrorToHTTPStatus(err), gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, compras)
}

// GetPurchaseById handles the retrieval of a single purchase by ID.
func (ctrl *PurchaseController) GetPurchaseById(c *gin.Context) {
	param := c.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil || id < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	userId, exist := c.Get("user_id")
	if !exist {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User ID is required"})
		return
	}
	compra, err := ctrl.service.GetPurchaseById(id, userId.(int))
	if err != nil {
		c.JSON(mapErrorToHTTPStatus(err), gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, compra)
}

// UpdatePurchase handles updating an existing purchase.
func (ctrl *PurchaseController) UpdatePurchase(c *gin.Context) {
	var compraUpdater models.CompraUpdate
	if err := c.BindJSON(&compraUpdater); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON input"})
		return
	}

	userId, exist := c.Get("user_id")
	if !exist {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User ID is required"})
		return
	}
	compraUpdater.UserID = userId.(int)

	compra, err := ctrl.service.UpdateCompra(compraUpdater)
	if err != nil {
		c.JSON(mapErrorToHTTPStatus(err), gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, compra)
}

// CreatePurchase handles creating a new purchase.
func (ctrl *PurchaseController) CreatePurchase(c *gin.Context) {
	var compra models.Purchase
	if err := c.BindJSON(&compra); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON input"})
		return
	}

	userId, exist := c.Get("user_id")
	if !exist {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User ID is required"})
		return
	}
	compra.UserID = userId.(int)

	compra, err := ctrl.service.CreatePurchase(compra)
	if err != nil {
		c.JSON(mapErrorToHTTPStatus(err), gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, compra)
}

// DeletePurchase handles deleting a purchase by ID.
func (ctrl *PurchaseController) DeletePurchase(c *gin.Context) {
	id := c.Param("id")
	compraId, err := strconv.Atoi(id)
	if err != nil || compraId < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	userId, exist := c.Get("user_id")
	if !exist {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User ID is required"})
		return
	}

	err = ctrl.service.DeletePurchase(compraId, userId.(int))
	if err != nil {
		c.JSON(mapErrorToHTTPStatus(err), gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Purchase deleted successfully"})
}
