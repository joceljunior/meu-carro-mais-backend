package json

type CupomProdutoResponse struct {
	ID                  uint     `json:"id"`
	IDLoja              *uint    `json:"id_loja,omitempty" example:"1" description:"ID da loja dona do cupom"`
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
	MoedasUtiliza       *int     `json:"moedas_utiliza,omitempty"`
	Distancia           *float64 `json:"distancia,omitempty"`
}

type CuponsProdutoResponse struct {
	Cupons []CupomProdutoResponse `json:"cupons"`
	Total  int                    `json:"total"`
}
