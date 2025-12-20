package json

import "time"

type UserResponse struct {
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
	Loja           *LojaUsuarioResponse `json:"loja,omitempty"`
	Mensagem       string               `json:"mensagem,omitempty"`
}
