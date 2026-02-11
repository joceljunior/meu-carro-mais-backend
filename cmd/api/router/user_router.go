package router

import (
	"meu-carro-mais/internal/handlers"

	"github.com/gin-gonic/gin"
)

type UserRouter struct{}

func (ur *UserRouter) RegisterRoutes(rg *gin.RouterGroup) {
	users := rg.Group("/users")
	{
		users.POST("", handlers.CreateUserHandler) // POST /users - Criar usuário (mobile)
		users.GET("", handlers.GetAllUsersHandler) // GET /users - Listar todos os usuários

		// =====================================================
		// ENDPOINTS ADMINISTRATIVO (criar usuários especiais)
		// =====================================================
		users.POST("/administrativo", handlers.CreateAdministrativoHandler) // POST /users/administrativo - Criar usuário administrativo
		users.POST("/executivo", handlers.CreateExecutivoHandler)           // POST /users/executivo - Criar usuário executivo

		// =====================================================
		// ENDPOINTS CUSTOMER
		// =====================================================
		users.POST("/customer", handlers.CreateCustomerHandler)                // POST /users/customer - Criar customer (pendente aprovação)
		users.GET("/customers", handlers.GetAllCustomersHandler)               // GET /users/customers - Listar todos customers (com filtro por status)
		users.GET("/customers/pendentes", handlers.GetCustomersPendentesHandler) // GET /users/customers/pendentes - Listar customers pendentes
		users.POST("/customers/:id/aprovar", handlers.AprovarCustomerHandler)  // POST /users/customers/:id/aprovar - Aprovar customer
		users.POST("/customers/:id/rejeitar", handlers.RejeitarCustomerHandler) // POST /users/customers/:id/rejeitar - Rejeitar customer

		// =====================================================
		// ENDPOINTS EXECUTIVO
		// =====================================================
		users.GET("/executivos", handlers.GetAllExecutivosHandler) // GET /users/executivos - Listar todos os executivos

		// =====================================================
		// ENDPOINTS SOLICITAÇÃO DE EXECUTIVO (para administrativo)
		// =====================================================
		users.GET("/solicitacoes-executivo", handlers.GetSolicitacoesExecutivoPendentesHandler) // GET /users/solicitacoes-executivo - Listar solicitações pendentes

		// =====================================================
		// ENDPOINTS ESPECÍFICOS DO USUÁRIO (devem vir antes dos endpoints com :id)
		// =====================================================
		users.GET("/:id/veiculos", handlers.GetVeiculosByUsuarioHandler)                      // GET /users/:id/veiculos - Veículos do usuário
		users.GET("/:id/historico", handlers.GetHistoricosByUsuarioHandler)                   // GET /users/:id/historico - Histórico do usuário
		users.GET("/:id/historicos-resgate", handlers.GetHistoricosResgateByUsuarioIDHandler) // GET /users/:id/historicos-resgate - Históricos de resgate do usuário
		users.GET("/:id/avaliacoes", handlers.GetAvaliacoesByUsuarioIDHandler)                // GET /users/:id/avaliacoes - Avaliações do usuário
		users.GET("/:id/plan-status", handlers.GetUserPlanStatusHandler)                      // GET /users/:id/plan-status - Status do plano do usuário
		users.POST("/:id/solicitar-executivo", handlers.SolicitarExecutivoHandler)            // POST /users/:id/solicitar-executivo - Usuário mobile solicita virar executivo
		users.POST("/:id/aprovar-executivo", handlers.AprovarSolicitacaoExecutivoHandler)     // POST /users/:id/aprovar-executivo - Aprovar solicitação de executivo
		users.POST("/:id/rejeitar-executivo", handlers.RejeitarSolicitacaoExecutivoHandler)   // POST /users/:id/rejeitar-executivo - Rejeitar solicitação de executivo
		users.POST("/:id/cancelar-executivo", handlers.CancelarExecutivoHandler)              // POST /users/:id/cancelar-executivo - Cancelar executivo aprovado

		// =====================================================
		// ENDPOINTS CRUD COM :id (devem vir por último)
		// =====================================================
		users.GET("/:id", handlers.GetUserHandler)              // GET /users/:id - Buscar usuário por ID
		users.PUT("/:id", handlers.UpdateUserHandler)           // PUT /users/:id - Atualizar usuário
		users.DELETE("/:id", handlers.SoftDeleteUserHandler)    // DELETE /users/:id - Soft delete usuário
		users.POST("/:id/restore", handlers.RestoreUserHandler) // POST /users/:id/restore - Restaurar usuário
	}
}
