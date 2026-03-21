package models

import "time"

type HistoricoResgate struct {
	ID              uint       `gorm:"primaryKey"`
	IDCupom         *uint      `gorm:"column:id_cupom"`
	IDUsuario         uint `gorm:"not null"`
	MoedasUtilizadas  int  `gorm:"not null;default:0"` // Moedas do app usadas no resgate (0 se não aplicável)
	DataResgate     time.Time  `gorm:"autoCreateTime"`
	DataAtualizacao time.Time  `gorm:"autoUpdateTime"`
	Status          string     `gorm:"size:20;default:'pendente'"` // "pendente", "efetivado"

	// Relacionamentos
	Cupom   *Cupom  `gorm:"foreignKey:IDCupom"`
	Usuario Usuario `gorm:"foreignKey:IDUsuario"`
}
