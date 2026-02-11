package json

import "time"

// HistoricoResgateClienteResponse representa um histórico de resgate simplificado para visualização do cliente
type HistoricoResgateClienteResponse struct {
	ID              uint           `json:"id"`
	IDCupom         *uint          `json:"id_cupom,omitempty"`
	IDUsuario       uint           `json:"id_usuario"`
	DataResgate     time.Time      `json:"data_resgate"`
	DataAtualizacao time.Time      `json:"data_atualizacao"`
	Status          string         `json:"status"`
	Cupom           *CupomResponse `json:"cupom,omitempty"`
}

// HistoricosResgateClienteResponse representa a lista de históricos de resgate do cliente
type HistoricosResgateClienteResponse struct {
	Historicos []HistoricoResgateClienteResponse `json:"historicos"`
	Total      int                               `json:"total"`
}
