package json

type LojaRequest struct {
	Nome           string  `json:"nome" binding:"required"`
	CNPJ           string  `json:"cnpj" binding:"required"`
	Imagem         string  `json:"imagem,omitempty"`
	Latitude       float64 `json:"latitude" binding:"required"`
	Longitude      float64 `json:"longitude" binding:"required"`
	Rating         int     `json:"rating,omitempty"`
	IsMeuCarroMais bool    `json:"is_meu_carro_mais,omitempty"`
	IDCategoria    uint    `json:"id_categoria" binding:"required"`
}
