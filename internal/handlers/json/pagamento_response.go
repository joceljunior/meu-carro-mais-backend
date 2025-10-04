package json

import "time"

type CheckoutResponse struct {
	SessionURL string `json:"session_url"`
	SessionID  string `json:"session_id"`
	Mensagem   string `json:"mensagem"`
}

type HistoricoPagamentoResponse struct {
	ID              uint       `json:"id"`
	IDUsuario       uint       `json:"id_usuario"`
	StripeSessionID string     `json:"stripe_session_id"`
	StripePaymentID string     `json:"stripe_payment_id,omitempty"`
	Status          string     `json:"status"`
	TipoPlano       string     `json:"tipo_plano"`
	Valor           float64    `json:"valor"`
	Moeda           string     `json:"moeda"`
	DataPagamento   *time.Time `json:"data_pagamento,omitempty"`
	DataVencimento  *time.Time `json:"data_vencimento,omitempty"`
	DataCriacao     time.Time  `json:"data_criacao"`
	DataAtualizacao time.Time  `json:"data_atualizacao"`

	// Dados do usuário
	Usuario UserResponse `json:"usuario,omitempty"`
}

type HistoricosPagamentoResponse struct {
	Historicos []HistoricoPagamentoResponse `json:"historicos"`
	Total      int                          `json:"total"`
}

type WebhookResponse struct {
	Mensagem string `json:"mensagem"`
	Status   string `json:"status"`
}

type CustomerPortalResponse struct {
	PortalURL string `json:"portal_url"`
	Mensagem  string `json:"mensagem"`
}
