package router

import (
	"meu-carro-mais/internal/handlers"

	"github.com/gin-gonic/gin"
)

type HistoricoResgateRouter struct{}

func (hrr *HistoricoResgateRouter) RegisterRoutes(rg *gin.RouterGroup) {
	// CRUD básico de históricos de resgate
	historicosResgate := rg.Group("/historicos-resgate")
	{
		historicosResgate.POST("", handlers.CreateHistoricoResgateHandler)                 // POST /historicos-resgate - Criar histórico de resgate
		historicosResgate.GET("", handlers.GetAllHistoricosResgateHandler)                 // GET /historicos-resgate - Listar todos os históricos

		// Endpoints específicos (devem vir antes dos endpoints com :id)
		historicosResgate.GET("/cliente/:id", handlers.GetHistoricosResgateClienteByUsuarioIDHandler) // GET /historicos-resgate/cliente/:id - Histórico simplificado do cliente
		historicosResgate.GET("/usuario/:id", handlers.GetHistoricosResgateByUsuarioIDDirectHandler)    // GET /historicos-resgate/usuario/:id - Históricos de resgate do usuário
		historicosResgate.GET("/cupom/:id", handlers.GetHistoricosResgateByCupomIDHandler)         // GET /historicos-resgate/cupom/:id - Histórico de um cupom específico
		historicosResgate.PUT("/:id/status", handlers.UpdateStatusHistoricoResgateHandler) // PUT /historicos-resgate/:id/status - Atualizar status
		historicosResgate.PUT("/:id/aprovar", handlers.AprovarResgateHandler)              // PUT /historicos-resgate/:id/aprovar - Aprovar resgate
		historicosResgate.PUT("/:id/rejeitar", handlers.RejeitarResgateHandler)            // PUT /historicos-resgate/:id/rejeitar - Rejeitar resgate
		historicosResgate.POST("/:id/restore", handlers.RestoreHistoricoResgateHandler)     // POST /historicos-resgate/:id/restore - Restaurar histórico

		// Endpoints CRUD com :id (devem vir por último)
		historicosResgate.GET("/:id", handlers.GetHistoricoResgateHandler)           // GET /historicos-resgate/:id - Buscar histórico por ID
		historicosResgate.PUT("/:id", handlers.UpdateHistoricoResgateHandler)          // PUT /historicos-resgate/:id - Atualizar histórico
		historicosResgate.DELETE("/:id", handlers.SoftDeleteHistoricoResgateHandler)  // DELETE /historicos-resgate/:id - Soft delete histórico
	}
}
