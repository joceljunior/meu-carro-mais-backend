package router

import (
	"meu-carro-mais/internal/handlers"

	"github.com/gin-gonic/gin"
)

type OfertaAutoMaisRouter struct{}

func (oar *OfertaAutoMaisRouter) RegisterRoutes(rg *gin.RouterGroup) {
	// CRUD básico de ofertas Auto Mais
	ofertas := rg.Group("/ofertas-auto-mais")
	{
		ofertas.POST("", handlers.CreateOfertaAutoMaisHandler)                    // POST /ofertas-auto-mais - Criar oferta
		ofertas.GET("", handlers.GetAllOfertasAutoMaisHandler)                    // GET /ofertas-auto-mais - Listar todas as ofertas
		
		// Endpoints específicos (devem vir antes dos endpoints com :id)
		ofertas.GET("/proximidade", handlers.GetOfertasAutoMaisByProximidadeHandler) // GET /ofertas-auto-mais/proximidade - Listar por proximidade
		ofertas.GET("/ativas", handlers.GetOfertasAutoMaisAtivasHandler)          // GET /ofertas-auto-mais/ativas - Listar ofertas ativas
		
		// Endpoints CRUD com :id (devem vir por último)
		ofertas.GET("/:id", handlers.GetOfertaAutoMaisHandler)                    // GET /ofertas-auto-mais/:id - Buscar oferta por ID
		ofertas.PUT("/:id", handlers.UpdateOfertaAutoMaisHandler)                 // PUT /ofertas-auto-mais/:id - Atualizar oferta
		ofertas.POST("/:id/desativar", handlers.DesativarOfertaAutoMaisHandler)   // POST /ofertas-auto-mais/:id/desativar - Desativar oferta
		ofertas.POST("/:id/ativar", handlers.AtivarOfertaAutoMaisHandler)         // POST /ofertas-auto-mais/:id/ativar - Ativar oferta
		ofertas.DELETE("/:id", handlers.SoftDeleteOfertaAutoMaisHandler)          // DELETE /ofertas-auto-mais/:id - Soft delete oferta
		ofertas.POST("/:id/restore", handlers.RestoreOfertaAutoMaisHandler)       // POST /ofertas-auto-mais/:id/restore - Restaurar oferta
	}

	// Rotas de ofertas via loja
	lojas := rg.Group("/lojas")
	{
		lojas.GET("/:id/ofertas-auto-mais", handlers.GetOfertasAutoMaisByLojaIDHandler)           // GET /lojas/:id/ofertas-auto-mais - Ofertas Auto Mais da loja
		lojas.GET("/:id/ofertas-auto-mais/ativas", handlers.GetOfertasAutoMaisAtivasByLojaIDHandler) // GET /lojas/:id/ofertas-auto-mais/ativas - Ofertas ativas da loja
	}
}

