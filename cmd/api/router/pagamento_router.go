package router

import (
	"meu-carro-mais/internal/handlers"

	"github.com/gin-gonic/gin"
)

type PagamentoRouter struct{}

func (pr *PagamentoRouter) RegisterRoutes(rg *gin.RouterGroup) {
	// Rotas de pagamento
	pagamentos := rg.Group("/pagamentos")
	{
		// Endpoints específicos (devem vir antes dos endpoints com :id)
		pagamentos.POST("/checkout", handlers.CreateCheckoutSessionHandler)      // POST /pagamentos/checkout - Criar sessão de checkout
		pagamentos.POST("/webhook", handlers.ProcessWebhookHandler)              // POST /pagamentos/webhook - Processar webhook do Stripe
		pagamentos.GET("/historicos", handlers.GetAllHistoricosPagamentoHandler) // GET /pagamentos/historicos - Listar todos os históricos

		// Endpoints CRUD com :id (devem vir por último)
		pagamentos.GET("/historicos/:id", handlers.GetHistoricoPagamentoHandler)              // GET /pagamentos/historicos/:id - Buscar histórico por ID
		pagamentos.DELETE("/historicos/:id", handlers.SoftDeleteHistoricoPagamentoHandler)    // DELETE /pagamentos/historicos/:id - Soft delete histórico
		pagamentos.POST("/historicos/:id/restore", handlers.RestoreHistoricoPagamentoHandler) // POST /pagamentos/historicos/:id/restore - Restaurar histórico
	}

	// Rotas para históricos de pagamento de usuários
	usuarios := rg.Group("/usuarios")
	{
		usuarios.GET("/:id_usuario/historicos-pagamento", handlers.GetHistoricosPagamentoByUsuarioIDHandler) // GET /usuarios/:id_usuario/historicos-pagamento - Históricos de pagamento do usuário
	}
}
