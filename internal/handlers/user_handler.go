package handlers

import (
	"meu-carro-mais/internal/handlers/json"
	"meu-carro-mais/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateUserHandler godoc
// @Summary      Criação do usuário completo
// @Description  Cria um novo usuário com todos os dados fornecidos
// @Tags         Usuários
// @Accept       json
// @Produce      json
// @Param        request body json.UserRequest true "Dados completos do usuário"
// @Success      201  {object}  json.UserResponse "Usuário criado com sucesso"
// @Failure      400  {object}  map[string]interface{} "Dados inválidos"
// @Failure      500  {object}  map[string]interface{} "Erro interno do servidor"
// @Router       /users [post]
func CreateUserHandler(c *gin.Context) {
	var req json.UserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Dados inválidos",
			"details": err.Error(),
		})
		return
	}

	resp, err := services.CreateUser(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, resp)
}
