package requests

import (
	"crypto/sha256"
	"fmt"

	"github.com/gabrieldebem/todo-api/packages/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type CreateUserRequest struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

func (r CreateUserRequest) ToUser() models.User {
	hash := sha256.New().Sum([]byte(r.Password))
	password := fmt.Sprintf("%x", hash)
	return models.User{
		Name:     r.Name,
		Email:    r.Email,
		Password: string(password),
	}
}

func (r CreateUserRequest) Validate(c *gin.Context) error {
    err := validator.New().Struct(r)

    if err != nil {
        c.JSON(400, gin.H{
            "message": "Invalid request",
            "errors":  err.Error(),
        })
    }

    return err
}
