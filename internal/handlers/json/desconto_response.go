package json

import "time"

type DescontoResponse struct {
	ID              uint         `json:"id"`
	IDLoja          uint         `json:"id_loja"`
	Porcentagem     float64      `json:"porcentagem"`
	Ativo           bool         `json:"ativo"`
	DataValidade    time.Time    `json:"data_validade"`
	DataCadastro    time.Time    `json:"data_cadastro"`
	DataAtualizacao time.Time    `json:"data_atualizacao"`
	Loja            LojaResponse `json:"loja,omitempty"`
}

type DescontosResponse struct {
	Descontos []DescontoResponse `json:"descontos"`
	Total     int                `json:"total"`
}

