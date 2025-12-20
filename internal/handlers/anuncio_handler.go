package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"meu-carro-mais/internal/handlers/json"
	"meu-carro-mais/internal/services"
)

// GetAnunciosHandler godoc
// @Summary      Lista todos os anúncios
// @Description  Retorna todos os anúncios disponíveis com informações da loja e categoria
// @Tags         Anúncios
// @Accept       json
// @Produce      json
// @Success      200  {object}  json.AnunciosResponse
// @Failure      500  {object}  map[string]interface{} "Erro interno do servidor"
// @Router       /anuncios [get]
func GetAnunciosHandler(c *gin.Context) {
	resp, err := services.GetAnuncios()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// CreateAnuncioHandler godoc
// @Summary      Criação do anúncio completo
// @Description  Cria um novo anúncio com todos os dados fornecidos
// @Tags         Anúncios
// @Accept       json
// @Produce      json
// @Param        request body json.AnuncioRequest true "Dados completos do anúncio"
// @Success      201  {object}  json.AnuncioResponse "Anúncio criado com sucesso"
// @Failure      400  {object}  map[string]interface{} "Dados inválidos"
// @Failure      500  {object}  map[string]interface{} "Erro interno do servidor"
// @Router       /anuncios [post]
func CreateAnuncioHandler(c *gin.Context) {
	var req json.AnuncioRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Dados inválidos",
			"details": err.Error(),
		})
		return
	}

	resp, err := services.CreateAnuncio(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Registra log da criação
	idAnuncio := uint(resp.ID)
	LogAction(c, "criar", "anuncio", &idAnuncio, 
		"Anúncio criado: "+resp.Titulo, nil, resp)

	c.JSON(http.StatusCreated, resp)
}

// GetAnuncioHandler godoc
// @Summary      Busca anúncio por ID
// @Description  Retorna os dados de um anúncio específico pelo ID
// @Tags         Anúncios
// @Accept       json
// @Produce      json
// @Param        id path int true "ID do anúncio"
// @Success      200 {object} json.AnuncioResponse "Anúncio encontrado"
// @Failure      400 {object} map[string]interface{} "ID inválido"
// @Failure      404 {object} map[string]interface{} "Anúncio não encontrado"
// @Failure      500 {object} map[string]interface{} "Erro interno do servidor"
// @Router       /anuncios/{id} [get]
func GetAnuncioHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID inválido",
		})
		return
	}

	resp, err := services.GetAnuncioByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Anúncio não encontrado",
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// GetAllAnunciosHandler godoc
// @Summary      Lista todos os anúncios
// @Description  Retorna uma lista com todos os anúncios ativos
// @Tags         Anúncios
// @Accept       json
// @Produce      json
// @Success      200 {array} json.AnuncioResponse "Lista de anúncios"
// @Failure      500 {object} map[string]interface{} "Erro interno do servidor"
// @Router       /anuncios [get]
func GetAllAnunciosHandler(c *gin.Context) {
	resp, err := services.GetAllAnuncios()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// UpdateAnuncioHandler godoc
// @Summary      Atualiza anúncio
// @Description  Atualiza os dados de um anúncio existente
// @Tags         Anúncios
// @Accept       json
// @Produce      json
// @Param        id path int true "ID do anúncio"
// @Param        request body json.AnuncioRequest true "Dados atualizados do anúncio"
// @Success      200 {object} json.AnuncioResponse "Anúncio atualizado com sucesso"
// @Failure      400 {object} map[string]interface{} "Dados inválidos"
// @Failure      404 {object} map[string]interface{} "Anúncio não encontrado"
// @Failure      500 {object} map[string]interface{} "Erro interno do servidor"
// @Router       /anuncios/{id} [put]
func UpdateAnuncioHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID inválido",
		})
		return
	}

	var req json.AnuncioRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Dados inválidos",
			"details": err.Error(),
		})
		return
	}

	// Busca dados antigos para o log
	anuncioAntigo, _ := services.GetAnuncioByID(uint(id))

	resp, err := services.UpdateAnuncio(uint(id), req)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Anúncio não encontrado",
		})
		return
	}

	// Registra log da atualização
	idAnuncio := uint(id)
	LogAction(c, "atualizar", "anuncio", &idAnuncio,
		"Anúncio atualizado: "+resp.Titulo, anuncioAntigo, resp)

	c.JSON(http.StatusOK, resp)
}

// SoftDeleteAnuncioHandler godoc
// @Summary      Exclui anúncio (soft delete)
// @Description  Realiza soft delete do anúncio, marcando como excluído sem remover do banco
// @Tags         Anúncios
// @Accept       json
// @Produce      json
// @Param        id path int true "ID do anúncio"
// @Success      200 {object} map[string]interface{} "Anúncio excluído com sucesso"
// @Failure      400 {object} map[string]interface{} "ID inválido"
// @Failure      404 {object} map[string]interface{} "Anúncio não encontrado"
// @Failure      500 {object} map[string]interface{} "Erro interno do servidor"
// @Router       /anuncios/{id} [delete]
func SoftDeleteAnuncioHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID inválido",
		})
		return
	}

	// Busca dados antes de deletar para o log
	anuncioAntigo, _ := services.GetAnuncioByID(uint(id))

	err = services.SoftDeleteAnuncio(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Anúncio não encontrado",
		})
		return
	}

	// Registra log da exclusão
	idAnuncio := uint(id)
	LogAction(c, "deletar", "anuncio", &idAnuncio,
		"Anúncio excluído (soft delete)", anuncioAntigo, nil)

	c.JSON(http.StatusOK, gin.H{
		"message": "Anúncio excluído com sucesso",
	})
}

// RestoreAnuncioHandler godoc
// @Summary      Restaura anúncio excluído
// @Description  Restaura um anúncio que foi soft deleted
// @Tags         Anúncios
// @Accept       json
// @Produce      json
// @Param        id path int true "ID do anúncio"
// @Success      200 {object} map[string]interface{} "Anúncio restaurado com sucesso"
// @Failure      400 {object} map[string]interface{} "ID inválido"
// @Failure      404 {object} map[string]interface{} "Anúncio não encontrado"
// @Failure      500 {object} map[string]interface{} "Erro interno do servidor"
// @Router       /anuncios/{id}/restore [post]
func RestoreAnuncioHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID inválido",
		})
		return
	}

	err = services.RestoreAnuncio(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Anúncio não encontrado ou não foi excluído",
		})
		return
	}

	// Busca dados restaurados para o log
	anuncioRestaurado, _ := services.GetAnuncioByID(uint(id))

	// Registra log da restauração
	idAnuncio := uint(id)
	LogAction(c, "restaurar", "anuncio", &idAnuncio,
		"Anúncio restaurado", nil, anuncioRestaurado)

	c.JSON(http.StatusOK, gin.H{
		"message": "Anúncio restaurado com sucesso",
	})
}

// GetCategoriasAnuncioHandler godoc
// @Summary      Lista categorias de anúncio
// @Description  Retorna todas as categorias de anúncio disponíveis
// @Tags         Anúncios
// @Accept       json
// @Produce      json
// @Success      200  {object}  json.CategoriasAnuncioResponse
// @Failure      500  {object}  map[string]interface{} "Erro interno do servidor"
// @Router       /anuncios/categorias [get]
func GetCategoriasAnuncioHandler(c *gin.Context) {
	resp, err := services.GetCategoriasAnuncio()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// GetAnunciosByLojaIDHandler godoc
// @Summary      Lista anúncios de uma loja
// @Description  Retorna todos os anúncios ativos de uma loja específica, ordenados por destaque e data
// @Tags         Anúncios
// @Accept       json
// @Produce      json
// @Param        loja_id path int true "ID da loja"
// @Success      200  {object}  json.AnunciosResponse
// @Failure      400  {object}  map[string]interface{} "ID de loja inválido"
// @Failure      500  {object}  map[string]interface{} "Erro interno do servidor"
// @Router       /anuncios/loja/{loja_id} [get]
func GetAnunciosByLojaIDHandler(c *gin.Context) {
	lojaIDStr := c.Param("loja_id")
	lojaID, err := strconv.ParseUint(lojaIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID de loja inválido",
		})
		return
	}

	resp, err := services.GetAnunciosByLojaID(uint(lojaID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// ResgatarAnuncioHandler godoc
// @Summary      Resgata um anúncio
// @Description  Cria um histórico de resgate com status pendente quando um usuário resgata um anúncio
// @Tags         Anúncios
// @Accept       json
// @Produce      json
// @Param        id path int true "ID do anúncio"
// @Param        request body json.ResgatarAnuncioRequest true "Dados do resgate (ID do usuário)"
// @Success      201  {object}  json.HistoricoResgateResponse "Histórico de resgate criado com sucesso"
// @Failure      400  {object}  map[string]interface{} "Dados inválidos"
// @Failure      404  {object}  map[string]interface{} "Anúncio não encontrado"
// @Failure      500  {object}  map[string]interface{} "Erro interno do servidor"
// @Router       /anuncios/{id}/resgatar [post]
func ResgatarAnuncioHandler(c *gin.Context) {
	idStr := c.Param("id")
	anuncioID, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID do anúncio inválido",
		})
		return
	}

	var req json.ResgatarAnuncioRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Dados inválidos",
			"details": err.Error(),
		})
		return
	}

	resp, err := services.CreateHistoricoResgateFromAnuncio(uint(anuncioID), req.IDUsuario)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Registra log do resgate
	idHistorico := uint(resp.ID)
	idAnuncio := uint(anuncioID)
	LogAction(c, "resgatar", "anuncio", &idAnuncio,
		"Anúncio resgatado - histórico de resgate criado", nil, resp)
	LogAction(c, "criar", "historico_resgate", &idHistorico,
		"Histórico de resgate criado para anúncio", nil, resp)

	c.JSON(http.StatusCreated, resp)
}