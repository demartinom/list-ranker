package handlers

import (
	"net/http"

	"github.com/demartinom/list-ranker-v2/pkg/models"
	"github.com/gin-gonic/gin"
)

// * TODO: Need to figure out how to reset this on page refresh
var battleList []*models.Item

func SendPremade(c *gin.Context) {
	battleList = []*models.Item{}

	premadeLists := models.PremadeList()
	c.JSON(http.StatusOK, gin.H{"premades": premadeLists})
}

func ReceiveChoice(c *gin.Context) {
	var req models.Choice

	// Bind JSON to struct
	if err := c.BindJSON(&req); err != nil {
		return
	}

	battleList = models.ReadCSV(req.Selection)
}
