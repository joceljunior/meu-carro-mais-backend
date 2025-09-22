package router

import (
	"meu-carro-mais/internal/handlers"

	"github.com/gin-gonic/gin"
)

type UserRouter struct{}

func (ur *UserRouter) RegisterRoutes(rg *gin.RouterGroup) {
	users := rg.Group("/users")
	{
		users.POST("", handlers.CreateUserHandler) // POST /users - Criar usuário
		users.GET("", handlers.GetAllUsersHandler) // GET /users - Listar todos os usuários

		// Endpoints específicos (devem vir antes dos endpoints com :id)
		users.GET("/:id/veiculos", handlers.GetVeiculosByUsuarioHandler)                      // GET /users/:id/veiculos - Veículos do usuário
		users.GET("/:id/historico", handlers.GetHistoricosByUsuarioHandler)                   // GET /users/:id/historico - Histórico do usuário
		users.GET("/:id/historicos-resgate", handlers.GetHistoricosResgateByUsuarioIDHandler) // GET /users/:id/historicos-resgate - Históricos de resgate do usuário
		users.GET("/:id/avaliacoes", handlers.GetAvaliacoesByUsuarioIDHandler)                // GET /users/:id/avaliacoes - Avaliações do usuário

		// Endpoints CRUD com :id (devem vir por último)
		users.GET("/:id", handlers.GetUserHandler)              // GET /users/:id - Buscar usuário por ID
		users.PUT("/:id", handlers.UpdateUserHandler)           // PUT /users/:id - Atualizar usuário
		users.DELETE("/:id", handlers.SoftDeleteUserHandler)    // DELETE /users/:id - Soft delete usuário
		users.POST("/:id/restore", handlers.RestoreUserHandler) // POST /users/:id/restore - Restaurar usuário
	}
}
