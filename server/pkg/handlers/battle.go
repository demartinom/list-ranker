package handlers

import (
	"net/http"

	"github.com/demartinom/list-ranker/pkg/models"
	"github.com/gin-gonic/gin"
)

// Assign BattleList state to local variable
var battleList = &models.BattleList
var RoundRobin = &models.RoundRobin

// TODO: Explicit mode flag
func SendBattlers(c *gin.Context) {
	round := models.BeginRound(models.BattleList.BattleList)
	itemsLeft := len(models.BattleList.BattleList)

	if itemsLeft == 4 {
		models.RoundRobinRounds(battleList.BattleList)
		c.JSON(http.StatusOK, gin.H{"battlers": []*models.Item{RoundRobin.FightList[RoundRobin.Current][0], RoundRobin.FightList[RoundRobin.Current][1]}, "itemsLeft": itemsLeft})

		RoundRobin.Current++

		return
	}

	if round != nil {
		c.JSON(http.StatusOK, gin.H{"results": models.FinalRanking.RankingsList, "itemsLeft": itemsLeft})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"battlers": battleList.CurrentCombatants, "itemsLeft": itemsLeft})
	}
}

func ReceiveBattlerChoice(c *gin.Context) {
	var req models.Choice

	if err := c.BindJSON(&req); err != nil {
		return
	}
	if !models.RoundRobinMode {
		models.BattleResult(battleList.CurrentCombatants, battleList.CurrentCombatants, battleList.CurrentIndexes, req.Selection)
	}
}
