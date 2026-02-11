package models

import "time"

type HistoricoResgate struct {
	ID              uint       `gorm:"primaryKey"`
	IDCupom         *uint      `gorm:"column:id_cupom"`
	IDUsuario       uint       `gorm:"not null"`
	DataResgate     time.Time  `gorm:"autoCreateTime"`
	DataAtualizacao time.Time  `gorm:"autoUpdateTime"`
	Status          string     `gorm:"size:20;default:'pendente'"` // "pendente", "efetivado"

	// Relacionamentos
	Cupom   *Cupom  `gorm:"foreignKey:IDCupom"`
	Usuario Usuario `gorm:"foreignKey:IDUsuario"`
}
