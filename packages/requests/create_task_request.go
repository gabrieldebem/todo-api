package requests

import (
	"time"

	"github.com/gabrieldebem/todo-api/packages/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

var defaultTimeFormat string = "2006-01-02 03:04:05"

type CreateTaskRequest struct {
	Title        string `json:"title" validate:"required"`
	Description  string `json:"description"`
	ProgrammedTo string `json:"programmed_to" validate:"required,datetime=2006-01-02 03:04:05"`
	UserID       string `json:"user_id" validate:"required,uuid"`
}

func (r CreateTaskRequest) ToTask() models.Task {
	date, err := time.Parse(defaultTimeFormat, r.ProgrammedTo)

	if err != nil {
		panic(err)
	}

	return models.Task{
		Title:        r.Title,
		Description:  r.Description,
		ProgrammedTo: date,
		UserID:       uuid.MustParse(r.UserID),
	}
}

func (r CreateTaskRequest) Validate(c *gin.Context) error {
	err := validator.New().Struct(r)

	if err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid request",
			"errors":  err.Error(),
		})
	}

	return err
}
