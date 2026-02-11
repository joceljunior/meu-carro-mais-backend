package services

import (
	"meu-carro-mais/internal/database/datasource"
	"meu-carro-mais/internal/handlers/json"
)

// getImagemVeiculoLoja busca a imagem principal de um veículo de loja
func getImagemVeiculoLoja(idVeiculoLoja uint) string {
	upload, err := datasource.GetUploadPrincipalByEntidade("veiculo_loja", idVeiculoLoja)
	if err != nil {
		return ""
	}
	return upload.URL
}

// getFotosVeiculoLoja busca todas as fotos de um veículo de loja
func getFotosVeiculoLoja(idVeiculoLoja uint) []json.VeiculoFotoResponse {
	uploads, err := datasource.GetUploadsByVeiculoLojaID(idVeiculoLoja)
	if err != nil {
		return nil
	}

	var fotos []json.VeiculoFotoResponse
	for _, upload := range uploads {
		if upload.Tipo == "Imagem" {
			foto := json.VeiculoFotoResponse{
				ID:          upload.ID,
				URL:         upload.URL,
				NomeArquivo: upload.NomeArquivo,
				Tamanho:     upload.Tamanho,
				TipoMime:    upload.TipoMime,
				Principal:   upload.Principal,
				Ordem:       upload.Ordem,
				DataUpload:  upload.DataUpload,
			}
			fotos = append(fotos, foto)
		}
	}

	return fotos
}

// CreateVeiculoLoja cria um novo veículo de loja
func CreateVeiculoLoja(req json.VeiculoLojaRequest) (*json.VeiculoLojaResponse, error) {
	veiculo, err := datasource.CreateVeiculoLoja(req)
	if err != nil {
		return nil, err
	}

	response := &json.VeiculoLojaResponse{
		ID:           veiculo.ID,
		Modelo:       veiculo.Modelo,
		Ano:          veiculo.Ano,
		Cor:          veiculo.Cor,
		Placa:        veiculo.Placa,
		IDLoja:       veiculo.IDLoja,
		Imagem:       getImagemVeiculoLoja(veiculo.ID),
		Fotos:        getFotosVeiculoLoja(veiculo.ID),
		DataCadastro: veiculo.DataCadastro,
		Ativo:        veiculo.Ativo,
		Loja: json.LojaResponse{
			ID:          veiculo.Loja.ID,
			Nome:        veiculo.Loja.Nome,
			CNPJ:        veiculo.Loja.CNPJ,
			Imagem:      veiculo.Loja.Imagem,
			Latitude:    veiculo.Loja.Latitude,
			Longitude:   veiculo.Loja.Longitude,
			Categoria: veiculo.Loja.Categoria,
			IDUsuario: veiculo.Loja.IDUsuario,
		},
	}

	return response, nil
}

// GetVeiculoLojaByID busca um veículo de loja por ID
func GetVeiculoLojaByID(id uint) (*json.VeiculoLojaResponse, error) {
	veiculo, err := datasource.GetVeiculoLojaByID(id)
	if err != nil {
		return nil, err
	}

	response := &json.VeiculoLojaResponse{
		ID:           veiculo.ID,
		Modelo:       veiculo.Modelo,
		Ano:          veiculo.Ano,
		Cor:          veiculo.Cor,
		Placa:        veiculo.Placa,
		IDLoja:       veiculo.IDLoja,
		Imagem:       getImagemVeiculoLoja(veiculo.ID),
		Fotos:        getFotosVeiculoLoja(veiculo.ID),
		DataCadastro: veiculo.DataCadastro,
		Ativo:        veiculo.Ativo,
		Loja: json.LojaResponse{
			ID:          veiculo.Loja.ID,
			Nome:        veiculo.Loja.Nome,
			CNPJ:        veiculo.Loja.CNPJ,
			Imagem:      veiculo.Loja.Imagem,
			Latitude:    veiculo.Loja.Latitude,
			Longitude:   veiculo.Loja.Longitude,
			Categoria: veiculo.Loja.Categoria,
			IDUsuario: veiculo.Loja.IDUsuario,
		},
	}

	return response, nil
}

// GetAllVeiculosLoja retorna todos os veículos de loja ativos
func GetAllVeiculosLoja() ([]json.VeiculoLojaResponse, error) {
	veiculos, err := datasource.GetAllVeiculosLoja()
	if err != nil {
		return nil, err
	}

	var responses []json.VeiculoLojaResponse
	for _, veiculo := range veiculos {
		response := json.VeiculoLojaResponse{
			ID:           veiculo.ID,
			Modelo:       veiculo.Modelo,
			Ano:          veiculo.Ano,
			Cor:          veiculo.Cor,
			Placa:        veiculo.Placa,
			IDLoja:       veiculo.IDLoja,
			Imagem:       getImagemVeiculoLoja(veiculo.ID),
			Fotos:        getFotosVeiculoLoja(veiculo.ID),
			DataCadastro: veiculo.DataCadastro,
			Ativo:        veiculo.Ativo,
			Loja: json.LojaResponse{
				ID:          veiculo.Loja.ID,
				Nome:        veiculo.Loja.Nome,
				CNPJ:        veiculo.Loja.CNPJ,
				Imagem:      veiculo.Loja.Imagem,
				Latitude:    veiculo.Loja.Latitude,
				Longitude:   veiculo.Loja.Longitude,
				Categoria: veiculo.Loja.Categoria,
				IDUsuario: veiculo.Loja.IDUsuario,
			},
		}
		responses = append(responses, response)
	}

	return responses, nil
}

// GetVeiculosLojaByLojaID retorna todos os veículos de uma loja específica
func GetVeiculosLojaByLojaID(idLoja uint) (*json.VeiculosLojaResponse, error) {
	veiculos, err := datasource.GetVeiculosLojaByLojaID(idLoja)
	if err != nil {
		return nil, err
	}

	var veiculosResponse []json.VeiculoLojaResponse
	for _, veiculo := range veiculos {
		veiculoResp := json.VeiculoLojaResponse{
			ID:           veiculo.ID,
			Modelo:       veiculo.Modelo,
			Ano:          veiculo.Ano,
			Cor:          veiculo.Cor,
			Placa:        veiculo.Placa,
			IDLoja:       veiculo.IDLoja,
			Imagem:       getImagemVeiculoLoja(veiculo.ID),
			Fotos:        getFotosVeiculoLoja(veiculo.ID),
			DataCadastro: veiculo.DataCadastro,
			Ativo:        veiculo.Ativo,
			Loja: json.LojaResponse{
				ID:          veiculo.Loja.ID,
				Nome:        veiculo.Loja.Nome,
				CNPJ:        veiculo.Loja.CNPJ,
				Imagem:      veiculo.Loja.Imagem,
				Latitude:    veiculo.Loja.Latitude,
				Longitude:   veiculo.Loja.Longitude,
				Categoria: veiculo.Loja.Categoria,
				IDUsuario: veiculo.Loja.IDUsuario,
			},
		}
		veiculosResponse = append(veiculosResponse, veiculoResp)
	}

	response := &json.VeiculosLojaResponse{
		Veiculos: veiculosResponse,
		Total:    len(veiculosResponse),
	}

	return response, nil
}

// UpdateVeiculoLoja atualiza um veículo de loja existente
func UpdateVeiculoLoja(id uint, req json.VeiculoLojaRequest) (*json.VeiculoLojaResponse, error) {
	veiculo, err := datasource.UpdateVeiculoLoja(id, req)
	if err != nil {
		return nil, err
	}

	response := &json.VeiculoLojaResponse{
		ID:           veiculo.ID,
		Modelo:       veiculo.Modelo,
		Ano:          veiculo.Ano,
		Cor:          veiculo.Cor,
		Placa:        veiculo.Placa,
		IDLoja:       veiculo.IDLoja,
		Imagem:       getImagemVeiculoLoja(veiculo.ID),
		Fotos:        getFotosVeiculoLoja(veiculo.ID),
		DataCadastro: veiculo.DataCadastro,
		Ativo:        veiculo.Ativo,
		Loja: json.LojaResponse{
			ID:          veiculo.Loja.ID,
			Nome:        veiculo.Loja.Nome,
			CNPJ:        veiculo.Loja.CNPJ,
			Imagem:      veiculo.Loja.Imagem,
			Latitude:    veiculo.Loja.Latitude,
			Longitude:   veiculo.Loja.Longitude,
			Categoria: veiculo.Loja.Categoria,
			IDUsuario: veiculo.Loja.IDUsuario,
		},
	}

	return response, nil
}

// SoftDeleteVeiculoLoja realiza soft delete do veículo de loja
func SoftDeleteVeiculoLoja(id uint) error {
	return datasource.SoftDeleteVeiculoLoja(id)
}

// RestoreVeiculoLoja restaura um veículo de loja que foi soft deleted
func RestoreVeiculoLoja(id uint) error {
	return datasource.RestoreVeiculoLoja(id)
}
