package handlers

import (
	"net/http"
	"strconv"

	"meu-carro-mais/internal/handlers/json"
	"meu-carro-mais/internal/services"

	"github.com/gin-gonic/gin"
)

// CreateDescontoHandler godoc
// @Summary      Cria um novo desconto para uma loja
// @Description  Cria um novo desconto para uma loja. Uma loja só pode ter um desconto ativo por vez.
// @Tags         Descontos
// @Accept       json
// @Produce      json
// @Param        desconto body json.DescontoRequest true "Dados do desconto"
// @Success      201  {object}  json.DescontoResponse
// @Failure      400  {object}  map[string]interface{} "Dados inválidos ou loja já possui desconto ativo"
// @Failure      500  {object}  map[string]interface{} "Erro interno do servidor"
// @Router       /descontos [post]
func CreateDescontoHandler(c *gin.Context) {
	var req json.DescontoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Dados inválidos: " + err.Error(),
		})
		return
	}

	resp, err := services.CreateDesconto(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, resp)
}

// GetDescontoHandler godoc
// @Summary      Busca desconto por ID
// @Description  Retorna um desconto específico pelo ID
// @Tags         Descontos
// @Accept       json
// @Produce      json
// @Param        id path int true "ID do desconto"
// @Success      200  {object}  json.DescontoResponse
// @Failure      400  {object}  map[string]interface{} "ID inválido"
// @Failure      404  {object}  map[string]interface{} "Desconto não encontrado"
// @Router       /descontos/{id} [get]
func GetDescontoHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID inválido",
		})
		return
	}

	resp, err := services.GetDescontoByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Desconto não encontrado",
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// GetAllDescontosHandler godoc
// @Summary      Lista todos os descontos
// @Description  Retorna todos os descontos do sistema
// @Tags         Descontos
// @Accept       json
// @Produce      json
// @Success      200  {object}  json.DescontosResponse
// @Failure      500  {object}  map[string]interface{} "Erro interno do servidor"
// @Router       /descontos [get]
func GetAllDescontosHandler(c *gin.Context) {
	resp, err := services.GetAllDescontos()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	response := json.DescontosResponse{
		Descontos: resp,
		Total:     len(resp),
	}

	c.JSON(http.StatusOK, response)
}

// GetDescontosAtivosHandler godoc
// @Summary      Lista todos os descontos ativos
// @Description  Retorna todos os descontos ativos e não expirados do sistema
// @Tags         Descontos
// @Accept       json
// @Produce      json
// @Success      200  {object}  json.DescontosResponse
// @Failure      500  {object}  map[string]interface{} "Erro interno do servidor"
// @Router       /descontos/ativos [get]
func GetDescontosAtivosHandler(c *gin.Context) {
	resp, err := services.GetAllDescontosAtivos()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	response := json.DescontosResponse{
		Descontos: resp,
		Total:     len(resp),
	}

	c.JSON(http.StatusOK, response)
}

// GetDescontosByLojaIDHandler godoc
// @Summary      Lista histórico de descontos de uma loja
// @Description  Retorna todos os descontos (ativos e inativos) de uma loja específica
// @Tags         Descontos
// @Accept       json
// @Produce      json
// @Param        id path int true "ID da loja"
// @Success      200  {object}  json.DescontosResponse
// @Failure      400  {object}  map[string]interface{} "ID de loja inválido"
// @Failure      500  {object}  map[string]interface{} "Erro interno do servidor"
// @Router       /lojas/{id}/descontos [get]
func GetDescontosByLojaIDHandler(c *gin.Context) {
	idLojaStr := c.Param("id")
	idLoja, err := strconv.ParseUint(idLojaStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID de loja inválido",
		})
		return
	}

	resp, err := services.GetDescontosByLojaID(uint(idLoja))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// GetDescontoAtivoByLojaIDHandler godoc
// @Summary      Busca desconto ativo de uma loja
// @Description  Retorna o desconto ativo atual de uma loja específica
// @Tags         Descontos
// @Accept       json
// @Produce      json
// @Param        id path int true "ID da loja"
// @Success      200  {object}  json.DescontoResponse
// @Failure      400  {object}  map[string]interface{} "ID de loja inválido"
// @Failure      404  {object}  map[string]interface{} "Nenhum desconto ativo encontrado"
// @Router       /lojas/{id}/desconto-ativo [get]
func GetDescontoAtivoByLojaIDHandler(c *gin.Context) {
	idLojaStr := c.Param("id")
	idLoja, err := strconv.ParseUint(idLojaStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID de loja inválido",
		})
		return
	}

	resp, err := services.GetDescontoAtivoByLojaID(uint(idLoja))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Nenhum desconto ativo encontrado para esta loja",
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// CancelarDescontoHandler godoc
// @Summary      Cancela um desconto
// @Description  Desativa um desconto ativo pelo ID
// @Tags         Descontos
// @Accept       json
// @Produce      json
// @Param        id path int true "ID do desconto"
// @Success      200  {object}  map[string]interface{} "Desconto cancelado com sucesso"
// @Failure      400  {object}  map[string]interface{} "ID inválido ou desconto já inativo"
// @Failure      404  {object}  map[string]interface{} "Desconto não encontrado"
// @Router       /descontos/{id}/cancelar [post]
func CancelarDescontoHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID inválido",
		})
		return
	}

	err = services.CancelarDesconto(uint(id))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Desconto cancelado com sucesso",
	})
}

// CancelarDescontoLojaHandler godoc
// @Summary      Cancela o desconto ativo de uma loja
// @Description  Desativa o desconto ativo atual de uma loja específica
// @Tags         Descontos
// @Accept       json
// @Produce      json
// @Param        id path int true "ID da loja"
// @Success      200  {object}  map[string]interface{} "Desconto cancelado com sucesso"
// @Failure      400  {object}  map[string]interface{} "ID de loja inválido"
// @Failure      404  {object}  map[string]interface{} "Nenhum desconto ativo encontrado"
// @Router       /lojas/{id}/desconto-ativo/cancelar [post]
func CancelarDescontoLojaHandler(c *gin.Context) {
	idLojaStr := c.Param("id")
	idLoja, err := strconv.ParseUint(idLojaStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID de loja inválido",
		})
		return
	}

	err = services.CancelarDescontoAtivoByLojaID(uint(idLoja))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Desconto cancelado com sucesso",
	})
}

// SoftDeleteDescontoHandler godoc
// @Summary      Remove desconto (soft delete)
// @Description  Realiza soft delete de um desconto, marcando-o como excluído
// @Tags         Descontos
// @Accept       json
// @Produce      json
// @Param        id path int true "ID do desconto"
// @Success      200  {object}  map[string]interface{} "Desconto removido com sucesso"
// @Failure      400  {object}  map[string]interface{} "ID inválido"
// @Failure      404  {object}  map[string]interface{} "Desconto não encontrado"
// @Router       /descontos/{id} [delete]
func SoftDeleteDescontoHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID inválido",
		})
		return
	}

	err = services.SoftDeleteDesconto(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Desconto não encontrado",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Desconto removido com sucesso",
	})
}

// RestoreDescontoHandler godoc
// @Summary      Restaura desconto
// @Description  Restaura um desconto que foi soft deleted
// @Tags         Descontos
// @Accept       json
// @Produce      json
// @Param        id path int true "ID do desconto"
// @Success      200  {object}  map[string]interface{} "Desconto restaurado com sucesso"
// @Failure      400  {object}  map[string]interface{} "ID inválido"
// @Failure      404  {object}  map[string]interface{} "Desconto não encontrado"
// @Router       /descontos/{id}/restore [post]
func RestoreDescontoHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID inválido",
		})
		return
	}

	err = services.RestoreDesconto(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Desconto não encontrado",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Desconto restaurado com sucesso",
	})
}

