package api

import (
	"log"
	"net/http"

	"github.com/datto27/goecom/microservices/core/internal/auth"
	"github.com/datto27/goecom/microservices/core/internal/product"
	"github.com/datto27/goecom/microservices/core/internal/user"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

type APIServer struct {
	port string
	db   *gorm.DB
}

func NewAPIServer(port string, db *gorm.DB) *APIServer {
	return &APIServer{
		port: port,
		db:   db,
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
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// init repositories
	userRepository := user.NewRepository(s.db)
	productRepository := product.NewRepository(s.db)

	// init handlers
	userHandler := user.NewHandler(userRepository)
	authHandler := auth.NewHandler(userRepository)
	productHandler := product.NewHandler(productRepository)

	// group routes. in this case path will be /api/v1/auth....
	v1 := router.Group("/api/v1")
	{
		// v1.Use()
		authHandler.Routes(v1.Group("/auth"))
		userHandler.Routes(v1.Group("/users"))
		productHandler.Routes(v1.Group("/products"))
	}

	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"error": "endpoint not found"})
	})

	router.Run(s.port)
}
