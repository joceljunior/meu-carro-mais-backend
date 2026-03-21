package services

import (
	"meu-carro-mais/internal/database/datasource"
	"meu-carro-mais/internal/handlers/json"
)

// MoedasUsuarioParaJSON retorna moedas gerais (carteira) e lista de moedas por loja (saldo > 0).
func MoedasUsuarioParaJSON(usuarioID uint) (gerais int, porLoja []json.MoedaLojaUsuarioItem, err error) {
	carteira, err := datasource.GetCarteiraByUsuarioID(usuarioID)
	if err == nil && carteira != nil {
		gerais = carteira.SaldoGeral
	}

	rows, err := datasource.GetUsuarioMoedasLojaByUsuarioID(usuarioID)
	if err != nil {
		return 0, nil, err
	}

	porLoja = make([]json.MoedaLojaUsuarioItem, 0, len(rows))
	for _, r := range rows {
		if r.Saldo <= 0 {
			continue
		}
		nome := ""
		if r.Loja.ID != 0 {
			nome = r.Loja.Nome
		}
		porLoja = append(porLoja, json.MoedaLojaUsuarioItem{
			IDLoja:   r.LojaID,
			NomeLoja: nome,
			Saldo:    r.Saldo,
		})
	}
	return gerais, porLoja, nil
}
