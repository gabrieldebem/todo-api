package routes

import (
	"github.com/gabrieldebem/todo-api/packages/controllers"
	"github.com/gabrieldebem/todo-api/packages/repositories"
	"github.com/gin-gonic/gin"
)

var userController = controllers.UserController{Repository: repositories.UserRepository{}}
var taskController = controllers.TaskController{Repository: repositories.TaskRepository{}}

func Init() *gin.Engine {
	r := gin.Default()

	api := r.Group("/api")
	{
		api.GET("/", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "Hello, World!",
			})
		})
		api.GET("/users", userController.ListUsers)
		api.GET("/users/:id", userController.FindUser)
		api.POST("/users", userController.CreateUser)

        api.GET("/tasks", taskController.ListTasks)
        api.GET("/tasks/:id", taskController.FindTask)
        api.POST("/tasks", taskController.CreateTask)
	}

	return r
}
