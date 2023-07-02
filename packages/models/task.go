package models

import (
	"time"

	"github.com/google/uuid"
)

type Task struct {
	ID           uuid.UUID `gorm:"<-;type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	Title        string    `gorm:"<-;type:varchar(255);not null" json:"title"`
	Description  string    `gorm:"<-;type:varchar(255)" json:"description"`
	Completed    bool      `gorm:"<-;type:boolean;default:false" json:"completed"`
	ProgrammedTo time.Time `gorm:"<-;type:timestamp without time zone;not null" json:"programmed_to"`
	User         User      `gorm:"foreignKey:UserID" json:"-"`
	UserID       uuid.UUID `gorm:"<-;type:uuid;not null" json:"user_id"`
	CreatedAt    time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt    time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
