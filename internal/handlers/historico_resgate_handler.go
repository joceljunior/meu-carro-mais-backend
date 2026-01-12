package handlers

import (
	"net/http"
	"strconv"

	"meu-carro-mais/internal/handlers/json"
	"meu-carro-mais/internal/services"

	"github.com/gin-gonic/gin"
)

// CreateHistoricoResgateHandler godoc
// @Summary      Cria um novo histórico de resgate
// @Description  Cria um novo histórico de resgate no sistema
// @Tags         Histórico de Resgates
// @Accept       json
// @Produce      json
// @Param        historico body json.HistoricoResgateRequest true "Dados do histórico de resgate"
// @Success      201  {object}  json.HistoricoResgateResponse
// @Failure      400  {object}  map[string]interface{} "Dados inválidos"
// @Failure      500  {object}  map[string]interface{} "Erro interno do servidor"
// @Router       /historicos-resgate [post]
func CreateHistoricoResgateHandler(c *gin.Context) {
	var req json.HistoricoResgateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Dados inválidos: " + err.Error(),
		})
		return
	}

	resp, err := services.CreateHistoricoResgate(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, resp)
}

// GetHistoricoResgateHandler godoc
// @Summary      Busca histórico de resgate por ID
// @Description  Retorna um histórico de resgate específico pelo ID
// @Tags         Histórico de Resgates
// @Accept       json
// @Produce      json
// @Param        id path int true "ID do histórico de resgate"
// @Success      200  {object}  json.HistoricoResgateResponse
// @Failure      400  {object}  map[string]interface{} "ID inválido"
// @Failure      404  {object}  map[string]interface{} "Histórico não encontrado"
// @Failure      500  {object}  map[string]interface{} "Erro interno do servidor"
// @Router       /historicos-resgate/{id} [get]
func GetHistoricoResgateHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID inválido",
		})
		return
	}

	resp, err := services.GetHistoricoResgateByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Histórico não encontrado",
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// GetAllHistoricosResgateHandler godoc
// @Summary      Lista todos os históricos de resgate
// @Description  Retorna todos os históricos de resgate ativos do sistema
// @Tags         Histórico de Resgates
// @Accept       json
// @Produce      json
// @Success      200  {object}  json.HistoricosResgateResponse
// @Failure      500  {object}  map[string]interface{} "Erro interno do servidor"
// @Router       /historicos-resgate [get]
func GetAllHistoricosResgateHandler(c *gin.Context) {
	resp, err := services.GetAllHistoricosResgate()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	response := json.HistoricosResgateResponse{
		Historicos: resp,
		Total:      len(resp),
	}

	c.JSON(http.StatusOK, response)
}

// GetHistoricosResgateByUsuarioIDHandler godoc
// @Summary      Lista históricos de resgate de um usuário
// @Description  Retorna todos os históricos de resgate de um usuário específico
// @Tags         Histórico de Resgates
// @Accept       json
// @Produce      json
// @Param        id path int true "ID do usuário"
// @Success      200  {object}  json.HistoricosResgateResponse
// @Failure      400  {object}  map[string]interface{} "ID de usuário inválido"
// @Failure      500  {object}  map[string]interface{} "Erro interno do servidor"
// @Router       /usuarios/{id}/historicos-resgate [get]
func GetHistoricosResgateByUsuarioIDHandler(c *gin.Context) {
	idUsuarioStr := c.Param("id")
	idUsuario, err := strconv.ParseUint(idUsuarioStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID de usuário inválido",
		})
		return
	}

	resp, err := services.GetHistoricosResgateByUsuarioID(uint(idUsuario))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// GetHistoricosResgateByLojaIDHandler godoc
// @Summary      Lista históricos de resgate de uma loja
// @Description  Retorna todos os históricos de resgate de uma loja específica
// @Tags         Histórico de Resgates
// @Accept       json
// @Produce      json
// @Param        id path int true "ID da loja"
// @Success      200  {object}  json.HistoricosResgateResponse
// @Failure      400  {object}  map[string]interface{} "ID de loja inválido"
// @Failure      500  {object}  map[string]interface{} "Erro interno do servidor"
// @Router       /lojas/{id}/historicos-resgate [get]
func GetHistoricosResgateByLojaIDHandler(c *gin.Context) {
	idLojaStr := c.Param("id")
	idLoja, err := strconv.ParseUint(idLojaStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID de loja inválido",
		})
		return
	}

	resp, err := services.GetHistoricosResgateByLojaID(uint(idLoja))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// UpdateHistoricoResgateHandler godoc
// @Summary      Atualiza histórico de resgate
// @Description  Atualiza os dados de um histórico de resgate existente
// @Tags         Histórico de Resgates
// @Accept       json
// @Produce      json
// @Param        id path int true "ID do histórico de resgate"
// @Param        historico body json.HistoricoResgateRequest true "Dados atualizados do histórico de resgate"
// @Success      200  {object}  json.HistoricoResgateResponse
// @Failure      400  {object}  map[string]interface{} "Dados inválidos"
// @Failure      404  {object}  map[string]interface{} "Histórico não encontrado"
// @Failure      500  {object}  map[string]interface{} "Erro interno do servidor"
// @Router       /historicos-resgate/{id} [put]
func UpdateHistoricoResgateHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID inválido",
		})
		return
	}

	var req json.HistoricoResgateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Dados inválidos: " + err.Error(),
		})
		return
	}

	resp, err := services.UpdateHistoricoResgate(uint(id), req)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Histórico não encontrado",
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// UpdateStatusHistoricoResgateHandler godoc
// @Summary      Atualiza status do histórico de resgate
// @Description  Atualiza apenas o status de um histórico de resgate
// @Tags         Histórico de Resgates
// @Accept       json
// @Produce      json
// @Param        id path int true "ID do histórico de resgate"
// @Param        status body map[string]string true "Novo status"
// @Success      200  {object}  map[string]interface{} "Status atualizado com sucesso"
// @Failure      400  {object}  map[string]interface{} "Dados inválidos"
// @Failure      404  {object}  map[string]interface{} "Histórico não encontrado"
// @Failure      500  {object}  map[string]interface{} "Erro interno do servidor"
// @Router       /historicos-resgate/{id}/status [put]
func UpdateStatusHistoricoResgateHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID inválido",
		})
		return
	}

	var req map[string]string
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Dados inválidos: " + err.Error(),
		})
		return
	}

	status, exists := req["status"]
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Campo 'status' é obrigatório",
		})
		return
	}

	// Valida o status
	validStatuses := map[string]bool{
		"pendente":   true,
		"confirmado": true,
		"cancelado":  true,
	}
	if !validStatuses[status] {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Status inválido. Valores aceitos: pendente, confirmado, cancelado",
		})
		return
	}

	err = services.UpdateStatusHistoricoResgate(uint(id), status)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Histórico não encontrado",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Status atualizado com sucesso",
		"status":   status,
	})
}

// AprovarResgateHandler godoc
// @Summary      Aprova um resgate
// @Description  Aprova um resgate pendente, alterando o status para confirmado
// @Tags         Histórico de Resgates
// @Accept       json
// @Produce      json
// @Param        id path int true "ID do histórico de resgate"
// @Success      200  {object}  json.HistoricoResgateResponse "Resgate aprovado com sucesso"
// @Failure      400  {object}  map[string]interface{} "Resgate não está pendente ou dados inválidos"
// @Failure      404  {object}  map[string]interface{} "Histórico não encontrado"
// @Failure      500  {object}  map[string]interface{} "Erro interno do servidor"
// @Router       /historicos-resgate/{id}/aprovar [put]
func AprovarResgateHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID inválido",
		})
		return
	}

	// Verifica se o resgate existe e está pendente
	historico, err := services.GetHistoricoResgateByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Histórico não encontrado",
		})
		return
	}

	if historico.Status != "pendente" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Apenas resgates com status 'pendente' podem ser aprovados",
		})
		return
	}

	err = services.UpdateStatusHistoricoResgate(uint(id), "confirmado")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Retorna o histórico atualizado
	historicoAtualizado, _ := services.GetHistoricoResgateByID(uint(id))
	
	// Registra log da aprovação
	idHistorico := uint(id)
	LogAction(c, "aprovar", "historico_resgate", &idHistorico,
		"Resgate aprovado pela loja", historico, historicoAtualizado)

	c.JSON(http.StatusOK, historicoAtualizado)
}

// RejeitarResgateHandler godoc
// @Summary      Rejeita um resgate
// @Description  Rejeita um resgate pendente, alterando o status para cancelado
// @Tags         Histórico de Resgates
// @Accept       json
// @Produce      json
// @Param        id path int true "ID do histórico de resgate"
// @Success      200  {object}  json.HistoricoResgateResponse "Resgate rejeitado com sucesso"
// @Failure      400  {object}  map[string]interface{} "Resgate não está pendente ou dados inválidos"
// @Failure      404  {object}  map[string]interface{} "Histórico não encontrado"
// @Failure      500  {object}  map[string]interface{} "Erro interno do servidor"
// @Router       /historicos-resgate/{id}/rejeitar [put]
func RejeitarResgateHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID inválido",
		})
		return
	}

	// Verifica se o resgate existe e está pendente
	historico, err := services.GetHistoricoResgateByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Histórico não encontrado",
		})
		return
	}

	if historico.Status != "pendente" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Apenas resgates com status 'pendente' podem ser rejeitados",
		})
		return
	}

	err = services.UpdateStatusHistoricoResgate(uint(id), "cancelado")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Retorna o histórico atualizado
	historicoAtualizado, _ := services.GetHistoricoResgateByID(uint(id))
	
	// Registra log da rejeição
	idHistorico := uint(id)
	LogAction(c, "rejeitar", "historico_resgate", &idHistorico,
		"Resgate rejeitado pela loja", historico, historicoAtualizado)

	c.JSON(http.StatusOK, historicoAtualizado)
}

// SoftDeleteHistoricoResgateHandler godoc
// @Summary      Remove histórico de resgate (soft delete)
// @Description  Realiza soft delete de um histórico de resgate, marcando-o como excluído
// @Tags         Histórico de Resgates
// @Accept       json
// @Produce      json
// @Param        id path int true "ID do histórico de resgate"
// @Success      200  {object}  map[string]interface{} "Histórico removido com sucesso"
// @Failure      400  {object}  map[string]interface{} "ID inválido"
// @Failure      404  {object}  map[string]interface{} "Histórico não encontrado"
// @Failure      500  {object}  map[string]interface{} "Erro interno do servidor"
// @Router       /historicos-resgate/{id} [delete]
func SoftDeleteHistoricoResgateHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID inválido",
		})
		return
	}

	err = services.SoftDeleteHistoricoResgate(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Histórico não encontrado",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Histórico removido com sucesso",
	})
}

// RestoreHistoricoResgateHandler godoc
// @Summary      Restaura histórico de resgate
// @Description  Restaura um histórico de resgate que foi soft deleted
// @Tags         Histórico de Resgates
// @Accept       json
// @Produce      json
// @Param        id path int true "ID do histórico de resgate"
// @Success      200  {object}  map[string]interface{} "Histórico restaurado com sucesso"
// @Failure      400  {object}  map[string]interface{} "ID inválido"
// @Failure      404  {object}  map[string]interface{} "Histórico não encontrado"
// @Failure      500  {object}  map[string]interface{} "Erro interno do servidor"
// @Router       /historicos-resgate/{id}/restore [post]
func RestoreHistoricoResgateHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID inválido",
		})
		return
	}

	err = services.RestoreHistoricoResgate(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Histórico não encontrado",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Histórico restaurado com sucesso",
	})
}

// GetHistoricosResgateClienteByUsuarioIDHandler godoc
// @Summary      Lista histórico de resgate simplificado do cliente
// @Description  Retorna histórico de resgate do cliente com campos simplificados (id, nome loja, imagem loja, data resgate, status, valor)
// @Tags         Histórico de Resgates
// @Accept       json
// @Produce      json
// @Param        id path int true "ID do usuário"
// @Success      200  {object}  json.HistoricosResgateClienteResponse
// @Failure      400  {object}  map[string]interface{} "ID de usuário inválido"
// @Failure      500  {object}  map[string]interface{} "Erro interno do servidor"
// @Router       /historicos-resgate/cliente/{id} [get]
func GetHistoricosResgateClienteByUsuarioIDHandler(c *gin.Context) {
	idUsuarioStr := c.Param("id")
	idUsuario, err := strconv.ParseUint(idUsuarioStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID de usuário inválido",
		})
		return
	}

	resp, err := services.GetHistoricosResgateClienteByUsuarioID(uint(idUsuario))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// GetHistoricosResgateByAnuncioIDHandler godoc
// @Summary      Lista histórico de resgate de um anúncio específico
// @Description  Retorna todos os históricos de resgate de um anúncio específico (utilização do anúncio)
// @Tags         Histórico de Resgates
// @Accept       json
// @Produce      json
// @Param        id path int true "ID do anúncio"
// @Success      200  {object}  json.HistoricosResgateResponse
// @Failure      400  {object}  map[string]interface{} "ID de anúncio inválido"
// @Failure      500  {object}  map[string]interface{} "Erro interno do servidor"
// @Router       /historicos-resgate/anuncio/{id} [get]
func GetHistoricosResgateByAnuncioIDHandler(c *gin.Context) {
	idAnuncioStr := c.Param("id")
	idAnuncio, err := strconv.ParseUint(idAnuncioStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID de anúncio inválido",
		})
		return
	}

	resp, err := services.GetHistoricosResgateByAnuncioID(uint(idAnuncio))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}