package json

import "time"

type RegistroInteresseResponse struct {
	ID              uint                  `json:"id"`
	IDAnuncio       uint                  `json:"id_anuncio"`
	Nome            string                `json:"nome"`
	Email           string                `json:"email"`
	Telefone        string                `json:"telefone"`
	Mensagem        string                `json:"mensagem"`
	DataCadastro    time.Time             `json:"data_cadastro"`
	DataAtualizacao time.Time             `json:"data_atualizacao"`
	Anuncio         *AnuncioResponse      `json:"anuncio,omitempty"`
}

type RegistrosInteresseResponse struct {
	RegistrosInteresse []RegistroInteresseResponse `json:"registros_interesse"`
	Total              int                        `json:"total"`
}

