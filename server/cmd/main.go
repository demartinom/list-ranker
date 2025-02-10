package main

import (
	"log"

	"github.com/demartinom/list-ranker-v2/pkg/routes"
	"github.com/gin-gonic/gin"
)

func main() {
router := gin.Default()

routes.SetupRoutes(router)

log.Println("Server running on http://localhost:8080")
router.Run(":8080")
}
