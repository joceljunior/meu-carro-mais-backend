package json

import "time"

// CustomerResponse representa a resposta com dados de um customer
type CustomerResponse struct {
	ID             uint                 `json:"id"`
	Nome           string               `json:"nome"`
	Email          string               `json:"email"`
	CPF            string               `json:"cpf"`
	Imagem         string               `json:"imagem,omitempty"`
	Telefone       string               `json:"telefone,omitempty"`
	Endereco       string               `json:"endereco,omitempty"`
	DataNascimento *time.Time           `json:"data_nascimento,omitempty"`
	DataCadastro   time.Time            `json:"data_cadastro"`
	Ativo          bool                 `json:"ativo"`
	Latitude       *float64             `json:"latitude,omitempty"`
	Longitude      *float64             `json:"longitude,omitempty"`
	IDPlano        uint                 `json:"id_plano"`
	IDLoja         *uint                `json:"id_loja,omitempty"`
	Tipo           string               `json:"tipo"`
	Status         string               `json:"status"`
	IDExecutivo    *uint                `json:"id_executivo,omitempty"`
	Loja           *LojaUsuarioResponse `json:"loja,omitempty"`
	Executivo      *ExecutivoInfo       `json:"executivo,omitempty"`
	Mensagem       string               `json:"mensagem,omitempty"`
}

// ExecutivoInfo representa informações básicas do executivo
type ExecutivoInfo struct {
	ID    uint   `json:"id"`
	Nome  string `json:"nome"`
	Email string `json:"email"`
}

// CustomersListResponse representa a resposta com lista de customers
type CustomersListResponse struct {
	Customers []CustomerResponse `json:"customers"`
	Total     int                `json:"total"`
	Mensagem  string             `json:"mensagem,omitempty"`
}

// AprovacaoResponse representa a resposta de aprovação/rejeição de customer
type AprovacaoResponse struct {
	ID                   uint    `json:"id"`
	Status               string  `json:"status"`
	Motivo               string  `json:"motivo,omitempty"`
	BonificacaoExecutivo *int    `json:"bonificacao_executivo,omitempty"` // Moedas bonificadas ao executivo (se aprovado)
	Mensagem             string  `json:"mensagem"`
}

// SolicitacaoExecutivoResponse representa a resposta de uma solicitação para virar executivo
type SolicitacaoExecutivoResponse struct {
	ID                         uint       `json:"id"`
	Nome                       string     `json:"nome"`
	Email                      string     `json:"email"`
	Tipo                       string     `json:"tipo"`
	SolicitacaoExecutivo       string     `json:"solicitacao_executivo"`
	DataSolicitacaoExecutivo   *time.Time `json:"data_solicitacao_executivo,omitempty"`
	MotivoSolicitacaoExecutivo string     `json:"motivo_solicitacao_executivo,omitempty"`
	Mensagem                   string     `json:"mensagem"`
}

// SolicitacoesExecutivoListResponse representa a lista de solicitações de executivo
type SolicitacoesExecutivoListResponse struct {
	Solicitacoes []SolicitacaoExecutivoResponse `json:"solicitacoes"`
	Total        int                            `json:"total"`
	Mensagem     string                         `json:"mensagem,omitempty"`
}

