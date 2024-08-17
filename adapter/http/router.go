package http

import (
	"delegacia-facil/adapter/http/router"
	"delegacia-facil/container"

	"github.com/gin-gonic/gin"
)

func SetupRouter(cont *container.Container) *gin.Engine {
	r := gin.Default()

	// Configuração das rotas com os handlers do container
	router.SetupUserRoutes(r, cont.UserHandler)
	router.SetupDelegaciasRoutes(r, cont.DelegaciaHandler)

	return r
}
