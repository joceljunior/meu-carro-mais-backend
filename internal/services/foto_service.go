package services

import (
	"meu-carro-mais/internal/database/datasource"
	"meu-carro-mais/internal/handlers/json"
)

// CreateFoto cria uma nova foto
func CreateFoto(req json.FotoRequest) (*json.FotoResponse, error) {
	foto, err := datasource.CreateFoto(req)
	if err != nil {
		return nil, err
	}

	response := &json.FotoResponse{
		ID:              foto.ID,
		IDVeiculo:       foto.IDVeiculo,
		IDVeiculoLoja:   foto.IDVeiculoLoja,
		IDProduto:       foto.IDProduto,
		IDServico:       foto.IDServico,
		IDLoja:          foto.IDLoja,
		TipoEntidade:    foto.TipoEntidade,
		URL:             foto.URL,
		NomeArquivo:     foto.NomeArquivo,
		Tamanho:         foto.Tamanho,
		TipoMime:        foto.TipoMime,
		Principal:       foto.Principal,
		Ordem:           foto.Ordem,
		DataUpload:      foto.DataUpload,
		DataAtualizacao: foto.DataAtualizacao,
	}

	// Adiciona dados do veículo se existir
	if foto.Veiculo != nil {
		veiculoResp := &json.VeiculoResponse{
			ID:           foto.Veiculo.ID,
			Modelo:       foto.Veiculo.Modelo,
			Ano:          foto.Veiculo.Ano,
			Cor:          foto.Veiculo.Cor,
			Placa:        foto.Veiculo.Placa,
			IDUsuario:    foto.Veiculo.IDUsuario,
			DataCadastro: foto.Veiculo.DataCadastro,
			Ativo:        foto.Veiculo.Ativo,
		}
		response.Veiculo = veiculoResp
	}

	// Adiciona dados do veículo de loja se existir
	if foto.VeiculoLoja != nil {
		veiculoLojaResp := &json.VeiculoLojaResponse{
			ID:           foto.VeiculoLoja.ID,
			Modelo:       foto.VeiculoLoja.Modelo,
			Ano:          foto.VeiculoLoja.Ano,
			Cor:          foto.VeiculoLoja.Cor,
			Placa:        foto.VeiculoLoja.Placa,
			IDLoja:       foto.VeiculoLoja.IDLoja,
			DataCadastro: foto.VeiculoLoja.DataCadastro,
			Ativo:        foto.VeiculoLoja.Ativo,
		}
		response.VeiculoLoja = veiculoLojaResp
	}

	// Adiciona dados do produto se existir
	if foto.Produto != nil {
		produtoResp := &json.ProdutoResponse{
			ID:           foto.Produto.ID,
			Nome:         foto.Produto.Nome,
			Descricao:    foto.Produto.Descricao,
			Preco:        foto.Produto.Preco,
			Imagem:       foto.Produto.Imagem,
			Estoque:      foto.Produto.Estoque,
			Ativo:        foto.Produto.Ativo,
			IDLoja:       foto.Produto.IDLoja,
			DataCadastro: foto.Produto.DataCadastro,
		}
		response.Produto = produtoResp
	}

	// Adiciona dados do serviço se existir
	if foto.Servico != nil {
		servicoResp := &json.ServicoResponse{
			ID:          foto.Servico.ID,
			Titulo:      foto.Servico.Titulo,
			Descricao:   foto.Servico.Descricao,
			Preco:       foto.Servico.Preco,
			Imagem:      foto.Servico.Imagem,
			Destaque:    foto.Servico.Destaque,
			IDCategoria: foto.Servico.IDCategoria,
			Categoria:   foto.Servico.Categoria.Nome,
		}
		response.Servico = servicoResp
	}

	// Adiciona dados da loja se existir
	if foto.Loja != nil {
		lojaResp := &json.LojaResponse{
			ID:          foto.Loja.ID,
			Nome:        foto.Loja.Nome,
			CNPJ:        foto.Loja.CNPJ,
			Imagem:      foto.Loja.Imagem,
			Latitude:    foto.Loja.Latitude,
			Longitude:   foto.Loja.Longitude,
			IDCategoria: foto.Loja.IDCategoria,
			Categoria:   foto.Loja.Categoria.Nome,
		}
		response.Loja = lojaResp
	}

	return response, nil
}

// GetFotoByID busca uma foto por ID
func GetFotoByID(id uint) (*json.FotoResponse, error) {
	foto, err := datasource.GetFotoByID(id)
	if err != nil {
		return nil, err
	}

	response := &json.FotoResponse{
		ID:              foto.ID,
		IDVeiculo:       foto.IDVeiculo,
		IDVeiculoLoja:   foto.IDVeiculoLoja,
		IDProduto:       foto.IDProduto,
		IDServico:       foto.IDServico,
		IDLoja:          foto.IDLoja,
		TipoEntidade:    foto.TipoEntidade,
		URL:             foto.URL,
		NomeArquivo:     foto.NomeArquivo,
		Tamanho:         foto.Tamanho,
		TipoMime:        foto.TipoMime,
		Principal:       foto.Principal,
		Ordem:           foto.Ordem,
		DataUpload:      foto.DataUpload,
		DataAtualizacao: foto.DataAtualizacao,
	}

	// Adiciona dados do veículo se existir
	if foto.Veiculo != nil {
		veiculoResp := &json.VeiculoResponse{
			ID:           foto.Veiculo.ID,
			Modelo:       foto.Veiculo.Modelo,
			Ano:          foto.Veiculo.Ano,
			Cor:          foto.Veiculo.Cor,
			Placa:        foto.Veiculo.Placa,
			IDUsuario:    foto.Veiculo.IDUsuario,
			DataCadastro: foto.Veiculo.DataCadastro,
			Ativo:        foto.Veiculo.Ativo,
		}
		response.Veiculo = veiculoResp
	}

	// Adiciona dados do veículo de loja se existir
	if foto.VeiculoLoja != nil {
		veiculoLojaResp := &json.VeiculoLojaResponse{
			ID:           foto.VeiculoLoja.ID,
			Modelo:       foto.VeiculoLoja.Modelo,
			Ano:          foto.VeiculoLoja.Ano,
			Cor:          foto.VeiculoLoja.Cor,
			Placa:        foto.VeiculoLoja.Placa,
			IDLoja:       foto.VeiculoLoja.IDLoja,
			DataCadastro: foto.VeiculoLoja.DataCadastro,
			Ativo:        foto.VeiculoLoja.Ativo,
		}
		response.VeiculoLoja = veiculoLojaResp
	}

	// Adiciona dados do produto se existir
	if foto.Produto != nil {
		produtoResp := &json.ProdutoResponse{
			ID:           foto.Produto.ID,
			Nome:         foto.Produto.Nome,
			Descricao:    foto.Produto.Descricao,
			Preco:        foto.Produto.Preco,
			Imagem:       foto.Produto.Imagem,
			Estoque:      foto.Produto.Estoque,
			Ativo:        foto.Produto.Ativo,
			IDLoja:       foto.Produto.IDLoja,
			DataCadastro: foto.Produto.DataCadastro,
		}
		response.Produto = produtoResp
	}

	// Adiciona dados do serviço se existir
	if foto.Servico != nil {
		servicoResp := &json.ServicoResponse{
			ID:          foto.Servico.ID,
			Titulo:      foto.Servico.Titulo,
			Descricao:   foto.Servico.Descricao,
			Preco:       foto.Servico.Preco,
			Imagem:      foto.Servico.Imagem,
			Destaque:    foto.Servico.Destaque,
			IDCategoria: foto.Servico.IDCategoria,
			Categoria:   foto.Servico.Categoria.Nome,
		}
		response.Servico = servicoResp
	}

	// Adiciona dados da loja se existir
	if foto.Loja != nil {
		lojaResp := &json.LojaResponse{
			ID:          foto.Loja.ID,
			Nome:        foto.Loja.Nome,
			CNPJ:        foto.Loja.CNPJ,
			Imagem:      foto.Loja.Imagem,
			Latitude:    foto.Loja.Latitude,
			Longitude:   foto.Loja.Longitude,
			IDCategoria: foto.Loja.IDCategoria,
			Categoria:   foto.Loja.Categoria.Nome,
		}
		response.Loja = lojaResp
	}

	return response, nil
}

// GetAllFotos retorna todas as fotos ativas
func GetAllFotos() ([]json.FotoResponse, error) {
	fotos, err := datasource.GetAllFotos()
	if err != nil {
		return nil, err
	}

	var responses []json.FotoResponse
	for _, foto := range fotos {
		response := json.FotoResponse{
			ID:              foto.ID,
			IDVeiculo:       foto.IDVeiculo,
			IDVeiculoLoja:   foto.IDVeiculoLoja,
			IDProduto:       foto.IDProduto,
			IDServico:       foto.IDServico,
			TipoEntidade:    foto.TipoEntidade,
			URL:             foto.URL,
			NomeArquivo:     foto.NomeArquivo,
			Tamanho:         foto.Tamanho,
			TipoMime:        foto.TipoMime,
			Principal:       foto.Principal,
			Ordem:           foto.Ordem,
			DataUpload:      foto.DataUpload,
			DataAtualizacao: foto.DataAtualizacao,
		}

		// Adiciona dados do veículo se existir
		if foto.Veiculo != nil {
			veiculoResp := &json.VeiculoResponse{
				ID:           foto.Veiculo.ID,
				Modelo:       foto.Veiculo.Modelo,
				Ano:          foto.Veiculo.Ano,
				Cor:          foto.Veiculo.Cor,
				Placa:        foto.Veiculo.Placa,
				IDUsuario:    foto.Veiculo.IDUsuario,
				DataCadastro: foto.Veiculo.DataCadastro,
				Ativo:        foto.Veiculo.Ativo,
			}
			response.Veiculo = veiculoResp
		}

		// Adiciona dados do veículo de loja se existir
		if foto.VeiculoLoja != nil {
			veiculoLojaResp := &json.VeiculoLojaResponse{
				ID:           foto.VeiculoLoja.ID,
				Modelo:       foto.VeiculoLoja.Modelo,
				Ano:          foto.VeiculoLoja.Ano,
				Cor:          foto.VeiculoLoja.Cor,
				Placa:        foto.VeiculoLoja.Placa,
				IDLoja:       foto.VeiculoLoja.IDLoja,
				DataCadastro: foto.VeiculoLoja.DataCadastro,
				Ativo:        foto.VeiculoLoja.Ativo,
			}
			response.VeiculoLoja = veiculoLojaResp
		}

		// Adiciona dados do produto se existir
		if foto.Produto != nil {
			produtoResp := &json.ProdutoResponse{
				ID:           foto.Produto.ID,
				Nome:         foto.Produto.Nome,
				Descricao:    foto.Produto.Descricao,
				Preco:        foto.Produto.Preco,
				Imagem:       foto.Produto.Imagem,
				Estoque:      foto.Produto.Estoque,
				Ativo:        foto.Produto.Ativo,
				IDLoja:       foto.Produto.IDLoja,
				DataCadastro: foto.Produto.DataCadastro,
			}
			response.Produto = produtoResp
		}

		// Adiciona dados do serviço se existir
		if foto.Servico != nil {
			servicoResp := &json.ServicoResponse{
				ID:          foto.Servico.ID,
				Titulo:      foto.Servico.Titulo,
				Descricao:   foto.Servico.Descricao,
				Preco:       foto.Servico.Preco,
				Imagem:      foto.Servico.Imagem,
				Destaque:    foto.Servico.Destaque,
				IDCategoria: foto.Servico.IDCategoria,
				Categoria:   foto.Servico.Categoria.Nome,
			}
			response.Servico = servicoResp
		}

		responses = append(responses, response)
	}

	return responses, nil
}

// GetFotosByVeiculoID retorna todas as fotos de um veículo específico
func GetFotosByVeiculoID(idVeiculo uint) (*json.FotosResponse, error) {
	fotos, err := datasource.GetFotosByVeiculoID(idVeiculo)
	if err != nil {
		return nil, err
	}

	var fotosResponse []json.FotoResponse
	for _, foto := range fotos {
		fotoResp := json.FotoResponse{
			ID:              foto.ID,
			IDVeiculo:       foto.IDVeiculo,
			IDVeiculoLoja:   foto.IDVeiculoLoja,
			IDProduto:       foto.IDProduto,
			IDServico:       foto.IDServico,
			TipoEntidade:    foto.TipoEntidade,
			URL:             foto.URL,
			NomeArquivo:     foto.NomeArquivo,
			Tamanho:         foto.Tamanho,
			TipoMime:        foto.TipoMime,
			Principal:       foto.Principal,
			Ordem:           foto.Ordem,
			DataUpload:      foto.DataUpload,
			DataAtualizacao: foto.DataAtualizacao,
		}

		// Adiciona dados do veículo se existir
		if foto.Veiculo != nil {
			veiculoResp := &json.VeiculoResponse{
				ID:           foto.Veiculo.ID,
				Modelo:       foto.Veiculo.Modelo,
				Ano:          foto.Veiculo.Ano,
				Cor:          foto.Veiculo.Cor,
				Placa:        foto.Veiculo.Placa,
				IDUsuario:    foto.Veiculo.IDUsuario,
				DataCadastro: foto.Veiculo.DataCadastro,
				Ativo:        foto.Veiculo.Ativo,
			}
			fotoResp.Veiculo = veiculoResp
		}

		fotosResponse = append(fotosResponse, fotoResp)
	}

	response := &json.FotosResponse{
		Fotos: fotosResponse,
		Total: len(fotosResponse),
	}

	return response, nil
}

// GetFotosByVeiculoLojaID retorna todas as fotos de um veículo de loja específico
func GetFotosByVeiculoLojaID(idVeiculoLoja uint) (*json.FotosResponse, error) {
	fotos, err := datasource.GetFotosByVeiculoLojaID(idVeiculoLoja)
	if err != nil {
		return nil, err
	}

	var fotosResponse []json.FotoResponse
	for _, foto := range fotos {
		fotoResp := json.FotoResponse{
			ID:              foto.ID,
			IDVeiculo:       foto.IDVeiculo,
			IDVeiculoLoja:   foto.IDVeiculoLoja,
			IDProduto:       foto.IDProduto,
			IDServico:       foto.IDServico,
			TipoEntidade:    foto.TipoEntidade,
			URL:             foto.URL,
			NomeArquivo:     foto.NomeArquivo,
			Tamanho:         foto.Tamanho,
			TipoMime:        foto.TipoMime,
			Principal:       foto.Principal,
			Ordem:           foto.Ordem,
			DataUpload:      foto.DataUpload,
			DataAtualizacao: foto.DataAtualizacao,
		}

		// Adiciona dados do veículo de loja se existir
		if foto.VeiculoLoja != nil {
			veiculoLojaResp := &json.VeiculoLojaResponse{
				ID:           foto.VeiculoLoja.ID,
				Modelo:       foto.VeiculoLoja.Modelo,
				Ano:          foto.VeiculoLoja.Ano,
				Cor:          foto.VeiculoLoja.Cor,
				Placa:        foto.VeiculoLoja.Placa,
				IDLoja:       foto.VeiculoLoja.IDLoja,
				DataCadastro: foto.VeiculoLoja.DataCadastro,
				Ativo:        foto.VeiculoLoja.Ativo,
			}
			fotoResp.VeiculoLoja = veiculoLojaResp
		}

		fotosResponse = append(fotosResponse, fotoResp)
	}

	response := &json.FotosResponse{
		Fotos: fotosResponse,
		Total: len(fotosResponse),
	}

	return response, nil
}

// GetFotosByProdutoID retorna todas as fotos de um produto específico
func GetFotosByProdutoID(idProduto uint) (*json.FotosResponse, error) {
	fotos, err := datasource.GetFotosByProdutoID(idProduto)
	if err != nil {
		return nil, err
	}

	var fotosResponse []json.FotoResponse
	for _, foto := range fotos {
		fotoResp := json.FotoResponse{
			ID:              foto.ID,
			IDVeiculo:       foto.IDVeiculo,
			IDVeiculoLoja:   foto.IDVeiculoLoja,
			IDProduto:       foto.IDProduto,
			IDServico:       foto.IDServico,
			TipoEntidade:    foto.TipoEntidade,
			URL:             foto.URL,
			NomeArquivo:     foto.NomeArquivo,
			Tamanho:         foto.Tamanho,
			TipoMime:        foto.TipoMime,
			Principal:       foto.Principal,
			Ordem:           foto.Ordem,
			DataUpload:      foto.DataUpload,
			DataAtualizacao: foto.DataAtualizacao,
		}

		// Adiciona dados do produto se existir
		if foto.Produto != nil {
			produtoResp := &json.ProdutoResponse{
				ID:           foto.Produto.ID,
				Nome:         foto.Produto.Nome,
				Descricao:    foto.Produto.Descricao,
				Preco:        foto.Produto.Preco,
				Imagem:       foto.Produto.Imagem,
				Estoque:      foto.Produto.Estoque,
				Ativo:        foto.Produto.Ativo,
				IDLoja:       foto.Produto.IDLoja,
				DataCadastro: foto.Produto.DataCadastro,
			}
			fotoResp.Produto = produtoResp
		}

		fotosResponse = append(fotosResponse, fotoResp)
	}

	response := &json.FotosResponse{
		Fotos: fotosResponse,
		Total: len(fotosResponse),
	}

	return response, nil
}

// GetFotosByServicoID retorna todas as fotos de um serviço específico
func GetFotosByServicoID(idServico uint) (*json.FotosResponse, error) {
	fotos, err := datasource.GetFotosByServicoID(idServico)
	if err != nil {
		return nil, err
	}

	var fotosResponse []json.FotoResponse
	for _, foto := range fotos {
		fotoResp := json.FotoResponse{
			ID:              foto.ID,
			IDVeiculo:       foto.IDVeiculo,
			IDVeiculoLoja:   foto.IDVeiculoLoja,
			IDProduto:       foto.IDProduto,
			IDServico:       foto.IDServico,
			IDLoja:          foto.IDLoja,
			TipoEntidade:    foto.TipoEntidade,
			URL:             foto.URL,
			NomeArquivo:     foto.NomeArquivo,
			Tamanho:         foto.Tamanho,
			TipoMime:        foto.TipoMime,
			Principal:       foto.Principal,
			Ordem:           foto.Ordem,
			DataUpload:      foto.DataUpload,
			DataAtualizacao: foto.DataAtualizacao,
		}

		// Adiciona dados do serviço se existir
		if foto.Servico != nil {
			servicoResp := &json.ServicoResponse{
				ID:          foto.Servico.ID,
				Titulo:      foto.Servico.Titulo,
				Descricao:   foto.Servico.Descricao,
				Preco:       foto.Servico.Preco,
				Imagem:      foto.Servico.Imagem,
				Destaque:    foto.Servico.Destaque,
				IDCategoria: foto.Servico.IDCategoria,
				Categoria:   foto.Servico.Categoria.Nome,
			}
			fotoResp.Servico = servicoResp
		}

		fotosResponse = append(fotosResponse, fotoResp)
	}

	response := &json.FotosResponse{
		Fotos: fotosResponse,
		Total: len(fotosResponse),
	}

	return response, nil
}

// GetFotosByLojaID retorna todas as fotos de uma loja específica
func GetFotosByLojaID(idLoja uint) (*json.FotosResponse, error) {
	fotos, err := datasource.GetFotosByLojaID(idLoja)
	if err != nil {
		return nil, err
	}

	var fotosResponse []json.FotoResponse
	for _, foto := range fotos {
		fotoResp := json.FotoResponse{
			ID:              foto.ID,
			IDVeiculo:       foto.IDVeiculo,
			IDVeiculoLoja:   foto.IDVeiculoLoja,
			IDProduto:       foto.IDProduto,
			IDServico:       foto.IDServico,
			IDLoja:          foto.IDLoja,
			TipoEntidade:    foto.TipoEntidade,
			URL:             foto.URL,
			NomeArquivo:     foto.NomeArquivo,
			Tamanho:         foto.Tamanho,
			TipoMime:        foto.TipoMime,
			Principal:       foto.Principal,
			Ordem:           foto.Ordem,
			DataUpload:      foto.DataUpload,
			DataAtualizacao: foto.DataAtualizacao,
		}

		// Adiciona dados da loja se existir
		if foto.Loja != nil {
			lojaResp := &json.LojaResponse{
				ID:          foto.Loja.ID,
				Nome:        foto.Loja.Nome,
				CNPJ:        foto.Loja.CNPJ,
				Imagem:      foto.Loja.Imagem,
				Latitude:    foto.Loja.Latitude,
				Longitude:   foto.Loja.Longitude,
				IDCategoria: foto.Loja.IDCategoria,
				Categoria:   foto.Loja.Categoria.Nome,
			}
			fotoResp.Loja = lojaResp
		}

		fotosResponse = append(fotosResponse, fotoResp)
	}

	response := &json.FotosResponse{
		Fotos: fotosResponse,
		Total: len(fotosResponse),
	}

	return response, nil
}

// GetFotoPrincipalByEntidade retorna a foto principal de uma entidade
func GetFotoPrincipalByEntidade(tipoEntidade string, idEntidade uint) (*json.FotoResponse, error) {
	foto, err := datasource.GetFotoPrincipalByEntidade(tipoEntidade, idEntidade)
	if err != nil {
		return nil, err
	}

	response := &json.FotoResponse{
		ID:              foto.ID,
		IDVeiculo:       foto.IDVeiculo,
		IDVeiculoLoja:   foto.IDVeiculoLoja,
		IDProduto:       foto.IDProduto,
		IDServico:       foto.IDServico,
		IDLoja:          foto.IDLoja,
		TipoEntidade:    foto.TipoEntidade,
		URL:             foto.URL,
		NomeArquivo:     foto.NomeArquivo,
		Tamanho:         foto.Tamanho,
		TipoMime:        foto.TipoMime,
		Principal:       foto.Principal,
		Ordem:           foto.Ordem,
		DataUpload:      foto.DataUpload,
		DataAtualizacao: foto.DataAtualizacao,
	}

	// Adiciona dados do veículo se existir
	if foto.Veiculo != nil {
		veiculoResp := &json.VeiculoResponse{
			ID:           foto.Veiculo.ID,
			Modelo:       foto.Veiculo.Modelo,
			Ano:          foto.Veiculo.Ano,
			Cor:          foto.Veiculo.Cor,
			Placa:        foto.Veiculo.Placa,
			IDUsuario:    foto.Veiculo.IDUsuario,
			DataCadastro: foto.Veiculo.DataCadastro,
			Ativo:        foto.Veiculo.Ativo,
		}
		response.Veiculo = veiculoResp
	}

	// Adiciona dados do veículo de loja se existir
	if foto.VeiculoLoja != nil {
		veiculoLojaResp := &json.VeiculoLojaResponse{
			ID:           foto.VeiculoLoja.ID,
			Modelo:       foto.VeiculoLoja.Modelo,
			Ano:          foto.VeiculoLoja.Ano,
			Cor:          foto.VeiculoLoja.Cor,
			Placa:        foto.VeiculoLoja.Placa,
			IDLoja:       foto.VeiculoLoja.IDLoja,
			DataCadastro: foto.VeiculoLoja.DataCadastro,
			Ativo:        foto.VeiculoLoja.Ativo,
		}
		response.VeiculoLoja = veiculoLojaResp
	}

	// Adiciona dados do produto se existir
	if foto.Produto != nil {
		produtoResp := &json.ProdutoResponse{
			ID:           foto.Produto.ID,
			Nome:         foto.Produto.Nome,
			Descricao:    foto.Produto.Descricao,
			Preco:        foto.Produto.Preco,
			Imagem:       foto.Produto.Imagem,
			Estoque:      foto.Produto.Estoque,
			Ativo:        foto.Produto.Ativo,
			IDLoja:       foto.Produto.IDLoja,
			DataCadastro: foto.Produto.DataCadastro,
		}
		response.Produto = produtoResp
	}

	// Adiciona dados do serviço se existir
	if foto.Servico != nil {
		servicoResp := &json.ServicoResponse{
			ID:          foto.Servico.ID,
			Titulo:      foto.Servico.Titulo,
			Descricao:   foto.Servico.Descricao,
			Preco:       foto.Servico.Preco,
			Imagem:      foto.Servico.Imagem,
			Destaque:    foto.Servico.Destaque,
			IDCategoria: foto.Servico.IDCategoria,
			Categoria:   foto.Servico.Categoria.Nome,
		}
		response.Servico = servicoResp
	}

	// Adiciona dados da loja se existir
	if foto.Loja != nil {
		lojaResp := &json.LojaResponse{
			ID:          foto.Loja.ID,
			Nome:        foto.Loja.Nome,
			CNPJ:        foto.Loja.CNPJ,
			Imagem:      foto.Loja.Imagem,
			Latitude:    foto.Loja.Latitude,
			Longitude:   foto.Loja.Longitude,
			IDCategoria: foto.Loja.IDCategoria,
			Categoria:   foto.Loja.Categoria.Nome,
		}
		response.Loja = lojaResp
	}

	return response, nil
}

// UpdateFoto atualiza uma foto existente
func UpdateFoto(id uint, req json.FotoRequest) (*json.FotoResponse, error) {
	foto, err := datasource.UpdateFoto(id, req)
	if err != nil {
		return nil, err
	}

	response := &json.FotoResponse{
		ID:              foto.ID,
		IDVeiculo:       foto.IDVeiculo,
		IDVeiculoLoja:   foto.IDVeiculoLoja,
		IDProduto:       foto.IDProduto,
		IDServico:       foto.IDServico,
		IDLoja:          foto.IDLoja,
		TipoEntidade:    foto.TipoEntidade,
		URL:             foto.URL,
		NomeArquivo:     foto.NomeArquivo,
		Tamanho:         foto.Tamanho,
		TipoMime:        foto.TipoMime,
		Principal:       foto.Principal,
		Ordem:           foto.Ordem,
		DataUpload:      foto.DataUpload,
		DataAtualizacao: foto.DataAtualizacao,
	}

	// Adiciona dados do veículo se existir
	if foto.Veiculo != nil {
		veiculoResp := &json.VeiculoResponse{
			ID:           foto.Veiculo.ID,
			Modelo:       foto.Veiculo.Modelo,
			Ano:          foto.Veiculo.Ano,
			Cor:          foto.Veiculo.Cor,
			Placa:        foto.Veiculo.Placa,
			IDUsuario:    foto.Veiculo.IDUsuario,
			DataCadastro: foto.Veiculo.DataCadastro,
			Ativo:        foto.Veiculo.Ativo,
		}
		response.Veiculo = veiculoResp
	}

	// Adiciona dados do veículo de loja se existir
	if foto.VeiculoLoja != nil {
		veiculoLojaResp := &json.VeiculoLojaResponse{
			ID:           foto.VeiculoLoja.ID,
			Modelo:       foto.VeiculoLoja.Modelo,
			Ano:          foto.VeiculoLoja.Ano,
			Cor:          foto.VeiculoLoja.Cor,
			Placa:        foto.VeiculoLoja.Placa,
			IDLoja:       foto.VeiculoLoja.IDLoja,
			DataCadastro: foto.VeiculoLoja.DataCadastro,
			Ativo:        foto.VeiculoLoja.Ativo,
		}
		response.VeiculoLoja = veiculoLojaResp
	}

	// Adiciona dados do produto se existir
	if foto.Produto != nil {
		produtoResp := &json.ProdutoResponse{
			ID:           foto.Produto.ID,
			Nome:         foto.Produto.Nome,
			Descricao:    foto.Produto.Descricao,
			Preco:        foto.Produto.Preco,
			Imagem:       foto.Produto.Imagem,
			Estoque:      foto.Produto.Estoque,
			Ativo:        foto.Produto.Ativo,
			IDLoja:       foto.Produto.IDLoja,
			DataCadastro: foto.Produto.DataCadastro,
		}
		response.Produto = produtoResp
	}

	// Adiciona dados do serviço se existir
	if foto.Servico != nil {
		servicoResp := &json.ServicoResponse{
			ID:          foto.Servico.ID,
			Titulo:      foto.Servico.Titulo,
			Descricao:   foto.Servico.Descricao,
			Preco:       foto.Servico.Preco,
			Imagem:      foto.Servico.Imagem,
			Destaque:    foto.Servico.Destaque,
			IDCategoria: foto.Servico.IDCategoria,
			Categoria:   foto.Servico.Categoria.Nome,
		}
		response.Servico = servicoResp
	}

	// Adiciona dados da loja se existir
	if foto.Loja != nil {
		lojaResp := &json.LojaResponse{
			ID:          foto.Loja.ID,
			Nome:        foto.Loja.Nome,
			CNPJ:        foto.Loja.CNPJ,
			Imagem:      foto.Loja.Imagem,
			Latitude:    foto.Loja.Latitude,
			Longitude:   foto.Loja.Longitude,
			IDCategoria: foto.Loja.IDCategoria,
			Categoria:   foto.Loja.Categoria.Nome,
		}
		response.Loja = lojaResp
	}

	return response, nil
}

// SetFotoPrincipal define uma foto como principal
func SetFotoPrincipal(id uint) error {
	return datasource.SetFotoPrincipal(id)
}

// SoftDeleteFoto realiza soft delete da foto
func SoftDeleteFoto(id uint) error {
	return datasource.SoftDeleteFoto(id)
}

// RestoreFoto restaura uma foto que foi soft deleted
func RestoreFoto(id uint) error {
	return datasource.RestoreFoto(id)
}
