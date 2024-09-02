package article

import (
	"github.com/PatricioRios/Compras/models"
	"gorm.io/gorm"
)

type ArticleRepositoryImpl struct {
	db *gorm.DB
}

func NewArticleRepositoryImpl(db *gorm.DB) ArticleRepository {
	return &ArticleRepositoryImpl{
		db: db,
	}
}

func (r *ArticleRepositoryImpl) CreateArticle(article models.Article) (models.Article, error) {
	tx := r.db.Create(&article)

	err := tx.Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return article, ErrArticleNotFound
		}
		return article, ErrInternalError
	}
	return article, nil
}
func (r *ArticleRepositoryImpl) UpdateArticle(article models.Article) (models.Article, error) {
	tx := r.db.Save(&article)

	err := tx.Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return article, ErrArticleNotFound
		}
		return article, ErrInternalError
	}
	return article, nil
}
func (r *ArticleRepositoryImpl) DeleteArticle(articleId int, userId int) error {
	tx := r.db.Where("id = ? AND user_id = ?", articleId, userId).Delete(&models.Article{})
	err := tx.Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return ErrArticleNotFound
		}
		return ErrInternalError
	}
	return nil
}

func (r *ArticleRepositoryImpl) GetArticleById(articleId int) (models.Article, error) {
	var article models.Article
	tx := r.db.Where("id = ?", articleId).First(&article)
	err := tx.Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return article, ErrArticleNotFound
		}
		return article, ErrInternalError
	}
	return article, nil
}

func (r *ArticleRepositoryImpl) GetArticlesByPurchaseId(purchaseId int) ([]models.Article, error) {
	var articles []models.Article
	tx := r.db.Where("compra_id = ?").Find(&articles)
	err := tx.Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return articles, ErrArticlesNotFound
		}
		return articles, ErrInternalError
	}
	return articles, nil
}
