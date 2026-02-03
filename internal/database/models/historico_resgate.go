package models

import "time"

type HistoricoResgate struct {
	ID               uint       `gorm:"primaryKey"`
	IDUsuario        uint       `gorm:"not null"`
	IDProduto        *uint      `gorm:"null"` // Pode ser null se for serviço
	IDServico        *uint      `gorm:"null"` // Pode ser null se for produto
	IDVeiculo        *uint      `gorm:"null"` // Pode ser null se for produto/serviço
	IDLoja           uint       `gorm:"not null"`
	TipoResgate      string     `gorm:"size:20;not null"` // "produto", "servico", "veiculo"
	Quantidade       int        `gorm:"default:1"`        // Quantidade do item
	ValorUnitario    float64    `gorm:"type:decimal(10,2);default:0"`
	ValorOriginal    float64    `gorm:"type:decimal(10,2);default:0"`    // Valor antes do desconto
	DescontoAplicado float64    `gorm:"type:decimal(10,2);default:0"`    // Valor do desconto aplicado
	Valor            float64    `gorm:"type:decimal(10,2);not null"`     // Valor total pago (após desconto)
	Status           string     `gorm:"size:20;default:'pendente'"`      // "pendente", "confirmado", "cancelado"
	DataResgate      time.Time  `gorm:"autoCreateTime"`
	DataAtualizacao  time.Time  `gorm:"autoUpdateTime"`
	DataExclusao     *time.Time `gorm:"index"`

	// Relacionamentos
	Usuario Usuario  `gorm:"foreignKey:IDUsuario"`
	Produto *Produto `gorm:"foreignKey:IDProduto"`
	Servico *Servico `gorm:"foreignKey:IDServico"`
	Veiculo *Veiculo `gorm:"foreignKey:IDVeiculo"`
	Loja    Loja     `gorm:"foreignKey:IDLoja"`
}
