package article

import (
	"errors"

	"github.com/PatricioRios/Compras/models"
)

type ArticleRepository interface {
	UpdateArticle(article models.Article) (models.Article, error)
	CreateArticle(article models.Article) (models.Article, error)
	DeleteArticle(articleId int, purchaseId int) error
	GetArticleById(ArticleId int) (models.Article, error)
	GetArticlesByPurchaseId(purchaseId int) ([]models.Article, error)
}

var (
	ErrArticleNotFound  error = errors.New("ArticleNotFound")
	ErrInternalError    error = errors.New("InternalError")
	ErrArticlesNotFound error = errors.New("ArticlesForPourchaseIdNotFound")
)
