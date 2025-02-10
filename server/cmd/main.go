package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
router := gin.Default()

log.Println("Server running on http://localhost:8080")
router.Run(":8080")
}
