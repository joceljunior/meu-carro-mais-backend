package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	jsonHandlers "meu-carro-mais/internal/handlers/json"
	"meu-carro-mais/internal/services"

	"github.com/gin-gonic/gin"
)

// CreateCheckoutSessionHandler godoc
// @Summary      Cria sessão de checkout
// @Description  Cria uma sessão de checkout no Stripe para pagamento premium
// @Tags         Pagamentos
// @Accept       json
// @Produce      json
// @Param        checkout body json.CheckoutRequest true "Dados do checkout"
// @Success      200  {object}  json.CheckoutResponse
// @Failure      400  {object}  map[string]interface{} "Dados inválidos"
// @Failure      500  {object}  map[string]interface{} "Erro interno do servidor"
// @Router       /pagamentos/checkout [post]
func CreateCheckoutSessionHandler(c *gin.Context) {
	var req jsonHandlers.CheckoutRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Dados inválidos: " + err.Error(),
		})
		return
	}

	resp, err := services.CreateCheckoutSession(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// ProcessWebhookHandler godoc
// @Summary      Processa webhook do Stripe
// @Description  Processa webhooks do Stripe para atualizar status de pagamentos
// @Tags         Pagamentos
// @Accept       json
// @Produce      json
// @Param        webhook body json.WebhookRequest true "Dados do webhook"
// @Success      200  {object}  json.WebhookResponse
// @Failure      400  {object}  map[string]interface{} "Dados inválidos"
// @Failure      500  {object}  map[string]interface{} "Erro interno do servidor"
// @Router       /pagamentos/webhook [post]
func ProcessWebhookHandler(c *gin.Context) {
	// Lê o body da requisição
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Erro ao ler body da requisição",
		})
		return
	}

	// Verifica a assinatura do webhook (opcional, mas recomendado)
	signature := c.GetHeader("Stripe-Signature")
	if signature != "" {
		// Aqui você pode verificar a assinatura do webhook
		// usando o webhook secret configurado
		// event, err := webhook.ConstructEvent(body, signature, webhookSecret)
		// if err != nil {
		//     c.JSON(http.StatusBadRequest, gin.H{"error": "Assinatura inválida"})
		//     return
		// }
	}

	// Parse do JSON
	var webhookReq jsonHandlers.WebhookRequest
	if err := json.Unmarshal(body, &webhookReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Erro ao fazer parse do JSON: " + err.Error(),
		})
		return
	}

	// Processa o webhook
	resp, err := services.ProcessWebhook(webhookReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// GetHistoricoPagamentoHandler godoc
// @Summary      Busca histórico de pagamento por ID
// @Description  Retorna um histórico de pagamento específico pelo ID
// @Tags         Pagamentos
// @Accept       json
// @Produce      json
// @Param        id path int true "ID do histórico de pagamento"
// @Success      200  {object}  json.HistoricoPagamentoResponse
// @Failure      400  {object}  map[string]interface{} "ID inválido"
// @Failure      404  {object}  map[string]interface{} "Histórico não encontrado"
// @Failure      500  {object}  map[string]interface{} "Erro interno do servidor"
// @Router       /pagamentos/historicos/{id} [get]
func GetHistoricoPagamentoHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID inválido",
		})
		return
	}

	resp, err := services.GetHistoricoPagamentoByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Histórico de pagamento não encontrado",
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// GetHistoricosPagamentoByUsuarioIDHandler godoc
// @Summary      Lista históricos de pagamento de um usuário
// @Description  Retorna todos os históricos de pagamento de um usuário específico
// @Tags         Pagamentos
// @Accept       json
// @Produce      json
// @Param        id path int true "ID do usuário"
// @Success      200  {object}  json.HistoricosPagamentoResponse
// @Failure      400  {object}  map[string]interface{} "ID de usuário inválido"
// @Failure      500  {object}  map[string]interface{} "Erro interno do servidor"
// @Router       /pagamentos/usuarios/{id}/historicos [get]
func GetHistoricosPagamentoByUsuarioIDHandler(c *gin.Context) {
	idUsuarioStr := c.Param("id_usuario")
	idUsuario, err := strconv.ParseUint(idUsuarioStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID de usuário inválido",
		})
		return
	}

	resp, err := services.GetHistoricosPagamentoByUsuarioID(uint(idUsuario))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// GetAllHistoricosPagamentoHandler godoc
// @Summary      Lista todos os históricos de pagamento
// @Description  Retorna todos os históricos de pagamento do sistema
// @Tags         Pagamentos
// @Accept       json
// @Produce      json
// @Success      200  {object}  json.HistoricosPagamentoResponse
// @Failure      500  {object}  map[string]interface{} "Erro interno do servidor"
// @Router       /pagamentos/historicos [get]
func GetAllHistoricosPagamentoHandler(c *gin.Context) {
	resp, err := services.GetAllHistoricosPagamento()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// SoftDeleteHistoricoPagamentoHandler godoc
// @Summary      Remove histórico de pagamento (soft delete)
// @Description  Realiza soft delete de um histórico de pagamento
// @Tags         Pagamentos
// @Accept       json
// @Produce      json
// @Param        id path int true "ID do histórico de pagamento"
// @Success      200  {object}  map[string]interface{} "Histórico removido com sucesso"
// @Failure      400  {object}  map[string]interface{} "ID inválido"
// @Failure      404  {object}  map[string]interface{} "Histórico não encontrado"
// @Failure      500  {object}  map[string]interface{} "Erro interno do servidor"
// @Router       /pagamentos/historicos/{id} [delete]
func SoftDeleteHistoricoPagamentoHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID inválido",
		})
		return
	}

	err = services.SoftDeleteHistoricoPagamento(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Histórico de pagamento não encontrado",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Histórico de pagamento removido com sucesso",
	})
}

// RestoreHistoricoPagamentoHandler godoc
// @Summary      Restaura histórico de pagamento
// @Description  Restaura um histórico de pagamento que foi soft deleted
// @Tags         Pagamentos
// @Accept       json
// @Produce      json
// @Param        id path int true "ID do histórico de pagamento"
// @Success      200  {object}  map[string]interface{} "Histórico restaurado com sucesso"
// @Failure      400  {object}  map[string]interface{} "ID inválido"
// @Failure      404  {object}  map[string]interface{} "Histórico não encontrado"
// @Failure      500  {object}  map[string]interface{} "Erro interno do servidor"
// @Router       /pagamentos/historicos/{id}/restore [post]
func RestoreHistoricoPagamentoHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID inválido",
		})
		return
	}

	err = services.RestoreHistoricoPagamento(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Histórico de pagamento não encontrado",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Histórico de pagamento restaurado com sucesso",
	})
}
