package article

import (
	"errors"

	"github.com/PatricioRios/Compras/models"
)

type Service interface {
	UpdateArticle(models.Article) (models.Article, error)
	CreateArticle(models.Article) (models.Article, error)
	DeleteArticle(articleId int, purchaseId int) error
}

var (
	ErrArticleName      error = errors.New("NameOfArticleIsEmpty")
	ErrArticlePrice     error = errors.New("PriceOfArticleIsNegative")
	ErrArticleNotFound  error = errors.New("ArticleNotFound")
	ErrArticleInvalidId error = errors.New("ArticleInvalidId")
	ErrArticlesNotFound error = errors.New("ArticlesForPourchaseIdNotFound")

	ErrPurchaseInvalidId error = errors.New("PurchaseInvalidId")
	ErrPurchaseNotFound  error = errors.New("PurcharseForArticleNotFound")

	ErrInternalError error = errors.New("InternalError")
)
