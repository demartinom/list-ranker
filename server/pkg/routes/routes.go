// Package includes list of routes to communicate with the frontend
package routes

import (
	"github.com/demartinom/list-ranker/pkg/handlers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	router.GET("/api/premades", handlers.SendPremade)
	router.POST("/api/listchoice", handlers.ReceiveChoice)
	router.POST("/api/battlers", handlers.SendBattlers)
	router.POST("/api/battlerChoice", handlers.ReceiveBattlerChoice)
}
