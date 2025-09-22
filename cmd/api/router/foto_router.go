package router

import (
	"meu-carro-mais/internal/handlers"

	"github.com/gin-gonic/gin"
)

type FotoRouter struct{}

func (fr *FotoRouter) RegisterRoutes(rg *gin.RouterGroup) {
	// CRUD básico de fotos
	fotos := rg.Group("/fotos")
	{
		fotos.POST("", handlers.CreateFotoHandler) // POST /fotos - Criar foto
		fotos.GET("", handlers.GetAllFotosHandler) // GET /fotos - Listar todas as fotos

		// Endpoints específicos (devem vir antes dos endpoints com :id)
		fotos.GET("/principal/:tipo/:id", handlers.GetFotoPrincipalByEntidadeHandler) // GET /fotos/principal/:tipo/:id - Foto principal de uma entidade

		// Endpoints CRUD com :id (devem vir por último)
		fotos.GET("/:id", handlers.GetFotoHandler)                    // GET /fotos/:id - Buscar foto por ID
		fotos.PUT("/:id", handlers.UpdateFotoHandler)                 // PUT /fotos/:id - Atualizar foto
		fotos.PUT("/:id/principal", handlers.SetFotoPrincipalHandler) // PUT /fotos/:id/principal - Definir como principal
		fotos.DELETE("/:id", handlers.SoftDeleteFotoHandler)          // DELETE /fotos/:id - Soft delete foto
		fotos.POST("/:id/restore", handlers.RestoreFotoHandler)       // POST /fotos/:id/restore - Restaurar foto
	}

	// Rotas para fotos de veículos
	veiculos := rg.Group("/veiculos")
	{
		veiculos.GET("/:id/fotos", handlers.GetFotosByVeiculoIDHandler) // GET /veiculos/:id/fotos - Fotos do veículo
	}

	// Rotas para fotos de veículos de loja
	veiculosLoja := rg.Group("/veiculos-loja")
	{
		veiculosLoja.GET("/:id/fotos", handlers.GetFotosByVeiculoLojaIDHandler) // GET /veiculos-loja/:id/fotos - Fotos do veículo de loja
	}

	// Rotas para fotos de produtos
	produtos := rg.Group("/produtos")
	{
		produtos.GET("/:id/fotos", handlers.GetFotosByProdutoIDHandler) // GET /produtos/:id/fotos - Fotos do produto
	}

	// Rotas para fotos de serviços
	servicos := rg.Group("/servicos")
	{
		servicos.GET("/:id/fotos", handlers.GetFotosByServicoIDHandler) // GET /servicos/:id/fotos - Fotos do serviço
	}

	// Rotas para fotos de lojas
	lojas := rg.Group("/lojas")
	{
		lojas.GET("/:id/fotos", handlers.GetFotosByLojaIDHandler) // GET /lojas/:id/fotos - Fotos da loja
	}
}
