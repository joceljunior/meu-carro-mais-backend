package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"meu-carro-mais/internal/handlers/json"
	"meu-carro-mais/internal/services"
)

// GetLojasByProximidadeHandler godoc
// @Summary      Lista lojas por proximidade
// @Description  Retorna lista de lojas ordenadas por proximidade do usuário
// @Tags         Lojas
// @Accept       json
// @Produce      json
// @Param        latitude  query     number  true  "Latitude do usuário"
// @Param        longitude query     number  true  "Longitude do usuário"
// @Success      200       {object}  json.LojasResponse
// @Failure      400       {object}  map[string]interface{} "Parâmetros inválidos"
// @Failure      500       {object}  map[string]interface{} "Erro interno do servidor"
// @Router       /lojas/proximidade [get]
func GetLojasByProximidadeHandler(c *gin.Context) {
	// Obtém os parâmetros da query string
	latStr := c.Query("latitude")
	lngStr := c.Query("longitude")

	if latStr == "" || lngStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Parâmetros latitude e longitude são obrigatórios",
		})
		return
	}

	// Converte para float64
	var latitude, longitude float64
	if _, err := fmt.Sscanf(latStr, "%f", &latitude); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Latitude deve ser um número válido",
		})
		return
	}

	if _, err := fmt.Sscanf(lngStr, "%f", &longitude); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Longitude deve ser um número válido",
		})
		return
	}

	// Valida as coordenadas
	if latitude < -90 || latitude > 90 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Latitude deve estar entre -90 e 90",
		})
		return
	}

	if longitude < -180 || longitude > 180 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Longitude deve estar entre -180 e 180",
		})
		return
	}

	// Busca as lojas
	resp, err := services.GetLojasByProximidade(latitude, longitude)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Erro ao buscar lojas: " + err.Error(),
		})
		return
	}

	// Retorna a resposta
	c.JSON(http.StatusOK, resp)
}

// CreateLojaHandler godoc
// @Summary      Criação da loja completa
// @Description  Cria uma nova loja com todos os dados fornecidos, incluindo rating e status premium
// @Tags         Lojas
// @Accept       json
// @Produce      json
// @Param        request body json.LojaRequest true "Dados completos da loja (rating e is_meu_carro_mais são opcionais); desconto_geral_porcentagem é obrigatório (0–100)"
// @Success      201  {object}  json.LojaResponse "Loja criada com sucesso"
// @Failure      400  {object}  map[string]interface{} "Dados inválidos"
// @Failure      500  {object}  map[string]interface{} "Erro interno do servidor"
// @Router       /lojas [post]
func CreateLojaHandler(c *gin.Context) {
	var req json.LojaRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Dados inválidos",
			"details": err.Error(),
		})
		return
	}

	resp, err := services.CreateLoja(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, resp)
}

// GetLojaHandler godoc
// @Summary      Busca loja por ID
// @Description  Retorna os dados de uma loja específica pelo ID
// @Tags         Lojas
// @Accept       json
// @Produce      json
// @Param        id path int true "ID da loja"
// @Success      200 {object} json.LojaResponse "Loja encontrada"
// @Failure      400 {object} map[string]interface{} "ID inválido"
// @Failure      404 {object} map[string]interface{} "Loja não encontrada"
// @Failure      500 {object} map[string]interface{} "Erro interno do servidor"
// @Router       /lojas/{id} [get]
func GetLojaHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID inválido",
		})
		return
	}

	resp, err := services.GetLojaByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Loja não encontrada",
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// GetAllLojasHandler godoc
// @Summary      Lista todas as lojas
// @Description  Retorna uma lista com todas as lojas ativas
// @Tags         Lojas
// @Accept       json
// @Produce      json
// @Success      200 {array} json.LojaResponse "Lista de lojas"
// @Failure      500 {object} map[string]interface{} "Erro interno do servidor"
// @Router       /lojas [get]
func GetAllLojasHandler(c *gin.Context) {
	resp, err := services.GetAllLojas()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// UpdateLojaHandler godoc
// @Summary      Atualiza loja
// @Description  Atualiza os dados de uma loja existente, incluindo rating e status premium
// @Tags         Lojas
// @Accept       json
// @Produce      json
// @Param        id path int true "ID da loja"
// @Param        request body json.LojaRequest true "Dados atualizados da loja (rating e is_meu_carro_mais são opcionais); desconto_geral_porcentagem é obrigatório (0–100)"
// @Success      200 {object} json.LojaResponse "Loja atualizada com sucesso"
// @Failure      400 {object} map[string]interface{} "Dados inválidos"
// @Failure      404 {object} map[string]interface{} "Loja não encontrada"
// @Failure      500 {object} map[string]interface{} "Erro interno do servidor"
// @Router       /lojas/{id} [put]
func UpdateLojaHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID inválido",
		})
		return
	}

	var req json.LojaRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Dados inválidos",
			"details": err.Error(),
		})
		return
	}

	resp, err := services.UpdateLoja(uint(id), req)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Loja não encontrada",
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// SoftDeleteLojaHandler godoc
// @Summary      Exclui loja (soft delete)
// @Description  Realiza soft delete da loja, marcando como excluída sem remover do banco
// @Tags         Lojas
// @Accept       json
// @Produce      json
// @Param        id path int true "ID da loja"
// @Success      200 {object} map[string]interface{} "Loja excluída com sucesso"
// @Failure      400 {object} map[string]interface{} "ID inválido"
// @Failure      404 {object} map[string]interface{} "Loja não encontrada"
// @Failure      500 {object} map[string]interface{} "Erro interno do servidor"
// @Router       /lojas/{id} [delete]
func SoftDeleteLojaHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID inválido",
		})
		return
	}

	err = services.SoftDeleteLoja(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Loja não encontrada",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Loja excluída com sucesso",
	})
}

// RestoreLojaHandler godoc
// @Summary      Restaura loja excluída
// @Description  Restaura uma loja que foi soft deleted
// @Tags         Lojas
// @Accept       json
// @Produce      json
// @Param        id path int true "ID da loja"
// @Success      200 {object} map[string]interface{} "Loja restaurada com sucesso"
// @Failure      400 {object} map[string]interface{} "ID inválido"
// @Failure      404 {object} map[string]interface{} "Loja não encontrada"
// @Failure      500 {object} map[string]interface{} "Erro interno do servidor"
// @Router       /lojas/{id}/restore [post]
func RestoreLojaHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID inválido",
		})
		return
	}

	err = services.RestoreLoja(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Loja não encontrada ou não foi excluída",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Loja restaurada com sucesso",
	})
}

// GetCategoriasLojistaHandler godoc
// @Summary      Lista categorias de lojista
// @Description  Retorna todas as categorias de lojista disponíveis
// @Tags         Lojas
// @Accept       json
// @Produce      json
// @Success      200  {object}  json.CategoriasLojistaResponse
// @Failure      500  {object}  map[string]interface{} "Erro interno do servidor"
// @Router       /lojas/categorias [get]
func GetCategoriasLojistaHandler(c *gin.Context) {
	resp, err := services.GetCategoriasLojista()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// GetLojasByUsuarioIDHandler godoc
// @Summary      Lista lojas por usuário
// @Description  Retorna todas as lojas de um usuário específico
// @Tags         Lojas
// @Accept       json
// @Produce      json
// @Param        id_usuario path int true "ID do usuário"
// @Success      200 {array} json.LojaResponse "Lista de lojas do usuário"
// @Failure      400 {object} map[string]interface{} "ID inválido"
// @Failure      500 {object} map[string]interface{} "Erro interno do servidor"
// @Router       /lojas/usuario/{id_usuario} [get]
func GetLojasByUsuarioIDHandler(c *gin.Context) {
	idStr := c.Param("id_usuario")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID do usuário inválido",
		})
		return
	}

	resp, err := services.GetLojasByUsuarioID(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}