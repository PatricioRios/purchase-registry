package routes

import (
	"github.com/PatricioRios/Compras/routes/auth"
	"github.com/PatricioRios/Compras/routes/user"
	"go.uber.org/fx"
)

// Module exports dependency to container
var Module = fx.Options(
	fx.Provide(NewRoutes),
	fx.Provide(NewMiscRoutes),
	fx.Provide(NewCompraRoutes),
	fx.Provide(NewCategoryRoutes),
	fx.Provide(user.NewUserRoutes),
	fx.Provide(auth.NewAuthRoutes),
)

// Routes contains multiple routes
type Routes []Route

// Route interface
type Route interface {
	Setup()
}

// NewRoutes sets up routes
func NewRoutes(
	miscRoutes MiscRoutes,
	compraRoutes PurchaseRoutes,
	categoryRoutes CategoryRoutes,
	userRoutes user.UserRoutes,
	authRoutes auth.AuthRoutes,
) Routes {
	return Routes{
		miscRoutes,
		compraRoutes,
		categoryRoutes,
		userRoutes,
		authRoutes,
	}
}

// Setup all the route
func (r Routes) Setup() {
	for _, route := range r {
		route.Setup()
	}
}
