package router

import (
	"meu-carro-mais/internal/handlers"

	"github.com/gin-gonic/gin"
)

type LogRouter struct{}

func (lr *LogRouter) RegisterRoutes(rg *gin.RouterGroup) {
	logs := rg.Group("/logs")
	{
		logs.GET("", handlers.GetAllLogsHandler)                      // GET /logs - Listar todos os logs
		logs.GET("/:id", handlers.GetLogByIDHandler)                  // GET /logs/:id - Buscar log por ID
		logs.GET("/usuario/:id", handlers.GetLogsByUsuarioIDHandler)   // GET /logs/usuario/:id - Logs de um usuário
		logs.GET("/entidade/:entidade/:id", handlers.GetLogsByEntidadeHandler) // GET /logs/entidade/:entidade/:id - Logs de uma entidade
		logs.GET("/acao/:tipo", handlers.GetLogsByTipoAcaoHandler)    // GET /logs/acao/:tipo - Logs por tipo de ação
	}
}

