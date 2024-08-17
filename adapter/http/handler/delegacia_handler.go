package handler

import (
	"delegacia-facil/core/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type DelegaciaHandler struct {
	delegaciaUseCase *usecase.DelegaciaUseCase
}

// NewDelegaciaHandler cria uma nova instância de DelegaciaHandler
func NewDelegaciaHandler(duc *usecase.DelegaciaUseCase) *DelegaciaHandler {
	return &DelegaciaHandler{delegaciaUseCase: duc}
}

func (dh *DelegaciaHandler) ListaDelegacias(c *gin.Context) {
	delegacias, err := dh.delegaciaUseCase.ListaDelegacias()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, delegacias)
}

func (dh *DelegaciaHandler) FiltroDelegacias(c *gin.Context) {
	horario24hStr := c.DefaultQuery("horario24h", "false")
	horario24h, err := strconv.ParseBool(horario24hStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid value for horario24h, expected a boolean"})
		return
	}

	delegacias, err := dh.delegaciaUseCase.FiltraDelegacias(horario24h)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, delegacias)
}
