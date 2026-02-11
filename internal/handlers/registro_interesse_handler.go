package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"meu-carro-mais/internal/handlers/json"
	"meu-carro-mais/internal/services"
)

// CreateRegistroInteresseHandler godoc
// @Summary      Registra interesse em um veículo
// @Description  Cria um novo registro de interesse em um anúncio de veículo
// @Tags         Registro de Interesse
// @Accept       json
// @Produce      json
// @Param        request body json.RegistroInteresseRequest true "Dados do registro de interesse"
// @Success      201  {object}  json.RegistroInteresseResponse "Registro de interesse criado com sucesso"
// @Failure      400  {object}  map[string]interface{} "Dados inválidos"
// @Failure      500  {object}  map[string]interface{} "Erro interno do servidor"
// @Router       /registro-interesse [post]
func CreateRegistroInteresseHandler(c *gin.Context) {
	var req json.RegistroInteresseRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Dados inválidos",
			"details": err.Error(),
		})
		return
	}

	resp, err := services.CreateRegistroInteresse(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Registra log do registro de interesse
	idRegistro := uint(resp.ID)
	idCupom := uint(req.IDCupom)
	LogAction(c, "criar", "registro_interesse", &idRegistro,
		"Registro de interesse criado para cupom", nil, resp)
	LogAction(c, "registrar_interesse", "cupom", &idCupom,
		"Interesse registrado no cupom por "+resp.Nome, nil, resp)

	c.JSON(http.StatusCreated, resp)
}

// GetRegistroInteresseHandler godoc
// @Summary      Busca registro de interesse por ID
// @Description  Retorna os dados de um registro de interesse específico pelo ID
// @Tags         Registro de Interesse
// @Accept       json
// @Produce      json
// @Param        id path int true "ID do registro de interesse"
// @Success      200 {object} json.RegistroInteresseResponse "Registro de interesse encontrado"
// @Failure      400 {object} map[string]interface{} "ID inválido"
// @Failure      404 {object} map[string]interface{} "Registro de interesse não encontrado"
// @Failure      500 {object} map[string]interface{} "Erro interno do servidor"
// @Router       /registro-interesse/{id} [get]
func GetRegistroInteresseHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID inválido",
		})
		return
	}

	resp, err := services.GetRegistroInteresseByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Registro de interesse não encontrado",
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// GetAllRegistroInteressesHandler godoc
// @Summary      Lista todos os registros de interesse
// @Description  Retorna uma lista com todos os registros de interesse ativos
// @Tags         Registro de Interesse
// @Accept       json
// @Produce      json
// @Success      200 {array} json.RegistroInteresseResponse "Lista de registros de interesse"
// @Failure      500 {object} map[string]interface{} "Erro interno do servidor"
// @Router       /registro-interesse [get]
func GetAllRegistroInteressesHandler(c *gin.Context) {
	resp, err := services.GetAllRegistroInteresses()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// GetRegistroInteressesByCupomHandler godoc
// @Summary      Lista registros de interesse por cupom
// @Description  Retorna todos os registros de interesse de um cupom específico
// @Tags         Registro de Interesse
// @Accept       json
// @Produce      json
// @Param        cupom_id path int true "ID do cupom"
// @Success      200 {array} json.RegistroInteresseResponse "Lista de registros de interesse do cupom"
// @Failure      400 {object} map[string]interface{} "ID inválido"
// @Failure      500 {object} map[string]interface{} "Erro interno do servidor"
// @Router       /registro-interesse/cupom/{cupom_id} [get]
func GetRegistroInteressesByCupomHandler(c *gin.Context) {
	cupomIDStr := c.Param("cupom_id")
	cupomID, err := strconv.ParseUint(cupomIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID do cupom inválido",
		})
		return
	}

	resp, err := services.GetRegistroInteressesByCupomID(uint(cupomID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// SoftDeleteRegistroInteresseHandler godoc
// @Summary      Exclui registro de interesse (soft delete)
// @Description  Realiza soft delete do registro de interesse, marcando como excluído sem remover do banco
// @Tags         Registro de Interesse
// @Accept       json
// @Produce      json
// @Param        id path int true "ID do registro de interesse"
// @Success      200 {object} map[string]interface{} "Registro de interesse excluído com sucesso"
// @Failure      400 {object} map[string]interface{} "ID inválido"
// @Failure      404 {object} map[string]interface{} "Registro de interesse não encontrado"
// @Failure      500 {object} map[string]interface{} "Erro interno do servidor"
// @Router       /registro-interesse/{id} [delete]
func SoftDeleteRegistroInteresseHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID inválido",
		})
		return
	}

	err = services.SoftDeleteRegistroInteresse(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Registro de interesse não encontrado",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Registro de interesse excluído com sucesso",
	})
}

// RestoreRegistroInteresseHandler godoc
// @Summary      Restaura registro de interesse excluído
// @Description  Restaura um registro de interesse que foi soft deleted
// @Tags         Registro de Interesse
// @Accept       json
// @Produce      json
// @Param        id path int true "ID do registro de interesse"
// @Success      200 {object} map[string]interface{} "Registro de interesse restaurado com sucesso"
// @Failure      400 {object} map[string]interface{} "ID inválido"
// @Failure      404 {object} map[string]interface{} "Registro de interesse não encontrado"
// @Failure      500 {object} map[string]interface{} "Erro interno do servidor"
// @Router       /registro-interesse/{id}/restore [post]
func RestoreRegistroInteresseHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID inválido",
		})
		return
	}

	err = services.RestoreRegistroInteresse(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Registro de interesse não encontrado ou não foi excluído",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Registro de interesse restaurado com sucesso",
	})
}

