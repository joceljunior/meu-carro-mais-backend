package handlers

import (
	"meu-carro-mais/internal/handlers/json"
	"meu-carro-mais/internal/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateCarteiraHandler godoc
// @Summary      Criação de carteira
// @Description  Cria uma nova carteira para um usuário
// @Tags         Carteiras
// @Accept       json
// @Produce      json
// @Param        request body json.CarteiraRequest true "Dados da carteira"
// @Success      201  {object}  json.CarteiraResponse "Carteira criada com sucesso"
// @Failure      400  {object}  map[string]interface{} "Dados inválidos"
// @Failure      500  {object}  map[string]interface{} "Erro interno do servidor"
// @Router       /carteiras [post]
func CreateCarteiraHandler(c *gin.Context) {
	var req json.CarteiraRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Dados inválidos",
			"details": err.Error(),
		})
		return
	}

	resp, err := services.CreateCarteira(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, resp)
}

// GetCarteiraHandler godoc
// @Summary      Busca carteira por ID
// @Description  Retorna os dados de uma carteira específica pelo ID
// @Tags         Carteiras
// @Accept       json
// @Produce      json
// @Param        id path int true "ID da carteira"
// @Success      200 {object} json.CarteiraComUsuarioResponse "Carteira encontrada"
// @Failure      400 {object} map[string]interface{} "ID inválido"
// @Failure      404 {object} map[string]interface{} "Carteira não encontrada"
// @Failure      500 {object} map[string]interface{} "Erro interno do servidor"
// @Router       /carteiras/{id} [get]
func GetCarteiraHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID inválido",
		})
		return
	}

	resp, err := services.GetCarteiraByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Carteira não encontrada",
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// GetCarteiraByUsuarioHandler godoc
// @Summary      Busca carteira por usuário
// @Description  Retorna a carteira de um usuário específico
// @Tags         Carteiras
// @Accept       json
// @Produce      json
// @Param        usuario_id path int true "ID do usuário"
// @Success      200 {object} json.CarteiraComUsuarioResponse "Carteira encontrada"
// @Failure      400 {object} map[string]interface{} "ID inválido"
// @Failure      404 {object} map[string]interface{} "Carteira não encontrada"
// @Failure      500 {object} map[string]interface{} "Erro interno do servidor"
// @Router       /carteiras/usuario/{usuario_id} [get]
func GetCarteiraByUsuarioHandler(c *gin.Context) {
	usuarioIDStr := c.Param("usuario_id")
	usuarioID, err := strconv.ParseUint(usuarioIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID do usuário inválido",
		})
		return
	}

	resp, err := services.GetCarteiraByUsuarioID(uint(usuarioID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Carteira não encontrada",
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// GetAllCarteirasHandler godoc
// @Summary      Lista todas as carteiras
// @Description  Retorna uma lista com todas as carteiras
// @Tags         Carteiras
// @Accept       json
// @Produce      json
// @Success      200 {object} json.CarteirasResponse "Lista de carteiras"
// @Failure      500 {object} map[string]interface{} "Erro interno do servidor"
// @Router       /carteiras [get]
func GetAllCarteirasHandler(c *gin.Context) {
	resp, err := services.GetAllCarteiras()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// UpdateCarteiraHandler godoc
// @Summary      Atualiza carteira
// @Description  Atualiza os dados de uma carteira existente
// @Tags         Carteiras
// @Accept       json
// @Produce      json
// @Param        id path int true "ID da carteira"
// @Param        request body json.CarteiraRequest true "Dados atualizados da carteira"
// @Success      200 {object} json.CarteiraResponse "Carteira atualizada com sucesso"
// @Failure      400 {object} map[string]interface{} "Dados inválidos"
// @Failure      404 {object} map[string]interface{} "Carteira não encontrada"
// @Failure      500 {object} map[string]interface{} "Erro interno do servidor"
// @Router       /carteiras/{id} [put]
func UpdateCarteiraHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID inválido",
		})
		return
	}

	var req json.CarteiraRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Dados inválidos",
			"details": err.Error(),
		})
		return
	}

	resp, err := services.UpdateCarteira(uint(id), req)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Carteira não encontrada",
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// UpdateCarteiraSaldoHandler godoc
// @Summary      Atualiza saldo da carteira
// @Description  Atualiza apenas o saldo de uma carteira existente
// @Tags         Carteiras
// @Accept       json
// @Produce      json
// @Param        id path int true "ID da carteira"
// @Param        request body json.CarteiraSaldoRequest true "Novo saldo da carteira"
// @Success      200 {object} json.CarteiraResponse "Saldo atualizado com sucesso"
// @Failure      400 {object} map[string]interface{} "Dados inválidos"
// @Failure      404 {object} map[string]interface{} "Carteira não encontrada"
// @Failure      500 {object} map[string]interface{} "Erro interno do servidor"
// @Router       /carteiras/{id}/saldo [put]
func UpdateCarteiraSaldoHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID inválido",
		})
		return
	}

	var req json.CarteiraSaldoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Dados inválidos",
			"details": err.Error(),
		})
		return
	}

	resp, err := services.UpdateCarteiraSaldo(uint(id), req)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Carteira não encontrada",
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// AdicionarSaldoHandler godoc
// @Summary      Adiciona saldo à carteira
// @Description  Adiciona um valor ao saldo atual da carteira
// @Tags         Carteiras
// @Accept       json
// @Produce      json
// @Param        id path int true "ID da carteira"
// @Param        request body json.CarteiraOperacaoRequest true "Valor a ser adicionado"
// @Success      200 {object} json.CarteiraOperacaoResponse "Saldo adicionado com sucesso"
// @Failure      400 {object} map[string]interface{} "Dados inválidos"
// @Failure      404 {object} map[string]interface{} "Carteira não encontrada"
// @Failure      500 {object} map[string]interface{} "Erro interno do servidor"
// @Router       /carteiras/{id}/adicionar [post]
func AdicionarSaldoHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID inválido",
		})
		return
	}

	var req json.CarteiraOperacaoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Dados inválidos",
			"details": err.Error(),
		})
		return
	}

	resp, err := services.AdicionarSaldo(uint(id), req)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Carteira não encontrada",
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// SubtrairSaldoHandler godoc
// @Summary      Subtrai saldo da carteira
// @Description  Subtrai um valor do saldo atual da carteira
// @Tags         Carteiras
// @Accept       json
// @Produce      json
// @Param        id path int true "ID da carteira"
// @Param        request body json.CarteiraOperacaoRequest true "Valor a ser subtraído"
// @Success      200 {object} json.CarteiraOperacaoResponse "Saldo subtraído com sucesso"
// @Failure      400 {object} map[string]interface{} "Dados inválidos"
// @Failure      404 {object} map[string]interface{} "Carteira não encontrada"
// @Failure      500 {object} map[string]interface{} "Erro interno do servidor"
// @Router       /carteiras/{id}/subtrair [post]
func SubtrairSaldoHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID inválido",
		})
		return
	}

	var req json.CarteiraOperacaoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Dados inválidos",
			"details": err.Error(),
		})
		return
	}

	resp, err := services.SubtrairSaldo(uint(id), req)
	if err != nil {
		if err.Error() == "saldo insuficiente" {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Saldo insuficiente",
			})
			return
		}
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Carteira não encontrada",
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// DeleteCarteiraHandler godoc
// @Summary      Exclui carteira
// @Description  Remove uma carteira do sistema
// @Tags         Carteiras
// @Accept       json
// @Produce      json
// @Param        id path int true "ID da carteira"
// @Success      200 {object} map[string]interface{} "Carteira excluída com sucesso"
// @Failure      400 {object} map[string]interface{} "ID inválido"
// @Failure      404 {object} map[string]interface{} "Carteira não encontrada"
// @Failure      500 {object} map[string]interface{} "Erro interno do servidor"
// @Router       /carteiras/{id} [delete]
func DeleteCarteiraHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID inválido",
		})
		return
	}

	err = services.DeleteCarteira(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Carteira não encontrada",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Carteira excluída com sucesso",
	})
}

// GetCarteirasBySaldoRangeHandler godoc
// @Summary      Busca carteiras por range de saldo
// @Description  Retorna carteiras com saldo dentro de um range específico
// @Tags         Carteiras
// @Accept       json
// @Produce      json
// @Param        saldo_min query number true "Saldo mínimo" example(100)
// @Param        saldo_max query number true "Saldo máximo" example(2000)
// @Success      200 {object} json.CarteirasResponse "Carteiras encontradas"
// @Failure      400 {object} map[string]interface{} "Parâmetros inválidos"
// @Failure      500 {object} map[string]interface{} "Erro interno do servidor"
// @Router       /carteiras/range [get]
func GetCarteirasBySaldoRangeHandler(c *gin.Context) {
	saldoMinStr := c.Query("saldo_min")
	saldoMaxStr := c.Query("saldo_max")

	if saldoMinStr == "" || saldoMaxStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Parâmetros saldo_min e saldo_max são obrigatórios",
		})
		return
	}

	saldoMin, err := strconv.ParseFloat(saldoMinStr, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Saldo mínimo inválido",
		})
		return
	}

	saldoMax, err := strconv.ParseFloat(saldoMaxStr, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Saldo máximo inválido",
		})
		return
	}

	resp, err := services.GetCarteirasBySaldoRange(saldoMin, saldoMax)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// GetCarteirasComSaldoMaiorHandler godoc
// @Summary      Busca carteiras com saldo maior
// @Description  Retorna carteiras com saldo maior que um valor específico
// @Tags         Carteiras
// @Accept       json
// @Produce      json
// @Param        valor query number true "Valor mínimo de saldo" example(1000)
// @Success      200 {object} json.CarteirasResponse "Carteiras encontradas"
// @Failure      400 {object} map[string]interface{} "Parâmetro inválido"
// @Failure      500 {object} map[string]interface{} "Erro interno do servidor"
// @Router       /carteiras/saldo-maior [get]
func GetCarteirasComSaldoMaiorHandler(c *gin.Context) {
	valorStr := c.Query("valor")

	if valorStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Parâmetro valor é obrigatório",
		})
		return
	}

	valor, err := strconv.ParseFloat(valorStr, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Valor inválido",
		})
		return
	}

	resp, err := services.GetCarteirasComSaldoMaior(valor)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}
