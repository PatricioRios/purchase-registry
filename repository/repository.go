package repository

import (
	"github.com/PatricioRios/Compras/repository/v1/category"
	"github.com/PatricioRios/Compras/repository/v1/purchase"
	"github.com/PatricioRios/Compras/repository/v1/user"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(purchase.NewRepositoryCompraImpl),
	fx.Provide(category.NewRepositoryCategoryImpl),
	fx.Provide(user.NewUserRepositoryImpl),
)
