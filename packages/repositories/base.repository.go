package repositories

import (
	"github.com/gabrieldebem/todo-api/database"
	"gorm.io/gorm"
)

type BaseRepository[T any] struct {
}

func (r BaseRepository[T]) GetDb() *gorm.DB {
	return database.Init()
}

func (r *BaseRepository[T]) FindAll() ([]T, error) {
	db := r.GetDb()

	var tasks []T
	err := db.Find(&tasks).Error

	return tasks, err
}

func (r *BaseRepository[T]) FindWhere(query interface{}, args ...interface{}) ([]T, error) {
	db := r.GetDb()

	var tasks []T
	err := db.Where(query, args...).Find(&tasks).Error

	return tasks, err
}

func (r *BaseRepository[T]) First(u *T) (T, error) {
	db := r.GetDb()

	err := db.First(u).Error

	return *u, err
}

func (r *BaseRepository[T]) Create(u *T) (T, error) {
	db := r.GetDb()

	err := db.Create(u).Error

	return *u, err
}

func (r *BaseRepository[T]) Update(u *T) (T, error) {
	db := r.GetDb()

	err := db.Save(u).Error

	return *u, err
}

func (r *BaseRepository[T]) Delete(u *T) (T, error) {
	db := r.GetDb()

	err := db.Delete(u).Error

	return *u, err
}
