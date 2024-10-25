package db

import (
	"log"

	"github.com/datto27/goecom/pkg/models"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB(connStr string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}
	
	db.Debug().AutoMigrate(&models.User{})
	
	return db
}
