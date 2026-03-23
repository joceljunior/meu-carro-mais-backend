package services

import (
	"meu-carro-mais/internal/database/datasource"
	"meu-carro-mais/internal/database/models"
	"meu-carro-mais/internal/handlers/json"
)

// LojaUsuarioResponseComCupom monta dados da loja para login/perfil, incluindo cupom em destaque quando existir.
func LojaUsuarioResponseComCupom(loja models.Loja) json.LojaUsuarioResponse {
	var cupom *json.CupomDestaqueResponse
	if loja.ID != 0 {
		if c, err := datasource.GetCupomDestaqueByLojaID(loja.ID); err == nil {
			cupom = &json.CupomDestaqueResponse{
				ID:        c.ID,
				Titulo:    c.Titulo,
				Descricao: c.Descricao,
				Preco:     c.Preco,
				Imagem:    c.Imagem,
				TipoCupom: c.TipoCupom,
			}
		}
	}
	return json.LojaUsuarioResponseFromModel(loja, cupom)
}
