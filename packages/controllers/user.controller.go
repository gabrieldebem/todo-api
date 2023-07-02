package controllers

import (
	"github.com/gabrieldebem/todo-api/packages/exceptions"
	"github.com/gabrieldebem/todo-api/packages/models"
	"github.com/gabrieldebem/todo-api/packages/repositories"
	"github.com/gabrieldebem/todo-api/packages/requests"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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
	id, err := uuid.Parse(c.Param("id"))

	if err != nil {
		exceptions.HttpException{
			Context: c,
			Error:   err,
			Message: "Invalid ID",
			Code:    400,
		}.Throw()
		return
	}

	user := models.User{ID: id}

	user, err = u.Repository.First(&user)

	if err != nil {
		exceptions.HttpException{
			Context: c,
			Error:   err,
			Message: "User not found",
			Code:    404,
		}.Throw()
		return
	}

	c.JSON(200, gin.H{
		"user": user,
	})
}

func (u UserController) ListUsers(c *gin.Context) {
	users, err := u.Repository.FindAll()

	if err != nil {
		exceptions.HttpException{
			Context: c,
			Error:   err,
		}.Throw()
		return
	}

	c.JSON(200, gin.H{
		"data": users,
	})
}
