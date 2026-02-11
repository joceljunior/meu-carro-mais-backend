package services

import (
	"meu-carro-mais/internal/database/datasource"
	"meu-carro-mais/internal/database/models"
	"meu-carro-mais/internal/handlers/json"
)

// getCupomDestaqueFromLoja extrai o cupom destaque de uma loja
func getCupomDestaqueFromLoja(loja models.Loja) *json.CupomDestaqueResponse {
	if len(loja.Cupons) == 0 {
		return nil
	}

	// Busca o primeiro cupom destaque (deve haver apenas um por loja)
	for _, cupom := range loja.Cupons {
		if cupom.Destaque {
			return &json.CupomDestaqueResponse{
				ID:        cupom.ID,
				Titulo:    cupom.Titulo,
				Descricao: cupom.Descricao,
				Preco:     cupom.Preco,
				Imagem:    cupom.Imagem,
				TipoCupom: cupom.TipoCupom,
			}
		}
	}

	return nil
}

// GetLojasByProximidade busca lojas ordenadas por proximidade do usuário
func GetLojasByProximidade(latitude, longitude float64) (*json.LojasResponse, error) {
	lojasComDistancia, err := datasource.GetLojasByProximidade(latitude, longitude)
	if err != nil {
		return nil, err
	}

	var lojasResponse []json.LojaResponse
	for _, lojaComDist := range lojasComDistancia {
		lojaResp := json.LojaResponse{
			ID:              lojaComDist.ID,
			Nome:            lojaComDist.Nome,
			CNPJ:            lojaComDist.CNPJ,
			Imagem:          lojaComDist.Imagem,
			Endereco:        lojaComDist.Endereco,
			Latitude:        lojaComDist.Latitude,
			Longitude:       lojaComDist.Longitude,
			Rating:          lojaComDist.Rating,
			IsMeuCarroMais:  lojaComDist.IsMeuCarroMais,
			Categoria:       lojaComDist.Categoria,
			IDUsuario:       lojaComDist.IDUsuario,
			CupomDestaque: getCupomDestaqueFromLoja(lojaComDist.Loja),
		}
		lojasResponse = append(lojasResponse, lojaResp)
	}

	return &json.LojasResponse{
		Lojas: lojasResponse,
		Total: len(lojasResponse),
	}, nil
}

// GetCategoriasLojista retorna todas as categorias de lojista
func GetCategoriasLojista() (*json.CategoriasLojistaResponse, error) {
	categorias, err := datasource.GetCategoriasLojista()
	if err != nil {
		return nil, err
	}

	var categoriasResponse []json.CategoriaLojistaResponse
	for _, categoria := range categorias {
		categoriaResp := json.CategoriaLojistaResponse{
			ID:   categoria.ID,
			Nome: categoria.Nome,
		}
		categoriasResponse = append(categoriasResponse, categoriaResp)
	}

	response := &json.CategoriasLojistaResponse{
		Categorias: categoriasResponse,
		Total:      len(categoriasResponse),
	}

	return response, nil
}

// CreateLoja cria uma nova loja
func CreateLoja(req json.LojaRequest) (*json.LojaResponse, error) {
	loja, err := datasource.CreateLoja(req)
	if err != nil {
		return nil, err
	}

	response := &json.LojaResponse{
		ID:              loja.ID,
		Nome:            loja.Nome,
		CNPJ:            loja.CNPJ,
		Imagem:          loja.Imagem,
		Endereco:        loja.Endereco,
		Latitude:        loja.Latitude,
		Longitude:       loja.Longitude,
		Rating:          loja.Rating,
		IsMeuCarroMais:  loja.IsMeuCarroMais,
		Categoria:       loja.Categoria,
		IDUsuario:       loja.IDUsuario,
		CupomDestaque: getCupomDestaqueFromLoja(*loja),
	}

	return response, nil
}

// GetLojaByID busca uma loja por ID
func GetLojaByID(id uint) (*json.LojaResponse, error) {
	loja, err := datasource.GetLojaByID(id)
	if err != nil {
		return nil, err
	}

	// Monta a resposta do usuário indicador (se existir)
	var usuarioIndicadorResponse *json.UsuarioIndicadorResponse
	if loja.UsuarioIndicador != nil && loja.UsuarioIndicador.ID != 0 {
		usuarioIndicadorResponse = &json.UsuarioIndicadorResponse{
			ID:     loja.UsuarioIndicador.ID,
			Nome:   loja.UsuarioIndicador.Nome,
			Email:  loja.UsuarioIndicador.Email,
			Imagem: loja.UsuarioIndicador.Imagem,
		}
	}

	response := &json.LojaResponse{
		ID:                 loja.ID,
		Nome:               loja.Nome,
		CNPJ:               loja.CNPJ,
		Imagem:             loja.Imagem,
		Endereco:           loja.Endereco,
		Latitude:           loja.Latitude,
		Longitude:          loja.Longitude,
		Rating:             loja.Rating,
		IsMeuCarroMais:     loja.IsMeuCarroMais,
		Categoria:          loja.Categoria,
		IDUsuario:          loja.IDUsuario,
		IDUsuarioIndicador: loja.IDUsuarioIndicador,
		DataVinculoUsuario: loja.DataVinculoUsuario,
		UsuarioIndicador:   usuarioIndicadorResponse,
		CupomDestaque:    getCupomDestaqueFromLoja(*loja),
	}

	return response, nil
}

// GetAllLojas retorna todas as lojas ativas
func GetAllLojas() ([]json.LojaResponse, error) {
	lojas, err := datasource.GetAllLojas()
	if err != nil {
		return nil, err
	}

	var responses []json.LojaResponse
	for _, loja := range lojas {
		response := json.LojaResponse{
			ID:              loja.ID,
			Nome:            loja.Nome,
			CNPJ:            loja.CNPJ,
			Imagem:          loja.Imagem,
			Endereco:        loja.Endereco,
			Latitude:        loja.Latitude,
			Longitude:       loja.Longitude,
			Rating:          loja.Rating,
			IsMeuCarroMais:  loja.IsMeuCarroMais,
			Categoria:       loja.Categoria,
			IDUsuario:       loja.IDUsuario,
			CupomDestaque: getCupomDestaqueFromLoja(loja),
		}
		responses = append(responses, response)
	}

	return responses, nil
}

// UpdateLoja atualiza uma loja existente
func UpdateLoja(id uint, req json.LojaRequest) (*json.LojaResponse, error) {
	loja, err := datasource.UpdateLoja(id, req)
	if err != nil {
		return nil, err
	}

	response := &json.LojaResponse{
		ID:              loja.ID,
		Nome:            loja.Nome,
		CNPJ:            loja.CNPJ,
		Imagem:          loja.Imagem,
		Endereco:        loja.Endereco,
		Latitude:        loja.Latitude,
		Longitude:       loja.Longitude,
		Rating:          loja.Rating,
		IsMeuCarroMais:  loja.IsMeuCarroMais,
		Categoria:       loja.Categoria,
		IDUsuario:       loja.IDUsuario,
		CupomDestaque: getCupomDestaqueFromLoja(*loja),
	}

	return response, nil
}

// SoftDeleteLoja realiza soft delete da loja
func SoftDeleteLoja(id uint) error {
	return datasource.SoftDeleteLoja(id)
}

// RestoreLoja restaura uma loja que foi soft deleted
func RestoreLoja(id uint) error {
	return datasource.RestoreLoja(id)
}

// GetLojasByUsuarioID retorna todas as lojas de um usuário
func GetLojasByUsuarioID(idUsuario uint) ([]json.LojaResponse, error) {
	lojas, err := datasource.GetLojasByUsuarioID(idUsuario)
	if err != nil {
		return nil, err
	}

	var responses []json.LojaResponse
	for _, loja := range lojas {
		response := json.LojaResponse{
			ID:              loja.ID,
			Nome:            loja.Nome,
			CNPJ:            loja.CNPJ,
			Imagem:          loja.Imagem,
			Endereco:        loja.Endereco,
			Latitude:        loja.Latitude,
			Longitude:       loja.Longitude,
			Rating:          loja.Rating,
			IsMeuCarroMais:  loja.IsMeuCarroMais,
			Categoria:       loja.Categoria,
			IDUsuario:       loja.IDUsuario,
			CupomDestaque: getCupomDestaqueFromLoja(loja),
		}
		responses = append(responses, response)
	}

	return responses, nil
}
