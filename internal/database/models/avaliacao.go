package models

import "time"

type Avaliacao struct {
	ID              uint       `gorm:"primaryKey"`
	IDUsuario       uint       `gorm:"not null"`
	IDLoja          uint       `gorm:"not null"`
	Nota            int        `gorm:"not null;check:nota >= 1 AND nota <= 5"` // Nota de 1 a 5
	Comentario      string     `gorm:"size:500"`
	DataAvaliacao   time.Time  `gorm:"autoCreateTime"`
	DataAtualizacao time.Time  `gorm:"autoUpdateTime"`
	DataExclusao    *time.Time `gorm:"index"`

	// Relacionamentos
	Usuario Usuario `gorm:"foreignKey:IDUsuario"`
	Loja    Loja    `gorm:"foreignKey:IDLoja"`
}
