package models

import "time"

type Cupom struct {
	ID                  uint       `gorm:"primaryKey"`
	Titulo              string     `gorm:"size:255"`
	Descricao           string     `gorm:"size:255"`
	Preco               float64    `gorm:"type:decimal(10,2)"`
	Imagem              string     `gorm:"size:255"`
	Destaque            bool       `gorm:"default:false"`
	Categoria           string     `gorm:"size:100"`
	DataCadastro        time.Time  `gorm:"autoCreateTime"`
	DataAtualizacao     time.Time  `gorm:"autoUpdateTime"`
	DataExclusao        *time.Time `gorm:"index"`
	IDLoja              *uint      `gorm:"null"`
	IDProduto           *uint      `gorm:"null"`
	IDServico           *uint      `gorm:"null"`
	IDVeiculo           *uint      `gorm:"null"`
	IDOfertaAutoMais    *uint      `gorm:"null;index"`
	IDUsuario           *uint      `gorm:"null;index"`
	TipoCupom           string     `gorm:"size:20;not null;column:tipo_cupom"`
	PorcentagemDesconto float64    `gorm:"type:decimal(5,2);default:0"`
	PrecoComDesconto    float64    `gorm:"type:decimal(10,2)"`
	Loja                *Loja           `gorm:"foreignKey:IDLoja"`
	Produto             *Produto        `gorm:"foreignKey:IDProduto"`
	Servico             *Servico        `gorm:"foreignKey:IDServico"`
	Veiculo             *Veiculo        `gorm:"foreignKey:IDVeiculo"`
	OfertaAutoMais      *OfertaAutoMais `gorm:"foreignKey:IDOfertaAutoMais"`
	Usuario             *Usuario        `gorm:"foreignKey:IDUsuario"`
	HistoricosVeiculo   []HistoricoVeiculo `gorm:"foreignKey:IDCupom"`
}

func (Cupom) TableName() string {
	return "cupons"
}
