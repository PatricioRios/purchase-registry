package auth

import (
	"net/http"
	"strings"

	"github.com/PatricioRios/Compras/models"
	"github.com/PatricioRios/Compras/services/v1/auth"
	"github.com/gin-gonic/gin"
)

type AuthController struct {
	authService auth.AuthService
}

func NewAuthController(authService auth.AuthService) AuthController {
	return AuthController{authService: authService}
}

func (ctrl *AuthController) LogInWithEmail(c *gin.Context) {

	var user models.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	token, refreshToken, err := ctrl.authService.LogInWithEmail(user)
	if err != nil {
		c.JSON(err.Code(), gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token, "refreshToken": refreshToken})

}

func (ctrl *AuthController) LogInWithUserName(c *gin.Context) {
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	token, refreshToken, err := ctrl.authService.LogInWithUserName(user)
	if err != nil {
		c.JSON(err.Code(), gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token, "refreshToken": refreshToken})

}

func (crtl *AuthController) RefreshToken(c *gin.Context) {

	bearerToken := c.Request.Header.Get("Authorization")

	splitToken := strings.Split(bearerToken, " ")

	if len(splitToken) != 2 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		c.Abort()
		return
	}
	token := splitToken[1]

	err := crtl.authService.ValidateRefreshToken(token)

	if err != nil {
		c.JSON(err.Code(), gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	newToken, err := crtl.authService.RefreshToken(token)
	if err != nil {
		c.JSON(err.Code(), gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": newToken})
}

func (cttl *AuthController) ProtectedHandler() gin.HandlerFunc {
	return func(c *gin.Context) {

		bearerToken := c.Request.Header.Get("Authorization")

		splitToken := strings.Split(bearerToken, " ")

		if len(splitToken) != 2 {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}
		token := splitToken[1]

		err := cttl.authService.ValidateToken(token)

		if err != nil {
			c.JSON(err.Code(), gin.H{"error": err.Error()})
			c.Abort()
			return
		}
	}
}
