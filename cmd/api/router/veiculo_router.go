package router

import (
	"meu-carro-mais/internal/handlers"

	"github.com/gin-gonic/gin"
)

type VeiculoRouter struct{}

func (vr *VeiculoRouter) RegisterRoutes(rg *gin.RouterGroup) {
	// Rotas para veículos de usuários
	usuarios := rg.Group("/usuarios")
	{
		usuarios.GET("/:id_usuario/veiculos", handlers.GetVeiculosByUsuarioHandler)
		usuarios.GET("/:id_usuario/historico", handlers.GetHistoricosByUsuarioHandler)
	}

	// Rotas para histórico de veículos específicos
	veiculos := rg.Group("/veiculos")
	{
		veiculos.GET("/:id_veiculo/historico", handlers.GetHistoricosByVeiculoHandler)
	}
}
