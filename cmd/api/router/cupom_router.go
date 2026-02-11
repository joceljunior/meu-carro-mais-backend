package router

import (
	"meu-carro-mais/internal/handlers"

	"github.com/gin-gonic/gin"
)

type CupomRouter struct{}

func (cr *CupomRouter) RegisterRoutes(rg *gin.RouterGroup) {
	cupons := rg.Group("/cupons")
	{
		// CRUD básico
		cupons.POST("", handlers.CreateCupomHandler)              // POST /cupons - Criar cupom
		cupons.GET("", handlers.GetAllCuponsHandler)              // GET /cupons - Listar todos os cupons
		cupons.GET("/:id", handlers.GetCupomHandler)              // GET /cupons/:id - Buscar cupom por ID
		cupons.PUT("/:id", handlers.UpdateCupomHandler)           // PUT /cupons/:id - Atualizar cupom
		cupons.DELETE("/:id", handlers.SoftDeleteCupomHandler)    // DELETE /cupons/:id - Soft delete cupom
		cupons.POST("/:id/restore", handlers.RestoreCupomHandler) // POST /cupons/:id/restore - Restaurar cupom

		// Endpoints específicos (devem vir antes dos endpoints com :id)
		cupons.GET("/produtos", handlers.GetCuponsProdutosHandler)        // GET /cupons/produtos - Listar cupons de produtos
		cupons.GET("/veiculos", handlers.GetCuponsVeiculosHandler)       // GET /cupons/veiculos - Listar cupons de veículos
		cupons.GET("/servicos", handlers.GetCuponsServicosHandler)       // GET /cupons/servicos - Listar cupons de serviços
		cupons.GET("/loja/:loja_id", handlers.GetCuponsByLojaIDHandler) // GET /cupons/loja/:loja_id - Listar cupons por loja

		// Endpoints com :id (devem vir por último)
		cupons.POST("/:id/resgatar", handlers.ResgatarCupomHandler) // POST /cupons/:id/resgatar - Resgatar cupom (cria histórico de resgate)
	}
}
