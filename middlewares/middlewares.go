package middlewares

import (
	"github.com/PatricioRios/Compras/middlewares/auth"
	"go.uber.org/fx"
)

// Module Middleware exported
var Module = fx.Options(
	fx.Provide(NewMiddlewares),
	fx.Provide(auth.NewAuthMidleware),
)

// IMiddleware middleware interface
type IMiddleware interface {
	Setup()
}

// Middlewares contains multiple middleware
type Middlewares []IMiddleware

// NewMiddlewares creates new middlewares
// Register the middleware that should be applied directly (globally)
func NewMiddlewares(
	AuthMiddleware auth.AuthMiddleware,
) Middlewares {
	return Middlewares{
		AuthMiddleware,
	}
}

// Setup sets up middlewares
func (m Middlewares) Setup() {
	for _, middleware := range m {
		middleware.Setup()
	}
}
