package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"meu-carro-mais/internal/handlers/json"
	"meu-carro-mais/internal/services"
)

// TransferirVeiculoHandler godoc
// @Summary      Transfere um veículo para outro usuário
// @Description  Realiza a transferência de propriedade de um veículo de um usuário para outro. O veículo passa a pertencer ao novo usuário.
// @Tags         Transferências
// @Accept       json
// @Produce      json
// @Param        id_usuario path int true "ID do usuário atual (dono do veículo)"
// @Param        request body json.TransferenciaVeiculoRequest true "Dados da transferência"
// @Success      200 {object} json.TransferenciaVeiculoResponse "Veículo transferido com sucesso"
// @Failure      400 {object} map[string]interface{} "Dados inválidos ou veículo não pertence ao usuário"
// @Failure      404 {object} map[string]interface{} "Veículo ou usuário não encontrado"
// @Failure      500 {object} map[string]interface{} "Erro interno do servidor"
// @Router       /usuarios/{id_usuario}/transferir-veiculo [post]
func TransferirVeiculoHandler(c *gin.Context) {
	idUsuarioStr := c.Param("id_usuario")
	idUsuario, err := strconv.ParseUint(idUsuarioStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID de usuário inválido",
		})
		return
	}

	var req json.TransferenciaVeiculoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Dados inválidos",
			"details": err.Error(),
		})
		return
	}

	resp, err := services.TransferirVeiculoManual(req, uint(idUsuario))
	if err != nil {
		errMsg := err.Error()
		switch errMsg {
		case "veículo não encontrado":
			c.JSON(http.StatusNotFound, gin.H{
				"error": errMsg,
			})
		case "usuário destino não encontrado":
			c.JSON(http.StatusNotFound, gin.H{
				"error": errMsg,
			})
		case "veículo não pertence ao usuário de origem":
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errMsg,
			})
		case "usuário de origem e destino não podem ser o mesmo":
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errMsg,
			})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":   "Erro ao transferir veículo",
				"details": errMsg,
			})
		}
		return
	}

	// Log da transferência
	idVeiculo := req.IDVeiculo
	LogAction(c, "transferir", "veiculo", &idVeiculo,
		"Veículo transferido manualmente", nil, resp)

	c.JSON(http.StatusOK, resp)
}

// GetTransferenciaHandler godoc
// @Summary      Busca transferência por ID
// @Description  Retorna os dados de uma transferência específica pelo ID
// @Tags         Transferências
// @Accept       json
// @Produce      json
// @Param        id path int true "ID da transferência"
// @Success      200 {object} json.TransferenciaVeiculoResponse "Transferência encontrada"
// @Failure      400 {object} map[string]interface{} "ID inválido"
// @Failure      404 {object} map[string]interface{} "Transferência não encontrada"
// @Router       /transferencias/{id} [get]
func GetTransferenciaHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID inválido",
		})
		return
	}

	resp, err := services.GetTransferenciaByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Transferência não encontrada",
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// GetTransferenciasByVeiculoHandler godoc
// @Summary      Lista transferências de um veículo
// @Description  Retorna o histórico de transferências de um veículo específico
// @Tags         Transferências
// @Accept       json
// @Produce      json
// @Param        id path int true "ID do veículo"
// @Success      200 {object} json.TransferenciasResponse "Lista de transferências"
// @Failure      400 {object} map[string]interface{} "ID inválido"
// @Failure      500 {object} map[string]interface{} "Erro interno do servidor"
// @Router       /veiculos/{id}/transferencias [get]
func GetTransferenciasByVeiculoHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID de veículo inválido",
		})
		return
	}

	resp, err := services.GetTransferenciasByVeiculo(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// GetTransferenciasByUsuarioHandler godoc
// @Summary      Lista transferências de um usuário
// @Description  Retorna todas as transferências envolvendo um usuário (como origem ou destino)
// @Tags         Transferências
// @Accept       json
// @Produce      json
// @Param        id_usuario path int true "ID do usuário"
// @Success      200 {object} json.TransferenciasResponse "Lista de transferências"
// @Failure      400 {object} map[string]interface{} "ID inválido"
// @Failure      500 {object} map[string]interface{} "Erro interno do servidor"
// @Router       /usuarios/{id_usuario}/transferencias [get]
func GetTransferenciasByUsuarioHandler(c *gin.Context) {
	idStr := c.Param("id_usuario")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID de usuário inválido",
		})
		return
	}

	resp, err := services.GetTransferenciasByUsuario(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// GetAllTransferenciasHandler godoc
// @Summary      Lista todas as transferências
// @Description  Retorna todas as transferências de veículos do sistema
// @Tags         Transferências
// @Accept       json
// @Produce      json
// @Success      200 {object} json.TransferenciasResponse "Lista de transferências"
// @Failure      500 {object} map[string]interface{} "Erro interno do servidor"
// @Router       /transferencias [get]
func GetAllTransferenciasHandler(c *gin.Context) {
	resp, err := services.GetAllTransferencias()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// BuscarUsuariosParaTransferenciaHandler godoc
// @Summary      Busca usuários para transferência
// @Description  Busca usuários ativos para selecionar como novo dono do veículo. Permite buscar por nome, email ou CPF.
// @Tags         Transferências
// @Accept       json
// @Produce      json
// @Param        termo query string false "Termo de busca (nome, email ou CPF)"
// @Success      200 {object} json.UsuariosBuscaResponse "Lista de usuários encontrados"
// @Failure      500 {object} map[string]interface{} "Erro interno do servidor"
// @Router       /transferencias/buscar-usuarios [get]
func BuscarUsuariosParaTransferenciaHandler(c *gin.Context) {
	termo := c.Query("termo")

	resp, err := services.BuscarUsuariosParaTransferencia(termo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}
