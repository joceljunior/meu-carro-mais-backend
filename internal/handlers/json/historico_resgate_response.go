package json

import "time"

type HistoricoResgateResponse struct {
	ID               uint      `json:"id"`
	IDUsuario        uint      `json:"id_usuario"`
	IDProduto        *uint     `json:"id_produto,omitempty"`
	IDServico        *uint     `json:"id_servico,omitempty"`
	IDVeiculo        *uint     `json:"id_veiculo,omitempty"`
	IDLoja           uint      `json:"id_loja"`
	TipoResgate      string    `json:"tipo_resgate"`
	Quantidade       int       `json:"quantidade"`
	ValorUnitario    float64   `json:"valor_unitario"`
	ValorOriginal       float64   `json:"valor_original"`
	DescontoAplicado    float64   `json:"desconto_aplicado"`
	PorcentagemDesconto float64   `json:"porcentagem_desconto"`
	Valor               float64   `json:"valor"`
	Status           string    `json:"status"`
	DataResgate      time.Time `json:"data_resgate"`
	DataAtualizacao  time.Time `json:"data_atualizacao"`

	// Dados relacionados
	Usuario UserResponse     `json:"usuario,omitempty"`
	Produto *ProdutoResponse `json:"produto,omitempty"`
	Servico *ServicoResponse `json:"servico,omitempty"`
	Veiculo *VeiculoResponse `json:"veiculo,omitempty"`
	Loja    LojaResponse     `json:"loja,omitempty"`
}

type HistoricosResgateResponse struct {
	Historicos []HistoricoResgateResponse `json:"historicos"`
	Total      int                        `json:"total"`
}
