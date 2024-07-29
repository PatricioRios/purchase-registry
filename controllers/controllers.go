package controllers

import (
	"github.com/PatricioRios/Compras/controllers/v1/auth"
	"github.com/PatricioRios/Compras/controllers/v1/category"
	misc_v1 "github.com/PatricioRios/Compras/controllers/v1/misc"
	"github.com/PatricioRios/Compras/controllers/v1/purchase"
	"github.com/PatricioRios/Compras/controllers/v1/user"
	"go.uber.org/fx"
)

// Module exported for initializing application
var Module = fx.Options(
	fx.Provide(misc_v1.NewMiscController),
	fx.Provide(purchase.NewPurchaseController),
	fx.Provide(category.NewCategoryController),
	fx.Provide(user.NewUserController),
	fx.Provide(auth.NewAuthController),
)
