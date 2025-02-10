package handlers

import (
	"net/http"

	"github.com/demartinom/list-ranker-v2/pkg/models"
	"github.com/gin-gonic/gin"
)

func SendPremade(c *gin.Context){
	premadeLists := models.PremadeList()
	
	c.JSON(http.StatusOK, gin.H{"premade options":premadeLists})
}