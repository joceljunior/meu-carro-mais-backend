package json

type LojaResponse struct {
	ID              uint                     `json:"id"`
	Nome            string                   `json:"nome"`
	CNPJ            string                   `json:"cnpj"`
	Imagem          string                   `json:"imagem"`
	Latitude        float64                  `json:"latitude"`
	Longitude       float64                  `json:"longitude"`
	Rating          int                      `json:"rating"`
	IsMeuCarroMais  bool                     `json:"is_meu_carro_mais"`
	Categoria       string                   `json:"categoria"`
	IDUsuario       uint                     `json:"id_usuario"`
	AnuncioDestaque *AnuncioDestaqueResponse `json:"anuncio_destaque,omitempty"`
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
