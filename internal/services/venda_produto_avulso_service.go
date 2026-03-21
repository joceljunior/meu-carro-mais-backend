package services

import (
	"errors"

	"meu-carro-mais/internal/database"
	"meu-carro-mais/internal/database/datasource"
	"meu-carro-mais/internal/database/models"
	"meu-carro-mais/internal/handlers/json"

	"gorm.io/gorm"
)

// VendaProdutoAvulsoModelToResponse converte o model para JSON.
func VendaProdutoAvulsoModelToResponse(v *models.VendaProdutoAvulso) json.VendaProdutoAvulsoResponse {
	return json.VendaProdutoAvulsoResponse{
		ID:               v.ID,
		IDUsuario:        v.IDUsuario,
		IDLoja:           v.IDLoja,
		Valor:            v.Valor,
		DescricaoProduto: v.DescricaoProduto,
		DataVenda:        v.DataVenda,
		Usuario:          usuarioModelToUserResponse(v.Usuario),
		Loja:             json.LojaFromModel(v.Loja),
	}
}

// VendasProdutoAvulsoModelsToResponses converte uma lista de models.
func VendasProdutoAvulsoModelsToResponses(vendas []models.VendaProdutoAvulso) []json.VendaProdutoAvulsoResponse {
	out := make([]json.VendaProdutoAvulsoResponse, 0, len(vendas))
	for i := range vendas {
		out = append(out, VendaProdutoAvulsoModelToResponse(&vendas[i]))
	}
	return out
}

// CreateVendaProdutoAvulso registra venda de produto não cadastrado; resolve o cliente pelo email e credita moedas por loja.
func CreateVendaProdutoAvulso(idLoja uint, req json.VendaProdutoAvulsoRequest) (*json.VendaProdutoAvulsoResponse, error) {
	usuario, err := datasource.GetUserByEmailOnly(req.EmailCliente)
	if err != nil {
		return nil, errors.New("cliente não encontrado para o email informado")
	}
	loja, err := datasource.GetLojaByID(idLoja)
	if err != nil {
		return nil, errors.New("loja não encontrada")
	}

	var v *models.VendaProdutoAvulso
	err = database.DB.Transaction(func(tx *gorm.DB) error {
		var e error
		v, e = datasource.CreateVendaProdutoAvulsoTx(tx, idLoja, usuario.ID, req.Valor, req.DescricaoProduto)
		if e != nil {
			return e
		}
		return AplicarMoedasCreditoVendaAvulsaTx(tx, usuario.ID, idLoja, req.Valor, loja.DescontoGeralPorcentagem)
	})
	if err != nil {
		return nil, err
	}

	r := VendaProdutoAvulsoModelToResponse(v)
	return &r, nil
}
