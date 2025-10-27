package handlers

import (
	"net/http"

	"github.com/demartinom/list-ranker/pkg/models"
	"github.com/gin-gonic/gin"
)

func SendPremade(c *gin.Context) {
	premadeLists := models.PremadeList()
	c.JSON(http.StatusOK, gin.H{"premades": premadeLists})
}

func ReceiveChoice(c *gin.Context) {
	var req models.Choice

	// Bind JSON to struct
	if err := c.BindJSON(&req); err != nil {
		return
	}
	// Result final ranking to ensure no items left over from previous game
	models.FinalRanking = models.Ranking{}
	switch req.Type {
	case "premade":
		models.BattleList.SetGame(models.ReadCSV(req.Selection))
	case "custom":
		models.BattleList.SetGame(models.ReadCustom(req.Selection))
	}
}
