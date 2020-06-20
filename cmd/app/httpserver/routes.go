package httpserver

import (
	"github.com/gin-gonic/gin"
)

func RegistriRoutes(engine *gin.Engine) {
	engine.POST("/mutant/", Dependencies.HumanHandler.Post)
	engine.GET("/stats", Dependencies.HumanHandler.Get)
}
