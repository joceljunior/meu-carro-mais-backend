package router

import (
	"meu-carro-mais/internal/handlers"

	"github.com/gin-gonic/gin"
)

type UserRouter struct{}

func (ur *UserRouter) RegisterRoutes(rg *gin.RouterGroup) {
	users := rg.Group("/users")
	{
		users.POST("", handlers.CreateUserHandler)              // POST /users - Criar usuário
		users.GET("", handlers.GetAllUsersHandler)              // GET /users - Listar todos os usuários
		users.GET("/:id", handlers.GetUserHandler)              // GET /users/:id - Buscar usuário por ID
		users.PUT("/:id", handlers.UpdateUserHandler)           // PUT /users/:id - Atualizar usuário
		users.DELETE("/:id", handlers.SoftDeleteUserHandler)    // DELETE /users/:id - Soft delete usuário
		users.POST("/:id/restore", handlers.RestoreUserHandler) // POST /users/:id/restore - Restaurar usuário
	}
}
