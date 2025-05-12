package db

import (
	"log"

	"github.com/datto27/goecom/microservices/core/models"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB(connStr string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})

	if err != nil {
		log.Fatalln("db error: ", err)
	}

	merr := db.Debug().AutoMigrate(&models.User{}, &models.Product{})

	if merr != nil {
		log.Fatalln("migration error: ", merr)
	}

	return db
}
