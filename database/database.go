package database

import (
	m "github.com/gabrieldebem/todo-api/packages/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	db, err := gorm.Open(postgres.Open(getDns()), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	runMigration(db)

	return db
}

func runMigration(db *gorm.DB) {
	err := db.AutoMigrate(
		&m.User{},
		&m.Task{},
	)

	if err != nil {
		panic(err)
	}
}

func getDns() string {
	return "host=localhost user=postgres password=password dbname=todo port=5432 sslmode=disable TimeZone=America/Sao_Paulo"
}
