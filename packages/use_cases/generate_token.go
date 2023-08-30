package usecases

import (
	"time"

	"github.com/gabrieldebem/todo-api/packages/models"
	"github.com/gabrieldebem/todo-api/packages/repositories"
	"github.com/gabrieldebem/todo-api/packages/utils"
	"github.com/google/uuid"
)

var r repositories.AccessTokenRepository = repositories.AccessTokenRepository{}

func GenerateToken(userId uuid.UUID) (models.AccessToken, error) {
	str := utils.RandStringBytes()

	expiration := time.Now().Add(1 * time.Hour)
	token := models.AccessToken{
		UserID:    userId,
		Token:     str,
		ExpiresAt: expiration,
	}

	token, err := r.Create(&token)

	return token, err
}
