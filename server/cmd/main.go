package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/demartinom/list-ranker/pkg/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	env := os.Getenv("GO_ENV")
	if env == "" {
		env = "development" // default
	}

	fmt.Printf("Running in %s environment\n", env)

	router := gin.Default()

	if env == "production" {
		router.Use(cors.New(cors.Config{
			AllowOrigins:     []string{"https://list-ranker-sage.vercel.app"},
			AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowHeaders:     []string{"Content-Type", "Authorization"},
			ExposeHeaders:    []string{"Content-Length"},
			AllowCredentials: true,
			MaxAge:           12 * time.Hour,
		}))
	} else {
		router.Use(cors.New(cors.Config{
			AllowOrigins:     []string{"http://localhost:5173"},
			AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowHeaders:     []string{"Content-Type", "Authorization"},
			ExposeHeaders:    []string{"Content-Length"},
			AllowCredentials: true,
			MaxAge:           12 * time.Hour,
		}))
	}

	routes.SetupRoutes(router)

	log.Println("Server running on http://localhost:8080")
	router.Run(":8080")
}
