package router

import (
	"meu-carro-mais/internal/handlers"

	"github.com/gin-gonic/gin"
)

type DescontoRouter struct{}

func (dr *DescontoRouter) RegisterRoutes(rg *gin.RouterGroup) {
	// CRUD básico de descontos
	descontos := rg.Group("/descontos")
	{
		descontos.POST("", handlers.CreateDescontoHandler)                // POST /descontos - Criar desconto
		descontos.GET("", handlers.GetAllDescontosHandler)                // GET /descontos - Listar todos os descontos
		descontos.GET("/ativos", handlers.GetDescontosAtivosHandler)      // GET /descontos/ativos - Listar descontos ativos
		descontos.GET("/:id", handlers.GetDescontoHandler)                // GET /descontos/:id - Buscar desconto por ID
		descontos.POST("/:id/cancelar", handlers.CancelarDescontoHandler) // POST /descontos/:id/cancelar - Cancelar desconto
		descontos.DELETE("/:id", handlers.SoftDeleteDescontoHandler)      // DELETE /descontos/:id - Soft delete desconto
		descontos.POST("/:id/restore", handlers.RestoreDescontoHandler)   // POST /descontos/:id/restore - Restaurar desconto
	}

	// Rotas de desconto via loja
	lojas := rg.Group("/lojas")
	{
		lojas.GET("/:id/descontos", handlers.GetDescontosByLojaIDHandler)                // GET /lojas/:id/descontos - Histórico de descontos da loja
		lojas.GET("/:id/desconto-ativo", handlers.GetDescontoAtivoByLojaIDHandler)       // GET /lojas/:id/desconto-ativo - Desconto ativo da loja
		lojas.POST("/:id/desconto-ativo/cancelar", handlers.CancelarDescontoLojaHandler) // POST /lojas/:id/desconto-ativo/cancelar - Cancelar desconto ativo da loja
	}
}

