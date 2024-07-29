package routes

import (
	misc_v1 "github.com/PatricioRios/Compras/controllers/v1/misc"
	"github.com/PatricioRios/Compras/utils"
)

// MiscRoutes struct
type MiscRoutes struct {
	logger         utils.Logger
	handler        utils.RequestHandler
	miscController misc_v1.MiscController
}

// Setup Misc routes
func (s MiscRoutes) Setup() {
	s.logger.Info("Setting up routes")
	api := s.handler.Gin.Group("/api/v1")
	{
		api.GET("/liveness", s.miscController.GetLiveness)
		api.GET("/readiness", s.miscController.GetReadiness)
		api.GET("/version", s.miscController.GetVersion)

	}
}

// NewMiscRoutes creates new Misc controller
func NewMiscRoutes(
	logger utils.Logger,
	handler utils.RequestHandler,
	miscController misc_v1.MiscController,
) MiscRoutes {
	return MiscRoutes{
		handler:        handler,
		logger:         logger,
		miscController: miscController,
	}
}
