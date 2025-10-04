package handlers

import (
	"meu-carro-mais/internal/handlers/json"
	"meu-carro-mais/internal/services"
	"net/http"
	"strconv"

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

// GetUserHandler godoc
// @Summary      Busca usuário por ID
// @Description  Retorna os dados de um usuário específico pelo ID
// @Tags         Usuários
// @Accept       json
// @Produce      json
// @Param        id path int true "ID do usuário"
// @Success      200 {object} json.UserResponse "Usuário encontrado"
// @Failure      400 {object} map[string]interface{} "ID inválido"
// @Failure      404 {object} map[string]interface{} "Usuário não encontrado"
// @Failure      500 {object} map[string]interface{} "Erro interno do servidor"
// @Router       /users/{id} [get]
func GetUserHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID inválido",
		})
		return
	}

	resp, err := services.GetUserByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Usuário não encontrado",
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// GetAllUsersHandler godoc
// @Summary      Lista todos os usuários
// @Description  Retorna uma lista com todos os usuários ativos
// @Tags         Usuários
// @Accept       json
// @Produce      json
// @Success      200 {array} json.UserResponse "Lista de usuários"
// @Failure      500 {object} map[string]interface{} "Erro interno do servidor"
// @Router       /users [get]
func GetAllUsersHandler(c *gin.Context) {
	resp, err := services.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// UpdateUserHandler godoc
// @Summary      Atualiza usuário
// @Description  Atualiza os dados de um usuário existente
// @Tags         Usuários
// @Accept       json
// @Produce      json
// @Param        id path int true "ID do usuário"
// @Param        request body json.UserRequest true "Dados atualizados do usuário"
// @Success      200 {object} json.UserResponse "Usuário atualizado com sucesso"
// @Failure      400 {object} map[string]interface{} "Dados inválidos"
// @Failure      404 {object} map[string]interface{} "Usuário não encontrado"
// @Failure      500 {object} map[string]interface{} "Erro interno do servidor"
// @Router       /users/{id} [put]
func UpdateUserHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID inválido",
		})
		return
	}

	var req json.UserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Dados inválidos",
			"details": err.Error(),
		})
		return
	}

	resp, err := services.UpdateUser(uint(id), req)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Usuário não encontrado",
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// SoftDeleteUserHandler godoc
// @Summary      Exclui usuário (soft delete)
// @Description  Realiza soft delete do usuário, marcando como excluído sem remover do banco
// @Tags         Usuários
// @Accept       json
// @Produce      json
// @Param        id path int true "ID do usuário"
// @Success      200 {object} map[string]interface{} "Usuário excluído com sucesso"
// @Failure      400 {object} map[string]interface{} "ID inválido"
// @Failure      404 {object} map[string]interface{} "Usuário não encontrado"
// @Failure      500 {object} map[string]interface{} "Erro interno do servidor"
// @Router       /users/{id} [delete]
func SoftDeleteUserHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID inválido",
		})
		return
	}

	err = services.SoftDeleteUser(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Usuário não encontrado",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Usuário excluído com sucesso",
	})
}

// RestoreUserHandler godoc
// @Summary      Restaura usuário excluído
// @Description  Restaura um usuário que foi soft deleted
// @Tags         Usuários
// @Accept       json
// @Produce      json
// @Param        id path int true "ID do usuário"
// @Success      200 {object} map[string]interface{} "Usuário restaurado com sucesso"
// @Failure      400 {object} map[string]interface{} "ID inválido"
// @Failure      404 {object} map[string]interface{} "Usuário não encontrado"
// @Failure      500 {object} map[string]interface{} "Erro interno do servidor"
// @Router       /users/{id}/restore [post]
func RestoreUserHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID inválido",
		})
		return
	}

	err = services.RestoreUser(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Usuário não encontrado ou não foi excluído",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Usuário restaurado com sucesso",
	})
}

// GetUserPlanStatusHandler godoc
// @Summary Verifica status do plano do usuário
// @Description Retorna informações sobre o plano atual do usuário e se ele é premium
// @Tags Usuários
// @Accept json
// @Produce json
// @Param id path int true "ID do usuário"
// @Success 200 {object} json.UserPlanStatusResponse "Status do plano do usuário"
// @Failure 400 {object} map[string]interface{} "ID inválido"
// @Failure 404 {object} map[string]interface{} "Usuário não encontrado"
// @Failure 500 {object} map[string]interface{} "Erro interno do servidor"
// @Router /users/{id}/plan-status [get]
func GetUserPlanStatusHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID inválido",
		})
		return
	}

	resp, err := services.GetUserPlanStatus(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Usuário não encontrado",
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}
