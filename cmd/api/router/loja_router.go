package router

import (
	"meu-carro-mais/internal/handlers"

	"github.com/gin-gonic/gin"
)

type LojaRouter struct{}

func (lr *LojaRouter) RegisterRoutes(rg *gin.RouterGroup) {
	lojas := rg.Group("/lojas")
	{
		// CRUD básico
		lojas.POST("", handlers.CreateLojaHandler) // POST /lojas - Criar loja
		lojas.GET("", handlers.GetAllLojasHandler) // GET /lojas - Listar todas as lojas

		// Endpoints específicos (devem vir antes dos endpoints com :id)
		lojas.GET("/:id/veiculos", handlers.GetVeiculosLojaByLojaIDHandler)                         // GET /lojas/:id/veiculos - Veículos da loja
		lojas.GET("/:id/produtos", handlers.GetProdutosByLojaIDHandler)                             // GET /lojas/:id/produtos - Produtos da loja
		lojas.GET("/:id/servicos", handlers.GetServicosByLojaIDHandler)                             // GET /lojas/:id/servicos - Serviços da loja
		lojas.GET("/:id/historicos-resgate", handlers.GetHistoricosResgateByLojaIDHandler)          // GET /lojas/:id/historicos-resgate - Históricos de resgate da loja
		lojas.POST("/:id/vendas-produto-avulso", handlers.CreateVendaProdutoAvulsoHandler)         // POST /lojas/:id/vendas-produto-avulso - Venda de produto não cadastrado
		lojas.GET("/:id/avaliacoes", handlers.GetAvaliacoesByLojaIDHandler)                         // GET /lojas/:id/avaliacoes - Avaliações da loja
		lojas.GET("/:id/avaliacoes/estatisticas", handlers.GetAvaliacaoEstatisticasByLojaIDHandler) // GET /lojas/:id/avaliacoes/estatisticas - Estatísticas de avaliações da loja
		lojas.GET("/proximidade", handlers.GetLojasByProximidadeHandler)                            // GET /lojas/proximidade - Listar por proximidade
		lojas.GET("/categorias", handlers.GetCategoriasLojistaHandler)                              // GET /lojas/categorias - Listar categorias
		lojas.GET("/usuario/:id_usuario", handlers.GetLojasByUsuarioIDHandler)                      // GET /lojas/usuario/:id_usuario - Listar lojas por usuário

		// Endpoints CRUD com :id (devem vir por último)
		lojas.GET("/:id", handlers.GetLojaHandler)              // GET /lojas/:id - Buscar loja por ID
		lojas.PUT("/:id", handlers.UpdateLojaHandler)           // PUT /lojas/:id - Atualizar loja
		lojas.DELETE("/:id", handlers.SoftDeleteLojaHandler)    // DELETE /lojas/:id - Soft delete loja
		lojas.POST("/:id/restore", handlers.RestoreLojaHandler) // POST /lojas/:id/restore - Restaurar loja
	}
}
