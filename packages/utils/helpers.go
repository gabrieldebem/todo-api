package utils

import (
	"github.com/gabrieldebem/todo-api/packages/models"
	"github.com/gin-gonic/gin"
)

func AuthUser(c *gin.Context) models.User {
	var user models.User
	user = c.Keys["user"].(models.User)

	return user
}
