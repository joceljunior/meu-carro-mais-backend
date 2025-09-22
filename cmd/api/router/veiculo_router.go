package router

import (
	"meu-carro-mais/internal/handlers"

	"github.com/gin-gonic/gin"
)

type VeiculoRouter struct{}

func (vr *VeiculoRouter) RegisterRoutes(rg *gin.RouterGroup) {
	// CRUD básico de veículos
	veiculos := rg.Group("/veiculos")
	{
		veiculos.POST("", handlers.CreateVeiculoHandler) // POST /veiculos - Criar veículo
		veiculos.GET("", handlers.GetAllVeiculosHandler) // GET /veiculos - Listar todos os veículos

		// Endpoints específicos (devem vir antes dos endpoints com :id)
		veiculos.GET("/:id/historico", handlers.GetHistoricosByVeiculoHandler) // GET /veiculos/:id/historico - Histórico do veículo

		// Endpoints CRUD com :id (devem vir por último)
		veiculos.GET("/:id", handlers.GetVeiculoHandler)              // GET /veiculos/:id - Buscar veículo por ID
		veiculos.PUT("/:id", handlers.UpdateVeiculoHandler)           // PUT /veiculos/:id - Atualizar veículo
		veiculos.DELETE("/:id", handlers.SoftDeleteVeiculoHandler)    // DELETE /veiculos/:id - Soft delete veículo
		veiculos.POST("/:id/restore", handlers.RestoreVeiculoHandler) // POST /veiculos/:id/restore - Restaurar veículo
	}

	// Rotas para veículos de usuários
	usuarios := rg.Group("/usuarios")
	{
		usuarios.GET("/:id_usuario/veiculos", handlers.GetVeiculosByUsuarioHandler)    // GET /usuarios/:id_usuario/veiculos - Veículos do usuário
		usuarios.GET("/:id_usuario/historico", handlers.GetHistoricosByUsuarioHandler) // GET /usuarios/:id_usuario/historico - Histórico do usuário
	}
}
