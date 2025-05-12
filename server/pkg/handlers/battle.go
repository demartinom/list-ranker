package handlers

import (
	"net/http"

	"github.com/demartinom/list-ranker/pkg/models"
	"github.com/gin-gonic/gin"
)

// Assign BattleList state to local variable
var battleList = &models.BattleList

func SendBattlers(c *gin.Context) {
	round := models.BeginRound(models.BattleList.BattleList)
	if round != nil {
		c.JSON(http.StatusOK, gin.H{"results": models.FinalRanking.RankingsList})
	} else {
		c.JSON(http.StatusOK, gin.H{"battlers": battleList.CurrentCombatants, "itemsLeft": len(models.BattleList.BattleList)})
	}
}

func ReceiveBattlerChoice(c *gin.Context) {
	var req models.Choice

	if err := c.BindJSON(&req); err != nil {
		return
	}
	models.BattleResult(battleList.CurrentCombatants, battleList.CurrentCombatants, battleList.CurrentIndexes, req.Selection)
}
