package routes

import (
	"github.com/PatricioRios/Compras/controllers/v1/auth"
	"github.com/PatricioRios/Compras/controllers/v1/category"

	"github.com/PatricioRios/Compras/utils"
)

// MiscRoutes struct
type CategoryRoutes struct {
	logger           utils.Logger
	handler          utils.RequestHandler
	CompraController category.CategoryController
	AuthController   auth.AuthController
}

// Setup Misc routes
func (s CategoryRoutes) Setup() {
	s.logger.Info("Setting up routes CATEGORY")
	api := s.handler.Gin.Group("/api/v1")
	api.Use(s.AuthController.ProtectedHandler())
	{
		api.GET("/category", s.CompraController.GetAllCategories)
		api.GET("/category/:id", s.CompraController.GetCategoryById)
		api.POST("/category", s.CompraController.CreateCategory)
		api.PUT("/category", s.CompraController.UpdateCategory)
		api.DELETE("/category/:id", s.CompraController.DeleteCategory)
	}
}

// NewMiscRoutes creates new Misc controller
func NewCategoryRoutes(
	logger utils.Logger,
	handler utils.RequestHandler,
	controller category.CategoryController,
	authController auth.AuthController,
) CategoryRoutes {
	return CategoryRoutes{
		handler:          handler,
		logger:           logger,
		CompraController: controller,
		AuthController:   authController,
	}
}
