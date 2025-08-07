package router

import (
	"meu-carro-mais/internal/handlers"

	"github.com/gin-gonic/gin"
)

type ServicoRouter struct{}

func (sr *ServicoRouter) RegisterRoutes(rg *gin.RouterGroup) {
	rg.GET("/servicos/proximidade", func(c *gin.Context) {
		handlers.GetServicosByProximidadeHandler(c.Writer, c.Request)
	})
}
