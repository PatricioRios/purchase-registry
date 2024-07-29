package utils

import (
	"github.com/PatricioRios/Compras/docs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// RequestHandler function
type RequestHandler struct {
	Gin *gin.Engine
}

type ResponseError struct {
	Message string `json:"message"`
} //@name ResponseError

type ResponseOk struct {
	Message string `json:"message"`
} //@name ResponseOk

// NewRequestHandler creates a new request handler
func NewRequestHandler(logger Logger, env Env) RequestHandler {
	gin.DefaultWriter = logger.GetGinLogger()
	engine := gin.New()
	docs.SwaggerInfo.Title = "Compras API"
	docs.SwaggerInfo.Description = "this is a api for registry yours purchases"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	docs.SwaggerInfo.BasePath = "/api"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.Version = "1.0"

	engine.GET(
		"/swagger/*any",
		ginSwagger.WrapHandler(
			swaggerFiles.Handler,
		),
	)

	return RequestHandler{Gin: engine}
}
