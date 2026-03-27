package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"meu-carro-mais/internal/handlers/json"
	"meu-carro-mais/internal/services"
)

// GetVeiculosByUsuarioHandler godoc
// @Summary      Lista veículos de um usuário
// @Description  Retorna todos os veículos ativos de um usuário específico, incluindo quilometragem (km), observações e imagem principal
// @Tags         Veículos
// @Accept       json
// @Produce      json
// @Param        id_usuario path int true "ID do usuário"
// @Success      200  {object}  json.VeiculosResponse "Lista de veículos (cada veículo inclui quilometragem, observacoes e imagem)"
// @Failure      400  {object}  map[string]interface{} "ID de usuário inválido"
// @Failure      500  {object}  map[string]interface{} "Erro interno do servidor"
// @Router       /usuarios/{id_usuario}/veiculos [get]
func GetVeiculosByUsuarioHandler(c *gin.Context) {
	idUsuarioStr := c.Param("id_usuario")
	idUsuario, err := strconv.ParseUint(idUsuarioStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID de usuário inválido",
		})
		return
	}

	resp, err := services.GetVeiculosByUsuario(uint(idUsuario))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// CreateVeiculoHandler godoc
// @Summary      Criação do veículo completo
// @Description  Cria um novo veículo com todos os dados fornecidos
// @Tags         Veículos
// @Accept       json
// @Produce      json
// @Param        request body json.VeiculoRequest true "Dados completos do veículo"
// @Success      201  {object}  json.VeiculoResponse "Veículo criado com sucesso"
// @Failure      400  {object}  map[string]interface{} "Dados inválidos"
// @Failure      500  {object}  map[string]interface{} "Erro interno do servidor"
// @Router       /veiculos [post]
func CreateVeiculoHandler(c *gin.Context) {
	var req json.VeiculoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Dados inválidos",
			"details": err.Error(),
		})
		return
	}

	resp, err := services.CreateVeiculo(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, resp)
}

// GetVeiculoHandler godoc
// @Summary      Busca veículo por ID
// @Description  Retorna os dados de um veículo específico pelo ID, incluindo quilometragem (km), observações e imagem principal
// @Tags         Veículos
// @Accept       json
// @Produce      json
// @Param        id path int true "ID do veículo"
// @Success      200 {object} json.VeiculoResponse "Veículo encontrado (inclui quilometragem, observacoes e imagem)"
// @Failure      400 {object} map[string]interface{} "ID inválido"
// @Failure      404 {object} map[string]interface{} "Veículo não encontrado"
// @Failure      500 {object} map[string]interface{} "Erro interno do servidor"
// @Router       /veiculos/{id} [get]
func GetVeiculoHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID inválido",
		})
		return
	}

	resp, err := services.GetVeiculoByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Veículo não encontrado",
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// GetAllVeiculosHandler godoc
// @Summary      Lista todos os veículos
// @Description  Retorna uma lista com todos os veículos ativos, incluindo quilometragem (km), observações e imagem principal
// @Tags         Veículos
// @Accept       json
// @Produce      json
// @Success      200 {array} json.VeiculoResponse "Lista de veículos (cada veículo inclui quilometragem, observacoes e imagem)"
// @Failure      500 {object} map[string]interface{} "Erro interno do servidor"
// @Router       /veiculos [get]
func GetAllVeiculosHandler(c *gin.Context) {
	resp, err := services.GetAllVeiculos()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// UpdateVeiculoHandler godoc
// @Summary      Atualiza veículo
// @Description  Atualiza os dados de um veículo existente
// @Tags         Veículos
// @Accept       json
// @Produce      json
// @Param        id path int true "ID do veículo"
// @Param        request body json.VeiculoRequest true "Dados atualizados do veículo"
// @Success      200 {object} json.VeiculoResponse "Veículo atualizado com sucesso"
// @Failure      400 {object} map[string]interface{} "Dados inválidos"
// @Failure      404 {object} map[string]interface{} "Veículo não encontrado"
// @Failure      500 {object} map[string]interface{} "Erro interno do servidor"
// @Router       /veiculos/{id} [put]
func UpdateVeiculoHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID inválido",
		})
		return
	}

	var req json.VeiculoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Dados inválidos",
			"details": err.Error(),
		})
		return
	}

	resp, err := services.UpdateVeiculo(uint(id), req)
	if err != nil {
		errMsg := err.Error()
		if errMsg == "veículo não encontrado" {
			c.JSON(http.StatusNotFound, gin.H{
				"error": errMsg,
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":   "Erro ao atualizar veículo",
				"details": errMsg,
			})
		}
		return
	}

	c.JSON(http.StatusOK, resp)
}

// SoftDeleteVeiculoHandler godoc
// @Summary      Exclui veículo (soft delete)
// @Description  Realiza soft delete do veículo, marcando como excluído sem remover do banco
// @Tags         Veículos
// @Accept       json
// @Produce      json
// @Param        id path int true "ID do veículo"
// @Success      200 {object} map[string]interface{} "Veículo excluído com sucesso"
// @Failure      400 {object} map[string]interface{} "ID inválido"
// @Failure      404 {object} map[string]interface{} "Veículo não encontrado"
// @Failure      500 {object} map[string]interface{} "Erro interno do servidor"
// @Router       /veiculos/{id} [delete]
func SoftDeleteVeiculoHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID inválido",
		})
		return
	}

	err = services.SoftDeleteVeiculo(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Veículo não encontrado",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Veículo excluído com sucesso",
	})
}

// RestoreVeiculoHandler godoc
// @Summary      Restaura veículo excluído
// @Description  Restaura um veículo que foi soft deleted
// @Tags         Veículos
// @Accept       json
// @Produce      json
// @Param        id path int true "ID do veículo"
// @Success      200 {object} map[string]interface{} "Veículo restaurado com sucesso"
// @Failure      400 {object} map[string]interface{} "ID inválido"
// @Failure      404 {object} map[string]interface{} "Veículo não encontrado"
// @Failure      500 {object} map[string]interface{} "Erro interno do servidor"
// @Router       /veiculos/{id}/restore [post]
func RestoreVeiculoHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID inválido",
		})
		return
	}

	err = services.RestoreVeiculo(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Veículo não encontrado ou não foi excluído",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Veículo restaurado com sucesso",
	})
}

// GetHistoricosByVeiculoHandler godoc
// @Summary      Lista histórico de um veículo
// @Description  Retorna o histórico do veículo: registros em `historico_veiculos` mais resgates já efetivados com `id_veiculo` igual ao informado na URL (sem duplicar quando já existe linha para o mesmo cupom). Ordenado por data (mais recente primeiro). O campo `id` de cada item pode ser o id da tabela de histórico do veículo ou, em registros derivados só do resgate, o id do histórico de resgate.
// @Tags         Veículos
// @Accept       json
// @Produce      json
// @Param        id path int true "ID do veículo"
// @Success      200  {object}  json.HistoricosVeiculoResponse
// @Failure      400  {object}  map[string]interface{} "ID de veículo inválido"
// @Failure      500  {object}  map[string]interface{} "Erro interno do servidor"
// @Router       /veiculos/{id}/historico [get]
func GetHistoricosByVeiculoHandler(c *gin.Context) {
	idVeiculoStr := c.Param("id")
	idVeiculo, err := strconv.ParseUint(idVeiculoStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID de veículo inválido",
		})
		return
	}

	resp, err := services.GetHistoricosByVeiculo(uint(idVeiculo))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// GetHistoricosByUsuarioHandler godoc
// @Summary      Lista histórico de todos os veículos de um usuário
// @Description  Retorna o histórico de todos os veículos de um usuário específico
// @Tags         Veículos
// @Accept       json
// @Produce      json
// @Param        id_usuario path int true "ID do usuário"
// @Success      200  {object}  json.HistoricosVeiculoResponse
// @Failure      400  {object}  map[string]interface{} "ID de usuário inválido"
// @Failure      500  {object}  map[string]interface{} "Erro interno do servidor"
// @Router       /usuarios/{id_usuario}/historico [get]
func GetHistoricosByUsuarioHandler(c *gin.Context) {
	idUsuarioStr := c.Param("id_usuario")
	idUsuario, err := strconv.ParseUint(idUsuarioStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID de usuário inválido",
		})
		return
	}

	resp, err := services.GetHistoricosByUsuario(uint(idUsuario))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}
