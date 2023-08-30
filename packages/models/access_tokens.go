package models

import (
	"time"

	"github.com/google/uuid"
)

type AccessToken struct {
	ID        uuid.UUID `gorm:"<-;type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	UserID    uuid.UUID `json:"user_id" gorm:"type:uuid;not null;"`
	User      User      `json:"user"`
	Token     string    `json:"-" gorm:"type:varchar(255);not null;"`
	ExpiresAt time.Time `json:"expires_at" gorm:"not null;"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
