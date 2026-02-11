package json

import "time"

// VeiculoFotoResponse representa uma foto do veículo na resposta
type VeiculoFotoResponse struct {
	ID          uint      `json:"id"`
	URL         string    `json:"url"`
	NomeArquivo string    `json:"nome_arquivo"`
	Tamanho     int64     `json:"tamanho"`
	TipoMime    string    `json:"tipo_mime"`
	Principal   bool      `json:"principal"`
	Ordem       int       `json:"ordem"`
	DataUpload  time.Time `json:"data_upload"`
}

type VeiculoResponse struct {
	ID                  uint                  `json:"id"`
	Marca               string                `json:"marca"`
	Modelo              string                `json:"modelo"`
	AnoFabricacao       int                   `json:"ano_fabricacao"`
	AnoModelo           int                   `json:"ano_modelo"`
	Cor                 string                `json:"cor"`
	Placa               string                `json:"placa"`
	Renavam             *string               `json:"renavam,omitempty"`
	Chassi              *string               `json:"chassi,omitempty"`
	TipoVeiculo         *string               `json:"tipo_veiculo,omitempty"`
	Combustivel         *string               `json:"combustivel,omitempty"`
	Quilometragem       *int                  `json:"quilometragem,omitempty"`
	Preco               *float64              `json:"preco,omitempty"`
	Licenciamento       *string               `json:"licenciamento,omitempty"`
	IPVAPago            *bool                 `json:"ipva_pago,omitempty"`
	PossuiFinanciamento *bool                 `json:"possui_financiamento,omitempty"`
	PossuiMultas        *bool                 `json:"possui_multas,omitempty"`
	Observacoes         *string               `json:"observacoes,omitempty"`
	Imagem              string                `json:"imagem,omitempty"`
	Fotos               []VeiculoFotoResponse `json:"fotos,omitempty"`
	IDUsuario           uint                  `json:"id_usuario"`
	IDCupom             *uint                 `json:"id_cupom,omitempty"`
	DataCadastro        time.Time             `json:"data_cadastro"`
	Ativo               bool                  `json:"ativo"`
}

type VeiculosResponse struct {
	Veiculos []VeiculoResponse `json:"veiculos"`
	Total    int               `json:"total"`
}

type HistoricoVeiculoResponse struct {
	ID           uint      `json:"id"`
	IDVeiculo    uint      `json:"id_veiculo"`
	IDCupom      uint      `json:"id_cupom"`
	Descricao    string    `json:"descricao"`
	Data         time.Time `json:"data"`
	DataCadastro time.Time `json:"data_cadastro"`
}

type HistoricosVeiculoResponse struct {
	Historicos []HistoricoVeiculoResponse `json:"historicos"`
	Total      int                        `json:"total"`
}
