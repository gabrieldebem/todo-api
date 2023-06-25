package routes

import (
	"github.com/gabrieldebem/todo-api/packages/actions"
	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	r := gin.Default()

	api := r.Group("/api")
	{
		api.GET("/", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "Hello, World!",
			})
		})
        api.GET("/users", actions.ListUsers)
        api.GET("/users/:id", actions.FindUser)
        api.POST("/users", actions.CreateUser)
	}

    return r
}
