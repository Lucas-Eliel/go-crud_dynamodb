package controllers

import (
	"crud_dynamodb/internal/models"
	"crud_dynamodb/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterHandlers(r *gin.Engine, service services.Service) {
	controller := UserController{service}

	r.GET("/users", controller.GetUsers)
	r.GET("/users/:id", controller.GetUser)
	r.POST("/users", controller.PostUser)
	r.DELETE("/users/:id", controller.DeleteUser)
}

type UserController struct {
	service services.Service
}

func (u UserController) GetUsers(c *gin.Context) {
	users := u.service.GetUsuarios()
	c.IndentedJSON(http.StatusOK, users)
}

func (u UserController) GetUser(c *gin.Context) {
	users := u.service.GetUsuario(c.Param("id"))
	c.IndentedJSON(http.StatusOK, users)
}

func (u UserController) PostUser(c *gin.Context) {
	var user models.User

	err := c.BindJSON(&user)

	if err != nil {
		panic(err)
	}

	u.service.PostUsuario(user)
	c.IndentedJSON(http.StatusCreated, user)
}

func (u UserController) DeleteUser(c *gin.Context) {

	id := c.Param("id")

	u.service.DeleteUsuario(id)

	c.IndentedJSON(http.StatusOK, id)
}
