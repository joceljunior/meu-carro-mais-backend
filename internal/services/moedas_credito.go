package services

import (
	"math"

	"meu-carro-mais/internal/database/datasource"
	"meu-carro-mais/internal/database/models"

	"gorm.io/gorm"
)

// CalcularMoedasGanhasPorDescontoGeralLoja aplica metade do percentual de desconto geral da loja sobre o valor total da venda e converte o resultado em moedas (1 real = 1 moeda).
// Ex.: desconto geral 10%, venda R$ 100 → 50% de 10% = 5% sobre R$ 100 → R$ 5 → 5 moedas.
func CalcularMoedasGanhasPorDescontoGeralLoja(valorTotalVendaReais float64, descontoGeralPct float64) int {
	if valorTotalVendaReais <= 0 || descontoGeralPct <= 0 {
		return 0
	}
	pctEfetivo := descontoGeralPct * 0.5
	reaisEmMoedas := valorTotalVendaReais * (pctEfetivo / 100.0)
	return int(math.Floor(reaisEmMoedas))
}

func valorBaseCupomParaDescontoGeral(c *models.Cupom) float64 {
	if c == nil {
		return 0
	}
	switch c.TipoCupom {
	case "produto":
		if c.Produto != nil {
			return c.Produto.Preco
		}
		return c.Preco
	case "servico":
		if c.Servico != nil {
			return c.Servico.Preco
		}
		return c.Preco
	case "veiculo":
		if c.Veiculo != nil && c.Veiculo.Preco != nil {
			return *c.Veiculo.Preco
		}
		return c.Preco
	default:
		return c.Preco
	}
}

func resolverLojaCupom(tx *gorm.DB, c *models.Cupom) (*models.Loja, error) {
	if c == nil || c.IDLoja == nil {
		return nil, nil
	}
	if c.Loja != nil && c.Loja.ID != 0 {
		return c.Loja, nil
	}
	var l models.Loja
	if err := tx.Where("id = ? AND data_exclusao IS NULL", *c.IDLoja).First(&l).Error; err != nil {
		return nil, err
	}
	return &l, nil
}

// AplicarMoedasCreditoResgateEfetivadoTx credita moedas por loja ao efetivar resgate (transação), com base no preço do cupom (valor da transação) e no desconto geral da loja.
func AplicarMoedasCreditoResgateEfetivadoTx(tx *gorm.DB, historicoID uint) error {
	h, err := datasource.GetHistoricoResgateByIDWithDB(tx, historicoID)
	if err != nil {
		return err
	}
	if h.MoedasLojaJaCreditadas {
		return nil
	}
	if h.Cupom == nil || h.Cupom.IDLoja == nil {
		return tx.Model(&models.HistoricoResgate{}).Where("id = ?", historicoID).Update("moedas_loja_ja_creditadas", true).Error
	}

	loja, err := resolverLojaCupom(tx, h.Cupom)
	if err != nil {
		return err
	}
	if loja == nil {
		return tx.Model(&models.HistoricoResgate{}).Where("id = ?", historicoID).Update("moedas_loja_ja_creditadas", true).Error
	}

	valorBase := valorBaseCupomParaDescontoGeral(h.Cupom)
	moedas := CalcularMoedasGanhasPorDescontoGeralLoja(valorBase, loja.DescontoGeralPorcentagem)
	if moedas > 0 {
		if err := datasource.AdicionarMoedasLojaUsuarioTx(tx, h.IDUsuario, *h.Cupom.IDLoja, moedas); err != nil {
			return err
		}
	}
	return tx.Model(&models.HistoricoResgate{}).Where("id = ?", historicoID).Update("moedas_loja_ja_creditadas", true).Error
}

// AplicarMoedasCreditoVendaAvulsaTx credita moedas por loja após venda de produto não cadastrado.
func AplicarMoedasCreditoVendaAvulsaTx(tx *gorm.DB, usuarioID, lojaID uint, valorVenda float64, descontoGeralPct float64) error {
	moedas := CalcularMoedasGanhasPorDescontoGeralLoja(valorVenda, descontoGeralPct)
	if moedas <= 0 {
		return nil
	}
	return datasource.AdicionarMoedasLojaUsuarioTx(tx, usuarioID, lojaID, moedas)
}
