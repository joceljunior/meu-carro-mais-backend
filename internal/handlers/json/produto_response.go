package json

import "time"

type ProdutoResponse struct {
	ID           uint         `json:"id"`
	Nome         string       `json:"nome"`
	Descricao    string       `json:"descricao"`
	Preco        float64      `json:"preco"`
	Imagem       string       `json:"imagem"`
	Estoque      int          `json:"estoque"`
	Ativo        bool         `json:"ativo"`
	Categoria    string       `json:"categoria"`
	IDLoja       uint         `json:"id_loja"`
	DataCadastro time.Time    `json:"data_cadastro"`
	Loja         LojaResponse `json:"loja,omitempty"`
}

type ProdutosResponse struct {
	Produtos []ProdutoResponse `json:"produtos"`
	Total    int               `json:"total"`
}
