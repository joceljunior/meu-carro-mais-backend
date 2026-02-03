package services

import (
	"errors"
	"meu-carro-mais/internal/database/datasource"
	"meu-carro-mais/internal/database/models"
	"meu-carro-mais/internal/handlers/json"
)

// Login realiza o login do usuário mobile
// IMPORTANTE: Apenas usuários do tipo "mobile" podem fazer login no app mobile
// Usuários do tipo customer, administrativo e executivo devem usar o login web
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

	// IMPORTANTE: Apenas usuários do tipo MOBILE podem fazer login no app
	// Outros tipos (customer, administrativo, executivo) devem usar /login/web
	if user.Tipo != models.TipoUsuarioMobile {
		switch user.Tipo {
		case models.TipoUsuarioCustomer:
			return nil, errors.New("usuários do tipo customer devem fazer login pela plataforma web")
		case models.TipoUsuarioAdministrativo:
			return nil, errors.New("usuários do tipo administrativo devem fazer login pela plataforma web")
		case models.TipoUsuarioExecutivo:
			return nil, errors.New("usuários do tipo executivo devem fazer login pela plataforma web")
		default:
			return nil, errors.New("tipo de usuário não permitido para login mobile")
		}
	}

	// Verifica se o usuário está aprovado
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
		ID:                         user.ID,
		Nome:                       user.Nome,
		Email:                      user.Email,
		CPF:                        user.CPF,
		Imagem:                     user.Imagem,
		Telefone:                   user.Telefone,
		Endereco:                   user.Endereco,
		DataNascimento:             user.DataNascimento,
		DataCadastro:               user.DataCadastro,
		Ativo:                      user.Ativo,
		Latitude:                   user.Latitude,
		Longitude:                  user.Longitude,
		IDPlano:                    user.IDPlano,
		IDLoja:                     user.IDLoja,
		Tipo:                       string(user.Tipo),
		Status:                     string(user.Status),
		NomePlano:                  user.Plano.Nome,
		SolicitacaoExecutivo:       string(user.SolicitacaoExecutivo),
		DataSolicitacaoExecutivo:   user.DataSolicitacaoExecutivo,
		MotivoSolicitacaoExecutivo: user.MotivoSolicitacaoExecutivo,
		LojaUsuarioResponse: json.LojaUsuarioResponse{
			Id:                      user.Loja.ID,
			Nome:                    user.Loja.Nome,
			Logo:                    user.Loja.Imagem,
			AnuncioDestaqueResponse: anuncioResp,
		},
	}
	return resp, nil
}

// LoginWeb realiza o login para a plataforma web
// Apenas usuários do tipo "executivo", "administrativo" e "customer" podem fazer login
// Customers precisam estar aprovados
func LoginWeb(req json.LoginRequest) (*json.LoginResponse, error) {
	// Busca o usuário por email
	user, err := datasource.GetUserByEmailOnly(req.Email)
	if err != nil {
		return nil, errors.New("usuário não encontrado")
	}

	// Verifica se o tipo de usuário é permitido para login web
	tiposPermitidos := map[models.TipoUsuario]bool{
		models.TipoUsuarioExecutivo:     true,
		models.TipoUsuarioAdministrativo: true,
		models.TipoUsuarioCustomer:      true,
	}

	if !tiposPermitidos[user.Tipo] {
		return nil, errors.New("usuários do tipo mobile não podem fazer login na plataforma web")
	}

	// Para customers, verifica se está aprovado
	if user.Tipo == models.TipoUsuarioCustomer {
		if user.Status == models.StatusUsuarioPendente {
			return nil, errors.New("sua conta está pendente de aprovação")
		}
		if user.Status == models.StatusUsuarioRejeitado {
			return nil, errors.New("sua conta foi rejeitada")
		}
	}

	// Monta a resposta
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
		ID:                         user.ID,
		Nome:                       user.Nome,
		Email:                      user.Email,
		CPF:                        user.CPF,
		Imagem:                     user.Imagem,
		Telefone:                   user.Telefone,
		Endereco:                   user.Endereco,
		DataNascimento:             user.DataNascimento,
		DataCadastro:               user.DataCadastro,
		Ativo:                      user.Ativo,
		Latitude:                   user.Latitude,
		Longitude:                  user.Longitude,
		IDPlano:                    user.IDPlano,
		IDLoja:                     user.IDLoja,
		Tipo:                       string(user.Tipo),
		Status:                     string(user.Status),
		NomePlano:                  user.Plano.Nome,
		SolicitacaoExecutivo:       string(user.SolicitacaoExecutivo),
		DataSolicitacaoExecutivo:   user.DataSolicitacaoExecutivo,
		MotivoSolicitacaoExecutivo: user.MotivoSolicitacaoExecutivo,
		LojaUsuarioResponse: json.LojaUsuarioResponse{
			Id:                      user.Loja.ID,
			Nome:                    user.Loja.Nome,
			Logo:                    user.Loja.Imagem,
			AnuncioDestaqueResponse: anuncioResp,
		},
	}
	return resp, nil
}
