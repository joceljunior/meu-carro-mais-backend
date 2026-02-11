package models

import "time"

type RegistroInteresse struct {
	ID              uint       `gorm:"primaryKey"`
	IDCupom         uint       `gorm:"not null;column:id_cupom"`
	Nome            string     `gorm:"size:255;not null"`
	Email           string     `gorm:"size:255;not null"`
	Telefone        string     `gorm:"size:20;not null"`
	Mensagem        string     `gorm:"size:1000"`
	DataCadastro    time.Time  `gorm:"autoCreateTime"`
	DataAtualizacao time.Time  `gorm:"autoUpdateTime"`
	DataExclusao    *time.Time `gorm:"index"`

	// Relacionamentos
	Cupom Cupom `gorm:"foreignKey:IDCupom"`
}
