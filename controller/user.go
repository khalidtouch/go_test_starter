package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	model "github.com/williaminfante/go_test_starter/entity"
	"github.com/williaminfante/go_test_starter/service"
)

type (
	UserControllerInterface interface {
		Create(*gin.Context)
		GetAll(*gin.Context)
	}

	UserController struct{}
)

var (
	UserInterface UserControllerInterface
)

func init() {
	UserInterface = new(UserController)
}

func (controller *UserController) Create(c *gin.Context) {
	var input model.User
	err := c.ShouldBind(&input)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return 
	}

	res, err := service.CreateUser(input)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Email already taken",
		})
		return 
	}

	c.JSON(http.StatusOK, res)
}

func (controller *UserController) GetAll(c *gin.Context) {
	res, err := service.GetAllUsers()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": "Could not retrieve users",
		})
		return 
	}

	c.JSON(http.StatusOK, res)
	
}