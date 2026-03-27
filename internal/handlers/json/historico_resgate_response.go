package json

import "time"

type HistoricoResgateResponse struct {
	ID               uint          `json:"id"`
	IDCupom          *uint         `json:"id_cupom,omitempty"`
	IDUsuario        uint          `json:"id_usuario"`
	IDVeiculo        *uint         `json:"id_veiculo,omitempty" example:"19" description:"Veículo do cliente no resgate (id_veiculo_usuario ao resgatar), quando informado"`
	MoedasUtilizadas int           `json:"moedas_utilizadas"`
	DataResgate      time.Time     `json:"data_resgate"`
	DataAtualizacao  time.Time     `json:"data_atualizacao"`
	Status           string        `json:"status"`

	// Dados relacionados
	Cupom   *CupomResponse `json:"cupom,omitempty"`
	Usuario UserResponse   `json:"usuario,omitempty"`
}

type HistoricosResgateResponse struct {
	Historicos               []HistoricoResgateResponse   `json:"historicos"`
	VendasProdutoAvulso      []VendaProdutoAvulsoResponse `json:"vendas_produto_avulso"`
	Total                    int                          `json:"total"`
	TotalVendasProdutoAvulso int                          `json:"total_vendas_produto_avulso"`
}
