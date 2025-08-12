package router

import (
	"meu-carro-mais/internal/handlers"

	"github.com/gin-gonic/gin"
)

type AnuncioRouter struct{}

func (ar *AnuncioRouter) RegisterRoutes(rg *gin.RouterGroup) {
	anuncios := rg.Group("/anuncios")
	{
		anuncios.GET("", handlers.GetAnunciosHandler)
		anuncios.GET("/categorias", handlers.GetCategoriasAnuncioHandler)
	}
}
