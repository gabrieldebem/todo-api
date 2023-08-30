package middlewares

import (
	usecases "github.com/gabrieldebem/todo-api/packages/use_cases"
	"github.com/gin-gonic/gin"
)

func AuthorizedMiddleware(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if token == "" {
		c.JSON(401, gin.H{
			"error": "Unauthorized",
			"data":  "You must be authenticated to access this route",
		})
		c.Abort()
		return
	}

    token = token[7:]

	isValid, model := usecases.ValidateToken(token)

	if !isValid {
		c.JSON(401, gin.H{
			"error": "Unauthorized",
			"data":  "You must be authenticated to access this route",
		})
		c.Abort()

		return
	}

	c.Set("user", model.User)

	c.Next()
}
