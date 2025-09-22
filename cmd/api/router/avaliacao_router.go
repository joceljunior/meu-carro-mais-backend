package router

import (
	"meu-carro-mais/internal/handlers"

	"github.com/gin-gonic/gin"
)

type AvaliacaoRouter struct{}

func (ar *AvaliacaoRouter) RegisterRoutes(rg *gin.RouterGroup) {
	// CRUD básico de avaliações
	avaliacoes := rg.Group("/avaliacoes")
	{
		avaliacoes.POST("", handlers.CreateAvaliacaoHandler)              // POST /avaliacoes - Criar avaliação
		avaliacoes.GET("", handlers.GetAllAvaliacoesHandler)              // GET /avaliacoes - Listar todas as avaliações
		avaliacoes.GET("/:id", handlers.GetAvaliacaoHandler)              // GET /avaliacoes/:id - Buscar avaliação por ID
		avaliacoes.PUT("/:id", handlers.UpdateAvaliacaoHandler)           // PUT /avaliacoes/:id - Atualizar avaliação
		avaliacoes.DELETE("/:id", handlers.SoftDeleteAvaliacaoHandler)    // DELETE /avaliacoes/:id - Soft delete avaliação
		avaliacoes.POST("/:id/restore", handlers.RestoreAvaliacaoHandler) // POST /avaliacoes/:id/restore - Restaurar avaliação
	}
}
