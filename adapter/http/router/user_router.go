package router

import (
	"test/adapter/http/handler"

	"github.com/gin-gonic/gin"
)

// SetupUserRoutes configura as rotas relacionadas a usuários
func SetupUserRoutes(router *gin.Engine, userHandler *handler.UserHandler) {
	userGroup := router.Group("/users")
	{
		userGroup.POST("/", userHandler.CreateUser)
		userGroup.GET("/:id", userHandler.GetUser)
		// Outras rotas relacionadas a usuários
	}
}
