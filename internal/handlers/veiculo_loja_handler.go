package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"meu-carro-mais/internal/handlers/json"
	"meu-carro-mais/internal/services"
)

// CreateVeiculoLojaHandler godoc
// @Summary      Criação do veículo de loja completo
// @Description  Cria um novo veículo de loja com todos os dados fornecidos
// @Tags         Veículos de Loja
// @Accept       json
// @Produce      json
// @Param        request body json.VeiculoLojaRequest true "Dados completos do veículo de loja"
// @Success      201  {object}  json.VeiculoLojaResponse "Veículo de loja criado com sucesso"
// @Failure      400  {object}  map[string]interface{} "Dados inválidos"
// @Failure      500  {object}  map[string]interface{} "Erro interno do servidor"
// @Router       /veiculos-loja [post]
func CreateVeiculoLojaHandler(c *gin.Context) {
	var req json.VeiculoLojaRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Dados inválidos",
			"details": err.Error(),
		})
		return
	}

	resp, err := services.CreateVeiculoLoja(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, resp)
}

// GetVeiculoLojaHandler godoc
// @Summary      Busca veículo de loja por ID
// @Description  Retorna os dados de um veículo de loja específico pelo ID
// @Tags         Veículos de Loja
// @Accept       json
// @Produce      json
// @Param        id path int true "ID do veículo de loja"
// @Success      200 {object} json.VeiculoLojaResponse "Veículo de loja encontrado"
// @Failure      400 {object} map[string]interface{} "ID inválido"
// @Failure      404 {object} map[string]interface{} "Veículo de loja não encontrado"
// @Failure      500 {object} map[string]interface{} "Erro interno do servidor"
// @Router       /veiculos-loja/{id} [get]
func GetVeiculoLojaHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID inválido",
		})
		return
	}

	resp, err := services.GetVeiculoLojaByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Veículo de loja não encontrado",
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// GetAllVeiculosLojaHandler godoc
// @Summary      Lista todos os veículos de loja
// @Description  Retorna uma lista com todos os veículos de loja ativos
// @Tags         Veículos de Loja
// @Accept       json
// @Produce      json
// @Success      200 {array} json.VeiculoLojaResponse "Lista de veículos de loja"
// @Failure      500 {object} map[string]interface{} "Erro interno do servidor"
// @Router       /veiculos-loja [get]
func GetAllVeiculosLojaHandler(c *gin.Context) {
	resp, err := services.GetAllVeiculosLoja()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// GetVeiculosLojaByLojaIDHandler godoc
// @Summary      Lista veículos de uma loja
// @Description  Retorna todos os veículos ativos de uma loja específica
// @Tags         Veículos de Loja
// @Accept       json
// @Produce      json
// @Param        id path int true "ID da loja"
// @Success      200  {object}  json.VeiculosLojaResponse
// @Failure      400  {object}  map[string]interface{} "ID de loja inválido"
// @Failure      500  {object}  map[string]interface{} "Erro interno do servidor"
// @Router       /lojas/{id}/veiculos [get]
func GetVeiculosLojaByLojaIDHandler(c *gin.Context) {
	idLojaStr := c.Param("id")
	idLoja, err := strconv.ParseUint(idLojaStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID de loja inválido",
		})
		return
	}

	resp, err := services.GetVeiculosLojaByLojaID(uint(idLoja))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// UpdateVeiculoLojaHandler godoc
// @Summary      Atualiza veículo de loja
// @Description  Atualiza os dados de um veículo de loja existente
// @Tags         Veículos de Loja
// @Accept       json
// @Produce      json
// @Param        id path int true "ID do veículo de loja"
// @Param        request body json.VeiculoLojaRequest true "Dados atualizados do veículo de loja"
// @Success      200 {object} json.VeiculoLojaResponse "Veículo de loja atualizado com sucesso"
// @Failure      400 {object} map[string]interface{} "Dados inválidos"
// @Failure      404 {object} map[string]interface{} "Veículo de loja não encontrado"
// @Failure      500 {object} map[string]interface{} "Erro interno do servidor"
// @Router       /veiculos-loja/{id} [put]
func UpdateVeiculoLojaHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID inválido",
		})
		return
	}

	var req json.VeiculoLojaRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Dados inválidos",
			"details": err.Error(),
		})
		return
	}

	resp, err := services.UpdateVeiculoLoja(uint(id), req)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Veículo de loja não encontrado",
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// SoftDeleteVeiculoLojaHandler godoc
// @Summary      Exclui veículo de loja (soft delete)
// @Description  Realiza soft delete do veículo de loja, marcando como excluído sem remover do banco
// @Tags         Veículos de Loja
// @Accept       json
// @Produce      json
// @Param        id path int true "ID do veículo de loja"
// @Success      200 {object} map[string]interface{} "Veículo de loja excluído com sucesso"
// @Failure      400 {object} map[string]interface{} "ID inválido"
// @Failure      404 {object} map[string]interface{} "Veículo de loja não encontrado"
// @Failure      500 {object} map[string]interface{} "Erro interno do servidor"
// @Router       /veiculos-loja/{id} [delete]
func SoftDeleteVeiculoLojaHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID inválido",
		})
		return
	}

	err = services.SoftDeleteVeiculoLoja(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Veículo de loja não encontrado",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Veículo de loja excluído com sucesso",
	})
}

// RestoreVeiculoLojaHandler godoc
// @Summary      Restaura veículo de loja excluído
// @Description  Restaura um veículo de loja que foi soft deleted
// @Tags         Veículos de Loja
// @Accept       json
// @Produce      json
// @Param        id path int true "ID do veículo de loja"
// @Success      200 {object} map[string]interface{} "Veículo de loja restaurado com sucesso"
// @Failure      400 {object} map[string]interface{} "ID inválido"
// @Failure      404 {object} map[string]interface{} "Veículo de loja não encontrado"
// @Failure      500 {object} map[string]interface{} "Erro interno do servidor"
// @Router       /veiculos-loja/{id}/restore [post]
func RestoreVeiculoLojaHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID inválido",
		})
		return
	}

	err = services.RestoreVeiculoLoja(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Veículo de loja não encontrado ou não foi excluído",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Veículo de loja restaurado com sucesso",
	})
}
