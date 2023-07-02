package repositories

import "github.com/gabrieldebem/todo-api/packages/models"

type TaskRepository struct {
	BaseRepository[models.Task]
}
