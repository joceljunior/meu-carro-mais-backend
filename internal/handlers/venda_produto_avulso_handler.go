package handlers

import (
	"net/http"
	"strconv"

	"meu-carro-mais/internal/handlers/json"
	"meu-carro-mais/internal/services"

	"github.com/gin-gonic/gin"
)

// CreateVendaProdutoAvulsoHandler godoc
// @Summary      Registra venda de produto não cadastrado
// @Description  A loja informa email do cliente, valor e descrição; o cliente é resolvido pelo email
// @Tags         Lojas
// @Accept       json
// @Produce      json
// @Param        id path int true "ID da loja"
// @Param        request body json.VendaProdutoAvulsoRequest true "Dados da venda"
// @Success      201  {object}  json.VendaProdutoAvulsoResponse
// @Failure      400  {object}  map[string]interface{}
// @Failure      404  {object}  map[string]interface{}
// @Failure      500  {object}  map[string]interface{}
// @Router       /lojas/{id}/vendas-produto-avulso [post]
func CreateVendaProdutoAvulsoHandler(c *gin.Context) {
	idStr := c.Param("id")
	idLoja, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID da loja inválido"})
		return
	}

	var req json.VendaProdutoAvulsoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos", "details": err.Error()})
		return
	}

	resp, err := services.CreateVendaProdutoAvulso(uint(idLoja), req)
	if err != nil {
		if err.Error() == "cliente não encontrado para o email informado" || err.Error() == "loja não encontrada" {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, resp)
}
