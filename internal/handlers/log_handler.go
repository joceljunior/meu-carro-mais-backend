package handlers

import (
	"net/http"
	"strconv"

	"meu-carro-mais/internal/handlers/json"
	"meu-carro-mais/internal/services"

	"github.com/gin-gonic/gin"
)

// GetAllLogsHandler godoc
// @Summary      Lista todos os logs
// @Description  Retorna todos os logs do sistema com paginação
// @Tags         Logs
// @Accept       json
// @Produce      json
// @Param        limit query int false "Limite de resultados (padrão: 50)"
// @Param        offset query int false "Offset para paginação (padrão: 0)"
// @Success      200  {object}  json.LogsResponse
// @Failure      500  {object}  map[string]interface{} "Erro interno do servidor"
// @Router       /logs [get]
func GetAllLogsHandler(c *gin.Context) {
	limitStr := c.DefaultQuery("limit", "50")
	offsetStr := c.DefaultQuery("offset", "0")

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit < 1 {
		limit = 50
	}
	if limit > 1000 {
		limit = 1000 // Máximo de 1000 registros por página
	}

	offset, err := strconv.Atoi(offsetStr)
	if err != nil || offset < 0 {
		offset = 0
	}

	logs, total, err := services.GetAllLogs(limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	responses := services.ConvertLogsToResponse(logs)
	c.JSON(http.StatusOK, json.LogsResponse{
		Logs:   responses,
		Total:  total,
		Limit:  limit,
		Offset: offset,
	})
}

// GetLogByIDHandler godoc
// @Summary      Busca log por ID
// @Description  Retorna um log específico pelo ID
// @Tags         Logs
// @Accept       json
// @Produce      json
// @Param        id path int true "ID do log"
// @Success      200  {object}  json.LogResponse
// @Failure      400  {object}  map[string]interface{} "ID inválido"
// @Failure      404  {object}  map[string]interface{} "Log não encontrado"
// @Failure      500  {object}  map[string]interface{} "Erro interno do servidor"
// @Router       /logs/{id} [get]
func GetLogByIDHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID inválido",
		})
		return
	}

	log, err := services.GetLogByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Log não encontrado",
		})
		return
	}

	response := services.ConvertLogToResponse(*log)
	c.JSON(http.StatusOK, response)
}

// GetLogsByUsuarioIDHandler godoc
// @Summary      Lista logs de um usuário
// @Description  Retorna todos os logs de um usuário específico
// @Tags         Logs
// @Accept       json
// @Produce      json
// @Param        id path int true "ID do usuário"
// @Success      200  {array}  json.LogResponse
// @Failure      400  {object}  map[string]interface{} "ID inválido"
// @Failure      500  {object}  map[string]interface{} "Erro interno do servidor"
// @Router       /logs/usuario/{id} [get]
func GetLogsByUsuarioIDHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID inválido",
		})
		return
	}

	logs, err := services.GetLogsByUsuarioID(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	responses := services.ConvertLogsToResponse(logs)
	c.JSON(http.StatusOK, responses)
}

// GetLogsByEntidadeHandler godoc
// @Summary      Lista logs de uma entidade
// @Description  Retorna todos os logs de uma entidade específica
// @Tags         Logs
// @Accept       json
// @Produce      json
// @Param        entidade path string true "Nome da entidade (cupom, produto, servico, etc.)"
// @Param        id path int true "ID da entidade"
// @Success      200  {array}  json.LogResponse
// @Failure      400  {object}  map[string]interface{} "Parâmetros inválidos"
// @Failure      500  {object}  map[string]interface{} "Erro interno do servidor"
// @Router       /logs/entidade/{entidade}/{id} [get]
func GetLogsByEntidadeHandler(c *gin.Context) {
	entidade := c.Param("entidade")
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID inválido",
		})
		return
	}

	logs, err := services.GetLogsByEntidade(entidade, uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	responses := services.ConvertLogsToResponse(logs)
	c.JSON(http.StatusOK, responses)
}

// GetLogsByTipoAcaoHandler godoc
// @Summary      Lista logs por tipo de ação
// @Description  Retorna todos os logs de um tipo de ação específico
// @Tags         Logs
// @Accept       json
// @Produce      json
// @Param        tipo path string true "Tipo de ação (criar, atualizar, deletar, resgatar, etc.)"
// @Success      200  {array}  json.LogResponse
// @Failure      500  {object}  map[string]interface{} "Erro interno do servidor"
// @Router       /logs/acao/{tipo} [get]
func GetLogsByTipoAcaoHandler(c *gin.Context) {
	tipoAcao := c.Param("tipo")

	logs, err := services.GetLogsByTipoAcao(tipoAcao)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	responses := services.ConvertLogsToResponse(logs)
	c.JSON(http.StatusOK, responses)
}

