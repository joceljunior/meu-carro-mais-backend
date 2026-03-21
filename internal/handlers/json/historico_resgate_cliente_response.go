package json

import "time"

// HistoricoResgateClienteResponse representa um histórico de resgate simplificado para visualização do cliente
type HistoricoResgateClienteResponse struct {
	ID               uint           `json:"id"`
	IDCupom          *uint          `json:"id_cupom,omitempty"`
	IDUsuario        uint           `json:"id_usuario"`
	MoedasUtilizadas int            `json:"moedas_utilizadas"`
	DataResgate      time.Time      `json:"data_resgate"`
	DataAtualizacao  time.Time      `json:"data_atualizacao"`
	Status           string         `json:"status"`
	Cupom            *CupomResponse `json:"cupom,omitempty"`
}

// HistoricosResgateClienteResponse representa a lista de históricos de resgate do cliente
type HistoricosResgateClienteResponse struct {
	Historicos               []HistoricoResgateClienteResponse `json:"historicos"`
	VendasProdutoAvulso      []VendaProdutoAvulsoResponse      `json:"vendas_produto_avulso"`
	Total                    int                               `json:"total"`
	TotalVendasProdutoAvulso int                               `json:"total_vendas_produto_avulso"`
}
