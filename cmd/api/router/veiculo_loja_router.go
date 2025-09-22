package router

import (
	"meu-carro-mais/internal/handlers"

	"github.com/gin-gonic/gin"
)

type VeiculoLojaRouter struct{}

func (vlr *VeiculoLojaRouter) RegisterRoutes(rg *gin.RouterGroup) {
	// CRUD básico de veículos de loja
	veiculosLoja := rg.Group("/veiculos-loja")
	{
		veiculosLoja.POST("", handlers.CreateVeiculoLojaHandler)              // POST /veiculos-loja - Criar veículo de loja
		veiculosLoja.GET("", handlers.GetAllVeiculosLojaHandler)              // GET /veiculos-loja - Listar todos os veículos de loja
		veiculosLoja.GET("/:id", handlers.GetVeiculoLojaHandler)              // GET /veiculos-loja/:id - Buscar veículo de loja por ID
		veiculosLoja.PUT("/:id", handlers.UpdateVeiculoLojaHandler)           // PUT /veiculos-loja/:id - Atualizar veículo de loja
		veiculosLoja.DELETE("/:id", handlers.SoftDeleteVeiculoLojaHandler)    // DELETE /veiculos-loja/:id - Soft delete veículo de loja
		veiculosLoja.POST("/:id/restore", handlers.RestoreVeiculoLojaHandler) // POST /veiculos-loja/:id/restore - Restaurar veículo de loja
	}

}
