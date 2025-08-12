package router

import (
	"meu-carro-mais/internal/handlers"

	"github.com/gin-gonic/gin"
)

type ServicoRouter struct{}

func (sr *ServicoRouter) RegisterRoutes(rg *gin.RouterGroup) {
	servicos := rg.Group("/servicos")
	{
		servicos.GET("/proximidade", handlers.GetServicosByProximidadeHandler)
		servicos.GET("/categorias", handlers.GetCategoriasServicoHandler)
	}
}
