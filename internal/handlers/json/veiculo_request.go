package json

type VeiculoRequest struct {
	Modelo    string `json:"modelo" binding:"required"`
	Ano       int    `json:"ano" binding:"required,min=1900,max=2030"`
	Cor       string `json:"cor" binding:"required"`
	Placa     string `json:"placa" binding:"required"`
	IDUsuario uint   `json:"id_usuario" binding:"required"`
}
