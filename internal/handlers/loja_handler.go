package handlers

import (
	"fmt"
	"meu-carro-mais/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
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