package handlers

import (
	"encoding/json"
	jr "meu-carro-mais/internal/handlers/json"
	"meu-carro-mais/internal/services"
	"net/http"
)

// CreateUserHandler godoc
// @Summary      Criação do usuário
// @Description  Cria um novo usuário
// @Tags         Criação de Usuário
// @Accept       json
// @Produce      json
// @Param        request body json.UserRequest true "Dados do usuário"
// @Success      201  {string}  string "Usuário criado com sucesso"
// @Failure      400  {string}  string "Dados inválidos"
// @Router       /createuser [post]
func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var req jr.UserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Dados inválidos", http.StatusBadRequest)
		return
	}

	resp, err := services.CreateUser(req)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
