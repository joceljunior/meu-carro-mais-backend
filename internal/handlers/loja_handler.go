package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"meu-carro-mais/internal/services"
)

// GetLojasByProximidadeHandler godoc
// @Summary      Lista lojas por proximidade
// @Description  Retorna lista de lojas ordenadas por proximidade do usuário
// @Tags         Lojas
// @Accept       json
// @Produce      json
// @Param        latitude query number true "Latitude do usuário"
// @Param        longitude query number true "Longitude do usuário"
// @Success      200  {object}  json.LojasResponse
// @Failure      400  {string}  string "Parâmetros inválidos"
// @Failure      500  {string}  string "Erro interno do servidor"
// @Router       /lojas/proximidade [get]
func GetLojasByProximidadeHandler(w http.ResponseWriter, r *http.Request) {
	// Obtém os parâmetros da query string
	latStr := r.URL.Query().Get("latitude")
	lngStr := r.URL.Query().Get("longitude")

	if latStr == "" || lngStr == "" {
		http.Error(w, "Parâmetros latitude e longitude são obrigatórios", http.StatusBadRequest)
		return
	}

	// Converte para float64
	latitude, err := strconv.ParseFloat(latStr, 64)
	if err != nil {
		http.Error(w, "Latitude inválida", http.StatusBadRequest)
		return
	}

	longitude, err := strconv.ParseFloat(lngStr, 64)
	if err != nil {
		http.Error(w, "Longitude inválida", http.StatusBadRequest)
		return
	}

	// Valida as coordenadas
	if latitude < -90 || latitude > 90 {
		http.Error(w, "Latitude deve estar entre -90 e 90", http.StatusBadRequest)
		return
	}

	if longitude < -180 || longitude > 180 {
		http.Error(w, "Longitude deve estar entre -180 e 180", http.StatusBadRequest)
		return
	}

	// Busca as lojas
	resp, err := services.GetLojasByProximidade(latitude, longitude)
	if err != nil {
		http.Error(w, "Erro ao buscar lojas: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Retorna a resposta
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
} 