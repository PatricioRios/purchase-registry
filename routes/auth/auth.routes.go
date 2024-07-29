package auth

import (
	"github.com/PatricioRios/Compras/controllers/v1/auth"

	"github.com/PatricioRios/Compras/utils"
)

// MiscRoutes struct
type AuthRoutes struct {
	logger         utils.Logger
	handler        utils.RequestHandler
	authController auth.AuthController
}

// Setup Misc routes
func (s AuthRoutes) Setup() {
	s.logger.Info("Setting up routes COMPRAS")
	api := s.handler.Gin.Group("/api/v1")
	{
		api.GET("/auth/email", s.authController.LogInWithEmail)
		api.GET("/auth/user_name", s.authController.LogInWithUserName)
		api.GET("/auth/refresh_token", s.authController.RefreshToken)
	}
}

// NewMiscRoutes creates new Misc controller
func NewAuthRoutes(
	logger utils.Logger,
	handler utils.RequestHandler,
	controller auth.AuthController,
) AuthRoutes {
	return AuthRoutes{
		handler:        handler,
		logger:         logger,
		authController: controller,
	}
}
