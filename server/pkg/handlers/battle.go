package handlers

import (
	"net/http"

	"github.com/demartinom/list-ranker-v2/pkg/models"
	"github.com/gin-gonic/gin"
)

// Assign BattleList state to local variable
var battleList = &models.BattleList

func SendBattlers(c *gin.Context) {
	models.BeginRound(models.BattleList.BattleList)

	c.JSON(http.StatusOK, gin.H{"battlers": battleList.CurrentCombatants})
}

func ReceiveBattlerChoice(c *gin.Context) {
	var req models.Choice

	if err := c.BindJSON(&req); err != nil {
		return
	}
	models.BattleResult(battleList.CurrentCombatants, battleList.CurrentCombatants, battleList.CurrentIndexes, req.Selection)
}
