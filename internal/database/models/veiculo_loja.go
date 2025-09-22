package models

import "time"

type VeiculoLoja struct {
	ID              uint       `gorm:"primaryKey"`
	Modelo          string     `gorm:"size:255;not null"`
	Ano             int        `gorm:"not null"`
	Cor             string     `gorm:"size:100;not null"`
	Placa           string     `gorm:"size:10;unique;not null"`
	IDLoja          uint       `gorm:"not null"`
	DataCadastro    time.Time  `gorm:"autoCreateTime"`
	DataAtualizacao time.Time  `gorm:"autoUpdateTime"`
	DataExclusao    *time.Time `gorm:"index"`
	Ativo           bool       `gorm:"default:true"`

	// Relacionamentos
	Loja       Loja               `gorm:"foreignKey:IDLoja"`
	Historicos []HistoricoVeiculo `gorm:"foreignKey:IDVeiculo"`
}
