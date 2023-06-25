package actions

import (
	"github.com/gabrieldebem/todo-api/packages/repositories"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func FindUser(c *gin.Context) {
    id, err := uuid.Parse(c.Param("id"))

    if err != nil {
        c.JSON(400, gin.H{
            "error": "Invalid UUID",
        })
        return
    }

    user, err := repositories.UserRepository{}.First(id)

    if err != nil {
        c.JSON(404, gin.H{
            "error": "User not found",
        })
        return
    }

    c.JSON(200, gin.H{
        "user": user,
    })
}
