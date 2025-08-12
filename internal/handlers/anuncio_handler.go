package handlers

import (
	"meu-carro-mais/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
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
