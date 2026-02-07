package json

type AnuncioProdutoResponse struct {
	ID                  uint     `json:"id"`
	NomeProduto         string   `json:"nome_produto"`
	NomeLoja            string   `json:"nome_loja"`
	EnderecoLoja        string   `json:"endereco_loja,omitempty"`
	Imagem              string   `json:"imagem"`
	PrecoOriginal       float64  `json:"preco_original"`
	PrecoComDesconto    float64  `json:"preco_com_desconto"`
	PorcentagemDesconto float64  `json:"porcentagem_desconto"`
	IsMeuCarroMais      bool     `json:"is_meu_carro_mais"`
	Categoria           string   `json:"categoria"`
	Descricao           string   `json:"descricao"`
	Rate                int      `json:"rate"`
	MoedasUtiliza       *int     `json:"moedas_utiliza,omitempty"` // Moedas da oferta Auto Mais, se houver
	Distancia           *float64 `json:"distancia,omitempty"`      // Distância em km da loja, se fornecida localização
}

type AnunciosProdutoResponse struct {
	Anuncios []AnuncioProdutoResponse `json:"anuncios"`
	Total    int                       `json:"total"`
}
