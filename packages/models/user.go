package models

import (
	"crypto/sha256"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `gorm:"<-;type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	Name      string    `json:"name" gorm:"type:varchar(255);not null;"`
    Email     string    `json:"email" gorm:"unique:true;index;not null"`
	Password  string    `json:"-" gorm:"type:varchar(255);not null;"`
	Tasks     []Task    `json:"-"`
    AccessTokens []AccessToken `json:"-"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

func (u *User) CheckPassword(password string) bool {
	hash := sha256.New().Sum([]byte(password))
	hashedPassword := fmt.Sprintf("%x", hash)

	return hashedPassword == u.Password
}
