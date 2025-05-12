package main

// @title           GoEcom
// @version         1.0
// @description     Golang + Gin api with basic ecomerce routes

// @license.name  Apache 2.0

// @host      localhost:4000
// @BasePath  /api/v1
import (
	"fmt"

	"github.com/datto27/goecom/microservices/core/api"
	"github.com/datto27/goecom/microservices/core/db"
	"github.com/datto27/goecom/shared/config"
)

func main() {
	connStr := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s", config.ENVS.Host, config.ENVS.Port, config.ENVS.User, config.ENVS.DBName, config.ENVS.Password, config.ENVS.SSLMode)

	db := db.InitDB(connStr)

	server := api.NewAPIServer(":8080", db)
	server.Run()
}
