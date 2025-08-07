package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"meu-carro-mais/internal/services"
)

// GetServicosByProximidadeHandler godoc
// @Summary      Lista serviços por proximidade
// @Description  Retorna lista de serviços ordenados por proximidade do usuário
// @Tags         Serviços
// @Accept       json
// @Produce      json
// @Param        latitude query number true "Latitude do usuário"
// @Param        longitude query number true "Longitude do usuário"
// @Success      200  {object}  json.ServicosResponse
// @Failure      400  {string}  string "Parâmetros inválidos"
// @Failure      500  {string}  string "Erro interno do servidor"
// @Router       /servicos/proximidade [get]
func GetServicosByProximidadeHandler(w http.ResponseWriter, r *http.Request) {
	// Obtém os parâmetros de query
	latitudeStr := r.URL.Query().Get("latitude")
	longitudeStr := r.URL.Query().Get("longitude")

	if latitudeStr == "" || longitudeStr == "" {
		http.Error(w, "Latitude e longitude são obrigatórios", http.StatusBadRequest)
		return
	}

	// Converte para float64
	latitude, err := strconv.ParseFloat(latitudeStr, 64)
	if err != nil {
		http.Error(w, "Latitude inválida", http.StatusBadRequest)
		return
	}

	longitude, err := strconv.ParseFloat(longitudeStr, 64)
	if err != nil {
		http.Error(w, "Longitude inválida", http.StatusBadRequest)
		return
	}

	// Chama o serviço
	resp, err := services.GetServicosByProximidade(latitude, longitude)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Retorna a resposta
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
