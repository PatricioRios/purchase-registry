package mysqlgorm

import (
	"fmt"

	"github.com/PatricioRios/Compras/models"
	"github.com/PatricioRios/Compras/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewMySQLORM(
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
