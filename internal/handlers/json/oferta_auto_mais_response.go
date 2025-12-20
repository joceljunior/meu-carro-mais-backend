package json

import "time"

type OfertaAutoMaisResponse struct {
	ID              uint         `json:"id"`
	IDLoja          uint         `json:"id_loja"`
	Nome            string       `json:"nome"`
	Descricao       string       `json:"descricao"`
	Moedas          int          `json:"moedas"`
	Porcentagem     float64      `json:"porcentagem"`
	Ativo           bool         `json:"ativo"`
	DataValidade    *time.Time   `json:"data_validade,omitempty"`
	DataCadastro    time.Time    `json:"data_cadastro"`
	DataAtualizacao time.Time    `json:"data_atualizacao"`
	Loja            LojaResponse `json:"loja,omitempty"`
}

type OfertasAutoMaisResponse struct {
	Ofertas []OfertaAutoMaisResponse `json:"ofertas"`
	Total   int                      `json:"total"`
}

