package category

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/PatricioRios/Compras/models"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func setupMockDB() (*gorm.DB, sqlmock.Sqlmock, error) {
	db, mock, err := sqlmock.New()
	if err != nil {
		return nil, nil, err
	}

	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn:                      db,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{})
	if err != nil {
		return nil, nil, err
	}

	return gormDB, mock, nil
}

func TestGetAllCategories(t *testing.T) {
	db, mock, err := setupMockDB()
	assert.NoError(t, err)

	repo := NewRepositoryCategoryImpl(db)

	userId := 1
	mock.ExpectQuery("SELECT \\* FROM `category_purchases` WHERE user_id = \\?").
		WithArgs(userId).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "user_id"}).
			AddRow(1, "Electronics", userId).
			AddRow(2, "Clothing", userId))

	categories, err := repo.GetAllCategories(userId)
	assert.NoError(t, err)
	assert.Len(t, categories, 2)
}

// hix.ai
func TestGetCategoryById(t *testing.T) {
	db, mock, err := setupMockDB()
	assert.NoError(t, err)

	repo := NewRepositoryCategoryImpl(db)

	id := 1
	userId := 1
	mock.ExpectQuery("SELECT \\* FROM `category_purchases` WHERE id = \\? AND user_id = \\? ORDER BY `category_purchases`.`id` LIMIT 1").
		WithArgs(id, userId).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "user_id"}).
			AddRow(id, "Electronics", userId))

	// Agregar mock para la consulta de purchases
	mock.ExpectQuery("SELECT \\* FROM `purchases` WHERE `purchases`.`categoria_id` = \\?").
		WithArgs(id).
		WillReturnRows(sqlmock.NewRows([]string{"id", "categoria_id", "amount"}).
			AddRow(1, id, 100).
			AddRow(2, id, 200))

	category, err := repo.GetCategoryById(id, userId)
	assert.NoError(t, err)
	assert.Equal(t, "Electronics", category.Name)
	assert.Len(t, category.Purchases, 2)
}

func TestCreateCategory(t *testing.T) {
	db, mock, err := setupMockDB()
	assert.NoError(t, err)

	repo := NewRepositoryCategoryImpl(db)

	category := models.CategoryPurchase{
		Name:   "Books",
		UserID: 1,
	}
	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO `category_purchases` \\(`name`,`user_id`,`created_at`,`updated_at`\\) VALUES \\(\\?,\\?,\\?,\\?\\)").
		WithArgs(category.Name, category.UserID, sqlmock.AnyArg(), sqlmock.AnyArg()).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	createdCategory, err := repo.CreateCategory(category)
	assert.NoError(t, err)
	assert.Equal(t, category.Name, createdCategory.Name)
}

func TestUpdateCategory(t *testing.T) {
	db, mock, err := setupMockDB()
	assert.NoError(t, err)

	repo := NewRepositoryCategoryImpl(db)

	category := models.CategoryPurchase{
		Id:     1,
		Name:   "Books Updated",
		UserID: 1,
	}
	mock.ExpectBegin()
	mock.ExpectExec("UPDATE `category_purchases` SET `name`=\\?,`user_id`=\\?,`created_at`=\\?,`updated_at`=\\? WHERE `id` = \\?").
		WithArgs(category.Name, category.UserID, sqlmock.AnyArg(), sqlmock.AnyArg(), category.Id).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	updatedCategory, err := repo.UpdateCategory(category)
	assert.NoError(t, err)
	assert.Equal(t, category.Name, updatedCategory.Name)
}

func TestDeleteCategory(t *testing.T) {
	db, mock, err := setupMockDB()
	assert.NoError(t, err)

	repo := NewRepositoryCategoryImpl(db)

	id := 1
	userId := 1
	mock.ExpectBegin()
	mock.ExpectExec("DELETE FROM `category_purchases` WHERE id = \\? AND user_id = \\?").
		WithArgs(id, userId).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	err = repo.DeleteCategory(id, userId)
	assert.NoError(t, err)
}

func TestGetCategoryByName(t *testing.T) {
	db, mock, err := setupMockDB()
	assert.NoError(t, err)

	repo := NewRepositoryCategoryImpl(db)

	name := "Electronics"
	userId := 1
	mock.ExpectQuery("SELECT \\* FROM `category_purchases` WHERE name = \\? AND user_id = \\? ORDER BY `category_purchases`.`id` LIMIT \\?").
		WithArgs(name, userId, 1).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "user_id"}).
			AddRow(1, name, userId))

	category, err := repo.GetCategoryByName(name, userId)
	assert.NoError(t, err)
	assert.Equal(t, name, category.Name)
}
