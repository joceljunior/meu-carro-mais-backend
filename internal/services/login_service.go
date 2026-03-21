package services

import (
	"errors"
	"meu-carro-mais/internal/database/datasource"
	"meu-carro-mais/internal/database/models"
	"meu-carro-mais/internal/handlers/json"
)

// Login realiza o login do usuário no app mobile
// Usuários do tipo "mobile" e "executivo" podem fazer login no app mobile
// Usuários do tipo "customer" e "administrativo" devem usar o login web (portal)
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

	// Verifica se o tipo de usuário é permitido para login mobile
	// Mobile e Executivo acessam o app mobile
	// Customer e Administrativo acessam o portal web
	tiposPermitidosMobile := map[models.TipoUsuario]bool{
		models.TipoUsuarioMobile:    true,
		models.TipoUsuarioExecutivo: true,
	}

	if !tiposPermitidosMobile[user.Tipo] {
		switch user.Tipo {
		case models.TipoUsuarioCustomer:
			return nil, errors.New("lojistas devem fazer login pela plataforma web")
		case models.TipoUsuarioAdministrativo:
			return nil, errors.New("administradores devem fazer login pela plataforma web")
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

	var cupomResp json.CupomDestaqueResponse
	if user.Loja.ID != 0 {
		cupom, err := datasource.GetCupomDestaqueByLojaID(user.Loja.ID)
		if err == nil {
			cupomResp = json.CupomDestaqueResponse{
				ID:        cupom.ID,
				Titulo:    cupom.Titulo,
				Descricao: cupom.Descricao,
				Preco:     cupom.Preco,
				Imagem:    cupom.Imagem,
				TipoCupom: cupom.TipoCupom,
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
		IDExecutivo:                user.IDExecutivo,
		SolicitacaoExecutivo:       string(user.SolicitacaoExecutivo),
		DataSolicitacaoExecutivo:   user.DataSolicitacaoExecutivo,
		MotivoSolicitacaoExecutivo: user.MotivoSolicitacaoExecutivo,
		LojaUsuarioResponse: json.LojaUsuarioResponse{
			Id:                    user.Loja.ID,
			Nome:                  user.Loja.Nome,
			Logo:                  user.Loja.Imagem,
			CupomDestaqueResponse: cupomResp,
		},
	}
	gerais, moedasLoja, err := MoedasUsuarioParaJSON(user.ID)
	if err != nil {
		return nil, err
	}
	resp.MoedasGerais = gerais
	resp.MoedasPorLoja = moedasLoja
	return resp, nil
}

// LoginWeb realiza o login para a plataforma web (portal)
// Apenas usuários do tipo "administrativo" e "customer" (lojistas) podem fazer login
// Customers precisam estar aprovados
func LoginWeb(req json.LoginRequest) (*json.LoginResponse, error) {
	// Busca o usuário por email
	user, err := datasource.GetUserByEmailOnly(req.Email)
	if err != nil {
		return nil, errors.New("usuário não encontrado")
	}

	// Verifica se o tipo de usuário é permitido para login web (portal)
	// Administrativo e Customer/Lojistas acessam o portal web
	// Mobile e Executivo acessam o app mobile
	tiposPermitidos := map[models.TipoUsuario]bool{
		models.TipoUsuarioAdministrativo: true,
		models.TipoUsuarioCustomer:       true,
	}

	if !tiposPermitidos[user.Tipo] {
		switch user.Tipo {
		case models.TipoUsuarioMobile:
			return nil, errors.New("usuários do app devem fazer login pelo aplicativo mobile")
		case models.TipoUsuarioExecutivo:
			return nil, errors.New("executivos devem fazer login pelo aplicativo mobile")
		default:
			return nil, errors.New("tipo de usuário não permitido para login web")
		}
	}

	// Para customers/lojistas, verifica se está aprovado
	if user.Tipo == models.TipoUsuarioCustomer {
		if user.Status == models.StatusUsuarioPendente {
			return nil, errors.New("sua conta está pendente de aprovação")
		}
		if user.Status == models.StatusUsuarioRejeitado {
			return nil, errors.New("sua conta foi rejeitada")
		}
	}

	// Monta a resposta
	var cupomResp json.CupomDestaqueResponse
	if user.Loja.ID != 0 {
		cupom, err := datasource.GetCupomDestaqueByLojaID(user.Loja.ID)
		if err == nil {
			cupomResp = json.CupomDestaqueResponse{
				ID:        cupom.ID,
				Titulo:    cupom.Titulo,
				Descricao: cupom.Descricao,
				Preco:     cupom.Preco,
				Imagem:    cupom.Imagem,
				TipoCupom: cupom.TipoCupom,
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
		IDExecutivo:                user.IDExecutivo,
		SolicitacaoExecutivo:       string(user.SolicitacaoExecutivo),
		DataSolicitacaoExecutivo:   user.DataSolicitacaoExecutivo,
		MotivoSolicitacaoExecutivo: user.MotivoSolicitacaoExecutivo,
		LojaUsuarioResponse: json.LojaUsuarioResponse{
			Id:                    user.Loja.ID,
			Nome:                  user.Loja.Nome,
			Logo:                  user.Loja.Imagem,
			CupomDestaqueResponse: cupomResp,
		},
	}
	gerais, moedasLoja, err := MoedasUsuarioParaJSON(user.ID)
	if err != nil {
		return nil, err
	}
	resp.MoedasGerais = gerais
	resp.MoedasPorLoja = moedasLoja
	return resp, nil
}
