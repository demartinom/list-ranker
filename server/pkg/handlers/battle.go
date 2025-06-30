package handlers

import (
	"net/http"

	"github.com/demartinom/list-ranker/pkg/models"
	"github.com/gin-gonic/gin"
)

// Assign BattleList state to local variable
var battleList = &models.BattleList
var RoundRobin = &models.RoundRobin

func SendBattlers(c *gin.Context) {
	round := models.BeginRound(models.BattleList.BattleList)
	itemsLeft := len(models.BattleList.BattleList)

	if itemsLeft == 5 && !models.RoundRobinMode {
		models.RoundRobinMode = true
		models.RoundRobinRounds(battleList.BattleList)
	}

	if round != nil {
		c.JSON(http.StatusOK, gin.H{"results": models.FinalRanking.RankingsList, "itemsLeft": itemsLeft})
		return
	} else if models.RoundRobinMode {
		currentRound := RoundRobin.Current

		c.JSON(http.StatusOK, gin.H{"battlers": []*models.Item{RoundRobin.FightList[currentRound][0], RoundRobin.FightList[currentRound][1]}, "itemsLeft": itemsLeft, "roundRobin": true})
		//Change above to something like Threshold - current?
	} else {
		c.JSON(http.StatusOK, gin.H{"battlers": battleList.CurrentCombatants, "itemsLeft": itemsLeft, "roundRobin": false})
	}
}

func ReceiveBattlerChoice(c *gin.Context) {
	var req models.Choice
	if err := c.BindJSON(&req); err != nil {
		return
	}
	if !models.RoundRobinMode {
		models.BattleResult(battleList.CurrentCombatants, battleList.CurrentCombatants, battleList.CurrentIndexes, req.Selection)
	} else {
		RoundRobin.RRRound(req.Selection)
	}
}
