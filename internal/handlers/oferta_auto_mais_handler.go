package handlers

import (
	"net/http"
	"strconv"

	"meu-carro-mais/internal/handlers/json"
	"meu-carro-mais/internal/services"

	"github.com/gin-gonic/gin"
)

// CreateOfertaAutoMaisHandler godoc
// @Summary      Cria uma nova oferta Auto Mais para uma loja
// @Description  Cria uma nova oferta Auto Mais. Uma loja pode ter várias ofertas Auto Mais. O campo moedas é obrigatório e indica quantas moedas do app são necessárias para usar essa oferta.
// @Tags         Ofertas Auto Mais
// @Accept       json
// @Produce      json
// @Param        oferta body json.OfertaAutoMaisRequest true "Dados da oferta Auto Mais"
// @Success      201  {object}  json.OfertaAutoMaisResponse
// @Failure      400  {object}  map[string]interface{} "Dados inválidos"
// @Failure      500  {object}  map[string]interface{} "Erro interno do servidor"
// @Router       /ofertas-auto-mais [post]
func CreateOfertaAutoMaisHandler(c *gin.Context) {
	var req json.OfertaAutoMaisRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Dados inválidos: " + err.Error(),
		})
		return
	}

	resp, err := services.CreateOfertaAutoMais(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, resp)
}

// GetOfertaAutoMaisHandler godoc
// @Summary      Busca oferta Auto Mais por ID
// @Description  Retorna uma oferta Auto Mais específica pelo ID
// @Tags         Ofertas Auto Mais
// @Accept       json
// @Produce      json
// @Param        id path int true "ID da oferta"
// @Success      200  {object}  json.OfertaAutoMaisResponse
// @Failure      400  {object}  map[string]interface{} "ID inválido"
// @Failure      404  {object}  map[string]interface{} "Oferta não encontrada"
// @Router       /ofertas-auto-mais/{id} [get]
func GetOfertaAutoMaisHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID inválido",
		})
		return
	}

	resp, err := services.GetOfertaAutoMaisByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Oferta não encontrada",
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// GetAllOfertasAutoMaisHandler godoc
// @Summary      Lista todas as ofertas Auto Mais
// @Description  Retorna todas as ofertas Auto Mais do sistema
// @Tags         Ofertas Auto Mais
// @Accept       json
// @Produce      json
// @Success      200  {object}  json.OfertasAutoMaisResponse
// @Failure      500  {object}  map[string]interface{} "Erro interno do servidor"
// @Router       /ofertas-auto-mais [get]
func GetAllOfertasAutoMaisHandler(c *gin.Context) {
	resp, err := services.GetAllOfertasAutoMais()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	response := json.OfertasAutoMaisResponse{
		Ofertas: resp,
		Total:   len(resp),
	}

	c.JSON(http.StatusOK, response)
}

// GetOfertasAutoMaisAtivasHandler godoc
// @Summary      Lista todas as ofertas Auto Mais ativas
// @Description  Retorna todas as ofertas Auto Mais ativas e não expiradas do sistema
// @Tags         Ofertas Auto Mais
// @Accept       json
// @Produce      json
// @Success      200  {object}  json.OfertasAutoMaisResponse
// @Failure      500  {object}  map[string]interface{} "Erro interno do servidor"
// @Router       /ofertas-auto-mais/ativas [get]
func GetOfertasAutoMaisAtivasHandler(c *gin.Context) {
	resp, err := services.GetAllOfertasAutoMaisAtivas()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	response := json.OfertasAutoMaisResponse{
		Ofertas: resp,
		Total:   len(resp),
	}

	c.JSON(http.StatusOK, response)
}

// GetOfertasAutoMaisByLojaIDHandler godoc
// @Summary      Lista ofertas Auto Mais de uma loja
// @Description  Retorna todas as ofertas Auto Mais de uma loja específica
// @Tags         Ofertas Auto Mais
// @Accept       json
// @Produce      json
// @Param        id path int true "ID da loja"
// @Success      200  {object}  json.OfertasAutoMaisResponse
// @Failure      400  {object}  map[string]interface{} "ID de loja inválido"
// @Failure      500  {object}  map[string]interface{} "Erro interno do servidor"
// @Router       /lojas/{id}/ofertas-auto-mais [get]
func GetOfertasAutoMaisByLojaIDHandler(c *gin.Context) {
	idLojaStr := c.Param("id")
	idLoja, err := strconv.ParseUint(idLojaStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID de loja inválido",
		})
		return
	}

	resp, err := services.GetOfertasAutoMaisByLojaID(uint(idLoja))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// GetOfertasAutoMaisAtivasByLojaIDHandler godoc
// @Summary      Lista ofertas Auto Mais ativas de uma loja
// @Description  Retorna apenas as ofertas Auto Mais ativas de uma loja específica
// @Tags         Ofertas Auto Mais
// @Accept       json
// @Produce      json
// @Param        id path int true "ID da loja"
// @Success      200  {object}  json.OfertasAutoMaisResponse
// @Failure      400  {object}  map[string]interface{} "ID de loja inválido"
// @Failure      500  {object}  map[string]interface{} "Erro interno do servidor"
// @Router       /lojas/{id}/ofertas-auto-mais/ativas [get]
func GetOfertasAutoMaisAtivasByLojaIDHandler(c *gin.Context) {
	idLojaStr := c.Param("id")
	idLoja, err := strconv.ParseUint(idLojaStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID de loja inválido",
		})
		return
	}

	resp, err := services.GetOfertasAutoMaisAtivasByLojaID(uint(idLoja))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// UpdateOfertaAutoMaisHandler godoc
// @Summary      Atualiza uma oferta Auto Mais
// @Description  Atualiza os dados de uma oferta Auto Mais existente
// @Tags         Ofertas Auto Mais
// @Accept       json
// @Produce      json
// @Param        id path int true "ID da oferta"
// @Param        oferta body json.OfertaAutoMaisUpdateRequest true "Dados para atualização"
// @Success      200  {object}  json.OfertaAutoMaisResponse
// @Failure      400  {object}  map[string]interface{} "Dados inválidos"
// @Failure      404  {object}  map[string]interface{} "Oferta não encontrada"
// @Router       /ofertas-auto-mais/{id} [put]
func UpdateOfertaAutoMaisHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID inválido",
		})
		return
	}

	var req json.OfertaAutoMaisUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Dados inválidos: " + err.Error(),
		})
		return
	}

	resp, err := services.UpdateOfertaAutoMais(uint(id), req)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// DesativarOfertaAutoMaisHandler godoc
// @Summary      Desativa uma oferta Auto Mais
// @Description  Desativa uma oferta Auto Mais pelo ID
// @Tags         Ofertas Auto Mais
// @Accept       json
// @Produce      json
// @Param        id path int true "ID da oferta"
// @Success      200  {object}  map[string]interface{} "Oferta desativada com sucesso"
// @Failure      400  {object}  map[string]interface{} "ID inválido ou oferta já inativa"
// @Failure      404  {object}  map[string]interface{} "Oferta não encontrada"
// @Router       /ofertas-auto-mais/{id}/desativar [post]
func DesativarOfertaAutoMaisHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID inválido",
		})
		return
	}

	err = services.DesativarOfertaAutoMais(uint(id))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Oferta desativada com sucesso",
	})
}

// AtivarOfertaAutoMaisHandler godoc
// @Summary      Ativa uma oferta Auto Mais
// @Description  Ativa uma oferta Auto Mais pelo ID
// @Tags         Ofertas Auto Mais
// @Accept       json
// @Produce      json
// @Param        id path int true "ID da oferta"
// @Success      200  {object}  map[string]interface{} "Oferta ativada com sucesso"
// @Failure      400  {object}  map[string]interface{} "ID inválido ou oferta já ativa"
// @Failure      404  {object}  map[string]interface{} "Oferta não encontrada"
// @Router       /ofertas-auto-mais/{id}/ativar [post]
func AtivarOfertaAutoMaisHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID inválido",
		})
		return
	}

	err = services.AtivarOfertaAutoMais(uint(id))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Oferta ativada com sucesso",
	})
}

// SoftDeleteOfertaAutoMaisHandler godoc
// @Summary      Remove oferta Auto Mais (soft delete)
// @Description  Realiza soft delete de uma oferta Auto Mais, marcando-a como excluída
// @Tags         Ofertas Auto Mais
// @Accept       json
// @Produce      json
// @Param        id path int true "ID da oferta"
// @Success      200  {object}  map[string]interface{} "Oferta removida com sucesso"
// @Failure      400  {object}  map[string]interface{} "ID inválido"
// @Failure      404  {object}  map[string]interface{} "Oferta não encontrada"
// @Router       /ofertas-auto-mais/{id} [delete]
func SoftDeleteOfertaAutoMaisHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID inválido",
		})
		return
	}

	err = services.SoftDeleteOfertaAutoMais(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Oferta não encontrada",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Oferta removida com sucesso",
	})
}

// RestoreOfertaAutoMaisHandler godoc
// @Summary      Restaura oferta Auto Mais
// @Description  Restaura uma oferta Auto Mais que foi soft deleted
// @Tags         Ofertas Auto Mais
// @Accept       json
// @Produce      json
// @Param        id path int true "ID da oferta"
// @Success      200  {object}  map[string]interface{} "Oferta restaurada com sucesso"
// @Failure      400  {object}  map[string]interface{} "ID inválido"
// @Failure      404  {object}  map[string]interface{} "Oferta não encontrada"
// @Router       /ofertas-auto-mais/{id}/restore [post]
func RestoreOfertaAutoMaisHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID inválido",
		})
		return
	}

	err = services.RestoreOfertaAutoMais(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Oferta não encontrada",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Oferta restaurada com sucesso",
	})
}

