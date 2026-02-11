package models

import "time"

type HistoricoVeiculo struct {
	ID           uint      `gorm:"primaryKey"`
	IDVeiculo    uint      `gorm:"not null"`
	IDCupom      uint      `gorm:"not null;column:id_cupom"`
	Descricao    string    `gorm:"size:500;not null"`
	Data         time.Time `gorm:"not null"`
	DataCadastro time.Time `gorm:"autoCreateTime"`

	// Relacionamentos
	Veiculo Veiculo `gorm:"foreignKey:IDVeiculo"`
	Cupom   Cupom   `gorm:"foreignKey:IDCupom"`
}
