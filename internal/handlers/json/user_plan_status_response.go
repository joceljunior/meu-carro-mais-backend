package json

import "time"

type UserPlanStatusResponse struct {
	IDUsuario       uint       `json:"id_usuario"`
	NomeUsuario     string     `json:"nome_usuario"`
	EmailUsuario    string     `json:"email_usuario"`
	IDPlano         uint       `json:"id_plano"`
	NomePlano       string     `json:"nome_plano"`
	IsPremium       bool       `json:"is_premium"`
	DataVencimento  *time.Time `json:"data_vencimento,omitempty"`
	StatusPagamento string     `json:"status_pagamento,omitempty"`
	Mensagem        string     `json:"mensagem"`
}
