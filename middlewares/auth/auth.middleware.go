package auth

import (
	"log"
	"time"

	"github.com/PatricioRios/Compras/controllers/v1/auth"
	"github.com/PatricioRios/Compras/utils"
	"github.com/gin-gonic/gin"
)

type AuthMiddleware struct {
	logger         utils.Logger
	handler        utils.RequestHandler
	authController auth.AuthController
}

func NewAuthMidleware(
	logger utils.Logger,
	handler utils.RequestHandler,
	authController auth.AuthController,
) AuthMiddleware {
	return AuthMiddleware{
		handler:        handler,
		logger:         logger,
		authController: authController,
	}
}

// general middleware example (?)
func (m AuthMiddleware) Setup() {

	m.logger.Info("Setup Middleware Auth")

	m.handler.Gin.Use(gin.Recovery(), gin.Logger(), gin.ErrorLogger())

}

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		c.Header("example", "12345")

		// Set example variable
		c.Set("example", "12345")

		// before request

		c.Next()

		// after request
		latency := time.Since(t)
		log.Print(latency)

		// access the status we are sending
		status := c.Writer.Status()
		log.Println(status)
	}
}
