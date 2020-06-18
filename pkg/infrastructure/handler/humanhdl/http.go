package humanhdl

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/matiasresca/xmen-meli/pkg/core/domain"
	"github.com/matiasresca/xmen-meli/pkg/core/ports"
	"net/http"
)

type HTTPHandler struct {
	service ports.HumanService
}

func NewHTTPHandler(service ports.HumanService) *HTTPHandler {
	return &HTTPHandler{service: service}
}

func (hdl *HTTPHandler) Post(ctx *gin.Context) {
	var human domain.Human
	err := json.NewDecoder(ctx.Request.Body).Decode(&human)
	if err != nil {
		panic("Error en el parseo de Human")
	}
	fmt.Println("human DNA => ", human.Dna)
	//Respuesta
	ctx.AbortWithStatus(http.StatusOK)
}
