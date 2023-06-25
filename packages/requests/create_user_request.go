package requests

import (
	"crypto/sha256"
	"fmt"

	"github.com/gabrieldebem/todo-api/packages/models"
)

type CreateUserRequest struct {
    Name     string `json:"name" binding:"required"`
    Email    string `json:"email" binding:"required"`
    Password string `json:"password" binding:"required"`
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
