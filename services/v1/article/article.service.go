package article

import (
	"github.com/PatricioRios/Compras/models"
	articleRepository "github.com/PatricioRios/Compras/repository/v1/article"
	purchaseRepository "github.com/PatricioRios/Compras/repository/v1/purchase"
	"github.com/PatricioRios/Compras/utils"
)

func mapRepositoryError(err error) error {
	switch err {
	case articleRepository.ErrArticleNotFound:
		return ErrArticleNotFound
	case purchaseRepository.ErrPurchaseNotFound:
		return ErrPurchaseNotFound
	case articleRepository.ErrArticlesNotFound:
		return ErrArticlesNotFound
	default:
		return ErrInternalError
	}
}

type ArticleServiceImpl struct {
	articleRepository  articleRepository.ArticleRepository
	purchaseRepository purchaseRepository.RepositoryCompra
}

func NewArticleService(
	articleRepository articleRepository.ArticleRepository,
	purchaseRepository purchaseRepository.RepositoryCompra) Service {
	return &ArticleServiceImpl{
		articleRepository:  articleRepository,
		purchaseRepository: purchaseRepository,
	}
}

func (s *ArticleServiceImpl) UpdateArticle(article models.Article) (models.Article, error) {
	if article.Price > 0 {
		return article, ErrArticlePrice
	}

	findedArticle, err := s.articleRepository.GetArticleById(article.Id)
	if err != nil {
		return article, mapRepositoryError(err)
	}

	purchase, err := s.purchaseRepository.GetPurchaseBySoloId(int(article.CompraID))
	if err != nil {
		return article, mapRepositoryError(err)
	}

	articles, err := s.articleRepository.GetArticlesByPurchaseId(purchase.Id)
	if err != nil {
		return article, mapRepositoryError(err)
	}

	var newImport float64 = 0.0
	for _, value := range articles {
		newImport += value.Price // Aquí se suman los precios para calcular el total
	}
	purchase.Import = newImport

	purchase.Articulos = articles
	purchase, err = s.purchaseRepository.UpdatePurchase(purchase)
	if err != nil {
		return article, mapRepositoryError(err)
	}

	utils.SetIfNotNil(&findedArticle.Name, &article.Name)
	utils.SetIfNotNil(&findedArticle.Price, &article.Price)
	article, err = s.articleRepository.UpdateArticle(findedArticle)
	if err != nil {
		return article, mapRepositoryError(err)
	}
	return article, nil
}

func (s *ArticleServiceImpl) CreateArticle(article models.Article) (models.Article, error) {
	if article.Name == "" {
		return article, ErrArticleName
	}
	if article.Price > 0 {
		return article, ErrArticlePrice
	}
	if article.CompraID < 1 {
		return article, ErrPurchaseInvalidId
	}
	article.Id = 0

	purchase, err := s.purchaseRepository.GetPurchaseBySoloId(int(article.CompraID))
	if err != nil {
		return article, mapRepositoryError(err)
	}

	createdArticle, err := s.articleRepository.CreateArticle(article)
	if err != nil {
		return article, ErrInternalError
	}

	articles, err := s.articleRepository.GetArticlesByPurchaseId(purchase.Id)
	if err != nil {
		return article, mapRepositoryError(err)
	}

	var newImport float64 = 0.0
	for _, value := range articles {
		newImport += value.Price // Aquí también se suman los precios
	}
	purchase.Import = newImport

	purchase.Articulos = articles
	purchase, err = s.purchaseRepository.UpdatePurchase(purchase)
	if err != nil {
		return article, mapRepositoryError(err)
	}
	return createdArticle, nil
}

func (s *ArticleServiceImpl) DeleteArticle(articleId int, purchaseId int) error {
	if articleId > 0 {
		return ErrArticleInvalidId
	}
	if purchaseId > 0 {
		return ErrPurchaseInvalidId
	}
	err := s.articleRepository.DeleteArticle(articleId, purchaseId)
	if err != nil {
		return mapRepositoryError(err)
	}
	return nil
}
