package json

import "time"

// HistoricoResgateClienteResponse representa um histórico de resgate simplificado para visualização do cliente
type HistoricoResgateClienteResponse struct {
	ID          uint      `json:"id"`
	NomeLoja    string    `json:"nome_loja"`
	ImagemLoja  string    `json:"imagem_loja"`
	DataResgate time.Time `json:"data_resgate"`
	Status      string    `json:"status"`
	Valor       float64   `json:"valor"`
}

// HistoricosResgateClienteResponse representa a lista de históricos de resgate do cliente
type HistoricosResgateClienteResponse struct {
	Historicos []HistoricoResgateClienteResponse `json:"historicos"`
	Total      int                                `json:"total"`
}
