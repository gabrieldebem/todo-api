package controllers

import (
	"github.com/gabrieldebem/todo-api/packages/repositories"
	usecases "github.com/gabrieldebem/todo-api/packages/use_cases"
	"github.com/gin-gonic/gin"
)

type AuthRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AuthController struct {
	UserRepository        *repositories.UserRepository
}

func (ct *AuthController) Login(c *gin.Context) {
	var request AuthRequest
	c.BindJSON(&request)

	users, _ := ct.UserRepository.FindWhere("email = ?", "", request.Username)

	if len(users) == 0 {
		c.JSON(401, gin.H{
			"error":   "Unauthorized",
			"message": "Invalid credentals",
		})
		return
	}

	user := users[0]
	if !user.CheckPassword(request.Password) {
		c.JSON(401, gin.H{
			"error": "Unauthorized",
			"data":  "Invalid password",
		})
		return
	}

	accessToken, authErr := usecases.GenerateToken(user.ID)

	if authErr != nil {
		c.JSON(401, gin.H{
			"error": "Unauthorized",
			"data":  authErr.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"token":      accessToken.Token,
		"expires_at": accessToken.ExpiresAt,
	})
}
