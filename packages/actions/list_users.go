package actions

import (
	"github.com/gabrieldebem/todo-api/packages/repositories"
	"github.com/gin-gonic/gin"
)

func ListUsers(c *gin.Context) {
    users := repositories.UserRepository{}.FindAll()

    c.JSON(200, gin.H{
        "users": users,
    })
}
