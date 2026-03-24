package handlers

import (
	"encoding/json"
	jsonresp "meu-carro-mais/internal/handlers/json"
	"meu-carro-mais/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// LoginHandler godoc
// @Summary      Login Mobile
// @Description  Login para o aplicativo mobile. Valida se o email existe. Se existir, retorna os dados do usuário. Se não existir, cria um novo usuário mobile. Usuários do tipo 'customer' não podem fazer login no mobile. A resposta inclui moedas_gerais (carteira) e moedas_por_loja (saldos por loja).
// @Tags         Autenticação
// @Accept       json
// @Produce      json
// @Param        request body json.LoginRequest true "Dados de login (email e senha)"
// @Success      200  {object}  json.LoginResponse
// @Failure      400  {string}  string "Dados inválidos"
// @Failure      401  {string}  string "Não autorizado"
// @Failure      500  {string}  string "Erro interno do servidor"
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

// LoginWebHandler godoc
// @Summary      Login Web
// @Description  Login para a plataforma web. Apenas usuários do tipo 'executivo', 'administrativo' e 'customer' podem fazer login. Customers precisam estar aprovados. A resposta inclui moedas_gerais (carteira) e moedas_por_loja (saldos por loja).
// @Tags         Autenticação
// @Accept       json
// @Produce      json
// @Param        request body json.LoginRequest true "Dados de login (email e senha)"
// @Success      200  {object}  json.LoginResponse "Login realizado com sucesso"
// @Failure      400  {object}  map[string]interface{} "Dados inválidos"
// @Failure      401  {object}  map[string]interface{} "Não autorizado"
// @Failure      404  {object}  map[string]interface{} "Usuário não encontrado"
// @Router       /login/web [post]
func LoginWebHandler(c *gin.Context) {
	var req jsonresp.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Dados inválidos",
			"details": err.Error(),
		})
		return
	}

	resp, err := services.LoginWeb(req)
	if err != nil {
		// Verifica o tipo de erro para retornar o status correto
		errorMsg := err.Error()
		if errorMsg == "usuário não encontrado" {
			c.JSON(http.StatusNotFound, gin.H{
				"error": errorMsg,
			})
			return
		}
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": errorMsg,
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}
