package json

type AnuncioResponse struct {
	ID          uint             `json:"id"`
	Titulo      string           `json:"titulo"`
	Descricao   string           `json:"descricao"`
	Preco       float64          `json:"preco"`
	Imagem      string           `json:"imagem"`
	Destaque    bool             `json:"destaque"`
	Categoria   string           `json:"categoria"`
	IDLoja      uint             `json:"id_loja"`
	IDProduto   *uint            `json:"id_produto,omitempty"`
	IDServico   *uint            `json:"id_servico,omitempty"`
	IDVeiculo   *uint            `json:"id_veiculo,omitempty"`
	TipoAnuncio string           `json:"tipo_anuncio"`
	Loja        LojaResponse     `json:"loja"`
	Produto     *ProdutoResponse `json:"produto,omitempty"`
	Servico     *ServicoResponse `json:"servico,omitempty"`
	Veiculo     *VeiculoResponse `json:"veiculo,omitempty"`
}

type AnunciosResponse struct {
	Anuncios []AnuncioResponse `json:"anuncios"`
	Total    int               `json:"total"`
}
