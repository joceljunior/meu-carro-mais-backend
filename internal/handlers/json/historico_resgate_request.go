package json

type HistoricoResgateRequest struct {
	IDUsuario        uint    `json:"id_usuario" binding:"required"`
	IDProduto        *uint   `json:"id_produto,omitempty"`
	IDServico        *uint   `json:"id_servico,omitempty"`
	IDVeiculo        *uint   `json:"id_veiculo,omitempty"`
	IDLoja           uint    `json:"id_loja" binding:"required"`
	TipoResgate      string  `json:"tipo_resgate" binding:"required,oneof=produto servico veiculo"`
	Quantidade       int     `json:"quantidade,omitempty"`
	ValorUnitario    float64 `json:"valor_unitario,omitempty"`
	ValorOriginal       float64 `json:"valor_original,omitempty"`
	DescontoAplicado    float64 `json:"desconto_aplicado,omitempty"`
	PorcentagemDesconto float64 `json:"porcentagem_desconto,omitempty"`
	Valor            float64 `json:"valor" binding:"required,min=0"`
	Status           string  `json:"status,omitempty" binding:"omitempty,oneof=pendente confirmado cancelado"`
}
