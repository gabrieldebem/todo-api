package requests

import (
	"time"

	"github.com/gabrieldebem/todo-api/packages/models"
	"github.com/google/uuid"
)

type CreateTaskRequest struct {
	Title        string `json:"title" validate:"required"`
	Description  string `json:"description"`
	ProgrammedTo string `json:"programmed_to" validate:"required"`
	UserID       string `json:"user_id" validate:"required"`
}

func (r CreateTaskRequest) ToTask() models.Task {
	date, err := time.Parse("2006-01-02 03:04:05", r.ProgrammedTo)

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
