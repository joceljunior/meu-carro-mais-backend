package router

import (
	"meu-carro-mais/internal/handlers"

	"github.com/gin-gonic/gin"
)

type AnuncioRouter struct{}

func (ar *AnuncioRouter) RegisterRoutes(rg *gin.RouterGroup) {
	anuncios := rg.Group("/anuncios")
	{
		// CRUD básico
		anuncios.POST("", handlers.CreateAnuncioHandler)              // POST /anuncios - Criar anúncio
		anuncios.GET("", handlers.GetAllAnunciosHandler)              // GET /anuncios - Listar todos os anúncios
		anuncios.GET("/:id", handlers.GetAnuncioHandler)              // GET /anuncios/:id - Buscar anúncio por ID
		anuncios.PUT("/:id", handlers.UpdateAnuncioHandler)           // PUT /anuncios/:id - Atualizar anúncio
		anuncios.DELETE("/:id", handlers.SoftDeleteAnuncioHandler)    // DELETE /anuncios/:id - Soft delete anúncio
		anuncios.POST("/:id/restore", handlers.RestoreAnuncioHandler) // POST /anuncios/:id/restore - Restaurar anúncio

		// Endpoints específicos (devem vir antes dos endpoints com :id)
		anuncios.GET("/produtos", handlers.GetAnunciosProdutosHandler)        // GET /anuncios/produtos - Listar anúncios de produtos
		anuncios.GET("/veiculos", handlers.GetAnunciosVeiculosHandler)       // GET /anuncios/veiculos - Listar anúncios de veículos
		anuncios.GET("/loja/:loja_id", handlers.GetAnunciosByLojaIDHandler) // GET /anuncios/loja/:loja_id - Listar anúncios por loja

		// Endpoints com :id (devem vir por último)
		anuncios.POST("/:id/resgatar", handlers.ResgatarAnuncioHandler) // POST /anuncios/:id/resgatar - Resgatar anúncio (cria histórico de resgate)
	}
}
