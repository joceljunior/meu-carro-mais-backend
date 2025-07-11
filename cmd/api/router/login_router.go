package router

import (
	"github.com/gin-gonic/gin"
	"meu-carro-mais/internal/handlers"
)

type LoginRouter struct{}

func (lr *LoginRouter) RegisterRoutes(rg *gin.RouterGroup) {
	rg.POST("/login", func(c *gin.Context) {
		handlers.LoginHandler(c.Writer, c.Request)
	})
} 