package json

type LojaRequest struct {
	Nome        string  `json:"nome" binding:"required"`
	CNPJ        string  `json:"cnpj" binding:"required"`
	Imagem      string  `json:"imagem,omitempty"`
	Latitude    float64 `json:"latitude" binding:"required"`
	Longitude   float64 `json:"longitude" binding:"required"`
	IDCategoria uint    `json:"id_categoria" binding:"required"`
}
