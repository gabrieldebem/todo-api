package controllers

import (
	"github.com/gabrieldebem/todo-api/packages/exceptions"
	"github.com/gabrieldebem/todo-api/packages/repositories"
	"github.com/gabrieldebem/todo-api/packages/requests"
	"github.com/gabrieldebem/todo-api/packages/utils"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	Repository repositories.UserRepository
}

func (u UserController) CreateUser(c *gin.Context) {
	var r requests.CreateUserRequest

	c.Bind(&r)
	var err error

	err = r.Validate(c)

	if err != nil {
		return
	}

	user := r.ToUser()

	user, err = u.Repository.Create(&user)

	if err != nil {
		exceptions.HttpException{
			Context: c,
			Error:   err,
		}.Throw()
		return
	}

	c.JSON(200, gin.H{
		"user": user,
	})
}

func (u UserController) FindUser(c *gin.Context) {
	user := utils.AuthUser(c)

	c.JSON(200, gin.H{
		"user": user,
	})
}
