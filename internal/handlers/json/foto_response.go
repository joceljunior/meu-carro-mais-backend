package json

import "time"

type FotoResponse struct {
	ID              uint      `json:"id"`
	IDVeiculo       *uint     `json:"id_veiculo,omitempty"`
	IDVeiculoLoja   *uint     `json:"id_veiculo_loja,omitempty"`
	IDProduto       *uint     `json:"id_produto,omitempty"`
	IDServico       *uint     `json:"id_servico,omitempty"`
	IDLoja          *uint     `json:"id_loja,omitempty"`
	TipoEntidade    string    `json:"tipo_entidade"`
	URL             string    `json:"url"`
	NomeArquivo     string    `json:"nome_arquivo"`
	Tamanho         int64     `json:"tamanho"`
	TipoMime        string    `json:"tipo_mime"`
	Principal       bool      `json:"principal"`
	Ordem           int       `json:"ordem"`
	DataUpload      time.Time `json:"data_upload"`
	DataAtualizacao time.Time `json:"data_atualizacao"`

	// Dados relacionados (opcionais)
	Veiculo     *VeiculoResponse     `json:"veiculo,omitempty"`
	VeiculoLoja *VeiculoLojaResponse `json:"veiculo_loja,omitempty"`
	Produto     *ProdutoResponse     `json:"produto,omitempty"`
	Servico     *ServicoResponse     `json:"servico,omitempty"`
	Loja        *LojaResponse        `json:"loja,omitempty"`
}

type FotosResponse struct {
	Fotos []FotoResponse `json:"fotos"`
	Total int            `json:"total"`
}

type FotoUploadResponse struct {
	Foto     FotoResponse `json:"foto"`
	Mensagem string       `json:"mensagem"`
}
