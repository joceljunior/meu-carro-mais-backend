package json

import "time"

type VeiculoResponse struct {
	ID           uint      `json:"id"`
	Modelo       string    `json:"modelo"`
	Ano          int       `json:"ano"`
	Cor          string    `json:"cor"`
	Placa        string    `json:"placa"`
	IDUsuario    uint      `json:"id_usuario"`
	DataCadastro time.Time `json:"data_cadastro"`
	Ativo        bool      `json:"ativo"`
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
