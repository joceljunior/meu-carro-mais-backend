package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"meu-carro-mais/internal/handlers/json"
	"meu-carro-mais/internal/services"
)

// GetServicosByProximidadeHandler godoc
// @Summary      Lista serviços por proximidade
// @Description  Retorna lista de serviços ordenados por proximidade do usuário
// @Tags         Serviços
// @Accept       json
// @Produce      json
// @Param        latitude  query     number  true  "Latitude do usuário"
// @Param        longitude query     number  true  "Longitude do usuário"
// @Success      200       {object}  json.ServicosResponse
// @Failure      400       {string}  string "Parâmetros inválidos"
// @Failure      500       {string}  string "Erro interno do servidor"
// @Router       /servicos/proximidade [get]
func GetServicosByProximidadeHandler(c *gin.Context) {
	latitudeStr := c.Query("latitude")
	longitudeStr := c.Query("longitude")

	if latitudeStr == "" || longitudeStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Latitude e longitude são obrigatórios",
		})
		return
	}

	var latitude, longitude float64
	if _, err := fmt.Sscanf(latitudeStr, "%f", &latitude); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Latitude deve ser um número válido",
		})
		return
	}

	if _, err := fmt.Sscanf(longitudeStr, "%f", &longitude); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Longitude deve ser um número válido",
		})
		return
	}

	resp, err := services.GetServicosByProximidade(latitude, longitude)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// CreateServicoHandler godoc
// @Summary      Criação do serviço completo
// @Description  Cria um novo serviço com todos os dados fornecidos
// @Tags         Serviços
// @Accept       json
// @Produce      json
// @Param        request body json.ServicoRequest true "Dados completos do serviço"
// @Success      201  {object}  json.ServicoResponse "Serviço criado com sucesso"
// @Failure      400  {object}  map[string]interface{} "Dados inválidos"
// @Failure      500  {object}  map[string]interface{} "Erro interno do servidor"
// @Router       /servicos [post]
func CreateServicoHandler(c *gin.Context) {
	var req json.ServicoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Dados inválidos",
			"details": err.Error(),
		})
		return
	}

	resp, err := services.CreateServico(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, resp)
}

// GetServicoHandler godoc
// @Summary      Busca serviço por ID
// @Description  Retorna os dados de um serviço específico pelo ID
// @Tags         Serviços
// @Accept       json
// @Produce      json
// @Param        id path int true "ID do serviço"
// @Success      200 {object} json.ServicoResponse "Serviço encontrado"
// @Failure      400 {object} map[string]interface{} "ID inválido"
// @Failure      404 {object} map[string]interface{} "Serviço não encontrado"
// @Failure      500 {object} map[string]interface{} "Erro interno do servidor"
// @Router       /servicos/{id} [get]
func GetServicoHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID inválido",
		})
		return
	}

	resp, err := services.GetServicoByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Serviço não encontrado",
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// GetAllServicosHandler godoc
// @Summary      Lista todos os serviços
// @Description  Retorna uma lista com todos os serviços ativos
// @Tags         Serviços
// @Accept       json
// @Produce      json
// @Success      200 {array} json.ServicoResponse "Lista de serviços"
// @Failure      500 {object} map[string]interface{} "Erro interno do servidor"
// @Router       /servicos [get]
func GetAllServicosHandler(c *gin.Context) {
	resp, err := services.GetAllServicos()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// UpdateServicoHandler godoc
// @Summary      Atualiza serviço
// @Description  Atualiza os dados de um serviço existente
// @Tags         Serviços
// @Accept       json
// @Produce      json
// @Param        id path int true "ID do serviço"
// @Param        request body json.ServicoRequest true "Dados atualizados do serviço"
// @Success      200 {object} json.ServicoResponse "Serviço atualizado com sucesso"
// @Failure      400 {object} map[string]interface{} "Dados inválidos"
// @Failure      404 {object} map[string]interface{} "Serviço não encontrado"
// @Failure      500 {object} map[string]interface{} "Erro interno do servidor"
// @Router       /servicos/{id} [put]
func UpdateServicoHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID inválido",
		})
		return
	}

	var req json.ServicoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Dados inválidos",
			"details": err.Error(),
		})
		return
	}

	resp, err := services.UpdateServico(uint(id), req)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Serviço não encontrado",
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// SoftDeleteServicoHandler godoc
// @Summary      Exclui serviço (soft delete)
// @Description  Realiza soft delete do serviço, marcando como excluído sem remover do banco
// @Tags         Serviços
// @Accept       json
// @Produce      json
// @Param        id path int true "ID do serviço"
// @Success      200 {object} map[string]interface{} "Serviço excluído com sucesso"
// @Failure      400 {object} map[string]interface{} "ID inválido"
// @Failure      404 {object} map[string]interface{} "Serviço não encontrado"
// @Failure      500 {object} map[string]interface{} "Erro interno do servidor"
// @Router       /servicos/{id} [delete]
func SoftDeleteServicoHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID inválido",
		})
		return
	}

	err = services.SoftDeleteServico(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Serviço não encontrado",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Serviço excluído com sucesso",
	})
}

// RestoreServicoHandler godoc
// @Summary      Restaura serviço excluído
// @Description  Restaura um serviço que foi soft deleted
// @Tags         Serviços
// @Accept       json
// @Produce      json
// @Param        id path int true "ID do serviço"
// @Success      200 {object} map[string]interface{} "Serviço restaurado com sucesso"
// @Failure      400 {object} map[string]interface{} "ID inválido"
// @Failure      404 {object} map[string]interface{} "Serviço não encontrado"
// @Failure      500 {object} map[string]interface{} "Erro interno do servidor"
// @Router       /servicos/{id}/restore [post]
func RestoreServicoHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID inválido",
		})
		return
	}

	err = services.RestoreServico(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Serviço não encontrado ou não foi excluído",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Serviço restaurado com sucesso",
	})
}

// GetServicosByLojaIDHandler godoc
// @Summary      Lista serviços de uma loja
// @Description  Retorna todos os serviços de uma loja específica
// @Tags         Serviços
// @Accept       json
// @Produce      json
// @Param        id path int true "ID da loja"
// @Success      200 {object} json.ServicosResponse "Lista de serviços da loja"
// @Failure      400 {object} map[string]interface{} "ID de loja inválido"
// @Failure      500 {object} map[string]interface{} "Erro interno do servidor"
// @Router       /lojas/{id}/servicos [get]
func GetServicosByLojaIDHandler(c *gin.Context) {
	idLojaStr := c.Param("id")
	idLoja, err := strconv.ParseUint(idLojaStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID de loja inválido",
		})
		return
	}

	resp, err := services.GetServicosByLojaID(uint(idLoja))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}
