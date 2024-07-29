package routes

import (
	compra_v1 "github.com/PatricioRios/Compras/controllers/v1/purchase"

	"github.com/PatricioRios/Compras/utils"
)

// MiscRoutes struct
type PurchaseRoutes struct {
	logger           utils.Logger
	handler          utils.RequestHandler
	CompraController compra_v1.PurchaseController
}

// Setup Misc routes
func (s PurchaseRoutes) Setup() {
	s.logger.Info("Setting up routes COMPRAS")
	api := s.handler.Gin.Group("/api/v1")
	{
		api.GET("/purchase", s.CompraController.GetAllPurchases)
		api.GET("/purchase/:id", s.CompraController.GetPurchaseById)
		api.POST("/purchase", s.CompraController.CreatePurchase)
		api.PUT("/purchase", s.CompraController.UpdatePurchase)
		api.DELETE("/purchase/:id", s.CompraController.DeletePurchase)
	}
}

// NewMiscRoutes creates new Misc controller
func NewCompraRoutes(
	logger utils.Logger,
	handler utils.RequestHandler,
	controller compra_v1.PurchaseController,
) PurchaseRoutes {
	return PurchaseRoutes{
		handler:          handler,
		logger:           logger,
		CompraController: controller,
	}
}
