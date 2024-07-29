package main

import (
	"github.com/PatricioRios/Compras/bootstrap"
	"github.com/PatricioRios/Compras/utils"

	"github.com/joho/godotenv"
	"go.uber.org/fx"
)

func main() {
	godotenv.Load()
	logger := utils.GetLogger().GetFxLogger()
	fx.New(bootstrap.Module, fx.Logger(logger)).Run()
}
