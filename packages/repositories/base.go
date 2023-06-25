package repositories

import (
	"github.com/gabrieldebem/todo-api/database"
	"gorm.io/gorm"
)

type BaseRepository struct {
}

func (r BaseRepository) GetDb() *gorm.DB {
    return database.Init()
}

