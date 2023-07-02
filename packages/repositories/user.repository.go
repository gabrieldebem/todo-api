package repositories

import (
	"github.com/gabrieldebem/todo-api/packages/models"
)

type UserRepository struct {
	BaseRepository[models.User]
}
