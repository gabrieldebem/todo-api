package routes

import (
	"github.com/gabrieldebem/todo-api/packages/controllers"
	"github.com/gabrieldebem/todo-api/packages/middlewares"
	"github.com/gabrieldebem/todo-api/packages/repositories"
	"github.com/gin-gonic/gin"
)

var userController = controllers.UserController{Repository: repositories.UserRepository{}}
var taskController = controllers.TaskController{Repository: repositories.TaskRepository{}}
var authController = controllers.AuthController{UserRepository: &repositories.UserRepository{}}

func Init() *gin.Engine {
	r := gin.Default()

	api := r.Group("/api")
	{
		api.GET("/", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "Hello, World!",
			})
		})
		api.POST("/login", authController.Login)
		api.POST("/users", userController.CreateUser)

		auth := api.Group("/")
		auth.Use(middlewares.AuthorizedMiddleware)
		{
			auth.GET("/users", userController.FindUser)

			auth.GET("/tasks", taskController.ListTasks)
			auth.GET("/tasks/:id", taskController.FindTask)
			auth.POST("/tasks", taskController.CreateTask)
		}
	}

	return r
}
