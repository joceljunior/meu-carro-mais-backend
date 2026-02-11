package models

import "time"

type Avaliacao struct {
	ID              uint       `gorm:"primaryKey"`
	IDUsuario       uint       `gorm:"not null"`
	IDLoja          *uint      `gorm:"null"`
	IDServico       *uint      `gorm:"null"`
	IDProduto       *uint      `gorm:"null"`
	IDCupom         *uint      `gorm:"null;column:id_cupom"`
	Nota            int        `gorm:"not null;check:nota >= 1 AND nota <= 5"`
	Comentario      string     `gorm:"size:500"`
	DataAvaliacao   time.Time  `gorm:"autoCreateTime"`
	DataAtualizacao time.Time  `gorm:"autoUpdateTime"`
	DataExclusao    *time.Time `gorm:"index"`

	// Relacionamentos
	Usuario Usuario  `gorm:"foreignKey:IDUsuario"`
	Loja    *Loja    `gorm:"foreignKey:IDLoja"`
	Servico *Servico `gorm:"foreignKey:IDServico"`
	Produto *Produto `gorm:"foreignKey:IDProduto"`
	Cupom   *Cupom   `gorm:"foreignKey:IDCupom"`
}
