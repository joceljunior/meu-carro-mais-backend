package handlers

import (
	"meu-carro-mais/internal/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetVeiculosByUsuarioHandler godoc
// @Summary      Lista veículos de um usuário
// @Description  Retorna todos os veículos ativos de um usuário específico
// @Tags         Veículos
// @Accept       json
// @Produce      json
// @Param        id_usuario path int true "ID do usuário"
// @Success      200  {object}  json.VeiculosResponse
// @Failure      400  {object}  map[string]interface{} "ID de usuário inválido"
// @Failure      500  {object}  map[string]interface{} "Erro interno do servidor"
// @Router       /usuarios/{id_usuario}/veiculos [get]
func GetVeiculosByUsuarioHandler(c *gin.Context) {
	idUsuarioStr := c.Param("id_usuario")
	idUsuario, err := strconv.ParseUint(idUsuarioStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID de usuário inválido",
		})
		return
	}

	resp, err := services.GetVeiculosByUsuario(uint(idUsuario))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// GetHistoricosByVeiculoHandler godoc
// @Summary      Lista histórico de um veículo
// @Description  Retorna o histórico completo de um veículo específico
// @Tags         Veículos
// @Accept       json
// @Produce      json
// @Param        id_veiculo path int true "ID do veículo"
// @Success      200  {object}  json.HistoricosVeiculoResponse
// @Failure      400  {object}  map[string]interface{} "ID de veículo inválido"
// @Failure      500  {object}  map[string]interface{} "Erro interno do servidor"
// @Router       /veiculos/{id_veiculo}/historico [get]
func GetHistoricosByVeiculoHandler(c *gin.Context) {
	idVeiculoStr := c.Param("id_veiculo")
	idVeiculo, err := strconv.ParseUint(idVeiculoStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID de veículo inválido",
		})
		return
	}

	resp, err := services.GetHistoricosByVeiculo(uint(idVeiculo))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// GetHistoricosByUsuarioHandler godoc
// @Summary      Lista histórico de todos os veículos de um usuário
// @Description  Retorna o histórico de todos os veículos de um usuário específico
// @Tags         Veículos
// @Accept       json
// @Produce      json
// @Param        id_usuario path int true "ID do usuário"
// @Success      200  {object}  json.HistoricosVeiculoResponse
// @Failure      400  {object}  map[string]interface{} "ID de usuário inválido"
// @Failure      500  {object}  map[string]interface{} "Erro interno do servidor"
// @Router       /usuarios/{id_usuario}/historico [get]
func GetHistoricosByUsuarioHandler(c *gin.Context) {
	idUsuarioStr := c.Param("id_usuario")
	idUsuario, err := strconv.ParseUint(idUsuarioStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID de usuário inválido",
		})
		return
	}

	resp, err := services.GetHistoricosByUsuario(uint(idUsuario))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}
