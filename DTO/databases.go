package DTO

import (
	mysqlgorm "github.com/PatricioRios/Compras/DTO/mysql-gorm"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(mysqlgorm.NewMySQLORM),
)

var TestModule = fx.Options(
	fx.Provide(mysqlgorm.NewMySQLORMTest),
)
