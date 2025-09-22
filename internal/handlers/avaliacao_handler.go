package handlers

import (
	"net/http"
	"strconv"

	"meu-carro-mais/internal/handlers/json"
	"meu-carro-mais/internal/services"

	"github.com/gin-gonic/gin"
)

// CreateAvaliacaoHandler godoc
// @Summary      Cria uma nova avaliação
// @Description  Cria uma nova avaliação de loja no sistema
// @Tags         Avaliações
// @Accept       json
// @Produce      json
// @Param        avaliacao body json.AvaliacaoRequest true "Dados da avaliação"
// @Success      201  {object}  json.AvaliacaoResponse
// @Failure      400  {object}  map[string]interface{} "Dados inválidos"
// @Failure      500  {object}  map[string]interface{} "Erro interno do servidor"
// @Router       /avaliacoes [post]
func CreateAvaliacaoHandler(c *gin.Context) {
	var req json.AvaliacaoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Dados inválidos: " + err.Error(),
		})
		return
	}

	resp, err := services.CreateAvaliacao(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, resp)
}

// GetAvaliacaoHandler godoc
// @Summary      Busca avaliação por ID
// @Description  Retorna uma avaliação específica pelo ID
// @Tags         Avaliações
// @Accept       json
// @Produce      json
// @Param        id path int true "ID da avaliação"
// @Success      200  {object}  json.AvaliacaoResponse
// @Failure      400  {object}  map[string]interface{} "ID inválido"
// @Failure      404  {object}  map[string]interface{} "Avaliação não encontrada"
// @Failure      500  {object}  map[string]interface{} "Erro interno do servidor"
// @Router       /avaliacoes/{id} [get]
func GetAvaliacaoHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID inválido",
		})
		return
	}

	resp, err := services.GetAvaliacaoByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Avaliação não encontrada",
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// GetAllAvaliacoesHandler godoc
// @Summary      Lista todas as avaliações
// @Description  Retorna todas as avaliações ativas do sistema
// @Tags         Avaliações
// @Accept       json
// @Produce      json
// @Success      200  {object}  json.AvaliacoesResponse
// @Failure      500  {object}  map[string]interface{} "Erro interno do servidor"
// @Router       /avaliacoes [get]
func GetAllAvaliacoesHandler(c *gin.Context) {
	resp, err := services.GetAllAvaliacoes()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	response := json.AvaliacoesResponse{
		Avaliacoes: resp,
		Total:      len(resp),
	}

	c.JSON(http.StatusOK, response)
}

// GetAvaliacoesByLojaIDHandler godoc
// @Summary      Lista avaliações de uma loja
// @Description  Retorna todas as avaliações de uma loja específica
// @Tags         Avaliações
// @Accept       json
// @Produce      json
// @Param        id path int true "ID da loja"
// @Success      200  {object}  json.AvaliacoesResponse
// @Failure      400  {object}  map[string]interface{} "ID de loja inválido"
// @Failure      500  {object}  map[string]interface{} "Erro interno do servidor"
// @Router       /lojas/{id}/avaliacoes [get]
func GetAvaliacoesByLojaIDHandler(c *gin.Context) {
	idLojaStr := c.Param("id")
	idLoja, err := strconv.ParseUint(idLojaStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID de loja inválido",
		})
		return
	}

	resp, err := services.GetAvaliacoesByLojaID(uint(idLoja))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// GetAvaliacoesByUsuarioIDHandler godoc
// @Summary      Lista avaliações de um usuário
// @Description  Retorna todas as avaliações feitas por um usuário específico
// @Tags         Avaliações
// @Accept       json
// @Produce      json
// @Param        id path int true "ID do usuário"
// @Success      200  {object}  json.AvaliacoesResponse
// @Failure      400  {object}  map[string]interface{} "ID de usuário inválido"
// @Failure      500  {object}  map[string]interface{} "Erro interno do servidor"
// @Router       /users/{id}/avaliacoes [get]
func GetAvaliacoesByUsuarioIDHandler(c *gin.Context) {
	idUsuarioStr := c.Param("id")
	idUsuario, err := strconv.ParseUint(idUsuarioStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID de usuário inválido",
		})
		return
	}

	resp, err := services.GetAvaliacoesByUsuarioID(uint(idUsuario))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// GetAvaliacaoEstatisticasByLojaIDHandler godoc
// @Summary      Estatísticas de avaliações de uma loja
// @Description  Retorna estatísticas detalhadas das avaliações de uma loja
// @Tags         Avaliações
// @Accept       json
// @Produce      json
// @Param        id path int true "ID da loja"
// @Success      200  {object}  json.AvaliacaoEstatisticasResponse
// @Failure      400  {object}  map[string]interface{} "ID de loja inválido"
// @Failure      500  {object}  map[string]interface{} "Erro interno do servidor"
// @Router       /lojas/{id}/avaliacoes/estatisticas [get]
func GetAvaliacaoEstatisticasByLojaIDHandler(c *gin.Context) {
	idLojaStr := c.Param("id")
	idLoja, err := strconv.ParseUint(idLojaStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID de loja inválido",
		})
		return
	}

	resp, err := services.GetAvaliacaoEstatisticasByLojaID(uint(idLoja))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// UpdateAvaliacaoHandler godoc
// @Summary      Atualiza avaliação
// @Description  Atualiza os dados de uma avaliação existente
// @Tags         Avaliações
// @Accept       json
// @Produce      json
// @Param        id path int true "ID da avaliação"
// @Param        avaliacao body json.AvaliacaoRequest true "Dados atualizados da avaliação"
// @Success      200  {object}  json.AvaliacaoResponse
// @Failure      400  {object}  map[string]interface{} "Dados inválidos"
// @Failure      404  {object}  map[string]interface{} "Avaliação não encontrada"
// @Failure      500  {object}  map[string]interface{} "Erro interno do servidor"
// @Router       /avaliacoes/{id} [put]
func UpdateAvaliacaoHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID inválido",
		})
		return
	}

	var req json.AvaliacaoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Dados inválidos: " + err.Error(),
		})
		return
	}

	resp, err := services.UpdateAvaliacao(uint(id), req)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Avaliação não encontrada",
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// SoftDeleteAvaliacaoHandler godoc
// @Summary      Remove avaliação (soft delete)
// @Description  Realiza soft delete de uma avaliação, marcando-a como excluída
// @Tags         Avaliações
// @Accept       json
// @Produce      json
// @Param        id path int true "ID da avaliação"
// @Success      200  {object}  map[string]interface{} "Avaliação removida com sucesso"
// @Failure      400  {object}  map[string]interface{} "ID inválido"
// @Failure      404  {object}  map[string]interface{} "Avaliação não encontrada"
// @Failure      500  {object}  map[string]interface{} "Erro interno do servidor"
// @Router       /avaliacoes/{id} [delete]
func SoftDeleteAvaliacaoHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID inválido",
		})
		return
	}

	err = services.SoftDeleteAvaliacao(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Avaliação não encontrada",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Avaliação removida com sucesso",
	})
}

// RestoreAvaliacaoHandler godoc
// @Summary      Restaura avaliação
// @Description  Restaura uma avaliação que foi soft deleted
// @Tags         Avaliações
// @Accept       json
// @Produce      json
// @Param        id path int true "ID da avaliação"
// @Success      200  {object}  map[string]interface{} "Avaliação restaurada com sucesso"
// @Failure      400  {object}  map[string]interface{} "ID inválido"
// @Failure      404  {object}  map[string]interface{} "Avaliação não encontrada"
// @Failure      500  {object}  map[string]interface{} "Erro interno do servidor"
// @Router       /avaliacoes/{id}/restore [post]
func RestoreAvaliacaoHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID inválido",
		})
		return
	}

	err = services.RestoreAvaliacao(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Avaliação não encontrada",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Avaliação restaurada com sucesso",
	})
}
