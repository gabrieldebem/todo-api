package repositories

import (
	m "github.com/gabrieldebem/todo-api/packages/models"
)

type AccessTokenRepository struct {
	BaseRepository[m.AccessToken]
}
