package main_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	authController "github.com/PatricioRios/Compras/controllers/v1/auth"
	misc_v1 "github.com/PatricioRios/Compras/controllers/v1/misc"
	"github.com/PatricioRios/Compras/models"
	authRepository "github.com/PatricioRios/Compras/repository/v1/auth"
	userRepository "github.com/PatricioRios/Compras/repository/v1/user"
	"github.com/PatricioRios/Compras/routes/auth"
	authService "github.com/PatricioRios/Compras/services/v1/auth"
	"github.com/PatricioRios/Compras/utils"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/alecthomas/assert"
	"github.com/gin-gonic/gin"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func NewMySQLORMTest(
	logger utils.Logger,
	env utils.Env,
) *gorm.DB {
	var err error

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		env.DB.User,
		env.DB.Pass,
		env.DB.Host,
		env.DB.Port,
		env.DB.Name,
	)

	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(fmt.Sprintf("Failed to connect to database because: %s!", err.Error()))
	}

	logger.Info("Database MySQL GORM connected")

	// Migrate the schema
	logger.Info("Migrating the schema...")
	err = DB.AutoMigrate(models.Purchase{}, models.Article{}, models.CategoryPurchase{}, models.User{})
	if err != nil {
		logger.Info("Schema NO migrated!")
	}
	logger.Info("Schema migrated!")
	return DB
}

func NewEnv(log utils.Logger) utils.Env {
	env := utils.Env{}
	viper.SetConfigFile(".env")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("☠️ cannot read configuration")
	}

	err = viper.Unmarshal(&env)
	if err != nil {
		log.Fatal("☠️ environment can't be loaded: ", err)
	}

	log.Infof("%+v \n", env)
	return env
}

func TestMiscVersion(t *testing.T) {
	r := SetUpRouter()
	miscController := misc_v1.NewMiscController(utils.GetLogger())
	r.GET("/version", miscController.GetVersion)
	req, _ := http.NewRequest("GET", "/version", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestLoginUserWithEmailController(t *testing.T) {
	logger := utils.GetLogger()
	env := NewEnv(logger)
	db := NewMySQLORMTest(logger, env)
	handler := utils.NewRequestHandler(logger, env)

	repositoryAuth := authRepository.NewAuthRepositoryImpl(env)
	userRepository, _ := userRepository.New(db)
	service := authService.NewAuthService(userRepository, repositoryAuth, env)
	controller := authController.NewAuthController(service)

	routes := auth.NewAuthRoutes(logger, handler, controller)
	routes.Setup()

	tests := []struct {
		name               string
		inputUser          models.User
		expectedStatusCode int
		expectedResponse   string
	}{
		{
			name: "Valid credentials",
			inputUser: models.User{
				Email:    "patricio@gmail.com",
				Password: "123",
			},
			expectedStatusCode: http.StatusOK,
			expectedResponse:   "Token",
		},
		{
			name: "Invalid email",
			inputUser: models.User{
				Email:    "invalid@gmail.com",
				Password: "123",
			},
			expectedStatusCode: http.StatusUnauthorized,
			expectedResponse:   authService.ErrUserNotFound.Error(),
		},
		{
			name: "Invalid password",
			inputUser: models.User{
				Email:    "patricio@gmail.com",
				Password: "wrongpassword",
			},
			expectedStatusCode: http.StatusUnauthorized,
			expectedResponse:   authService.ErrInvalidPassword.Error(),
		},
		{
			name: "Invalid jsonInput",
			inputUser: models.User{
				Password: "wrongpassword",
			},
			expectedStatusCode: http.StatusBadRequest,
			expectedResponse:   "error",
		},
		{
			name: "Ivalid email format",
			inputUser: models.User{
				Email:    "patricio@gmail",
				Password: "wrongpassword",
			},
			expectedStatusCode: http.StatusBadRequest,
			expectedResponse:   authService.ErrInvalidEmailFormat.Error(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reqBody, _ := json.Marshal(tt.inputUser)

			req := httptest.NewRequest(http.MethodGet, "/api/v1/auth/email", bytes.NewBuffer(reqBody))
			req.Header.Set("Content-Type", "application/json")

			w := httptest.NewRecorder()
			handler.Gin.ServeHTTP(w, req)

			stringss := fmt.Sprintf("Código de estado incorrecto expected %f, got %f", float32(tt.expectedStatusCode), float32(w.Code))
			responseBody := w.Body.String()
			assert.Equal(t, tt.expectedStatusCode, w.Code, stringss, responseBody)
			assert.Contains(t, w.Body.String(), tt.expectedResponse)
		})
	}
}
