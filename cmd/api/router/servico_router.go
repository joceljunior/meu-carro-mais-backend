package router

import (
	"meu-carro-mais/internal/handlers"

	"github.com/gin-gonic/gin"
)

type ServicoRouter struct{}

func (sr *ServicoRouter) RegisterRoutes(rg *gin.RouterGroup) {
	servicos := rg.Group("/servicos")
	{
		// CRUD básico
		servicos.POST("", handlers.CreateServicoHandler)              // POST /servicos - Criar serviço
		servicos.GET("", handlers.GetAllServicosHandler)              // GET /servicos - Listar todos os serviços
		servicos.GET("/:id", handlers.GetServicoHandler)              // GET /servicos/:id - Buscar serviço por ID
		servicos.PUT("/:id", handlers.UpdateServicoHandler)           // PUT /servicos/:id - Atualizar serviço
		servicos.DELETE("/:id", handlers.SoftDeleteServicoHandler)    // DELETE /servicos/:id - Soft delete serviço
		servicos.POST("/:id/restore", handlers.RestoreServicoHandler) // POST /servicos/:id/restore - Restaurar serviço

		// Endpoints específicos
		servicos.GET("/proximidade", handlers.GetServicosByProximidadeHandler) // GET /servicos/proximidade - Listar por proximidade
		servicos.GET("/categorias", handlers.GetCategoriasServicoHandler)      // GET /servicos/categorias - Listar categorias
	}
}
