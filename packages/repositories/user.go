package repositories

import (
	"github.com/gabrieldebem/todo-api/packages/models"
	"github.com/google/uuid"
)

type IUserRepository interface {
    FindAll() []models.User
    First(u *models.User) models.User
    Create(u *models.User) models.User
    Update(u *models.User) models.User
    Delete(u *models.User) models.User
}

type UserRepository struct {
    BaseRepository 
}

func(r UserRepository) FindAll() []models.User {
    db := r.GetDb()

    var users = &[]models.User{}
    db.Find(&users)

    return *users
}

func (r UserRepository) First(id uuid.UUID) (models.User, error) {
    db := r.GetDb()

    u := &models.User{}

    err := db.Where("id = ?", id).First(&u).Error

    return *u, err
}

func (r UserRepository) Create(u *models.User) (models.User, error) {
    db := r.GetDb()

    err := db.Create(&u).Error

    return *u, err
}

func (r UserRepository) Update(u *models.User) (models.User, error) {
    db := r.GetDb()

    err := db.Save(&u).Error

    return *u, err
}

func (r UserRepository) Delete(u *models.User) (models.User, error) {
    db := r.GetDb()

    err := db.Delete(&u).Error

    return *u, err
}
