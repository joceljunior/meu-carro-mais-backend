package router

import (
	"github.com/gin-gonic/gin"
	"meu-carro-mais/internal/handlers"
)

type LojaRouter struct{}

func (lr *LojaRouter) RegisterRoutes(rg *gin.RouterGroup) {
	rg.GET("/lojas/proximidade", func(c *gin.Context) {
		handlers.GetLojasByProximidadeHandler(c.Writer, c.Request)
	})
} 