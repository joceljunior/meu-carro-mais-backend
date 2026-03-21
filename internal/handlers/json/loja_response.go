package json

import (
	"time"

	"meu-carro-mais/internal/database/models"
)

// UsuarioIndicadorResponse representa o usuário que indicou a loja (versão simplificada)
type UsuarioIndicadorResponse struct {
	ID     uint   `json:"id"`
	Nome   string `json:"nome"`
	Email  string `json:"email"`
	Imagem string `json:"imagem,omitempty"`
}

type LojaResponse struct {
	ID                 uint                      `json:"id"`
	Nome               string                    `json:"nome"`
	CNPJ               string                    `json:"cnpj"`
	Imagem             string                    `json:"imagem"`
	Endereco           string                    `json:"endereco,omitempty"`
	Latitude           float64                   `json:"latitude"`
	Longitude          float64                   `json:"longitude"`
	Rating             int                       `json:"rating"`
	IsMeuCarroMais           bool                      `json:"is_meu_carro_mais"`
	Categoria                string                    `json:"categoria"`
	DescontoGeralPorcentagem float64                   `json:"desconto_geral_porcentagem"`
	IDUsuario                uint                      `json:"id_usuario"`
	IDUsuarioIndicador *uint                     `json:"id_usuario_indicador,omitempty"`
	DataVinculoUsuario *time.Time                `json:"data_vinculo_usuario,omitempty"`
	UsuarioIndicador   *UsuarioIndicadorResponse `json:"usuario_indicador,omitempty"`
	CupomDestaque      *CupomDestaqueResponse    `json:"cupom_destaque,omitempty"`
}

type LojasResponse struct {
	Lojas []LojaResponse `json:"lojas"`
	Total int            `json:"total"`
}

type CategoriaLojistaResponse struct {
	ID   uint   `json:"id"`
	Nome string `json:"nome"`
}

type CategoriasLojistaResponse struct {
	Categorias []CategoriaLojistaResponse `json:"categorias"`
	Total      int                        `json:"total"`
}

// LojaFromModel monta LojaResponse a partir do model (sem cupom destaque nem usuário indicador).
func LojaFromModel(loja models.Loja) LojaResponse {
	return LojaResponse{
		ID:                        loja.ID,
		Nome:                      loja.Nome,
		CNPJ:                      loja.CNPJ,
		Imagem:                    loja.Imagem,
		Endereco:                  loja.Endereco,
		Latitude:                  loja.Latitude,
		Longitude:                 loja.Longitude,
		Rating:                    loja.Rating,
		IsMeuCarroMais:            loja.IsMeuCarroMais,
		Categoria:                 loja.Categoria,
		DescontoGeralPorcentagem:  loja.DescontoGeralPorcentagem,
		IDUsuario:                 loja.IDUsuario,
		IDUsuarioIndicador:        loja.IDUsuarioIndicador,
		DataVinculoUsuario:        loja.DataVinculoUsuario,
	}
}
