package router

import (
	"meu-carro-mais/internal/handlers"

	"github.com/gin-gonic/gin"
)

type TransferenciaRouter struct{}

func (tr *TransferenciaRouter) RegisterRoutes(rg *gin.RouterGroup) {
	// Rotas de transferências
	transferencias := rg.Group("/transferencias")
	{
		transferencias.GET("", handlers.GetAllTransferenciasHandler)                       // GET /transferencias - Lista todas as transferências
		transferencias.GET("/buscar-usuarios", handlers.BuscarUsuariosParaTransferenciaHandler) // GET /transferencias/buscar-usuarios - Busca usuários para transferência
		transferencias.GET("/:id", handlers.GetTransferenciaHandler)                       // GET /transferencias/:id - Busca transferência por ID
	}

	// Rotas de transferências por veículo
	veiculos := rg.Group("/veiculos")
	{
		veiculos.GET("/:id/transferencias", handlers.GetTransferenciasByVeiculoHandler) // GET /veiculos/:id/transferencias - Lista transferências do veículo
	}

	// Rotas de transferências por usuário
	usuarios := rg.Group("/usuarios")
	{
		usuarios.GET("/:id_usuario/transferencias", handlers.GetTransferenciasByUsuarioHandler) // GET /usuarios/:id_usuario/transferencias - Lista transferências do usuário
		usuarios.POST("/:id_usuario/transferir-veiculo", handlers.TransferirVeiculoHandler)     // POST /usuarios/:id_usuario/transferir-veiculo - Transfere veículo
	}
}
