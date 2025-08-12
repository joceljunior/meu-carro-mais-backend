package router

import (
	"github.com/gin-gonic/gin"
	"meu-carro-mais/internal/handlers"
)

type LojaRouter struct{}

func (lr *LojaRouter) RegisterRoutes(rg *gin.RouterGroup) {
	lojas := rg.Group("/lojas")
	{
		lojas.GET("/proximidade", handlers.GetLojasByProximidadeHandler)
		lojas.GET("/categorias", handlers.GetCategoriasLojistaHandler)
	}
} 