package controllers

import (
	"github.com/gabrieldebem/todo-api/packages/exceptions"
	"github.com/gabrieldebem/todo-api/packages/models"
	"github.com/gabrieldebem/todo-api/packages/repositories"
	"github.com/gabrieldebem/todo-api/packages/requests"
	"github.com/gabrieldebem/todo-api/packages/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type TaskController struct {
	Repository repositories.TaskRepository
}

func (t TaskController) ListTasks(c *gin.Context) {
	tasks, err := t.Repository.FindAndLoadRelationship("user_id = ?", "User", utils.AuthUser(c).ID.String())

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

	var task []models.Task
	task, err = t.Repository.FindAndLoadRelationship(
		"id = ? AND user_id = ?",
        "User",
		id.String(),
		utils.AuthUser(c).ID.String(),
	)

	if err != nil {
		exceptions.HttpException{
			Context: c,
			Error:   err,
		}.Throw()
		return
	}

	c.JSON(200, gin.H{
		"task": task[0],
	})
}

func (t TaskController) CreateTask(c *gin.Context) {
	user := utils.AuthUser(c)
	var taskDto requests.CreateTaskRequest

	c.Bind(&taskDto)

	var err error

	err = taskDto.Validate(c)

	if err != nil {
		return
	}

	var task models.Task = taskDto.ToTask(user)

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
