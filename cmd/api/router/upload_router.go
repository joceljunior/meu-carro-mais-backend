package router

import (
	"meu-carro-mais/internal/handlers"

	"github.com/gin-gonic/gin"
)

type UploadRouter struct{}

func (ur *UploadRouter) RegisterRoutes(rg *gin.RouterGroup) {
	// CRUD básico de uploads
	uploads := rg.Group("/uploads")
	{
		uploads.POST("", handlers.CreateUploadHandler) // POST /uploads - Criar upload
		uploads.GET("", handlers.GetAllUploadsHandler)  // GET /uploads - Listar todos os uploads

		// Endpoints específicos (devem vir antes dos endpoints com :id)
		uploads.GET("/principal/:tipo/:id", handlers.GetUploadPrincipalByEntidadeHandler) // GET /uploads/principal/:tipo/:id - Upload principal de uma entidade

		// Endpoints CRUD com :id (devem vir por último)
		uploads.GET("/:id", handlers.GetUploadHandler)                    // GET /uploads/:id - Buscar upload por ID
		uploads.PUT("/:id", handlers.UpdateUploadHandler)                // PUT /uploads/:id - Atualizar upload
		uploads.PUT("/:id/principal", handlers.SetUploadPrincipalHandler) // PUT /uploads/:id/principal - Definir como principal (apenas imagens)
		uploads.DELETE("/:id", handlers.SoftDeleteUploadHandler)         // DELETE /uploads/:id - Soft delete upload
		uploads.POST("/:id/restore", handlers.RestoreUploadHandler)     // POST /uploads/:id/restore - Restaurar upload
	}

	// Rotas para uploads de usuários
	usuarios := rg.Group("/usuarios")
	{
		usuarios.GET("/:id_usuario/uploads", handlers.GetUploadsByUsuarioIDHandler) // GET /usuarios/:id_usuario/uploads - Uploads do usuário
	}

	// Rotas para uploads de veículos
	veiculos := rg.Group("/veiculos")
	{
		veiculos.GET("/:id/uploads", handlers.GetUploadsByVeiculoIDHandler) // GET /veiculos/:id/uploads - Uploads do veículo
	}

	// Rotas para uploads de veículos de loja
	veiculosLoja := rg.Group("/veiculos-loja")
	{
		veiculosLoja.GET("/:id/uploads", handlers.GetUploadsByVeiculoLojaIDHandler) // GET /veiculos-loja/:id/uploads - Uploads do veículo de loja
	}

	// Rotas para uploads de produtos
	produtos := rg.Group("/produtos")
	{
		produtos.GET("/:id/uploads", handlers.GetUploadsByProdutoIDHandler) // GET /produtos/:id/uploads - Uploads do produto
	}

	// Rotas para uploads de serviços
	servicos := rg.Group("/servicos")
	{
		servicos.GET("/:id/uploads", handlers.GetUploadsByServicoIDHandler) // GET /servicos/:id/uploads - Uploads do serviço
	}

	// Rotas para uploads de lojas
	lojas := rg.Group("/lojas")
	{
		lojas.GET("/:id/uploads", handlers.GetUploadsByLojaIDHandler) // GET /lojas/:id/uploads - Uploads da loja
	}
}

