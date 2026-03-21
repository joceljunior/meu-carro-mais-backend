package json

import "time"

// VendaProdutoAvulsoRequest body para registrar venda de produto não cadastrado na loja.
type VendaProdutoAvulsoRequest struct {
	EmailCliente     string  `json:"email_cliente" binding:"required,email"`
	Valor            float64 `json:"valor" binding:"required,gt=0"`
	DescricaoProduto string  `json:"descricao_produto" binding:"required,min=1,max=500"`
}

// VendaProdutoAvulsoResponse dados da venda avulsa para exibição (ex.: histórico).
type VendaProdutoAvulsoResponse struct {
	ID               uint         `json:"id"`
	IDUsuario        uint         `json:"id_usuario"`
	IDLoja           uint         `json:"id_loja"`
	Valor            float64      `json:"valor"`
	DescricaoProduto string       `json:"descricao_produto"`
	DataVenda        time.Time    `json:"data_venda"`
	Usuario          UserResponse `json:"usuario"`
	Loja             LojaResponse `json:"loja"`
}
