package json

type AnuncioVeiculoResponse struct {
	ID                  uint     `json:"id"`
	NomeVeiculo         string   `json:"nome_veiculo"`
	KM                  *int     `json:"km,omitempty"`
	AnoModelo           int      `json:"ano_modelo"`
	AnoFabricacao       *int     `json:"ano_fabricacao,omitempty"`
	IsMeuCarroMais      bool     `json:"is_meu_carro_mais"`
	Preco               float64  `json:"preco"`
	Imagem              string   `json:"imagem"`
	Modelo              string   `json:"modelo"`
	Marca               *string `json:"marca,omitempty"`
	Placa               string   `json:"placa"`
	Renavam             *string  `json:"renavam,omitempty"`
	Chassi              *string  `json:"chassi,omitempty"`
	Cor                 string   `json:"cor"`
	TipoVeiculo         *string  `json:"tipo_veiculo,omitempty"`
	Licenciamento       *string  `json:"licenciamento,omitempty"`
	IPVAPago            *bool    `json:"ipva_pago,omitempty"`
	PossuiFinanciamento *bool    `json:"possui_financiamento,omitempty"`
	PossuiMultas        *bool    `json:"possui_multas,omitempty"`
	Observacoes         *string  `json:"observacoes,omitempty"`
	Combustivel         *string  `json:"combustivel,omitempty"`
	MoedasUtiliza       *int     `json:"moedas_utiliza,omitempty"`
	Distancia           *float64 `json:"distancia,omitempty"` // Distância em km da loja, se fornecida localização
}

type AnunciosVeiculoResponse struct {
	Anuncios []AnuncioVeiculoResponse `json:"anuncios"`
	Total    int                       `json:"total"`
}
