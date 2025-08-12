package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

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

// GetCategoriasServicoHandler godoc
// @Summary      Lista categorias de serviço
// @Description  Retorna todas as categorias de serviço disponíveis
// @Tags         Serviços
// @Accept       json
// @Produce      json
// @Success      200  {object}  json.CategoriasServicoResponse
// @Failure      500  {object}  map[string]interface{} "Erro interno do servidor"
// @Router       /servicos/categorias [get]
func GetCategoriasServicoHandler(c *gin.Context) {
	resp, err := services.GetCategoriasServico()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}
