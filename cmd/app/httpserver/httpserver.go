package httpserver

import (
	"github.com/gin-gonic/gin"
	"log"
)

func Start() {
	Dependencies.Initialize()
	engine := gin.Default()
	RegistriRoutes(engine)
	log.Fatal(engine.Run(":8080"))
}
