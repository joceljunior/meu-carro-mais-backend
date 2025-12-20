package router

import (
	"github.com/gin-gonic/gin"
	"meu-carro-mais/internal/handlers"
)

type LoginRouter struct{}

func (lr *LoginRouter) RegisterRoutes(rg *gin.RouterGroup) {
	// Login Mobile - para app mobile (cria usuário se não existir)
	rg.POST("/login", func(c *gin.Context) {
		handlers.LoginHandler(c.Writer, c.Request)
	})

	// Login Web - para plataforma web (apenas executivo, administrativo e customer)
	rg.POST("/login/web", handlers.LoginWebHandler)
}