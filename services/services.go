package services

import (
	"github.com/PatricioRios/Compras/services/v1/auth"
	"github.com/PatricioRios/Compras/services/v1/category"
	"github.com/PatricioRios/Compras/services/v1/purchase"
	"github.com/PatricioRios/Compras/services/v1/user"
	"go.uber.org/fx"
)

// Module exports services present
var Module = fx.Options(
	// fx.Provide()
	fx.Provide(purchase.NewCompraService),
	fx.Provide(category.NewCategoryService),
	fx.Provide(user.NewUserService),
	fx.Provide(auth.NewAuthService),
)
