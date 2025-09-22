package json

type FotoRequest struct {
	IDVeiculo     *uint  `json:"id_veiculo,omitempty"`
	IDVeiculoLoja *uint  `json:"id_veiculo_loja,omitempty"`
	IDProduto     *uint  `json:"id_produto,omitempty"`
	IDServico     *uint  `json:"id_servico,omitempty"`
	IDLoja        *uint  `json:"id_loja,omitempty"`
	TipoEntidade  string `json:"tipo_entidade" binding:"required,oneof=veiculo veiculo_loja produto servico loja"`
	URL           string `json:"url" binding:"required"`
	NomeArquivo   string `json:"nome_arquivo" binding:"required"`
	Tamanho       int64  `json:"tamanho" binding:"required,min=1"`
	TipoMime      string `json:"tipo_mime" binding:"required"`
	Principal     bool   `json:"principal,omitempty"`
	Ordem         int    `json:"ordem,omitempty"`
}
