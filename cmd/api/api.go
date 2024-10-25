package api

import (
	"log"
	"net/http"

	"github.com/datto27/goecom/internal/auth"
	"github.com/datto27/goecom/internal/user"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type APIServer struct {
	port string
	db *gorm.DB
}

func NewAPIServer(port string, db *gorm.DB) *APIServer {
	return &APIServer {
		port: port,
		db: db,
	}
}

func (s *APIServer) Run() {
	router := gin.Default()

	gin.SetMode(gin.DebugMode)

	log.Println("API server running on port: ", s.port)

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// init repositories
	userRepository := user.NewRepository(s.db)

	// init handlers
	userHandler := user.NewHandler(userRepository)
	authHandler := auth.NewHandler(userRepository)

	// group routes. in this case path will be /api/v1/auth....
	v1 := router.Group("/api/v1")
	{
		// v1.Use()
		authHandler.Routes(v1.Group("/auth"))
		userHandler.Routes(v1.Group("/user"))
	}

	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"error": "endpoint not found"})
	})

	router.Run(s.port)
}
