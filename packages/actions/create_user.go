package actions

import (
	"github.com/gabrieldebem/todo-api/packages/repositories"
	"github.com/gabrieldebem/todo-api/packages/requests"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
    var r requests.CreateUserRequest

    c.Bind(&r)
    user := r.ToUser()
    var err error

    user, err = repositories.UserRepository{}.Create(&user)

    if err != nil {
        c.JSON(400, gin.H{
            "error": err.Error(),
        })
        return
    }

    c.JSON(200, gin.H{
        "user": user,
    })
}
