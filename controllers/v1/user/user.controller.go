package user

import (
	"net/http"
	"strconv"

	"github.com/PatricioRios/Compras/models"
	"github.com/PatricioRios/Compras/services/v1/user"
	"github.com/PatricioRios/Compras/utils"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService user.UserService
}

func NewUserController(userService user.UserService) UserController {
	return UserController{userService: userService}
}

func (ctrl *UserController) GetAllUsers(c *gin.Context) {
	users, err := ctrl.userService.GetAllUsers()
	if err != nil {
		c.JSON(err.Code(), utils.ResponseError{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}

func (ctrl *UserController) GetUserById(c *gin.Context) {
	param := c.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ResponseError{Message: err.Error()})
		return
	}
	user, errSrvc := ctrl.userService.GetUserById(id)
	if errSrvc != nil {
		c.JSON(errSrvc.Code(), utils.ResponseError{Message: errSrvc.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (ctrl *UserController) CreateUser(c *gin.Context) {
	user := models.User{}
	err := c.Bind(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ResponseError{Message: err.Error()})
		return
	}
	user, errSrvc := ctrl.userService.CreateUser(user)
	if errSrvc != nil {
		c.JSON(errSrvc.Code(), utils.ResponseError{Message: errSrvc.Error()})
		return
	}
	c.JSON(http.StatusCreated, user)
}
