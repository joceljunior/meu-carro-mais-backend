package handlers

import (
	"net/http"
	"strconv"

	"meu-carro-mais/internal/handlers/json"
	"meu-carro-mais/internal/services"

	"github.com/gin-gonic/gin"
)

// CreateProdutoHandler godoc
// @Summary      Cria um novo produto
// @Description  Cria um novo produto no sistema
// @Tags         Produtos
// @Accept       json
// @Produce      json
// @Param        produto body json.ProdutoRequest true "Dados do produto"
// @Success      201  {object}  json.ProdutoResponse
// @Failure      400  {object}  map[string]interface{} "Dados inválidos"
// @Failure      500  {object}  map[string]interface{} "Erro interno do servidor"
// @Router       /produtos [post]
func CreateProdutoHandler(c *gin.Context) {
	var req json.ProdutoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Dados inválidos: " + err.Error(),
		})
		return
	}

	resp, err := services.CreateProduto(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, resp)
}

// GetProdutoHandler godoc
// @Summary      Busca produto por ID
// @Description  Retorna um produto específico pelo ID
// @Tags         Produtos
// @Accept       json
// @Produce      json
// @Param        id path int true "ID do produto"
// @Success      200  {object}  json.ProdutoResponse
// @Failure      400  {object}  map[string]interface{} "ID inválido"
// @Failure      404  {object}  map[string]interface{} "Produto não encontrado"
// @Failure      500  {object}  map[string]interface{} "Erro interno do servidor"
// @Router       /produtos/{id} [get]
func GetProdutoHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID inválido",
		})
		return
	}

	resp, err := services.GetProdutoByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Produto não encontrado",
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// GetAllProdutosHandler godoc
// @Summary      Lista todos os produtos
// @Description  Retorna todos os produtos ativos do sistema
// @Tags         Produtos
// @Accept       json
// @Produce      json
// @Success      200  {object}  json.ProdutosResponse
// @Failure      500  {object}  map[string]interface{} "Erro interno do servidor"
// @Router       /produtos [get]
func GetAllProdutosHandler(c *gin.Context) {
	resp, err := services.GetAllProdutos()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	response := json.ProdutosResponse{
		Produtos: resp,
		Total:    len(resp),
	}

	c.JSON(http.StatusOK, response)
}

// GetProdutosByLojaIDHandler godoc
// @Summary      Lista produtos de uma loja
// @Description  Retorna todos os produtos ativos de uma loja específica
// @Tags         Produtos
// @Accept       json
// @Produce      json
// @Param        id path int true "ID da loja"
// @Success      200  {object}  json.ProdutosResponse
// @Failure      400  {object}  map[string]interface{} "ID de loja inválido"
// @Failure      500  {object}  map[string]interface{} "Erro interno do servidor"
// @Router       /lojas/{id}/produtos [get]
func GetProdutosByLojaIDHandler(c *gin.Context) {
	idLojaStr := c.Param("id")
	idLoja, err := strconv.ParseUint(idLojaStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID de loja inválido",
		})
		return
	}

	resp, err := services.GetProdutosByLojaID(uint(idLoja))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// UpdateProdutoHandler godoc
// @Summary      Atualiza produto
// @Description  Atualiza os dados de um produto existente
// @Tags         Produtos
// @Accept       json
// @Produce      json
// @Param        id path int true "ID do produto"
// @Param        produto body json.ProdutoRequest true "Dados atualizados do produto"
// @Success      200  {object}  json.ProdutoResponse
// @Failure      400  {object}  map[string]interface{} "Dados inválidos"
// @Failure      404  {object}  map[string]interface{} "Produto não encontrado"
// @Failure      500  {object}  map[string]interface{} "Erro interno do servidor"
// @Router       /produtos/{id} [put]
func UpdateProdutoHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID inválido",
		})
		return
	}

	var req json.ProdutoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Dados inválidos: " + err.Error(),
		})
		return
	}

	resp, err := services.UpdateProduto(uint(id), req)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Produto não encontrado",
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// SoftDeleteProdutoHandler godoc
// @Summary      Remove produto (soft delete)
// @Description  Realiza soft delete de um produto, marcando-o como excluído
// @Tags         Produtos
// @Accept       json
// @Produce      json
// @Param        id path int true "ID do produto"
// @Success      200  {object}  map[string]interface{} "Produto removido com sucesso"
// @Failure      400  {object}  map[string]interface{} "ID inválido"
// @Failure      404  {object}  map[string]interface{} "Produto não encontrado"
// @Failure      500  {object}  map[string]interface{} "Erro interno do servidor"
// @Router       /produtos/{id} [delete]
func SoftDeleteProdutoHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID inválido",
		})
		return
	}

	err = services.SoftDeleteProduto(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Produto não encontrado",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Produto removido com sucesso",
	})
}

// RestoreProdutoHandler godoc
// @Summary      Restaura produto
// @Description  Restaura um produto que foi soft deleted
// @Tags         Produtos
// @Accept       json
// @Produce      json
// @Param        id path int true "ID do produto"
// @Success      200  {object}  map[string]interface{} "Produto restaurado com sucesso"
// @Failure      400  {object}  map[string]interface{} "ID inválido"
// @Failure      404  {object}  map[string]interface{} "Produto não encontrado"
// @Failure      500  {object}  map[string]interface{} "Erro interno do servidor"
// @Router       /produtos/{id}/restore [post]
func RestoreProdutoHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID inválido",
		})
		return
	}

	err = services.RestoreProduto(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Produto não encontrado",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Produto restaurado com sucesso",
	})
}
