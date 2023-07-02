package controllers

import (
	"github.com/gabrieldebem/todo-api/packages/exceptions"
	"github.com/gabrieldebem/todo-api/packages/models"
	"github.com/gabrieldebem/todo-api/packages/repositories"
	"github.com/gabrieldebem/todo-api/packages/requests"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type TaskController struct {
    Repository repositories.TaskRepository
}

func (t TaskController) ListTasks(c *gin.Context) {
    tasks, err := t.Repository.FindAll()

    if err != nil {
        exceptions.HttpException{
            Context: c,
            Error:   err,
        }.Throw()
        return
    }

    c.JSON(200, gin.H{
        "data": tasks,
    })
}

func (t TaskController) FindTask(c *gin.Context) {
    id, err := uuid.Parse(c.Param("id"))

    if err != nil {
        exceptions.HttpException{
            Context: c,
            Error:   err,
            Message: "Invalid ID",
            Code:    400,
        }.Throw()
        return
    }

    task := models.Task{ID: id}

    task, err = t.Repository.First(&task)

    if err != nil {
        exceptions.HttpException{
            Context: c,
            Error:   err,
        }.Throw()
        return
    }

    c.JSON(200, gin.H{
        "task": task,
    })
}

func (t TaskController) CreateTask(c *gin.Context) {
    var taskDto requests.CreateTaskRequest

    c.Bind(&taskDto)
    var task models.Task = taskDto.ToTask()
    var err error

    task, err = t.Repository.Create(&task)

    if err != nil {
        exceptions.HttpException{
            Context: c,
            Error:   err,
        }.Throw()
        return
    }

    c.JSON(200, gin.H{
        "task": task,
    })
}
