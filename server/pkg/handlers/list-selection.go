package handlers

import (
	"fmt"
	"net/http"

	"github.com/demartinom/list-ranker-v2/pkg/models"
	"github.com/gin-gonic/gin"
)

type Choice struct {
	Selection string `json:"selection"`
}


func SendPremade(c *gin.Context){
	premadeLists := models.PremadeList()
	
	c.JSON(http.StatusOK, gin.H{"premades":premadeLists})
}

func ReceiveChoice(c *gin.Context) {
	var req Choice

	// Bind JSON to struct
	if err := c.BindJSON(&req); err != nil {
		return
	}

	fmt.Println("Received choice:", req.Selection)
}