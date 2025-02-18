package handlers

import (
	"net/http"

	"github.com/demartinom/list-ranker-v2/pkg/models"
	"github.com/gin-gonic/gin"
)

func SendBattlers( c *gin.Context)  {
	fighters, _ := models.ChooseBattlers(battleList)

	c.JSON(http.StatusOK, gin.H{"battlers":fighters})
}