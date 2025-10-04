package services

import (
	"encoding/json"
	"errors"
	"meu-carro-mais/internal/config"
	"meu-carro-mais/internal/database/datasource"
	"meu-carro-mais/internal/database/models"
	jsonHandlers "meu-carro-mais/internal/handlers/json"

	"github.com/stripe/stripe-go/v83"
)

// CreateCheckoutSession cria uma sessão de checkout no Stripe
func CreateCheckoutSession(req jsonHandlers.CheckoutRequest) (*jsonHandlers.CheckoutResponse, error) {
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

	response := &jsonHandlers.CheckoutResponse{
		SessionURL: session.URL,
		SessionID:  session.ID,
		Mensagem:   "Sessão de checkout criada com sucesso",
	}

	return response, nil
}

// CreateSubscriptionCheckoutSession cria uma sessão de checkout para assinatura
func CreateSubscriptionCheckoutSession(req jsonHandlers.SubscriptionCheckoutRequest) (*jsonHandlers.CheckoutResponse, error) {
	// Busca o usuário para obter o email
	usuario, err := datasource.GetUserByID(req.IDUsuario)
	if err != nil {
		return nil, errors.New("usuário não encontrado")
	}

	// Cria a sessão de checkout para assinatura no Stripe
	session, err := config.Stripe.CreateSubscriptionCheckoutSession(
		req.IDUsuario,
		usuario.Email,
		req.LookupKey,
		req.SuccessURL,
		req.CancelURL,
	)
	if err != nil {
		return nil, errors.New("erro ao criar sessão de checkout para assinatura: " + err.Error())
	}

	// Cria o histórico de pagamento
	checkoutReq := jsonHandlers.CheckoutRequest{
		IDUsuario:  req.IDUsuario,
		TipoPlano:  "subscription", // Tipo especial para assinatura
		SuccessURL: req.SuccessURL,
		CancelURL:  req.CancelURL,
	}
	_, err = datasource.CreateHistoricoPagamento(checkoutReq, session.ID)
	if err != nil {
		return nil, errors.New("erro ao criar histórico de pagamento: " + err.Error())
	}

	response := &jsonHandlers.CheckoutResponse{
		SessionURL: session.URL,
		SessionID:  session.ID,
		Mensagem:   "Sessão de checkout para assinatura criada com sucesso",
	}

	return response, nil
}

// CreateCustomerPortalSession cria uma sessão do portal de cobrança
func CreateCustomerPortalSession(req jsonHandlers.CustomerPortalRequest) (*jsonHandlers.CustomerPortalResponse, error) {
	// Busca a sessão do Stripe para obter o customer ID
	session, err := config.Stripe.GetSession(req.SessionID)
	if err != nil {
		return nil, errors.New("sessão não encontrada: " + err.Error())
	}

	// Cria a sessão do portal de cobrança
	portalSession, err := config.Stripe.CreateCustomerPortalSession(
		session.Customer.ID,
		config.Stripe.Domain,
	)
	if err != nil {
		return nil, errors.New("erro ao criar sessão do portal: " + err.Error())
	}

	response := &jsonHandlers.CustomerPortalResponse{
		PortalURL: portalSession.URL,
		Mensagem:  "Sessão do portal de cobrança criada com sucesso",
	}

	return response, nil
}

// ProcessWebhookWithSignature processa webhook com verificação de assinatura
func ProcessWebhookWithSignature(payload []byte, signature string) (stripe.Event, error) {
	return config.Stripe.ConstructWebhookEvent(payload, signature)
}

// ProcessStripeEvent processa eventos do Stripe
func ProcessStripeEvent(event stripe.Event) (*jsonHandlers.WebhookResponse, error) {
	switch event.Type {
	case "customer.subscription.created":
		return processSubscriptionCreated(event)
	case "customer.subscription.updated":
		return processSubscriptionUpdated(event)
	case "customer.subscription.deleted":
		return processSubscriptionDeleted(event)
	case "customer.subscription.trial_will_end":
		return processSubscriptionTrialWillEnd(event)
	case "checkout.session.completed":
		return processCheckoutSessionCompletedFromEvent(event)
	case "checkout.session.expired":
		return processCheckoutSessionExpiredFromEvent(event)
	case "payment_intent.succeeded":
		return processPaymentIntentSucceededFromEvent(event)
	case "payment_intent.payment_failed":
		return processPaymentIntentFailedFromEvent(event)
	default:
		return &jsonHandlers.WebhookResponse{
			Mensagem: "Evento não processado: " + string(event.Type),
			Status:   "ignored",
		}, nil
	}
}

// ProcessWebhook processa o webhook do Stripe
func ProcessWebhook(req jsonHandlers.WebhookRequest) (*jsonHandlers.WebhookResponse, error) {
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
		return &jsonHandlers.WebhookResponse{
			Mensagem: "Evento não processado: " + req.Type,
			Status:   "ignored",
		}, nil
	}
}

// processCheckoutSessionCompleted processa quando uma sessão de checkout é completada
func processCheckoutSessionCompleted(req jsonHandlers.WebhookRequest) (*jsonHandlers.WebhookResponse, error) {
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

	return &jsonHandlers.WebhookResponse{
		Mensagem: "Checkout session completed processado com sucesso",
		Status:   "success",
	}, nil
}

// processCheckoutSessionExpired processa quando uma sessão de checkout expira
func processCheckoutSessionExpired(req jsonHandlers.WebhookRequest) (*jsonHandlers.WebhookResponse, error) {
	sessionID := req.Data.Object.ID

	// Atualiza o status para canceled
	err := datasource.UpdateStatusHistoricoPagamentoBySessionID(sessionID, models.StatusPagamentoCanceled)
	if err != nil {
		return nil, errors.New("erro ao atualizar histórico de pagamento: " + err.Error())
	}

	return &jsonHandlers.WebhookResponse{
		Mensagem: "Checkout session expired processado com sucesso",
		Status:   "success",
	}, nil
}

// processPaymentIntentSucceeded processa quando um pagamento é bem-sucedido
func processPaymentIntentSucceeded(req jsonHandlers.WebhookRequest) (*jsonHandlers.WebhookResponse, error) {
	// Este evento é processado principalmente pelo checkout.session.completed
	// Mas pode ser útil para logs adicionais
	return &jsonHandlers.WebhookResponse{
		Mensagem: "Payment intent succeeded processado com sucesso",
		Status:   "success",
	}, nil
}

// processPaymentIntentFailed processa quando um pagamento falha
func processPaymentIntentFailed(req jsonHandlers.WebhookRequest) (*jsonHandlers.WebhookResponse, error) {
	// Este evento é processado principalmente pelo checkout.session.completed
	// Mas pode ser útil para logs adicionais
	return &jsonHandlers.WebhookResponse{
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

// processSubscriptionCreated processa quando uma assinatura é criada
func processSubscriptionCreated(event stripe.Event) (*jsonHandlers.WebhookResponse, error) {
	var subscription stripe.Subscription
	err := json.Unmarshal(event.Data.Raw, &subscription)
	if err != nil {
		return nil, errors.New("erro ao fazer parse da assinatura: " + err.Error())
	}

	// Aqui você pode implementar lógica específica para quando uma assinatura é criada
	// Por exemplo, ativar funcionalidades premium do usuário

	return &jsonHandlers.WebhookResponse{
		Mensagem: "Assinatura criada processada com sucesso",
		Status:   "success",
	}, nil
}

// processSubscriptionUpdated processa quando uma assinatura é atualizada
func processSubscriptionUpdated(event stripe.Event) (*jsonHandlers.WebhookResponse, error) {
	var subscription stripe.Subscription
	err := json.Unmarshal(event.Data.Raw, &subscription)
	if err != nil {
		return nil, errors.New("erro ao fazer parse da assinatura: " + err.Error())
	}

	// Aqui você pode implementar lógica específica para quando uma assinatura é atualizada
	// Por exemplo, atualizar status do usuário baseado no status da assinatura

	return &jsonHandlers.WebhookResponse{
		Mensagem: "Assinatura atualizada processada com sucesso",
		Status:   "success",
	}, nil
}

// processSubscriptionDeleted processa quando uma assinatura é cancelada
func processSubscriptionDeleted(event stripe.Event) (*jsonHandlers.WebhookResponse, error) {
	var subscription stripe.Subscription
	err := json.Unmarshal(event.Data.Raw, &subscription)
	if err != nil {
		return nil, errors.New("erro ao fazer parse da assinatura: " + err.Error())
	}

	// Aqui você pode implementar lógica específica para quando uma assinatura é cancelada
	// Por exemplo, desativar funcionalidades premium do usuário

	return &jsonHandlers.WebhookResponse{
		Mensagem: "Assinatura cancelada processada com sucesso",
		Status:   "success",
	}, nil
}

// processSubscriptionTrialWillEnd processa quando o trial está prestes a acabar
func processSubscriptionTrialWillEnd(event stripe.Event) (*jsonHandlers.WebhookResponse, error) {
	var subscription stripe.Subscription
	err := json.Unmarshal(event.Data.Raw, &subscription)
	if err != nil {
		return nil, errors.New("erro ao fazer parse da assinatura: " + err.Error())
	}

	// Aqui você pode implementar lógica específica para notificar o usuário
	// Por exemplo, enviar email de aviso sobre o fim do trial

	return &jsonHandlers.WebhookResponse{
		Mensagem: "Aviso de fim de trial processado com sucesso",
		Status:   "success",
	}, nil
}

// processCheckoutSessionCompletedFromEvent processa checkout session completed a partir de evento
func processCheckoutSessionCompletedFromEvent(event stripe.Event) (*jsonHandlers.WebhookResponse, error) {
	var session stripe.CheckoutSession
	err := json.Unmarshal(event.Data.Raw, &session)
	if err != nil {
		return nil, errors.New("erro ao fazer parse da sessão: " + err.Error())
	}

	// Busca o histórico de pagamento
	historico, err := datasource.GetHistoricoPagamentoBySessionID(session.ID)
	if err != nil {
		return nil, errors.New("histórico de pagamento não encontrado")
	}

	// Atualiza o status baseado no payment_status
	var status string
	switch session.PaymentStatus {
	case stripe.CheckoutSessionPaymentStatusPaid:
		status = models.StatusPagamentoCompleted
		// Torna o usuário premium
		err = makeUserPremium(historico.IDUsuario, historico.TipoPlano)
		if err != nil {
			return nil, errors.New("erro ao tornar usuário premium: " + err.Error())
		}
	case stripe.CheckoutSessionPaymentStatusUnpaid:
		status = models.StatusPagamentoFailed
	default:
		status = models.StatusPagamentoPending
	}

	// Atualiza o histórico de pagamento
	err = datasource.UpdateStatusHistoricoPagamentoBySessionID(session.ID, status)
	if err != nil {
		return nil, errors.New("erro ao atualizar histórico de pagamento: " + err.Error())
	}

	return &jsonHandlers.WebhookResponse{
		Mensagem: "Checkout session completed processado com sucesso",
		Status:   "success",
	}, nil
}

// processCheckoutSessionExpiredFromEvent processa checkout session expired a partir de evento
func processCheckoutSessionExpiredFromEvent(event stripe.Event) (*jsonHandlers.WebhookResponse, error) {
	var session stripe.CheckoutSession
	err := json.Unmarshal(event.Data.Raw, &session)
	if err != nil {
		return nil, errors.New("erro ao fazer parse da sessão: " + err.Error())
	}

	// Atualiza o status para canceled
	err = datasource.UpdateStatusHistoricoPagamentoBySessionID(session.ID, models.StatusPagamentoCanceled)
	if err != nil {
		return nil, errors.New("erro ao atualizar histórico de pagamento: " + err.Error())
	}

	return &jsonHandlers.WebhookResponse{
		Mensagem: "Checkout session expired processado com sucesso",
		Status:   "success",
	}, nil
}

// processPaymentIntentSucceededFromEvent processa payment intent succeeded a partir de evento
func processPaymentIntentSucceededFromEvent(event stripe.Event) (*jsonHandlers.WebhookResponse, error) {
	// Este evento é processado principalmente pelo checkout.session.completed
	// Mas pode ser útil para logs adicionais
	return &jsonHandlers.WebhookResponse{
		Mensagem: "Payment intent succeeded processado com sucesso",
		Status:   "success",
	}, nil
}

// processPaymentIntentFailedFromEvent processa payment intent failed a partir de evento
func processPaymentIntentFailedFromEvent(event stripe.Event) (*jsonHandlers.WebhookResponse, error) {
	// Este evento é processado principalmente pelo checkout.session.completed
	// Mas pode ser útil para logs adicionais
	return &jsonHandlers.WebhookResponse{
		Mensagem: "Payment intent failed processado com sucesso",
		Status:   "success",
	}, nil
}

// GetHistoricoPagamentoByID busca um histórico de pagamento por ID
func GetHistoricoPagamentoByID(id uint) (*jsonHandlers.HistoricoPagamentoResponse, error) {
	historico, err := datasource.GetHistoricoPagamentoByID(id)
	if err != nil {
		return nil, err
	}

	response := &jsonHandlers.HistoricoPagamentoResponse{
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
		usuarioResp := jsonHandlers.UserResponse{
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
func GetHistoricosPagamentoByUsuarioID(idUsuario uint) (*jsonHandlers.HistoricosPagamentoResponse, error) {
	historicos, err := datasource.GetHistoricosPagamentoByUsuarioID(idUsuario)
	if err != nil {
		return nil, err
	}

	var historicosResponse []jsonHandlers.HistoricoPagamentoResponse
	for _, historico := range historicos {
		historicoResp := jsonHandlers.HistoricoPagamentoResponse{
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
			usuarioResp := jsonHandlers.UserResponse{
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

	response := &jsonHandlers.HistoricosPagamentoResponse{
		Historicos: historicosResponse,
		Total:      len(historicosResponse),
	}

	return response, nil
}

// GetAllHistoricosPagamento retorna todos os históricos de pagamento
func GetAllHistoricosPagamento() (*jsonHandlers.HistoricosPagamentoResponse, error) {
	historicos, err := datasource.GetAllHistoricosPagamento()
	if err != nil {
		return nil, err
	}

	var historicosResponse []jsonHandlers.HistoricoPagamentoResponse
	for _, historico := range historicos {
		historicoResp := jsonHandlers.HistoricoPagamentoResponse{
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
			usuarioResp := jsonHandlers.UserResponse{
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

	response := &jsonHandlers.HistoricosPagamentoResponse{
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
