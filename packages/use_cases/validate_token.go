package usecases

import (
	"time"

	"github.com/gabrieldebem/todo-api/packages/models"
)

func ValidateToken(token string) (bool, models.AccessToken) {
	tokens, err := r.FindAndLoadRelationship("token = ?", "User", token)

	if err != nil {
		return false, models.AccessToken{}
	}

	if len(tokens) == 0 {
		return false, models.AccessToken{}
	}

	if tokens[0].ExpiresAt.Before(time.Now()) {
		return false, models.AccessToken{}
	}

	return true, tokens[0]
}
