package services

import (
	"meu-carro-mais/internal/database/datasource"
	"meu-carro-mais/internal/handlers/json"
)

// GetAnuncios retorna todos os anúncios
func GetAnuncios() (*json.AnunciosResponse, error) {
	anuncios, err := datasource.GetAnuncios()
	if err != nil {
		return nil, err
	}

	var anunciosResponse []json.AnuncioResponse
	for _, anuncio := range anuncios {
		anuncioResp := json.AnuncioResponse{
			ID:          anuncio.ID,
			Titulo:      anuncio.Titulo,
			Descricao:   anuncio.Descricao,
			Preco:       anuncio.Preco,
			Imagem:      anuncio.Imagem,
			Destaque:    anuncio.Destaque,
			IDLoja:      anuncio.IDLoja,
			IDCategoria: anuncio.IDCategoria,
			Categoria:   anuncio.Categoria.Nome,
			Loja: json.LojaResponse{
				ID:          anuncio.Loja.ID,
				Nome:        anuncio.Loja.Nome,
				CNPJ:        anuncio.Loja.CNPJ,
				Imagem:      anuncio.Loja.Imagem,
				Latitude:    anuncio.Loja.Latitude,
				Longitude:   anuncio.Loja.Longitude,
				IDCategoria: anuncio.Loja.IDCategoria,
				Categoria:   anuncio.Loja.Categoria.Nome,
			},
		}
		anunciosResponse = append(anunciosResponse, anuncioResp)
	}

	response := &json.AnunciosResponse{
		Anuncios: anunciosResponse,
		Total:    len(anunciosResponse),
	}

	return response, nil
}

// GetCategoriasAnuncio retorna todas as categorias de anúncio
func GetCategoriasAnuncio() (*json.CategoriasAnuncioResponse, error) {
	categorias, err := datasource.GetCategoriasAnuncio()
	if err != nil {
		return nil, err
	}

	var categoriasResponse []json.CategoriaAnuncioResponse
	for _, categoria := range categorias {
		categoriaResp := json.CategoriaAnuncioResponse{
			ID:   categoria.ID,
			Nome: categoria.Nome,
		}
		categoriasResponse = append(categoriasResponse, categoriaResp)
	}

	response := &json.CategoriasAnuncioResponse{
		Categorias: categoriasResponse,
		Total:      len(categoriasResponse),
	}

	return response, nil
}

// CreateAnuncio cria um novo anúncio
func CreateAnuncio(req json.AnuncioRequest) (*json.AnuncioResponse, error) {
	anuncio, err := datasource.CreateAnuncio(req)
	if err != nil {
		return nil, err
	}

	response := &json.AnuncioResponse{
		ID:          anuncio.ID,
		Titulo:      anuncio.Titulo,
		Descricao:   anuncio.Descricao,
		Preco:       anuncio.Preco,
		Imagem:      anuncio.Imagem,
		Destaque:    anuncio.Destaque,
		IDLoja:      anuncio.IDLoja,
		IDCategoria: anuncio.IDCategoria,
		Categoria:   anuncio.Categoria.Nome,
		Loja: json.LojaResponse{
			ID:          anuncio.Loja.ID,
			Nome:        anuncio.Loja.Nome,
			CNPJ:        anuncio.Loja.CNPJ,
			Imagem:      anuncio.Loja.Imagem,
			Latitude:    anuncio.Loja.Latitude,
			Longitude:   anuncio.Loja.Longitude,
			IDCategoria: anuncio.Loja.IDCategoria,
			Categoria:   anuncio.Loja.Categoria.Nome,
		},
	}

	return response, nil
}

// GetAnuncioByID busca um anúncio por ID
func GetAnuncioByID(id uint) (*json.AnuncioResponse, error) {
	anuncio, err := datasource.GetAnuncioByID(id)
	if err != nil {
		return nil, err
	}

	response := &json.AnuncioResponse{
		ID:          anuncio.ID,
		Titulo:      anuncio.Titulo,
		Descricao:   anuncio.Descricao,
		Preco:       anuncio.Preco,
		Imagem:      anuncio.Imagem,
		Destaque:    anuncio.Destaque,
		IDLoja:      anuncio.IDLoja,
		IDCategoria: anuncio.IDCategoria,
		Categoria:   anuncio.Categoria.Nome,
		Loja: json.LojaResponse{
			ID:          anuncio.Loja.ID,
			Nome:        anuncio.Loja.Nome,
			CNPJ:        anuncio.Loja.CNPJ,
			Imagem:      anuncio.Loja.Imagem,
			Latitude:    anuncio.Loja.Latitude,
			Longitude:   anuncio.Loja.Longitude,
			IDCategoria: anuncio.Loja.IDCategoria,
			Categoria:   anuncio.Loja.Categoria.Nome,
		},
	}

	return response, nil
}

// GetAllAnuncios retorna todos os anúncios ativos
func GetAllAnuncios() ([]json.AnuncioResponse, error) {
	anuncios, err := datasource.GetAllAnuncios()
	if err != nil {
		return nil, err
	}

	var responses []json.AnuncioResponse
	for _, anuncio := range anuncios {
		response := json.AnuncioResponse{
			ID:          anuncio.ID,
			Titulo:      anuncio.Titulo,
			Descricao:   anuncio.Descricao,
			Preco:       anuncio.Preco,
			Imagem:      anuncio.Imagem,
			Destaque:    anuncio.Destaque,
			IDLoja:      anuncio.IDLoja,
			IDCategoria: anuncio.IDCategoria,
			Categoria:   anuncio.Categoria.Nome,
			Loja: json.LojaResponse{
				ID:          anuncio.Loja.ID,
				Nome:        anuncio.Loja.Nome,
				CNPJ:        anuncio.Loja.CNPJ,
				Imagem:      anuncio.Loja.Imagem,
				Latitude:    anuncio.Loja.Latitude,
				Longitude:   anuncio.Loja.Longitude,
				IDCategoria: anuncio.Loja.IDCategoria,
				Categoria:   anuncio.Loja.Categoria.Nome,
			},
		}
		responses = append(responses, response)
	}

	return responses, nil
}

// UpdateAnuncio atualiza um anúncio existente
func UpdateAnuncio(id uint, req json.AnuncioRequest) (*json.AnuncioResponse, error) {
	anuncio, err := datasource.UpdateAnuncio(id, req)
	if err != nil {
		return nil, err
	}

	response := &json.AnuncioResponse{
		ID:          anuncio.ID,
		Titulo:      anuncio.Titulo,
		Descricao:   anuncio.Descricao,
		Preco:       anuncio.Preco,
		Imagem:      anuncio.Imagem,
		Destaque:    anuncio.Destaque,
		IDLoja:      anuncio.IDLoja,
		IDCategoria: anuncio.IDCategoria,
		Categoria:   anuncio.Categoria.Nome,
		Loja: json.LojaResponse{
			ID:          anuncio.Loja.ID,
			Nome:        anuncio.Loja.Nome,
			CNPJ:        anuncio.Loja.CNPJ,
			Imagem:      anuncio.Loja.Imagem,
			Latitude:    anuncio.Loja.Latitude,
			Longitude:   anuncio.Loja.Longitude,
			IDCategoria: anuncio.Loja.IDCategoria,
			Categoria:   anuncio.Loja.Categoria.Nome,
		},
	}

	return response, nil
}

// SoftDeleteAnuncio realiza soft delete do anúncio
func SoftDeleteAnuncio(id uint) error {
	return datasource.SoftDeleteAnuncio(id)
}

// RestoreAnuncio restaura um anúncio que foi soft deleted
func RestoreAnuncio(id uint) error {
	return datasource.RestoreAnuncio(id)
}
