package models

import "time"

type Avaliacao struct {
	ID              uint       `gorm:"primaryKey"`
	IDUsuario       uint       `gorm:"not null"`
	IDLoja          *uint      `gorm:"null"`                                   // Opcional - avaliar loja
	IDServico       *uint      `gorm:"null"`                                   // Opcional - avaliar serviço
	IDProduto       *uint      `gorm:"null"`                                   // Opcional - avaliar produto
	IDAnuncio       *uint      `gorm:"null"`                                   // Opcional - referência ao anúncio que gerou a avaliação
	Nota            int        `gorm:"not null;check:nota >= 1 AND nota <= 5"` // Nota de 1 a 5
	Comentario      string     `gorm:"size:500"`
	DataAvaliacao   time.Time  `gorm:"autoCreateTime"`
	DataAtualizacao time.Time  `gorm:"autoUpdateTime"`
	DataExclusao    *time.Time `gorm:"index"`

	// Relacionamentos
	Usuario Usuario  `gorm:"foreignKey:IDUsuario"`
	Loja    *Loja    `gorm:"foreignKey:IDLoja"`
	Servico *Servico `gorm:"foreignKey:IDServico"`
	Produto *Produto `gorm:"foreignKey:IDProduto"`
	Anuncio *Anuncio `gorm:"foreignKey:IDAnuncio"`
}
