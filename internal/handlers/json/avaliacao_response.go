package json

import "time"

type AvaliacaoResponse struct {
	ID              uint      `json:"id"`
	IDUsuario       uint      `json:"id_usuario"`
	IDLoja          uint      `json:"id_loja"`
	Nota            int       `json:"nota"`
	Comentario      string    `json:"comentario"`
	DataAvaliacao   time.Time `json:"data_avaliacao"`
	DataAtualizacao time.Time `json:"data_atualizacao"`

	// Dados relacionados
	Usuario UserResponse `json:"usuario,omitempty"`
	Loja    LojaResponse `json:"loja,omitempty"`
}

type AvaliacoesResponse struct {
	Avaliacoes []AvaliacaoResponse `json:"avaliacoes"`
	Total      int                 `json:"total"`
	MediaNota  float64             `json:"media_nota,omitempty"`
}

type AvaliacaoEstatisticasResponse struct {
	TotalAvaliacoes int     `json:"total_avaliacoes"`
	MediaNota       float64 `json:"media_nota"`
	Nota1           int     `json:"nota_1"`
	Nota2           int     `json:"nota_2"`
	Nota3           int     `json:"nota_3"`
	Nota4           int     `json:"nota_4"`
	Nota5           int     `json:"nota_5"`
}
