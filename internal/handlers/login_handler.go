package handlers

import (
	"encoding/json"
	jsonresp "meu-carro-mais/internal/handlers/json"
	"meu-carro-mais/internal/services"
	"net/http"
)

// LoginHandler godoc
// @Summary      Login do usuário
// @Description  Autentica um usuário e retorna seus dados
// @Tags         Autenticação
// @Accept       json
// @Produce      json
// @Param        request body json.LoginRequest true "Dados de login"
// @Success      200  {object}  json.LoginResponse
// @Failure      400  {string}  string "Dados inválidos"
// @Failure      401  {string}  string "Não autorizado"
// @Router       /login [post]
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var req jsonresp.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Dados inválidos", http.StatusBadRequest)
		return
	}
	resp, err := services.Login(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
