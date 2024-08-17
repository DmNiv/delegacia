package router

import (
	"delegacia-facil/adapter/http/handler"

	"github.com/gin-gonic/gin"
)

func SetupDelegaciasRoutes(router *gin.Engine, delegaciaHandler *handler.DelegaciaHandler) {
	delegaciasRoutes := router.Group("/delegacias")
	{
		delegaciasRoutes.GET("/lista-delegacias", delegaciaHandler.ListaDelegacias)
		delegaciasRoutes.GET("/filtro", delegaciaHandler.FiltroDelegacias)
	}
}
