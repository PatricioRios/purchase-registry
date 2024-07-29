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
	service *compra.CompraService
}

func NewPurchaseController(logger utils.Logger, service *compra.CompraService) PurchaseController {
	return PurchaseController{
		logger:  logger,
		service: service,
	}
}

// GetVersion
//
//	@Summary	Get version of app application
//	@Tags		Compras
//	@Accept		json
//	@Produce	json
//	@Success	200	{object}	models.Purchase
//	@Failure	500	{object}	utils.ResponseError
//	@Router		/v1/compra [get]
func (ctrl *PurchaseController) GetAllPurchases(c *gin.Context) {

	compras, err := ctrl.service.GetAllPurchases()
	if err != nil {
		c.JSON(err.Code(), gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, compras)
}

func (crl *PurchaseController) GetPurchaseById(c *gin.Context) {
	param := c.Param("id")
	id, _ := strconv.Atoi(param)
	compra, err := crl.service.GetPurchaseById(id)
	if err != nil {
		c.JSON(err.Code(), gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, compra)
}

func (crl *PurchaseController) UpdatePurchase(c *gin.Context) {

	var compraUpdater models.CompraUpdate
	if err := c.BindJSON(&compraUpdater); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	compra, err := crl.service.UpdateCompra(compraUpdater)
	if err != nil {
		c.JSON(err.Code(), gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, compra)
}

func (crl *PurchaseController) CreatePurchase(c *gin.Context) {
	var compra models.Purchase
	if err := c.BindJSON(&compra); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	compra, err := crl.service.CreatePurchase(compra)
	if err != nil {
		c.JSON(err.Code(), gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, compra)
}

func (crl *PurchaseController) DeletePurchase(c *gin.Context) {
	id := c.Param("id")
	compraId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

	}
	errSrvc := crl.service.DeletePurchase(compraId)
	if errSrvc != nil {
		c.JSON(errSrvc.Code(), gin.H{"error": errSrvc.Error()})
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Compra deleted"})
}
