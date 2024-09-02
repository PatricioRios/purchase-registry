package category_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	database "github.com/PatricioRios/Compras/DTO"
	"github.com/PatricioRios/Compras/bootstrap"
	"github.com/PatricioRios/Compras/controllers"
	"github.com/PatricioRios/Compras/middlewares"
	"github.com/PatricioRios/Compras/models"
	"github.com/PatricioRios/Compras/repository"
	"github.com/PatricioRios/Compras/routes"
	"github.com/PatricioRios/Compras/services"
	"github.com/PatricioRios/Compras/utils"
	"github.com/alecthomas/assert"

	"go.uber.org/fx"
)

var Module = fx.Options(
	middlewares.Module,
	services.Module,
	controllers.Module,
	routes.Module,
	utils.Module,
	database.TestModule,
	repository.Module,
	fx.Invoke(bootstrapTest),
)

func bootstrapTest(
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

func TestCreateArticle(t *testing.T) {
	// Cargar el archivo .env.test

	handler := utils.NewRequestHandler(utils.GetLogger(), utils.NewEnv(utils.GetLogger()))

	logger := utils.GetLogger().GetFxLogger()
	app := fx.New(fx.Options(fx.Provide(func() utils.RequestHandler {
		return handler
	})), bootstrap.Module, fx.Logger(logger))

	// Iniciar la aplicación en un contexto separado
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	startCtx, cancelStart := context.WithCancel(ctx)
	defer cancelStart()

	go func() {
		app.Run()
		cancelStart()
	}()

	<-startCtx.Done()

	// Realizar la petición HTTP de prueba
	w := httptest.NewRecorder()
	exampleCategory := models.CategoryPurchase{
		Name: "categoria",
	}
	categoryJSON, _ := json.Marshal(exampleCategory)
	req, _ := http.NewRequest("POST", "/api/v1/category", strings.NewReader(string(categoryJSON)))
	req.Header.Set("Content-Type", "application/json")
	handler.Gin.ServeHTTP(w, req)

	// Verificar la respuesta
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "categoria")
}
