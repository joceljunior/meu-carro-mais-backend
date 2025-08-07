package json

type LojaResponse struct {
	ID          uint    `json:"id"`
	Nome        string  `json:"nome"`
	CNPJ        string  `json:"cnpj"`
	Imagem      string  `json:"imagem"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
	IDCategoria uint    `json:"id_categoria"`
	Categoria   string  `json:"categoria"`
	Distancia   float64 `json:"distancia,omitempty"` // Distância em km
}

type LojasResponse struct {
	Lojas []LojaResponse `json:"lojas"`
	Total int            `json:"total"`
} 