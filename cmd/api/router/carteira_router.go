package router

import (
	"meu-carro-mais/internal/handlers"

	"github.com/gin-gonic/gin"
)

type CarteiraRouter struct{}

func (cr *CarteiraRouter) RegisterRoutes(rg *gin.RouterGroup) {
	carteiras := rg.Group("/carteiras")
	{
		carteiras.POST("", handlers.CreateCarteiraHandler) // POST /carteiras - Criar carteira
		carteiras.GET("", handlers.GetAllCarteirasHandler) // GET /carteiras - Listar todas as carteiras

		// Endpoints específicos (devem vir antes dos endpoints com :id)
		carteiras.GET("/usuario/:usuario_id", handlers.GetCarteiraByUsuarioHandler) // GET /carteiras/usuario/:usuario_id - Carteira por usuário
		carteiras.GET("/range", handlers.GetCarteirasBySaldoRangeHandler)           // GET /carteiras/range?saldo_min=100&saldo_max=2000 - Carteiras por range de saldo
		carteiras.GET("/saldo-maior", handlers.GetCarteirasComSaldoMaiorHandler)    // GET /carteiras/saldo-maior?valor=1000 - Carteiras com saldo maior

		// Endpoints CRUD com :id (devem vir por último)
		carteiras.GET("/:id", handlers.GetCarteiraHandler)       // GET /carteiras/:id - Buscar carteira por ID
		carteiras.PUT("/:id", handlers.UpdateCarteiraHandler)    // PUT /carteiras/:id - Atualizar carteira
		carteiras.DELETE("/:id", handlers.DeleteCarteiraHandler) // DELETE /carteiras/:id - Excluir carteira

		// Endpoints de operações específicas
		carteiras.PUT("/:id/saldo", handlers.UpdateCarteiraSaldoHandler) // PUT /carteiras/:id/saldo - Atualizar apenas saldo
		carteiras.POST("/:id/adicionar", handlers.AdicionarSaldoHandler) // POST /carteiras/:id/adicionar - Adicionar saldo
		carteiras.POST("/:id/subtrair", handlers.SubtrairSaldoHandler)   // POST /carteiras/:id/subtrair - Subtrair saldo
	}
}
