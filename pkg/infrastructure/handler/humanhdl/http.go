package humanhdl

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/matiasresca/xmen-meli/pkg/core/domain"
	"github.com/matiasresca/xmen-meli/pkg/core/ports"
)

type HTTPHandler struct {
	service ports.HumanService
}

func NewHTTPHandler(service ports.HumanService) *HTTPHandler {
	return &HTTPHandler{service: service}
}

func (hdl *HTTPHandler) Post(ctx *gin.Context) {
	//Parseo del Json al Modelo.-
	var human domain.Human
	err := json.NewDecoder(ctx.Request.Body).Decode(&human)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": "Error al convertir Json a Model",
		})
		return
	}
	//Llamada al servicio.-
	isMutant, err := hdl.service.CheckMutant(human.Dna)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	//Respuestas.-
	if isMutant {
		ctx.String(http.StatusOK, "")
	} else {
		ctx.String(http.StatusForbidden, "")
	}
}

func (hdl *HTTPHandler) Get(ctx *gin.Context) {
	//Obtengo las estadisticas.-
	stats, err := hdl.service.GetStats()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	//Respuesta.-
	ctx.JSON(http.StatusOK, stats)
}
