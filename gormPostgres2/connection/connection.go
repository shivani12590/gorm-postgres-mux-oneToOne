package connection

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"postgres/gormPostgres2/models"
)

func DatabaseConnection() *gorm.DB {
	url := os.Getenv("url")
	var err error
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})
	if db != nil {
		log.Println("database connected")
	}
	if err != nil {
		log.Println(err)
	}

	db.AutoMigrate(models.Employee{}, models.Department{})
	return db

}
