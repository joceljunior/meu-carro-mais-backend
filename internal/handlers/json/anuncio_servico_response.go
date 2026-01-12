package json

type AnuncioServicoResponse struct {
	ID                  uint     `json:"id"`
	NomeServico         string   `json:"nome_servico"`
	NomeLoja            string   `json:"nome_loja"`
	Imagem              string   `json:"imagem"`
	PrecoOriginal       float64  `json:"preco_original"`
	PrecoComDesconto    float64  `json:"preco_com_desconto"`
	PorcentagemDesconto float64  `json:"porcentagem_desconto"`
	IsMeuCarroMais      bool     `json:"is_meu_carro_mais"`
	Categoria           string   `json:"categoria"`
	Descricao           string   `json:"descricao"`
	Rate                int      `json:"rate"`
	MoedasUtiliza       *int     `json:"moedas_utiliza,omitempty"`
	Distancia           *float64 `json:"distancia,omitempty"` // Distância em km da loja, se fornecida localização
}

type AnunciosServicoResponse struct {
	Anuncios []AnuncioServicoResponse `json:"anuncios"`
	Total    int                       `json:"total"`
}
