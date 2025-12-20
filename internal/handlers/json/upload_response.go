package json

import "time"

type UploadResponse struct {
	ID              uint      `json:"id"`
	IDUsuario       *uint     `json:"id_usuario,omitempty"`
	IDVeiculo       *uint     `json:"id_veiculo,omitempty"`
	IDVeiculoLoja   *uint     `json:"id_veiculo_loja,omitempty"`
	IDProduto       *uint     `json:"id_produto,omitempty"`
	IDServico       *uint     `json:"id_servico,omitempty"`
	IDLoja          *uint     `json:"id_loja,omitempty"`
	TipoEntidade    string    `json:"tipo_entidade"`
	Tipo            string    `json:"tipo"` // "Imagem" ou "Documento"
	URL             string    `json:"url"`
	NomeArquivo     string    `json:"nome_arquivo"`
	Tamanho         int64     `json:"tamanho"`
	TipoMime        string    `json:"tipo_mime"`
	Principal       bool      `json:"principal"` // Apenas para imagens
	Ordem           int       `json:"ordem"`
	DataUpload      time.Time `json:"data_upload"`
	DataAtualizacao time.Time `json:"data_atualizacao"`

	// Dados relacionados (opcionais)
	Usuario   *UserResponse     `json:"usuario,omitempty"`
	Veiculo   *VeiculoResponse  `json:"veiculo,omitempty"`
	VeiculoLoja *VeiculoLojaResponse `json:"veiculo_loja,omitempty"`
	Produto   *ProdutoResponse  `json:"produto,omitempty"`
	Servico   *ServicoResponse  `json:"servico,omitempty"`
	Loja      *LojaResponse     `json:"loja,omitempty"`
}

type UploadsResponse struct {
	Uploads []UploadResponse `json:"uploads"`
	Total   int              `json:"total"`
}

type UploadUploadResponse struct {
	Upload   UploadResponse `json:"upload"`
	Mensagem string         `json:"mensagem"`
}

