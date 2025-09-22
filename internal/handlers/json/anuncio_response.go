package json

type AnuncioResponse struct {
	ID          uint             `json:"id"`
	Titulo      string           `json:"titulo"`
	Descricao   string           `json:"descricao"`
	Preco       float64          `json:"preco"`
	Imagem      string           `json:"imagem"`
	Destaque    bool             `json:"destaque"`
	IDLoja      uint             `json:"id_loja"`
	IDCategoria uint             `json:"id_categoria"`
	IDProduto   *uint            `json:"id_produto,omitempty"`
	IDServico   *uint            `json:"id_servico,omitempty"`
	IDVeiculo   *uint            `json:"id_veiculo,omitempty"`
	TipoAnuncio string           `json:"tipo_anuncio"`
	Categoria   string           `json:"categoria"`
	Loja        LojaResponse     `json:"loja"`
	Produto     *ProdutoResponse `json:"produto,omitempty"`
	Servico     *ServicoResponse `json:"servico,omitempty"`
	Veiculo     *VeiculoResponse `json:"veiculo,omitempty"`
}

type AnunciosResponse struct {
	Anuncios []AnuncioResponse `json:"anuncios"`
	Total    int               `json:"total"`
}

type CategoriaAnuncioResponse struct {
	ID   uint   `json:"id"`
	Nome string `json:"nome"`
}

type CategoriasAnuncioResponse struct {
	Categorias []CategoriaAnuncioResponse `json:"categorias"`
	Total      int                        `json:"total"`
}
