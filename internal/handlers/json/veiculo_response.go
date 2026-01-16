package json

import "time"

type VeiculoResponse struct {
	ID                  uint      `json:"id"`
	Marca               string    `json:"marca"`                          // Marca do veículo
	Modelo              string    `json:"modelo"`                          // Modelo do veículo
	AnoFabricacao       int       `json:"ano_fabricacao"`                 // Ano de fabricação
	AnoModelo           int       `json:"ano_modelo"`                     // Ano modelo
	Cor                 string    `json:"cor"`                              // Cor do veículo
	Placa               string    `json:"placa"`                            // Placa do veículo
	Renavam             *string   `json:"renavam,omitempty"`                // RENAVAM do veículo
	Chassi              *string   `json:"chassi,omitempty"`                 // Chassi do veículo
	TipoVeiculo         *string   `json:"tipo_veiculo,omitempty"`          // Tipo do veículo
	Combustivel         *string   `json:"combustivel,omitempty"`            // Tipo de combustível
	Quilometragem       *int      `json:"quilometragem,omitempty"`          // KM do veículo
	Preco               *float64  `json:"preco,omitempty"`                 // Preço do veículo
	Licenciamento       *string   `json:"licenciamento,omitempty"`          // Status do licenciamento
	IPVAPago            *bool     `json:"ipva_pago,omitempty"`              // Se o IPVA está pago
	PossuiFinanciamento *bool     `json:"possui_financiamento,omitempty"`   // Se possui financiamento
	PossuiMultas        *bool     `json:"possui_multas,omitempty"`         // Se possui multas
	Observacoes         *string   `json:"observacoes,omitempty"`           // Observações do veículo
	Imagem              string    `json:"imagem,omitempty"`                 // URL da imagem principal do veículo
	IDUsuario           uint      `json:"id_usuario"`
	DataCadastro        time.Time `json:"data_cadastro"`
	Ativo               bool      `json:"ativo"`
}

type VeiculosResponse struct {
	Veiculos []VeiculoResponse `json:"veiculos"`
	Total    int               `json:"total"`
}

type HistoricoVeiculoResponse struct {
	ID           uint      `json:"id"`
	IDVeiculo    uint      `json:"id_veiculo"`
	IDAnuncio    uint      `json:"id_anuncio"`
	Descricao    string    `json:"descricao"`
	Data         time.Time `json:"data"`
	DataCadastro time.Time `json:"data_cadastro"`
}

type HistoricosVeiculoResponse struct {
	Historicos []HistoricoVeiculoResponse `json:"historicos"`
	Total      int                        `json:"total"`
}
