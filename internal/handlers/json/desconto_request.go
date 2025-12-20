package json

import "time"

type DescontoRequest struct {
	IDLoja       uint      `json:"id_loja" binding:"required"`
	Porcentagem  float64   `json:"porcentagem" binding:"required,min=0,max=100"`
	DataValidade time.Time `json:"data_validade" binding:"required"`
}

type CancelarDescontoRequest struct {
	IDLoja uint `json:"id_loja" binding:"required"`
}

