package json

type LojaRequest struct {
	Nome               string  `json:"nome" binding:"required"`
	CNPJ               string  `json:"cnpj" binding:"required"`
	Imagem             string  `json:"imagem,omitempty"`
	Endereco           string  `json:"endereco,omitempty"`
	Latitude           float64 `json:"latitude" binding:"required"`
	Longitude          float64 `json:"longitude" binding:"required"`
	Rating             int     `json:"rating,omitempty"`
	IsMeuCarroMais     bool    `json:"is_meu_carro_mais,omitempty"`
	Categoria          string  `json:"categoria" binding:"required"`
	IDUsuario          uint    `json:"id_usuario" binding:"required"`
	IDUsuarioIndicador *uint   `json:"id_usuario_indicador,omitempty"` // ID do usuário que indicou esta loja (opcional)
}
