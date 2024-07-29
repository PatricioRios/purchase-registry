package bootstrap

import (
	"context"

	database "github.com/PatricioRios/Compras/DTO"
	"github.com/PatricioRios/Compras/controllers"
	"github.com/PatricioRios/Compras/middlewares"
	"github.com/PatricioRios/Compras/repository"
	"github.com/PatricioRios/Compras/routes"
	"github.com/PatricioRios/Compras/services"
	"github.com/PatricioRios/Compras/utils"
	"go.uber.org/fx"
)

// Module exported for initializing application
var Module = fx.Options(
	middlewares.Module,
	services.Module,
	controllers.Module,
	routes.Module,
	utils.Module,
	database.Module,
	repository.Module,
	fx.Invoke(bootstrap),
)

func bootstrap(
	lifecycle fx.Lifecycle,
	handler utils.RequestHandler,
	routes routes.Routes,
	env utils.Env,
	logger utils.Logger,
	middlewares middlewares.Middlewares,
) {

	lifecycle.Append(fx.Hook{
		OnStart: func(context.Context) error {

			go func() {
				middlewares.Setup()
				routes.Setup()
				host := "127.0.0.1"
				if env.Environment == "development" {
					host = "127.0.0.1"
				}
				handler.Gin.Run(host + ":" + env.ServerPort)
			}()
			return nil
		},
		OnStop: func(context.Context) error {
			logger.Info("Stopping Application")
			return nil
		},
	})
}
