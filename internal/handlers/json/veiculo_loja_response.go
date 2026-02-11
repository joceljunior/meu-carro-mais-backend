package json

import "time"

type VeiculoLojaResponse struct {
	ID           uint                  `json:"id"`
	Modelo       string                `json:"modelo"`
	Ano          int                   `json:"ano"`
	Cor          string                `json:"cor"`
	Placa        string                `json:"placa"`
	IDLoja       uint                  `json:"id_loja"`
	Imagem       string                `json:"imagem,omitempty"`
	Fotos        []VeiculoFotoResponse `json:"fotos,omitempty"`
	DataCadastro time.Time             `json:"data_cadastro"`
	Ativo        bool                  `json:"ativo"`
	Loja         LojaResponse          `json:"loja,omitempty"`
}

type VeiculosLojaResponse struct {
	Veiculos []VeiculoLojaResponse `json:"veiculos"`
	Total    int                   `json:"total"`
}
