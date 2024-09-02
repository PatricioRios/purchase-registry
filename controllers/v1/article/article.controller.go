package article

import (
	"net/http"
	"strconv"

	"github.com/PatricioRios/Compras/models"
	"github.com/PatricioRios/Compras/services/v1/article"
	"github.com/PatricioRios/Compras/utils"
	"github.com/gin-gonic/gin"
)

type ArticleController struct {
	articleService article.Service
}

func NewArticleController(articleService article.Service) *ArticleController {
	return &ArticleController{
		articleService: articleService,
	}
}

func mapStauts(err error) int {
	switch err {
	case article.ErrArticleName:
		return http.StatusBadRequest
	case article.ErrArticlePrice:
		return http.StatusBadRequest
	case article.ErrArticleNotFound:
		return http.StatusNoContent
	case article.ErrArticleInvalidId:
		return http.StatusBadRequest
	case article.ErrArticlesNotFound:
		return http.StatusBadRequest
	case article.ErrPurchaseInvalidId:
		return http.StatusBadRequest
	case article.ErrPurchaseNotFound:
		return http.StatusNoContent
	default:
		return http.StatusInternalServerError
	}
}

/*

func (ctrl *CategoryController) GetCategoryById(c *gin.Context) {
	param := c.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ResponseError{Message: "ID inválido"})
		return
	}
	userID, _ := c.Get("user_id")

	categoria, err := ctrl.service.GetCategoryById(id, userID.(int))
	if err != nil {
		if err == category.ErrInternalError {
			c.JSON(http.StatusInternalServerError, utils.ResponseError{Message: err.Error()})
			return
		}
		c.JSON(http.StatusBadRequest, utils.ResponseError{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, categoria)
}

*/

func (ctrl *ArticleController) UpdateArticle(c *gin.Context) {

	var article models.Article
	err := c.BindJSON(&article)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ResponseError{Message: err.Error()})
		return
	}
	article, err = ctrl.articleService.UpdateArticle(article)
	if err != nil {
		c.JSON(mapStauts(err), utils.ResponseError{Message: err.Error()})
		return
	}
}
func (ctrl *ArticleController) CreateArticle(c *gin.Context) {

	var article models.Article
	err := c.BindJSON(&article)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ResponseError{Message: err.Error()})
		return
	}
	article, err = ctrl.articleService.CreateArticle(article)
	if err != nil {
		c.JSON(mapStauts(err), utils.ResponseError{Message: err.Error()})
		return
	}

}
func (ctrl *ArticleController) DeleteArticle(c *gin.Context) {
	param := c.Param("purchase_id")
	purchaseId, err := strconv.Atoi(param)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ResponseError{Message: "ID inválido"})
		return
	}
	param = c.Param("article_id")
	articleId, err := strconv.Atoi(param)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ResponseError{Message: "ID inválido"})
		return
	}
	err = ctrl.articleService.DeleteArticle(articleId, purchaseId)

	if err != nil {
		c.JSON(mapStauts(err), utils.ResponseError{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Articulo eliminado"})
}
