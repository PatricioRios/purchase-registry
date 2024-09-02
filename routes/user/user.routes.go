package user

import (
	"github.com/PatricioRios/Compras/controllers/v1/auth"
	"github.com/PatricioRios/Compras/controllers/v1/user"

	"github.com/PatricioRios/Compras/utils"
)

// MiscRoutes struct
type UserRoutes struct {
	logger         utils.Logger
	handler        utils.RequestHandler
	userController user.UserController
	AuthController auth.AuthController
}

// Setup Misc routes
func (s UserRoutes) Setup() {
	s.logger.Info("Setting up routes COMPRAS")
	api := s.handler.Gin.Group("/api/v1")
	{
		api.GET("/user", s.AuthController.ProtectedHandler(), s.userController.GetAllUsers)
		api.GET("/user/:id", s.AuthController.ProtectedHandler(), s.userController.GetUserById)
		api.POST("/user", s.userController.CreateUser)
		//api.PUT("/user")
		//api.DELETE("/user/:id")
	}
}

// NewMiscRoutes creates new Misc controller
func NewUserRoutes(
	logger utils.Logger,
	handler utils.RequestHandler,
	controller user.UserController,
	AuthController auth.AuthController,
) UserRoutes {
	return UserRoutes{
		handler:        handler,
		logger:         logger,
		userController: controller,
		AuthController: AuthController,
	}
}
