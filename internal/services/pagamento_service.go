package services

import (
	"errors"
	"meu-carro-mais/internal/config"
	"meu-carro-mais/internal/database/datasource"
	"meu-carro-mais/internal/database/models"
	"meu-carro-mais/internal/handlers/json"
)

// CreateCheckoutSession cria uma sessão de checkout no Stripe
func CreateCheckoutSession(req json.CheckoutRequest) (*json.CheckoutResponse, error) {
	// Busca o usuário para obter o email
	usuario, err := datasource.GetUserByID(req.IDUsuario)
	if err != nil {
		return nil, errors.New("usuário não encontrado")
	}

	// Cria a sessão de checkout no Stripe
	session, err := config.Stripe.CreateCheckoutSession(
		req.IDUsuario,
		usuario.Email,
		req.TipoPlano,
		req.SuccessURL,
		req.CancelURL,
	)
	if err != nil {
		return nil, errors.New("erro ao criar sessão de checkout: " + err.Error())
	}

	// Cria o histórico de pagamento
	_, err = datasource.CreateHistoricoPagamento(req, session.ID)
	if err != nil {
		return nil, errors.New("erro ao criar histórico de pagamento: " + err.Error())
	}

	response := &json.CheckoutResponse{
		SessionURL: session.URL,
		SessionID:  session.ID,
		Mensagem:   "Sessão de checkout criada com sucesso",
	}

	return response, nil
}

// ProcessWebhook processa o webhook do Stripe
func ProcessWebhook(req json.WebhookRequest) (*json.WebhookResponse, error) {
	// Processa diferentes tipos de eventos
	switch req.Type {
	case "checkout.session.completed":
		return processCheckoutSessionCompleted(req)
	case "checkout.session.expired":
		return processCheckoutSessionExpired(req)
	case "payment_intent.succeeded":
		return processPaymentIntentSucceeded(req)
	case "payment_intent.payment_failed":
		return processPaymentIntentFailed(req)
	default:
		return &json.WebhookResponse{
			Mensagem: "Evento não processado: " + req.Type,
			Status:   "ignored",
		}, nil
	}
}

// processCheckoutSessionCompleted processa quando uma sessão de checkout é completada
func processCheckoutSessionCompleted(req json.WebhookRequest) (*json.WebhookResponse, error) {
	sessionID := req.Data.Object.ID
	paymentStatus := req.Data.Object.PaymentStatus

	// Busca o histórico de pagamento
	historico, err := datasource.GetHistoricoPagamentoBySessionID(sessionID)
	if err != nil {
		return nil, errors.New("histórico de pagamento não encontrado")
	}

	// Atualiza o status baseado no payment_status
	var status string
	switch paymentStatus {
	case "paid":
		status = models.StatusPagamentoCompleted
		// Torna o usuário premium
		err = makeUserPremium(historico.IDUsuario, historico.TipoPlano)
		if err != nil {
			return nil, errors.New("erro ao tornar usuário premium: " + err.Error())
		}
	case "unpaid":
		status = models.StatusPagamentoFailed
	default:
		status = models.StatusPagamentoPending
	}

	// Atualiza o histórico de pagamento
	err = datasource.UpdateStatusHistoricoPagamentoBySessionID(sessionID, status)
	if err != nil {
		return nil, errors.New("erro ao atualizar histórico de pagamento: " + err.Error())
	}

	return &json.WebhookResponse{
		Mensagem: "Checkout session completed processado com sucesso",
		Status:   "success",
	}, nil
}

// processCheckoutSessionExpired processa quando uma sessão de checkout expira
func processCheckoutSessionExpired(req json.WebhookRequest) (*json.WebhookResponse, error) {
	sessionID := req.Data.Object.ID

	// Atualiza o status para canceled
	err := datasource.UpdateStatusHistoricoPagamentoBySessionID(sessionID, models.StatusPagamentoCanceled)
	if err != nil {
		return nil, errors.New("erro ao atualizar histórico de pagamento: " + err.Error())
	}

	return &json.WebhookResponse{
		Mensagem: "Checkout session expired processado com sucesso",
		Status:   "success",
	}, nil
}

// processPaymentIntentSucceeded processa quando um pagamento é bem-sucedido
func processPaymentIntentSucceeded(req json.WebhookRequest) (*json.WebhookResponse, error) {
	// Este evento é processado principalmente pelo checkout.session.completed
	// Mas pode ser útil para logs adicionais
	return &json.WebhookResponse{
		Mensagem: "Payment intent succeeded processado com sucesso",
		Status:   "success",
	}, nil
}

// processPaymentIntentFailed processa quando um pagamento falha
func processPaymentIntentFailed(req json.WebhookRequest) (*json.WebhookResponse, error) {
	// Este evento é processado principalmente pelo checkout.session.completed
	// Mas pode ser útil para logs adicionais
	return &json.WebhookResponse{
		Mensagem: "Payment intent failed processado com sucesso",
		Status:   "success",
	}, nil
}

// makeUserPremium torna um usuário premium
func makeUserPremium(userID uint, tipoPlano string) error {
	// Verifica se o usuário existe
	_, err := datasource.GetUserByID(userID)
	if err != nil {
		return errors.New("usuário não encontrado")
	}

	// Define o ID do plano baseado no tipo
	var planoID uint
	switch tipoPlano {
	case models.TipoPlanoMonthly:
		planoID = 2 // ID do plano mensal (assumindo que existe)
	case models.TipoPlanoYearly:
		planoID = 3 // ID do plano anual (assumindo que existe)
	default:
		return errors.New("tipo de plano inválido")
	}

	// Atualiza o plano do usuário diretamente no banco
	err = datasource.UpdateUserPlano(userID, planoID)
	if err != nil {
		return errors.New("erro ao atualizar usuário: " + err.Error())
	}

	return nil
}

// GetHistoricoPagamentoByID busca um histórico de pagamento por ID
func GetHistoricoPagamentoByID(id uint) (*json.HistoricoPagamentoResponse, error) {
	historico, err := datasource.GetHistoricoPagamentoByID(id)
	if err != nil {
		return nil, err
	}

	response := &json.HistoricoPagamentoResponse{
		ID:              historico.ID,
		IDUsuario:       historico.IDUsuario,
		StripeSessionID: historico.StripeSessionID,
		StripePaymentID: historico.StripePaymentID,
		Status:          historico.Status,
		TipoPlano:       historico.TipoPlano,
		Valor:           historico.Valor,
		Moeda:           historico.Moeda,
		DataPagamento:   historico.DataPagamento,
		DataVencimento:  historico.DataVencimento,
		DataCriacao:     historico.DataCriacao,
		DataAtualizacao: historico.DataAtualizacao,
	}

	// Adiciona dados do usuário
	if historico.Usuario.ID != 0 {
		usuarioResp := json.UserResponse{
			ID:             historico.Usuario.ID,
			Nome:           historico.Usuario.Nome,
			Email:          historico.Usuario.Email,
			CPF:            historico.Usuario.CPF,
			Imagem:         historico.Usuario.Imagem,
			Telefone:       historico.Usuario.Telefone,
			Endereco:       historico.Usuario.Endereco,
			DataNascimento: historico.Usuario.DataNascimento,
			DataCadastro:   historico.Usuario.DataCadastro,
			Ativo:          historico.Usuario.Ativo,
			Latitude:       historico.Usuario.Latitude,
			Longitude:      historico.Usuario.Longitude,
			IDPlano:        historico.Usuario.IDPlano,
			IDLoja:         historico.Usuario.IDLoja,
		}
		response.Usuario = usuarioResp
	}

	return response, nil
}

// GetHistoricosPagamentoByUsuarioID retorna todos os históricos de pagamento de um usuário
func GetHistoricosPagamentoByUsuarioID(idUsuario uint) (*json.HistoricosPagamentoResponse, error) {
	historicos, err := datasource.GetHistoricosPagamentoByUsuarioID(idUsuario)
	if err != nil {
		return nil, err
	}

	var historicosResponse []json.HistoricoPagamentoResponse
	for _, historico := range historicos {
		historicoResp := json.HistoricoPagamentoResponse{
			ID:              historico.ID,
			IDUsuario:       historico.IDUsuario,
			StripeSessionID: historico.StripeSessionID,
			StripePaymentID: historico.StripePaymentID,
			Status:          historico.Status,
			TipoPlano:       historico.TipoPlano,
			Valor:           historico.Valor,
			Moeda:           historico.Moeda,
			DataPagamento:   historico.DataPagamento,
			DataVencimento:  historico.DataVencimento,
			DataCriacao:     historico.DataCriacao,
			DataAtualizacao: historico.DataAtualizacao,
		}

		// Adiciona dados do usuário
		if historico.Usuario.ID != 0 {
			usuarioResp := json.UserResponse{
				ID:             historico.Usuario.ID,
				Nome:           historico.Usuario.Nome,
				Email:          historico.Usuario.Email,
				CPF:            historico.Usuario.CPF,
				Imagem:         historico.Usuario.Imagem,
				Telefone:       historico.Usuario.Telefone,
				Endereco:       historico.Usuario.Endereco,
				DataNascimento: historico.Usuario.DataNascimento,
				DataCadastro:   historico.Usuario.DataCadastro,
				Ativo:          historico.Usuario.Ativo,
				Latitude:       historico.Usuario.Latitude,
				Longitude:      historico.Usuario.Longitude,
				IDPlano:        historico.Usuario.IDPlano,
				IDLoja:         historico.Usuario.IDLoja,
			}
			historicoResp.Usuario = usuarioResp
		}

		historicosResponse = append(historicosResponse, historicoResp)
	}

	response := &json.HistoricosPagamentoResponse{
		Historicos: historicosResponse,
		Total:      len(historicosResponse),
	}

	return response, nil
}

// GetAllHistoricosPagamento retorna todos os históricos de pagamento
func GetAllHistoricosPagamento() (*json.HistoricosPagamentoResponse, error) {
	historicos, err := datasource.GetAllHistoricosPagamento()
	if err != nil {
		return nil, err
	}

	var historicosResponse []json.HistoricoPagamentoResponse
	for _, historico := range historicos {
		historicoResp := json.HistoricoPagamentoResponse{
			ID:              historico.ID,
			IDUsuario:       historico.IDUsuario,
			StripeSessionID: historico.StripeSessionID,
			StripePaymentID: historico.StripePaymentID,
			Status:          historico.Status,
			TipoPlano:       historico.TipoPlano,
			Valor:           historico.Valor,
			Moeda:           historico.Moeda,
			DataPagamento:   historico.DataPagamento,
			DataVencimento:  historico.DataVencimento,
			DataCriacao:     historico.DataCriacao,
			DataAtualizacao: historico.DataAtualizacao,
		}

		// Adiciona dados do usuário
		if historico.Usuario.ID != 0 {
			usuarioResp := json.UserResponse{
				ID:             historico.Usuario.ID,
				Nome:           historico.Usuario.Nome,
				Email:          historico.Usuario.Email,
				CPF:            historico.Usuario.CPF,
				Imagem:         historico.Usuario.Imagem,
				Telefone:       historico.Usuario.Telefone,
				Endereco:       historico.Usuario.Endereco,
				DataNascimento: historico.Usuario.DataNascimento,
				DataCadastro:   historico.Usuario.DataCadastro,
				Ativo:          historico.Usuario.Ativo,
				Latitude:       historico.Usuario.Latitude,
				Longitude:      historico.Usuario.Longitude,
				IDPlano:        historico.Usuario.IDPlano,
				IDLoja:         historico.Usuario.IDLoja,
			}
			historicoResp.Usuario = usuarioResp
		}

		historicosResponse = append(historicosResponse, historicoResp)
	}

	response := &json.HistoricosPagamentoResponse{
		Historicos: historicosResponse,
		Total:      len(historicosResponse),
	}

	return response, nil
}

// SoftDeleteHistoricoPagamento realiza soft delete do histórico de pagamento
func SoftDeleteHistoricoPagamento(id uint) error {
	return datasource.SoftDeleteHistoricoPagamento(id)
}

// RestoreHistoricoPagamento restaura um histórico de pagamento
func RestoreHistoricoPagamento(id uint) error {
	return datasource.RestoreHistoricoPagamento(id)
}
