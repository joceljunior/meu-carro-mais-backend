package json

import "time"

type OfertaAutoMaisRequest struct {
	IDLoja       uint       `json:"id_loja" binding:"required"`
	Nome         string     `json:"nome" binding:"required,max=255"`
	Descricao    string     `json:"descricao" binding:"max=500"`
	Moedas       int        `json:"moedas" binding:"required,min=1"` // Quantidade de moedas necessárias (obrigatório)
	Porcentagem  float64    `json:"porcentagem" binding:"required,min=0,max=100"`
	DataValidade *time.Time `json:"data_validade"`
}

type OfertaAutoMaisUpdateRequest struct {
	Nome         string     `json:"nome" binding:"max=255"`
	Descricao    string     `json:"descricao" binding:"max=500"`
	Moedas       *int       `json:"moedas" binding:"omitempty,min=1"`
	Porcentagem  *float64   `json:"porcentagem" binding:"omitempty,min=0,max=100"`
	Ativo        *bool      `json:"ativo"`
	DataValidade *time.Time `json:"data_validade"`
}

