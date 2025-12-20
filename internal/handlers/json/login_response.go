package json

type LoginResponse struct {
	ID                  uint   `json:"id"`
	Nome                string `json:"nome"`
	Email               string `json:"email"`
	Tipo                string `json:"tipo"`
	Status              string `json:"status"`
	NomePlano           string `json:"nome_plano,omitempty"`
	LojaUsuarioResponse `json:"loja,omitempty"`
}

type AnuncioDestaqueResponse struct {
	ID          uint    `json:"id"`
	Titulo      string  `json:"titulo"`
	Descricao   string  `json:"descricao"`
	Preco       float64 `json:"preco"`
	Imagem      string  `json:"imagem"`
	TipoAnuncio string  `json:"tipo_anuncio"`
}
type LojaUsuarioResponse struct {
	Id                      uint   `json:"id"`
	Nome                    string `json:"nome"`
	Logo                    string `json:"logo,omitempty"`
	AnuncioDestaqueResponse `json:"anuncio_destaque,omitempty"`
}
