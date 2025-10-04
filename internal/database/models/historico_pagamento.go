package models

import "time"

type HistoricoPagamento struct {
	ID              uint       `gorm:"primaryKey"`
	IDUsuario       uint       `gorm:"not null"`
	StripeSessionID string     `gorm:"size:255;unique;not null"`
	StripePaymentID string     `gorm:"size:255"`
	Status          string     `gorm:"size:50;not null;default:'pending'"` // pending, completed, failed, canceled
	TipoPlano       string     `gorm:"size:50;not null"`                   // monthly, yearly
	Valor           float64    `gorm:"type:decimal(10,2);not null"`
	Moeda           string     `gorm:"size:3;not null;default:'BRL'"`
	DataPagamento   *time.Time `gorm:"null"`
	DataVencimento  *time.Time `gorm:"null"`
	DataCriacao     time.Time  `gorm:"autoCreateTime"`
	DataAtualizacao time.Time  `gorm:"autoUpdateTime"`
	DataExclusao    *time.Time `gorm:"index"`

	// Relacionamentos
	Usuario Usuario `gorm:"foreignKey:IDUsuario"`
}

// StatusPagamento representa os possíveis status de um pagamento
const (
	StatusPagamentoPending   = "pending"
	StatusPagamentoCompleted = "completed"
	StatusPagamentoFailed    = "failed"
	StatusPagamentoCanceled  = "canceled"
)

// Tipos de plano
const (
	TipoPlanoMonthly      = "monthly"
	TipoPlanoYearly       = "yearly"
	TipoPlanoSubscription = "subscription"
)
