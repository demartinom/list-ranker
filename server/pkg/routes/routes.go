package routes

import (
	"github.com/demartinom/list-ranker-v2/pkg/handlers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	router.GET("/api/premades", handlers.SendPremade)
	router.POST("/api/listchoice", handlers.ReceiveChoice)
	router.POST("/api/battlers", handlers.SendBattlers)
	router.POST("/api/battlerChoice", handlers.ReceiveBattlerChoice)
}
