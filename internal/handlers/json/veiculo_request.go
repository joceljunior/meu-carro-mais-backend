package json

type VeiculoRequest struct {
	Marca               string   `json:"marca" binding:"required" example:"Toyota"`
	Modelo              string   `json:"modelo" binding:"required" example:"Corolla"`
	AnoFabricacao       int      `json:"ano_fabricacao" binding:"required,min=1900,max=2030" example:"2020"`
	AnoModelo           int      `json:"ano_modelo" binding:"required,min=1900,max=2030" example:"2020"`
	Cor                 string   `json:"cor" binding:"required" example:"Branco"`
	Placa               string   `json:"placa" binding:"required" example:"ABC-1234"`
	Renavam             *string  `json:"renavam,omitempty" example:"12345678901"`
	Chassi              *string  `json:"chassi,omitempty" example:"9BW12345678901234"`
	TipoVeiculo         *string  `json:"tipo_veiculo,omitempty" example:"Carro"`
	Combustivel         *string  `json:"combustivel,omitempty" example:"Flex"`
	Quilometragem       *int     `json:"quilometragem,omitempty" example:"50000"`
	Preco               *float64 `json:"preco,omitempty" example:"45000.00"`
	Licenciamento       *string  `json:"licenciamento,omitempty" example:"Pago"`
	IPVAPago            *bool    `json:"ipva_pago,omitempty" example:"true"`
	PossuiFinanciamento *bool    `json:"possui_financiamento,omitempty" example:"false"`
	PossuiMultas        *bool    `json:"possui_multas,omitempty" example:"false"`
	Observacoes         *string  `json:"observacoes,omitempty" example:"Veículo em ótimo estado"`
	IDUsuario           uint     `json:"id_usuario" binding:"required"`
}
