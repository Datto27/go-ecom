package main

// @title           GoEcom
// @version         1.0
// @description     Golang + Gin api with basic ecomerce routes

// @license.name  Apache 2.0

// @host      localhost:4000
// @BasePath  /api/v1
import (
	"fmt"

	"github.com/datto27/goecom/cmd/api"
	"github.com/datto27/goecom/config"
	"github.com/datto27/goecom/db"
)

func main() {
	connStr := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s", config.ENVS.Host, config.ENVS.Port, config.ENVS.User, config.ENVS.DBName, config.ENVS.Password, config.ENVS.SSLMode)

	db := db.InitDB(connStr)

	server := api.NewAPIServer(":4000", db)
	server.Run()
}