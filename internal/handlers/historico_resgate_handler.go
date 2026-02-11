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
// @Router       /users/{id}/historicos-resgate [get]
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

// GetHistoricosResgateByUsuarioIDDirectHandler godoc
// @Summary      Lista históricos de resgate de um usuário (endpoint direto)
// @Description  Retorna todos os históricos de resgate de um usuário específico através do endpoint direto
// @Tags         Histórico de Resgates
// @Accept       json
// @Produce      json
// @Param        id path int true "ID do usuário"
// @Success      200  {object}  json.HistoricosResgateResponse
// @Failure      400  {object}  map[string]interface{} "ID de usuário inválido"
// @Failure      500  {object}  map[string]interface{} "Erro interno do servidor"
// @Router       /historicos-resgate/usuario/{id} [get]
func GetHistoricosResgateByUsuarioIDDirectHandler(c *gin.Context) {
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
		"efetivado":  true,
	}
	if !validStatuses[status] {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Status inválido. Valores aceitos: pendente, efetivado",
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

// EfetivarResgateHandler godoc
// @Summary      Efetiva um resgate
// @Description  Efetiva um resgate pendente, alterando o status para efetivado. Se o resgate tiver um veículo do usuário vinculado (para produtos/serviços), o histórico é automaticamente registrado no veículo. Se for venda de veículo, o veículo é automaticamente transferido para o comprador.
// @Tags         Histórico de Resgates
// @Accept       json
// @Produce      json
// @Param        id path int true "ID do histórico de resgate"
// @Success      200  {object}  json.HistoricoResgateResponse "Resgate efetivado com sucesso e histórico do veículo criado (se aplicável)"
// @Failure      400  {object}  map[string]interface{} "Resgate não está pendente ou dados inválidos"
// @Failure      404  {object}  map[string]interface{} "Histórico não encontrado"
// @Failure      500  {object}  map[string]interface{} "Erro interno do servidor"
// @Router       /historicos-resgate/{id}/efetivar [put]
func EfetivarResgateHandler(c *gin.Context) {
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
			"error": "Apenas resgates com status 'pendente' podem ser efetivados",
		})
		return
	}

	err = services.UpdateStatusHistoricoResgate(uint(id), "efetivado")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Se é uma venda de veículo, transfere automaticamente o veículo para o comprador
	if historico.Cupom != nil && historico.Cupom.TipoCupom == "veiculo" && historico.Cupom.IDVeiculo != nil && historico.Cupom.Veiculo != nil {
		idHistoricoResgate := uint(id)
		_, errTransf := services.TransferirVeiculoVendaLoja(
			*historico.Cupom.IDVeiculo,
			historico.Cupom.Veiculo.IDUsuario, // Dono atual do veículo
			historico.IDUsuario,                // Comprador (usuário que fez o resgate)
			historico.Cupom.IDLoja,             // Loja que realizou a venda
			&idHistoricoResgate,                // Histórico de resgate que originou a transferência
		)
		if errTransf != nil {
			// Log do erro mas não falha a aprovação (o status já foi atualizado)
			LogAction(c, "erro", "transferencia_veiculo", nil,
				"Erro ao transferir veículo após aprovação de venda: "+errTransf.Error(), nil, nil)
		} else {
			LogAction(c, "transferir", "veiculo", historico.Cupom.IDVeiculo,
				"Veículo transferido automaticamente por venda em loja", nil, nil)
		}
	}

	// Se o cupom tem um veículo vinculado, cria o histórico do veículo
	if historico.Cupom != nil && historico.IDCupom != nil && historico.Cupom.IDVeiculo != nil {
		// Monta a descrição do histórico baseado no tipo de cupom
		var descricao string
		switch historico.Cupom.TipoCupom {
		case "produto":
			if historico.Cupom.Produto != nil {
				descricao = "Produto resgatado: " + historico.Cupom.Produto.Nome
			} else {
				descricao = "Produto resgatado"
			}
		case "servico":
			if historico.Cupom.Servico != nil {
				descricao = "Serviço resgatado: " + historico.Cupom.Servico.Titulo
			} else {
				descricao = "Serviço resgatado"
			}
		default:
			descricao = "Resgate efetivado"
		}

		// Cria o registro no histórico do veículo
		_, errHist := services.CreateHistoricoVeiculoFromResgate(
			*historico.Cupom.IDVeiculo,
			*historico.IDCupom,
			descricao,
		)
		if errHist != nil {
			// Log do erro mas não falha a aprovação (o registro foi criado com sucesso)
			LogAction(c, "erro", "historico_veiculo", nil,
				"Erro ao criar histórico do veículo após aprovação: "+errHist.Error(), nil, nil)
		}
	}

	// Retorna o histórico atualizado
	historicoAtualizado, _ := services.GetHistoricoResgateByID(uint(id))
	
	// Registra log da efetivação
	idHistorico := uint(id)
	LogAction(c, "efetivar", "historico_resgate", &idHistorico,
		"Resgate efetivado pela loja", historico, historicoAtualizado)

	c.JSON(http.StatusOK, historicoAtualizado)
}

// ReverterResgateHandler godoc
// @Summary      Reverte um resgate efetivado
// @Description  Reverte um resgate efetivado, alterando o status de volta para pendente
// @Tags         Histórico de Resgates
// @Accept       json
// @Produce      json
// @Param        id path int true "ID do histórico de resgate"
// @Success      200  {object}  json.HistoricoResgateResponse "Resgate revertido com sucesso"
// @Failure      400  {object}  map[string]interface{} "Resgate não está efetivado ou dados inválidos"
// @Failure      404  {object}  map[string]interface{} "Histórico não encontrado"
// @Failure      500  {object}  map[string]interface{} "Erro interno do servidor"
// @Router       /historicos-resgate/{id}/reverter [put]
func ReverterResgateHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID inválido",
		})
		return
	}

	// Verifica se o resgate existe e está efetivado
	historico, err := services.GetHistoricoResgateByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Histórico não encontrado",
		})
		return
	}

	if historico.Status != "efetivado" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Apenas resgates com status 'efetivado' podem ser revertidos",
		})
		return
	}

	err = services.UpdateStatusHistoricoResgate(uint(id), "pendente")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Retorna o histórico atualizado
	historicoAtualizado, _ := services.GetHistoricoResgateByID(uint(id))
	
	// Registra log da reversão
	idHistorico := uint(id)
	LogAction(c, "reverter", "historico_resgate", &idHistorico,
		"Resgate revertido pela loja", historico, historicoAtualizado)

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

// GetHistoricosResgateByCupomIDHandler godoc
// @Summary      Lista histórico de resgate de um cupom específico
// @Description  Retorna todos os históricos de resgate de um cupom específico (utilização do cupom)
// @Tags         Histórico de Resgates
// @Accept       json
// @Produce      json
// @Param        id path int true "ID do cupom"
// @Success      200  {object}  json.HistoricosResgateResponse
// @Failure      400  {object}  map[string]interface{} "ID de cupom inválido"
// @Failure      500  {object}  map[string]interface{} "Erro interno do servidor"
// @Router       /historicos-resgate/cupom/{id} [get]
func GetHistoricosResgateByCupomIDHandler(c *gin.Context) {
	idCupomStr := c.Param("id")
	idCupom, err := strconv.ParseUint(idCupomStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID de cupom inválido",
		})
		return
	}

	resp, err := services.GetHistoricosResgateByCupomID(uint(idCupom))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}