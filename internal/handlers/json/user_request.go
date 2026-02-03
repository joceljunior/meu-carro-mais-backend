package json

import "time"

type UserRequest struct {
	Nome           string     `json:"nome" binding:"required"`
	Email          string     `json:"email" binding:"required,email"`
	Senha          string     `json:"senha,omitempty"`
	CPF            string     `json:"cpf" binding:"required"`
	Imagem         string     `json:"imagem,omitempty"`
	Telefone       string     `json:"telefone,omitempty"`
	Endereco       string     `json:"endereco,omitempty"`
	DataNascimento *time.Time `json:"data_nascimento,omitempty"`
	Latitude       *float64   `json:"latitude,omitempty"`
	Longitude      *float64   `json:"longitude,omitempty"`
}
