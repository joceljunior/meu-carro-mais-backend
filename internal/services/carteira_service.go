package services

import (
	"meu-carro-mais/internal/database/datasource"
	"meu-carro-mais/internal/handlers/json"
)

// CreateCarteira cria uma nova carteira
func CreateCarteira(req json.CarteiraRequest) (*json.CarteiraResponse, error) {
	carteira, err := datasource.CreateCarteira(req)
	if err != nil {
		return nil, err
	}

	response := &json.CarteiraResponse{
		ID:              carteira.ID,
		UsuarioID:       carteira.UsuarioID,
		Saldo:           carteira.Saldo,
		DataCriacao:     carteira.DataCriacao,
		DataAtualizacao: carteira.DataAtualizacao,
		Mensagem:        "Carteira criada com sucesso",
	}

	return response, nil
}

// GetCarteiraByID busca uma carteira por ID
func GetCarteiraByID(id uint) (*json.CarteiraComUsuarioResponse, error) {
	carteira, err := datasource.GetCarteiraByID(id)
	if err != nil {
		return nil, err
	}

	response := &json.CarteiraComUsuarioResponse{
		ID:              carteira.ID,
		UsuarioID:       carteira.UsuarioID,
		Saldo:           carteira.Saldo,
		DataCriacao:     carteira.DataCriacao,
		DataAtualizacao: carteira.DataAtualizacao,
		Mensagem:        "Carteira encontrada com sucesso",
	}

	// Se o usuário foi carregado, adiciona os dados
	if carteira.Usuario != nil {
		response.Usuario.ID = carteira.Usuario.ID
		response.Usuario.Nome = carteira.Usuario.Nome
		response.Usuario.Email = carteira.Usuario.Email
	}

	return response, nil
}

// GetCarteiraByUsuarioID busca a carteira de um usuário específico
func GetCarteiraByUsuarioID(usuarioID uint) (*json.CarteiraComUsuarioResponse, error) {
	carteira, err := datasource.GetCarteiraByUsuarioID(usuarioID)
	if err != nil {
		return nil, err
	}

	response := &json.CarteiraComUsuarioResponse{
		ID:              carteira.ID,
		UsuarioID:       carteira.UsuarioID,
		Saldo:           carteira.Saldo,
		DataCriacao:     carteira.DataCriacao,
		DataAtualizacao: carteira.DataAtualizacao,
		Mensagem:        "Carteira encontrada com sucesso",
	}

	// Se o usuário foi carregado, adiciona os dados
	if carteira.Usuario != nil {
		response.Usuario.ID = carteira.Usuario.ID
		response.Usuario.Nome = carteira.Usuario.Nome
		response.Usuario.Email = carteira.Usuario.Email
	}

	return response, nil
}

// GetAllCarteiras retorna todas as carteiras
func GetAllCarteiras() (*json.CarteirasResponse, error) {
	carteiras, err := datasource.GetAllCarteiras()
	if err != nil {
		return nil, err
	}

	var responses []json.CarteiraComUsuarioResponse
	for _, carteira := range carteiras {
		response := json.CarteiraComUsuarioResponse{
			ID:              carteira.ID,
			UsuarioID:       carteira.UsuarioID,
			Saldo:           carteira.Saldo,
			DataCriacao:     carteira.DataCriacao,
			DataAtualizacao: carteira.DataAtualizacao,
		}

		// Se o usuário foi carregado, adiciona os dados
		if carteira.Usuario != nil {
			response.Usuario.ID = carteira.Usuario.ID
			response.Usuario.Nome = carteira.Usuario.Nome
			response.Usuario.Email = carteira.Usuario.Email
		}

		responses = append(responses, response)
	}

	return &json.CarteirasResponse{
		Carteiras: responses,
		Total:     len(responses),
		Mensagem:  "Carteiras listadas com sucesso",
	}, nil
}

// UpdateCarteira atualiza uma carteira existente
func UpdateCarteira(id uint, req json.CarteiraRequest) (*json.CarteiraResponse, error) {
	carteira, err := datasource.UpdateCarteira(id, req)
	if err != nil {
		return nil, err
	}

	response := &json.CarteiraResponse{
		ID:              carteira.ID,
		UsuarioID:       carteira.UsuarioID,
		Saldo:           carteira.Saldo,
		DataCriacao:     carteira.DataCriacao,
		DataAtualizacao: carteira.DataAtualizacao,
		Mensagem:        "Carteira atualizada com sucesso",
	}

	return response, nil
}

// UpdateCarteiraSaldo atualiza apenas o saldo de uma carteira
func UpdateCarteiraSaldo(id uint, req json.CarteiraSaldoRequest) (*json.CarteiraResponse, error) {
	carteira, err := datasource.UpdateCarteiraSaldo(id, req.Saldo)
	if err != nil {
		return nil, err
	}

	response := &json.CarteiraResponse{
		ID:              carteira.ID,
		UsuarioID:       carteira.UsuarioID,
		Saldo:           carteira.Saldo,
		DataCriacao:     carteira.DataCriacao,
		DataAtualizacao: carteira.DataAtualizacao,
		Mensagem:        "Saldo atualizado com sucesso",
	}

	return response, nil
}

// AdicionarSaldo adiciona valor ao saldo da carteira
func AdicionarSaldo(id uint, req json.CarteiraOperacaoRequest) (*json.CarteiraOperacaoResponse, error) {
	// Busca a carteira atual para obter o saldo anterior
	carteiraAtual, err := datasource.GetCarteiraByID(id)
	if err != nil {
		return nil, err
	}

	saldoAnterior := carteiraAtual.Saldo

	// Adiciona o saldo
	carteira, err := datasource.AdicionarSaldo(id, req.Valor)
	if err != nil {
		return nil, err
	}

	response := &json.CarteiraOperacaoResponse{
		ID:              carteira.ID,
		UsuarioID:       carteira.UsuarioID,
		SaldoAnterior:   saldoAnterior,
		SaldoAtual:      carteira.Saldo,
		ValorOperacao:   req.Valor,
		TipoOperacao:    "adicao",
		DataAtualizacao: carteira.DataAtualizacao,
		Mensagem:        "Saldo adicionado com sucesso",
	}

	return response, nil
}

// SubtrairSaldo subtrai valor do saldo da carteira
func SubtrairSaldo(id uint, req json.CarteiraOperacaoRequest) (*json.CarteiraOperacaoResponse, error) {
	// Busca a carteira atual para obter o saldo anterior
	carteiraAtual, err := datasource.GetCarteiraByID(id)
	if err != nil {
		return nil, err
	}

	saldoAnterior := carteiraAtual.Saldo

	// Subtrai o saldo
	carteira, err := datasource.SubtrairSaldo(id, req.Valor)
	if err != nil {
		return nil, err
	}

	response := &json.CarteiraOperacaoResponse{
		ID:              carteira.ID,
		UsuarioID:       carteira.UsuarioID,
		SaldoAnterior:   saldoAnterior,
		SaldoAtual:      carteira.Saldo,
		ValorOperacao:   req.Valor,
		TipoOperacao:    "subtracao",
		DataAtualizacao: carteira.DataAtualizacao,
		Mensagem:        "Saldo subtraído com sucesso",
	}

	return response, nil
}

// DeleteCarteira remove uma carteira
func DeleteCarteira(id uint) error {
	return datasource.DeleteCarteira(id)
}

// GetCarteirasBySaldoRange busca carteiras com saldo dentro de um range
func GetCarteirasBySaldoRange(saldoMin, saldoMax float64) (*json.CarteirasResponse, error) {
	carteiras, err := datasource.GetCarteirasBySaldoRange(saldoMin, saldoMax)
	if err != nil {
		return nil, err
	}

	var responses []json.CarteiraComUsuarioResponse
	for _, carteira := range carteiras {
		response := json.CarteiraComUsuarioResponse{
			ID:              carteira.ID,
			UsuarioID:       carteira.UsuarioID,
			Saldo:           carteira.Saldo,
			DataCriacao:     carteira.DataCriacao,
			DataAtualizacao: carteira.DataAtualizacao,
		}

		// Se o usuário foi carregado, adiciona os dados
		if carteira.Usuario != nil {
			response.Usuario.ID = carteira.Usuario.ID
			response.Usuario.Nome = carteira.Usuario.Nome
			response.Usuario.Email = carteira.Usuario.Email
		}

		responses = append(responses, response)
	}

	return &json.CarteirasResponse{
		Carteiras: responses,
		Total:     len(responses),
		Mensagem:  "Carteiras encontradas com sucesso",
	}, nil
}

// GetCarteirasComSaldoMaior busca carteiras com saldo maior que um valor
func GetCarteirasComSaldoMaior(valor float64) (*json.CarteirasResponse, error) {
	carteiras, err := datasource.GetCarteirasComSaldoMaior(valor)
	if err != nil {
		return nil, err
	}

	var responses []json.CarteiraComUsuarioResponse
	for _, carteira := range carteiras {
		response := json.CarteiraComUsuarioResponse{
			ID:              carteira.ID,
			UsuarioID:       carteira.UsuarioID,
			Saldo:           carteira.Saldo,
			DataCriacao:     carteira.DataCriacao,
			DataAtualizacao: carteira.DataAtualizacao,
		}

		// Se o usuário foi carregado, adiciona os dados
		if carteira.Usuario != nil {
			response.Usuario.ID = carteira.Usuario.ID
			response.Usuario.Nome = carteira.Usuario.Nome
			response.Usuario.Email = carteira.Usuario.Email
		}

		responses = append(responses, response)
	}

	return &json.CarteirasResponse{
		Carteiras: responses,
		Total:     len(responses),
		Mensagem:  "Carteiras encontradas com sucesso",
	}, nil
}
