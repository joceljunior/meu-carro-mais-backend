package router

import (
	"meu-carro-mais/internal/handlers"

	"github.com/gin-gonic/gin"
)

type ProdutoRouter struct{}

func (pr *ProdutoRouter) RegisterRoutes(rg *gin.RouterGroup) {
	// CRUD básico de produtos
	produtos := rg.Group("/produtos")
	{
		produtos.POST("", handlers.CreateProdutoHandler)              // POST /produtos - Criar produto
		produtos.GET("", handlers.GetAllProdutosHandler)              // GET /produtos - Listar todos os produtos
		produtos.GET("/:id", handlers.GetProdutoHandler)              // GET /produtos/:id - Buscar produto por ID
		produtos.PUT("/:id", handlers.UpdateProdutoHandler)           // PUT /produtos/:id - Atualizar produto
		produtos.DELETE("/:id", handlers.SoftDeleteProdutoHandler)    // DELETE /produtos/:id - Soft delete produto
		produtos.POST("/:id/restore", handlers.RestoreProdutoHandler) // POST /produtos/:id/restore - Restaurar produto
	}
}
