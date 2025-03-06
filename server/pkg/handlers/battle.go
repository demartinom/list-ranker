package handlers

import (
	"fmt"
	"net/http"

	"github.com/demartinom/list-ranker-v2/pkg/models"
	"github.com/gin-gonic/gin"
)

func SendBattlers(c *gin.Context) {
	models.ChooseBattlers(models.BattleList.BattleList)

	c.JSON(http.StatusOK, gin.H{"battlers": models.BattleList.CurrentCombatants})
}

func ReceiveBattlerChoice(c *gin.Context) {
	var req models.Choice

	if err := c.BindJSON(&req); err != nil {
		return
	}
	fmt.Println(req)
}
