package json

import "time"

type HistoricoResgateResponse struct {
	ID              uint          `json:"id"`
	IDCupom         *uint         `json:"id_cupom,omitempty"`
	IDUsuario       uint          `json:"id_usuario"`
	DataResgate     time.Time     `json:"data_resgate"`
	DataAtualizacao time.Time     `json:"data_atualizacao"`
	Status          string        `json:"status"`

	// Dados relacionados
	Cupom   *CupomResponse `json:"cupom,omitempty"`
	Usuario UserResponse   `json:"usuario,omitempty"`
}

type HistoricosResgateResponse struct {
	Historicos []HistoricoResgateResponse `json:"historicos"`
	Total      int                        `json:"total"`
}
