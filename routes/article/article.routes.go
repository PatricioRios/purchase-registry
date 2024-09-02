package article

import (
	"github.com/PatricioRios/Compras/controllers/v1/article"

	"github.com/PatricioRios/Compras/utils"
)

// MiscRoutes struct
type AuthRoutes struct {
	logger            utils.Logger
	handler           utils.RequestHandler
	articleController article.ArticleController
}

// Setup Misc routes
func (s AuthRoutes) Setup() {
	s.logger.Info("Setting up routes COMPRAS")
	api := s.handler.Gin.Group("/api/v1/article")
	{
		api.POST("", s.articleController.CreateArticle)
		api.PUT("/user_name", s.articleController.UpdateArticle)
		api.DELETE("/:purchase_id/:article_id", s.articleController.DeleteArticle)
	}
}

// NewMiscRoutes creates new Misc controller
func NewAuthRoutes(
	logger utils.Logger,
	handler utils.RequestHandler,
	controller article.ArticleController,
) AuthRoutes {
	return AuthRoutes{
		handler:           handler,
		logger:            logger,
		articleController: controller,
	}
}
