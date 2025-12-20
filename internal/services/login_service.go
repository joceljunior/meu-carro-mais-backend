package services

import (
	"errors"
	"meu-carro-mais/internal/database/datasource"
	"meu-carro-mais/internal/database/models"
	"meu-carro-mais/internal/handlers/json"
)

// Login realiza o login do usuário mobile
// Importante: usuários do tipo "customer" não podem fazer login no mobile
func Login(req json.LoginRequest) (*json.LoginResponse, error) {
	// Primeiro, tenta buscar o usuário apenas por email
	user, err := datasource.GetUserByEmailOnly(req.Email)
	if err != nil {
		// Se não encontrou o usuário, cria um novo (tipo mobile)
		user, err = datasource.CreateUserFromLogin(req)
		if err != nil {
			return nil, err
		}
	}

	// IMPORTANTE: Verifica se o usuário é do tipo customer
	// Customers não podem fazer login no mobile
	if user.Tipo == models.TipoUsuarioCustomer {
		return nil, errors.New("usuários do tipo customer não podem fazer login no aplicativo mobile")
	}

	// Verifica se o usuário está aprovado (para casos de outros tipos que possam ter status)
	if user.Status == models.StatusUsuarioPendente {
		return nil, errors.New("usuário pendente de aprovação")
	}

	if user.Status == models.StatusUsuarioRejeitado {
		return nil, errors.New("usuário rejeitado")
	}

	var anuncioResp json.AnuncioDestaqueResponse
	if user.Loja.ID != 0 {
		anuncio, err := datasource.GetAnuncioDestaqueByLojaID(user.Loja.ID)
		if err == nil {
			anuncioResp = json.AnuncioDestaqueResponse{
				ID:        anuncio.ID,
				Titulo:    anuncio.Titulo,
				Descricao: anuncio.Descricao,
				Preco:     anuncio.Preco,
				Imagem:    anuncio.Imagem,
			}
		}
	}

	resp := &json.LoginResponse{
		ID:        user.ID,
		Nome:      user.Nome,
		Email:     user.Email,
		Tipo:      string(user.Tipo),
		Status:    string(user.Status),
		NomePlano: user.Plano.Nome,
		LojaUsuarioResponse: json.LojaUsuarioResponse{
			Id:                      user.Loja.ID,
			Nome:                    user.Loja.Nome,
			Logo:                    user.Loja.Imagem,
			AnuncioDestaqueResponse: anuncioResp,
		},
	}
	return resp, nil
}
