package services

import (
	"meu-carro-mais/internal/database/datasource"
	"meu-carro-mais/internal/database/models"
	"meu-carro-mais/internal/handlers/json"
)

// modelToOfertaAutoMaisResponse converte o model para response
func modelToOfertaAutoMaisResponse(oferta *models.OfertaAutoMais) *json.OfertaAutoMaisResponse {
	return &json.OfertaAutoMaisResponse{
		ID:              oferta.ID,
		IDLoja:          oferta.IDLoja,
		Nome:            oferta.Nome,
		Descricao:       oferta.Descricao,
		Moedas:          oferta.Moedas,
		Porcentagem:     oferta.Porcentagem,
		Ativo:           oferta.Ativo,
		DataValidade:    oferta.DataValidade,
		DataCadastro:    oferta.DataCadastro,
		DataAtualizacao: oferta.DataAtualizacao,
		Loja: json.LojaResponse{
			ID:             oferta.Loja.ID,
			Nome:           oferta.Loja.Nome,
			CNPJ:           oferta.Loja.CNPJ,
			Imagem:         oferta.Loja.Imagem,
			Latitude:       oferta.Loja.Latitude,
			Longitude:      oferta.Loja.Longitude,
			Rating:         oferta.Loja.Rating,
			IsMeuCarroMais: oferta.Loja.IsMeuCarroMais,
			Categoria:      oferta.Loja.Categoria,
			IDUsuario:      oferta.Loja.IDUsuario,
		},
	}
}

// CreateOfertaAutoMais cria uma nova oferta Auto Mais para uma loja
func CreateOfertaAutoMais(req json.OfertaAutoMaisRequest) (*json.OfertaAutoMaisResponse, error) {
	oferta, err := datasource.CreateOfertaAutoMais(req)
	if err != nil {
		return nil, err
	}

	return modelToOfertaAutoMaisResponse(oferta), nil
}

// GetOfertaAutoMaisByID busca uma oferta Auto Mais por ID
func GetOfertaAutoMaisByID(id uint) (*json.OfertaAutoMaisResponse, error) {
	oferta, err := datasource.GetOfertaAutoMaisByID(id)
	if err != nil {
		return nil, err
	}

	return modelToOfertaAutoMaisResponse(oferta), nil
}

// GetAllOfertasAutoMais retorna todas as ofertas Auto Mais
func GetAllOfertasAutoMais() ([]json.OfertaAutoMaisResponse, error) {
	ofertas, err := datasource.GetAllOfertasAutoMais()
	if err != nil {
		return nil, err
	}

	var responses []json.OfertaAutoMaisResponse
	for _, oferta := range ofertas {
		responses = append(responses, *modelToOfertaAutoMaisResponse(&oferta))
	}

	return responses, nil
}

// GetAllOfertasAutoMaisAtivas retorna todas as ofertas Auto Mais ativas
// Se latitude e longitude forem fornecidos, ordena por proximidade
func GetAllOfertasAutoMaisAtivas(latitude, longitude *float64) ([]json.OfertaAutoMaisResponse, error) {
	var responses []json.OfertaAutoMaisResponse

	if latitude != nil && longitude != nil {
		// Busca ofertas ordenadas por proximidade
		ofertasComDistancia, err := datasource.GetOfertasAutoMaisAtivasByProximidade(*latitude, *longitude)
		if err != nil {
			return nil, err
		}

		for _, ofertaComDistancia := range ofertasComDistancia {
			response := modelToOfertaAutoMaisResponse(&ofertaComDistancia.Oferta)
			distancia := ofertaComDistancia.Distancia
			response.Distancia = &distancia
			responses = append(responses, *response)
		}
	} else {
		// Busca ofertas sem ordenação por proximidade
		ofertas, err := datasource.GetAllOfertasAutoMaisAtivas()
		if err != nil {
			return nil, err
		}

		for _, oferta := range ofertas {
			responses = append(responses, *modelToOfertaAutoMaisResponse(&oferta))
		}
	}

	return responses, nil
}

// GetOfertasAutoMaisByLojaID retorna todas as ofertas Auto Mais de uma loja
func GetOfertasAutoMaisByLojaID(idLoja uint) (*json.OfertasAutoMaisResponse, error) {
	ofertas, err := datasource.GetOfertasAutoMaisByLojaID(idLoja)
	if err != nil {
		return nil, err
	}

	var ofertasResponse []json.OfertaAutoMaisResponse
	for _, oferta := range ofertas {
		ofertasResponse = append(ofertasResponse, *modelToOfertaAutoMaisResponse(&oferta))
	}

	response := &json.OfertasAutoMaisResponse{
		Ofertas: ofertasResponse,
		Total:   len(ofertasResponse),
	}

	return response, nil
}

// GetOfertasAutoMaisAtivasByLojaID retorna apenas as ofertas Auto Mais ativas de uma loja
func GetOfertasAutoMaisAtivasByLojaID(idLoja uint) (*json.OfertasAutoMaisResponse, error) {
	ofertas, err := datasource.GetOfertasAutoMaisAtivasByLojaID(idLoja)
	if err != nil {
		return nil, err
	}

	var ofertasResponse []json.OfertaAutoMaisResponse
	for _, oferta := range ofertas {
		ofertasResponse = append(ofertasResponse, *modelToOfertaAutoMaisResponse(&oferta))
	}

	response := &json.OfertasAutoMaisResponse{
		Ofertas: ofertasResponse,
		Total:   len(ofertasResponse),
	}

	return response, nil
}

// UpdateOfertaAutoMais atualiza uma oferta Auto Mais existente
func UpdateOfertaAutoMais(id uint, req json.OfertaAutoMaisUpdateRequest) (*json.OfertaAutoMaisResponse, error) {
	oferta, err := datasource.UpdateOfertaAutoMais(id, req)
	if err != nil {
		return nil, err
	}

	return modelToOfertaAutoMaisResponse(oferta), nil
}

// DesativarOfertaAutoMais desativa uma oferta Auto Mais
func DesativarOfertaAutoMais(id uint) error {
	return datasource.DesativarOfertaAutoMais(id)
}

// AtivarOfertaAutoMais ativa uma oferta Auto Mais
func AtivarOfertaAutoMais(id uint) error {
	return datasource.AtivarOfertaAutoMais(id)
}

// SoftDeleteOfertaAutoMais realiza soft delete da oferta
func SoftDeleteOfertaAutoMais(id uint) error {
	return datasource.SoftDeleteOfertaAutoMais(id)
}

// RestoreOfertaAutoMais restaura uma oferta que foi soft deleted
func RestoreOfertaAutoMais(id uint) error {
	return datasource.RestoreOfertaAutoMais(id)
}

