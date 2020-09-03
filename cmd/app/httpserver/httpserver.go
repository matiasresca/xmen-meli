package httpserver

import (
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Start() {
	Dependencies.Initialize()
	engine := gin.Default()
	engine.Use(cors.Default())
	RegistriRoutes(engine)
	log.Fatal(engine.Run(":8080"))
}
