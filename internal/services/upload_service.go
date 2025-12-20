package services

import (
	"meu-carro-mais/internal/database/datasource"
	"meu-carro-mais/internal/database/models"
	"meu-carro-mais/internal/handlers/json"
)

// convertUploadToResponse converte um modelo Upload para UploadResponse
func convertUploadToResponse(upload *models.Upload) *json.UploadResponse {
	response := &json.UploadResponse{
		ID:              upload.ID,
		IDUsuario:       upload.IDUsuario,
		IDVeiculo:       upload.IDVeiculo,
		IDVeiculoLoja:   upload.IDVeiculoLoja,
		IDProduto:       upload.IDProduto,
		IDServico:       upload.IDServico,
		IDLoja:          upload.IDLoja,
		TipoEntidade:    upload.TipoEntidade,
		Tipo:            upload.Tipo,
		URL:             upload.URL,
		NomeArquivo:     upload.NomeArquivo,
		Tamanho:         upload.Tamanho,
		TipoMime:        upload.TipoMime,
		Principal:       upload.Principal,
		Ordem:           upload.Ordem,
		DataUpload:      upload.DataUpload,
		DataAtualizacao: upload.DataAtualizacao,
	}

	// Adiciona dados do usuário se existir
	if upload.Usuario != nil {
		usuarioResp := &json.UserResponse{
			ID:    upload.Usuario.ID,
			Nome:  upload.Usuario.Nome,
			Email: upload.Usuario.Email,
			CPF:   upload.Usuario.CPF,
			Imagem: upload.Usuario.Imagem,
		}
		response.Usuario = usuarioResp
	}

	// Adiciona dados do veículo se existir
	if upload.Veiculo != nil {
		veiculoResp := &json.VeiculoResponse{
			ID:           upload.Veiculo.ID,
			Modelo:       upload.Veiculo.Modelo,
			Ano:          upload.Veiculo.Ano,
			Cor:          upload.Veiculo.Cor,
			Placa:        upload.Veiculo.Placa,
			IDUsuario:    upload.Veiculo.IDUsuario,
			DataCadastro: upload.Veiculo.DataCadastro,
			Ativo:        upload.Veiculo.Ativo,
		}
		response.Veiculo = veiculoResp
	}

	// Adiciona dados do veículo de loja se existir
	if upload.VeiculoLoja != nil {
		veiculoLojaResp := &json.VeiculoLojaResponse{
			ID:           upload.VeiculoLoja.ID,
			Modelo:       upload.VeiculoLoja.Modelo,
			Ano:          upload.VeiculoLoja.Ano,
			Cor:          upload.VeiculoLoja.Cor,
			Placa:        upload.VeiculoLoja.Placa,
			IDLoja:       upload.VeiculoLoja.IDLoja,
			DataCadastro: upload.VeiculoLoja.DataCadastro,
			Ativo:        upload.VeiculoLoja.Ativo,
		}
		response.VeiculoLoja = veiculoLojaResp
	}

	// Adiciona dados do produto se existir
	if upload.Produto != nil {
		produtoResp := &json.ProdutoResponse{
			ID:           upload.Produto.ID,
			Nome:         upload.Produto.Nome,
			Descricao:    upload.Produto.Descricao,
			Preco:        upload.Produto.Preco,
			Imagem:       upload.Produto.Imagem,
			Estoque:      upload.Produto.Estoque,
			Ativo:        upload.Produto.Ativo,
			IDLoja:       upload.Produto.IDLoja,
			DataCadastro: upload.Produto.DataCadastro,
		}
		response.Produto = produtoResp
	}

	// Adiciona dados do serviço se existir
	if upload.Servico != nil {
		servicoResp := &json.ServicoResponse{
			ID:        upload.Servico.ID,
			Titulo:    upload.Servico.Titulo,
			Descricao: upload.Servico.Descricao,
			Preco:     upload.Servico.Preco,
			Imagem:    upload.Servico.Imagem,
			Destaque:  upload.Servico.Destaque,
			Categoria: upload.Servico.Categoria,
		}
		response.Servico = servicoResp
	}

	// Adiciona dados da loja se existir
	if upload.Loja != nil {
		lojaResp := &json.LojaResponse{
			ID:          upload.Loja.ID,
			Nome:        upload.Loja.Nome,
			CNPJ:        upload.Loja.CNPJ,
			Imagem:      upload.Loja.Imagem,
			Latitude:    upload.Loja.Latitude,
			Longitude:   upload.Loja.Longitude,
			Categoria: upload.Loja.Categoria,
			IDUsuario: upload.Loja.IDUsuario,
		}
		response.Loja = lojaResp
	}

	return response
}

// CreateUpload cria um novo upload
func CreateUpload(req json.UploadRequest) (*json.UploadResponse, error) {
	upload, err := datasource.CreateUpload(req)
	if err != nil {
		return nil, err
	}
	return convertUploadToResponse(upload), nil
}

// GetUploadByID busca um upload por ID
func GetUploadByID(id uint) (*json.UploadResponse, error) {
	upload, err := datasource.GetUploadByID(id)
	if err != nil {
		return nil, err
	}
	return convertUploadToResponse(upload), nil
}

// GetAllUploads retorna todos os uploads ativos
func GetAllUploads() ([]json.UploadResponse, error) {
	uploads, err := datasource.GetAllUploads()
	if err != nil {
		return nil, err
	}

	var responses []json.UploadResponse
	for _, upload := range uploads {
		responses = append(responses, *convertUploadToResponse(&upload))
	}

	return responses, nil
}

// GetAllUploadsByTipo retorna todos os uploads ativos filtrados por tipo (Imagem ou Documento)
func GetAllUploadsByTipo(tipo string) ([]json.UploadResponse, error) {
	uploads, err := datasource.GetAllUploadsByTipo(tipo)
	if err != nil {
		return nil, err
	}

	var responses []json.UploadResponse
	for _, upload := range uploads {
		responses = append(responses, *convertUploadToResponse(&upload))
	}

	return responses, nil
}

// GetUploadsByUsuarioID retorna todos os uploads de um usuário específico
func GetUploadsByUsuarioID(idUsuario uint) (*json.UploadsResponse, error) {
	uploads, err := datasource.GetUploadsByUsuarioID(idUsuario)
	if err != nil {
		return nil, err
	}

	var uploadsResponse []json.UploadResponse
	for _, upload := range uploads {
		uploadsResponse = append(uploadsResponse, *convertUploadToResponse(&upload))
	}

	response := &json.UploadsResponse{
		Uploads: uploadsResponse,
		Total:   len(uploadsResponse),
	}

	return response, nil
}

// GetUploadsByVeiculoID retorna todos os uploads de um veículo específico
func GetUploadsByVeiculoID(idVeiculo uint) (*json.UploadsResponse, error) {
	uploads, err := datasource.GetUploadsByVeiculoID(idVeiculo)
	if err != nil {
		return nil, err
	}

	var uploadsResponse []json.UploadResponse
	for _, upload := range uploads {
		uploadsResponse = append(uploadsResponse, *convertUploadToResponse(&upload))
	}

	response := &json.UploadsResponse{
		Uploads: uploadsResponse,
		Total:   len(uploadsResponse),
	}

	return response, nil
}

// GetUploadsByVeiculoLojaID retorna todos os uploads de um veículo de loja específico
func GetUploadsByVeiculoLojaID(idVeiculoLoja uint) (*json.UploadsResponse, error) {
	uploads, err := datasource.GetUploadsByVeiculoLojaID(idVeiculoLoja)
	if err != nil {
		return nil, err
	}

	var uploadsResponse []json.UploadResponse
	for _, upload := range uploads {
		uploadsResponse = append(uploadsResponse, *convertUploadToResponse(&upload))
	}

	response := &json.UploadsResponse{
		Uploads: uploadsResponse,
		Total:   len(uploadsResponse),
	}

	return response, nil
}

// GetUploadsByProdutoID retorna todos os uploads de um produto específico
func GetUploadsByProdutoID(idProduto uint) (*json.UploadsResponse, error) {
	uploads, err := datasource.GetUploadsByProdutoID(idProduto)
	if err != nil {
		return nil, err
	}

	var uploadsResponse []json.UploadResponse
	for _, upload := range uploads {
		uploadsResponse = append(uploadsResponse, *convertUploadToResponse(&upload))
	}

	response := &json.UploadsResponse{
		Uploads: uploadsResponse,
		Total:   len(uploadsResponse),
	}

	return response, nil
}

// GetUploadsByServicoID retorna todos os uploads de um serviço específico
func GetUploadsByServicoID(idServico uint) (*json.UploadsResponse, error) {
	uploads, err := datasource.GetUploadsByServicoID(idServico)
	if err != nil {
		return nil, err
	}

	var uploadsResponse []json.UploadResponse
	for _, upload := range uploads {
		uploadsResponse = append(uploadsResponse, *convertUploadToResponse(&upload))
	}

	response := &json.UploadsResponse{
		Uploads: uploadsResponse,
		Total:   len(uploadsResponse),
	}

	return response, nil
}

// GetUploadsByLojaID retorna todos os uploads de uma loja específica
func GetUploadsByLojaID(idLoja uint) (*json.UploadsResponse, error) {
	uploads, err := datasource.GetUploadsByLojaID(idLoja)
	if err != nil {
		return nil, err
	}

	var uploadsResponse []json.UploadResponse
	for _, upload := range uploads {
		uploadsResponse = append(uploadsResponse, *convertUploadToResponse(&upload))
	}

	response := &json.UploadsResponse{
		Uploads: uploadsResponse,
		Total:   len(uploadsResponse),
	}

	return response, nil
}

// GetUploadPrincipalByEntidade retorna o upload principal (imagem) de uma entidade
func GetUploadPrincipalByEntidade(tipoEntidade string, idEntidade uint) (*json.UploadResponse, error) {
	upload, err := datasource.GetUploadPrincipalByEntidade(tipoEntidade, idEntidade)
	if err != nil {
		return nil, err
	}
	return convertUploadToResponse(upload), nil
}

// UpdateUpload atualiza um upload existente
func UpdateUpload(id uint, req json.UploadRequest) (*json.UploadResponse, error) {
	upload, err := datasource.UpdateUpload(id, req)
	if err != nil {
		return nil, err
	}
	return convertUploadToResponse(upload), nil
}

// SetUploadPrincipal define um upload como principal (apenas para imagens)
func SetUploadPrincipal(id uint) error {
	return datasource.SetUploadPrincipal(id)
}

// SoftDeleteUpload realiza soft delete do upload
func SoftDeleteUpload(id uint) error {
	return datasource.SoftDeleteUpload(id)
}

// RestoreUpload restaura um upload que foi soft deleted
func RestoreUpload(id uint) error {
	return datasource.RestoreUpload(id)
}

