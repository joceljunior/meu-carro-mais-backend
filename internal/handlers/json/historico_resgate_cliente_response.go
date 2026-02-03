package json

import "time"

// ItemCompraResponse representa um item comprado no histórico
type ItemCompraResponse struct {
	ID            uint    `json:"id"`
	Nome          string  `json:"nome"`
	Descricao     string  `json:"descricao,omitempty"`
	Imagem        string  `json:"imagem,omitempty"`
	TipoItem      string  `json:"tipo_item"` // "produto", "servico", "veiculo"
	Quantidade    int     `json:"quantidade"`
	ValorUnitario float64 `json:"valor_unitario"`
}

// HistoricoResgateClienteResponse representa um histórico de resgate simplificado para visualização do cliente
type HistoricoResgateClienteResponse struct {
	ID               uint               `json:"id"`
	NomeLoja         string             `json:"nome_loja"`
	ImagemLoja       string             `json:"imagem_loja"`
	DataResgate      time.Time          `json:"data_resgate"`
	Status           string             `json:"status"`
	Itens            []ItemCompraResponse `json:"itens"`
	Quantidade       int                `json:"quantidade"`
	ValorUnitario    float64            `json:"valor_unitario"`
	ValorOriginal    float64            `json:"valor_original"`
	DescontoAplicado float64            `json:"desconto_aplicado"`
	ValorTotal       float64            `json:"valor_total"`
}

// HistoricosResgateClienteResponse representa a lista de históricos de resgate do cliente
type HistoricosResgateClienteResponse struct {
	Historicos []HistoricoResgateClienteResponse `json:"historicos"`
	Total      int                               `json:"total"`
}
