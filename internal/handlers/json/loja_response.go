package json

import "time"

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
	IsMeuCarroMais     bool                      `json:"is_meu_carro_mais"`
	Categoria          string                    `json:"categoria"`
	IDUsuario          uint                      `json:"id_usuario"`
	IDUsuarioIndicador *uint                     `json:"id_usuario_indicador,omitempty"` // ID do usuário que indicou esta loja (opcional)
	DataVinculoUsuario *time.Time                `json:"data_vinculo_usuario,omitempty"` // Data do vínculo com o usuário indicador
	UsuarioIndicador   *UsuarioIndicadorResponse `json:"usuario_indicador,omitempty"`    // Dados do usuário que indicou
	AnuncioDestaque    *AnuncioDestaqueResponse  `json:"anuncio_destaque,omitempty"`
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
