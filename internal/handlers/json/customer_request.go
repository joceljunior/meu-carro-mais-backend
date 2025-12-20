package json

import "time"

// CustomerRequest representa os dados para criação de um usuário customer
type CustomerRequest struct {
	Nome           string     `json:"nome" binding:"required"`
	Email          string     `json:"email" binding:"required,email"`
	Senha          string     `json:"senha" binding:"required,min=6"`
	CPF            string     `json:"cpf" binding:"required"`
	Imagem         string     `json:"imagem,omitempty"`
	Telefone       string     `json:"telefone,omitempty"`
	Endereco       string     `json:"endereco,omitempty"`
	DataNascimento *time.Time `json:"data_nascimento,omitempty"`
	Latitude       *float64   `json:"latitude,omitempty"`
	Longitude      *float64   `json:"longitude,omitempty"`
	IDExecutivo    *uint      `json:"id_executivo,omitempty"` // ID do executivo que está criando o customer (opcional)
}

// AdministrativoRequest representa os dados para criação de um usuário administrativo
type AdministrativoRequest struct {
	Nome           string     `json:"nome" binding:"required"`
	Email          string     `json:"email" binding:"required,email"`
	Senha          string     `json:"senha" binding:"required,min=6"`
	CPF            string     `json:"cpf" binding:"required"`
	Imagem         string     `json:"imagem,omitempty"`
	Telefone       string     `json:"telefone,omitempty"`
	Endereco       string     `json:"endereco,omitempty"`
	DataNascimento *time.Time `json:"data_nascimento,omitempty"`
	Latitude       *float64   `json:"latitude,omitempty"`
	Longitude      *float64   `json:"longitude,omitempty"`
}

// ExecutivoRequest representa os dados para criação de um usuário executivo
type ExecutivoRequest struct {
	Nome           string     `json:"nome" binding:"required"`
	Email          string     `json:"email" binding:"required,email"`
	Senha          string     `json:"senha" binding:"required,min=6"`
	CPF            string     `json:"cpf" binding:"required"`
	Imagem         string     `json:"imagem,omitempty"`
	Telefone       string     `json:"telefone,omitempty"`
	Endereco       string     `json:"endereco,omitempty"`
	DataNascimento *time.Time `json:"data_nascimento,omitempty"`
	Latitude       *float64   `json:"latitude,omitempty"`
	Longitude      *float64   `json:"longitude,omitempty"`
}

// AprovarCustomerRequest representa a requisição para aprovar um customer
type AprovarCustomerRequest struct {
	Motivo string `json:"motivo,omitempty"` // Motivo opcional da aprovação
}

// RejeitarCustomerRequest representa a requisição para rejeitar um customer
type RejeitarCustomerRequest struct {
	Motivo string `json:"motivo" binding:"required"` // Motivo obrigatório da rejeição
}

// SolicitarExecutivoRequest representa a requisição de um usuário mobile para virar executivo
type SolicitarExecutivoRequest struct {
	Motivo string `json:"motivo" binding:"required"` // Justificativa/motivo da solicitação
}

// AprovarSolicitacaoExecutivoRequest representa a requisição para aprovar solicitação de executivo
type AprovarSolicitacaoExecutivoRequest struct {
	Motivo string `json:"motivo,omitempty"` // Motivo opcional da aprovação
}

// RejeitarSolicitacaoExecutivoRequest representa a requisição para rejeitar solicitação de executivo
type RejeitarSolicitacaoExecutivoRequest struct {
	Motivo string `json:"motivo" binding:"required"` // Motivo obrigatório da rejeição
}

