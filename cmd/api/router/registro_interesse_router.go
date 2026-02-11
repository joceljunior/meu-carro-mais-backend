package router

import (
	"meu-carro-mais/internal/handlers"

	"github.com/gin-gonic/gin"
)

type RegistroInteresseRouter struct{}

func (rir *RegistroInteresseRouter) RegisterRoutes(rg *gin.RouterGroup) {
	registroInteresse := rg.Group("/registro-interesse")
	{
		// CRUD básico
		registroInteresse.POST("", handlers.CreateRegistroInteresseHandler)              // POST /registro-interesse - Criar registro de interesse
		registroInteresse.GET("", handlers.GetAllRegistroInteressesHandler)              // GET /registro-interesse - Listar todos os registros

		// Endpoints específicos (devem vir antes dos endpoints com :id)
		registroInteresse.GET("/cupom/:cupom_id", handlers.GetRegistroInteressesByCupomHandler) // GET /registro-interesse/cupom/:cupom_id - Listar por cupom

		// Endpoints CRUD com :id (devem vir por último)
		registroInteresse.GET("/:id", handlers.GetRegistroInteresseHandler)              // GET /registro-interesse/:id - Buscar registro por ID
		registroInteresse.DELETE("/:id", handlers.SoftDeleteRegistroInteresseHandler)    // DELETE /registro-interesse/:id - Soft delete registro
		registroInteresse.POST("/:id/restore", handlers.RestoreRegistroInteresseHandler) // POST /registro-interesse/:id/restore - Restaurar registro
	}
}

