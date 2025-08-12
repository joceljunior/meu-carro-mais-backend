package router

import (
	"meu-carro-mais/internal/handlers"

	"github.com/gin-gonic/gin"
)

type UserRouter struct{}

func (ur *UserRouter) RegisterRoutes(rg *gin.RouterGroup) {
	users := rg.Group("/users")
	{
		users.POST("", handlers.CreateUserHandler)
	}
}
