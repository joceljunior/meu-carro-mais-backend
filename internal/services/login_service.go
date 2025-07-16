package services

import (
	"meu-carro-mais/internal/database/datasource"
	"meu-carro-mais/internal/handlers/json"
)

func Login(req json.LoginRequest) (*json.LoginResponse, error) {
	user, err := datasource.GetUserByEmail(req.Email, req.Senha)
	if err != nil {
		return nil, err
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
		NomePlano: user.Plano.Nome,
		LojaUsuarioResponse: json.LojaUsuarioResponse{
			Id:                      user.Loja.ID,
			Nome:                    user.Loja.Nome,
			AnuncioDestaqueResponse: anuncioResp,
		},
	}
	return resp, nil
}
