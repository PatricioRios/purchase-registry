package user

import (
	"github.com/PatricioRios/Compras/controllers/v1/user"

	"github.com/PatricioRios/Compras/utils"
)

// MiscRoutes struct
type UserRoutes struct {
	logger           utils.Logger
	handler          utils.RequestHandler
	CompraController user.UserController
}

// Setup Misc routes
func (s UserRoutes) Setup() {
	s.logger.Info("Setting up routes COMPRAS")
	api := s.handler.Gin.Group("/api/v1")
	{
		api.GET("/user", s.CompraController.GetAllUsers)
		api.GET("/user/:id", s.CompraController.GetUserById)
		api.POST("/user", s.CompraController.CreateUser)
		//api.PUT("/user")
		//api.DELETE("/user/:id")
	}
}

// NewMiscRoutes creates new Misc controller
func NewUserRoutes(
	logger utils.Logger,
	handler utils.RequestHandler,
	controller user.UserController,
) UserRoutes {
	return UserRoutes{
		handler:          handler,
		logger:           logger,
		CompraController: controller,
	}
}
