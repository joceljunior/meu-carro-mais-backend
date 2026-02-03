package handlers

import (
	"meu-carro-mais/internal/handlers/json"
	"meu-carro-mais/internal/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateUserHandler godoc
// @Summary      Criação do usuário completo
// @Description  Cria um novo usuário com todos os dados fornecidos
// @Tags         Usuários
// @Accept       json
// @Produce      json
// @Param        request body json.UserRequest true "Dados completos do usuário"
// @Success      201  {object}  json.UserResponse "Usuário criado com sucesso"
// @Failure      400  {object}  map[string]interface{} "Dados inválidos"
// @Failure      500  {object}  map[string]interface{} "Erro interno do servidor"
// @Router       /users [post]
func CreateUserHandler(c *gin.Context) {
	var req json.UserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Dados inválidos",
			"details": err.Error(),
		})
		return
	}

	resp, err := services.CreateUser(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, resp)
}

// GetUserHandler godoc
// @Summary      Busca usuário por ID
// @Description  Retorna os dados de um usuário específico pelo ID
// @Tags         Usuários
// @Accept       json
// @Produce      json
// @Param        id path int true "ID do usuário"
// @Success      200 {object} json.UserResponse "Usuário encontrado"
// @Failure      400 {object} map[string]interface{} "ID inválido"
// @Failure      404 {object} map[string]interface{} "Usuário não encontrado"
// @Failure      500 {object} map[string]interface{} "Erro interno do servidor"
// @Router       /users/{id} [get]
func GetUserHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID inválido",
		})
		return
	}

	resp, err := services.GetUserByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Usuário não encontrado",
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// GetAllUsersHandler godoc
// @Summary      Lista todos os usuários
// @Description  Retorna uma lista com todos os usuários ativos
// @Tags         Usuários
// @Accept       json
// @Produce      json
// @Success      200 {array} json.UserResponse "Lista de usuários"
// @Failure      500 {object} map[string]interface{} "Erro interno do servidor"
// @Router       /users [get]
func GetAllUsersHandler(c *gin.Context) {
	resp, err := services.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// UpdateUserHandler godoc
// @Summary      Atualiza usuário
// @Description  Atualiza os dados de um usuário existente
// @Tags         Usuários
// @Accept       json
// @Produce      json
// @Param        id path int true "ID do usuário"
// @Param        request body json.UserRequest true "Dados atualizados do usuário"
// @Success      200 {object} json.UserResponse "Usuário atualizado com sucesso"
// @Failure      400 {object} map[string]interface{} "Dados inválidos"
// @Failure      404 {object} map[string]interface{} "Usuário não encontrado"
// @Failure      500 {object} map[string]interface{} "Erro interno do servidor"
// @Router       /users/{id} [put]
func UpdateUserHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID inválido",
		})
		return
	}

	var req json.UserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Dados inválidos",
			"details": err.Error(),
		})
		return
	}

	resp, err := services.UpdateUser(uint(id), req)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Usuário não encontrado",
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// SoftDeleteUserHandler godoc
// @Summary      Exclui usuário (soft delete)
// @Description  Realiza soft delete do usuário, marcando como excluído sem remover do banco
// @Tags         Usuários
// @Accept       json
// @Produce      json
// @Param        id path int true "ID do usuário"
// @Success      200 {object} map[string]interface{} "Usuário excluído com sucesso"
// @Failure      400 {object} map[string]interface{} "ID inválido"
// @Failure      404 {object} map[string]interface{} "Usuário não encontrado"
// @Failure      500 {object} map[string]interface{} "Erro interno do servidor"
// @Router       /users/{id} [delete]
func SoftDeleteUserHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID inválido",
		})
		return
	}

	err = services.SoftDeleteUser(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Usuário não encontrado",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Usuário excluído com sucesso",
	})
}

// RestoreUserHandler godoc
// @Summary      Restaura usuário excluído
// @Description  Restaura um usuário que foi soft deleted
// @Tags         Usuários
// @Accept       json
// @Produce      json
// @Param        id path int true "ID do usuário"
// @Success      200 {object} map[string]interface{} "Usuário restaurado com sucesso"
// @Failure      400 {object} map[string]interface{} "ID inválido"
// @Failure      404 {object} map[string]interface{} "Usuário não encontrado"
// @Failure      500 {object} map[string]interface{} "Erro interno do servidor"
// @Router       /users/{id}/restore [post]
func RestoreUserHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID inválido",
		})
		return
	}

	err = services.RestoreUser(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Usuário não encontrado ou não foi excluído",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Usuário restaurado com sucesso",
	})
}

// GetUserPlanStatusHandler godoc
// @Summary Verifica status do plano do usuário
// @Description Retorna informações sobre o plano atual do usuário e se ele é premium
// @Tags Usuários
// @Accept json
// @Produce json
// @Param id path int true "ID do usuário"
// @Success 200 {object} json.UserPlanStatusResponse "Status do plano do usuário"
// @Failure 400 {object} map[string]interface{} "ID inválido"
// @Failure 404 {object} map[string]interface{} "Usuário não encontrado"
// @Failure 500 {object} map[string]interface{} "Erro interno do servidor"
// @Router /users/{id}/plan-status [get]
func GetUserPlanStatusHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID inválido",
		})
		return
	}

	resp, err := services.GetUserPlanStatus(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Usuário não encontrado",
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// =====================================================
// ENDPOINTS ADMINISTRATIVO
// =====================================================

// CreateAdministrativoHandler godoc
// @Summary      Criar usuário administrativo
// @Description  Cria um novo usuário do tipo administrativo com todos os poderes do sistema
// @Tags         Administrativo
// @Accept       json
// @Produce      json
// @Param        request body json.AdministrativoRequest true "Dados do usuário administrativo"
// @Success      201  {object}  json.CustomerResponse "Usuário administrativo criado com sucesso"
// @Failure      400  {object}  map[string]interface{} "Dados inválidos"
// @Failure      500  {object}  map[string]interface{} "Erro interno do servidor"
// @Router       /users/administrativo [post]
func CreateAdministrativoHandler(c *gin.Context) {
	var req json.AdministrativoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Dados inválidos",
			"details": err.Error(),
		})
		return
	}

	resp, err := services.CreateAdministrativo(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, resp)
}

// CreateExecutivoHandler godoc
// @Summary      Criar usuário executivo
// @Description  Cria um novo usuário do tipo executivo, que pode criar customers e receber bonificação quando aprovados
// @Tags         Administrativo
// @Accept       json
// @Produce      json
// @Param        request body json.ExecutivoRequest true "Dados do usuário executivo"
// @Success      201  {object}  json.CustomerResponse "Usuário executivo criado com sucesso"
// @Failure      400  {object}  map[string]interface{} "Dados inválidos"
// @Failure      500  {object}  map[string]interface{} "Erro interno do servidor"
// @Router       /users/executivo [post]
func CreateExecutivoHandler(c *gin.Context) {
	var req json.ExecutivoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Dados inválidos",
			"details": err.Error(),
		})
		return
	}

	resp, err := services.CreateExecutivo(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, resp)
}

// GetAllCustomersHandler godoc
// @Summary      Lista todos os customers
// @Description  Retorna uma lista com todos os customers (usuários que podem criar lojas/produtos)
// @Tags         Administrativo
// @Accept       json
// @Produce      json
// @Param        status query string false "Filtrar por status: pendente, aprovado, rejeitado"
// @Success      200 {object} json.CustomersListResponse "Lista de customers"
// @Failure      500 {object} map[string]interface{} "Erro interno do servidor"
// @Router       /users/customers [get]
func GetAllCustomersHandler(c *gin.Context) {
	status := c.Query("status")

	var resp *json.CustomersListResponse
	var err error

	if status != "" {
		resp, err = services.GetCustomersByStatus(status)
	} else {
		resp, err = services.GetAllCustomers()
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// GetAllExecutivosHandler godoc
// @Summary      Lista todos os executivos
// @Description  Retorna uma lista com todos os usuários do tipo executivo
// @Tags         Administrativo
// @Accept       json
// @Produce      json
// @Success      200 {object} json.ExecutivosListResponse "Lista de executivos"
// @Failure      500 {object} map[string]interface{} "Erro interno do servidor"
// @Router       /users/executivos [get]
func GetAllExecutivosHandler(c *gin.Context) {
	resp, err := services.GetAllExecutivos()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// GetCustomersPendentesHandler godoc
// @Summary      Lista customers pendentes
// @Description  Retorna uma lista com todos os customers aguardando aprovação
// @Tags         Administrativo
// @Accept       json
// @Produce      json
// @Success      200 {object} json.CustomersListResponse "Lista de customers pendentes"
// @Failure      500 {object} map[string]interface{} "Erro interno do servidor"
// @Router       /users/customers/pendentes [get]
func GetCustomersPendentesHandler(c *gin.Context) {
	resp, err := services.GetCustomersPendentes()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// AprovarCustomerHandler godoc
// @Summary      Aprova um customer
// @Description  Aprova um customer pendente. Se o customer foi criado por um executivo, o executivo recebe bonificação em moedas
// @Tags         Administrativo
// @Accept       json
// @Produce      json
// @Param        id path int true "ID do customer"
// @Param        request body json.AprovarCustomerRequest false "Dados da aprovação (motivo opcional)"
// @Success      200 {object} json.AprovacaoResponse "Customer aprovado com sucesso"
// @Failure      400 {object} map[string]interface{} "ID inválido ou customer não está pendente"
// @Failure      404 {object} map[string]interface{} "Customer não encontrado"
// @Failure      500 {object} map[string]interface{} "Erro interno do servidor"
// @Router       /users/customers/{id}/aprovar [post]
func AprovarCustomerHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID inválido",
		})
		return
	}

	var req json.AprovarCustomerRequest
	// Ignora erro se body estiver vazio (motivo é opcional)
	c.ShouldBindJSON(&req)

	resp, err := services.AprovarCustomer(uint(id), req.Motivo)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// RejeitarCustomerHandler godoc
// @Summary      Rejeita um customer
// @Description  Rejeita um customer pendente com motivo obrigatório
// @Tags         Administrativo
// @Accept       json
// @Produce      json
// @Param        id path int true "ID do customer"
// @Param        request body json.RejeitarCustomerRequest true "Dados da rejeição (motivo obrigatório)"
// @Success      200 {object} json.AprovacaoResponse "Customer rejeitado"
// @Failure      400 {object} map[string]interface{} "ID inválido, motivo não informado ou customer não está pendente"
// @Failure      404 {object} map[string]interface{} "Customer não encontrado"
// @Failure      500 {object} map[string]interface{} "Erro interno do servidor"
// @Router       /users/customers/{id}/rejeitar [post]
func RejeitarCustomerHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID inválido",
		})
		return
	}

	var req json.RejeitarCustomerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Motivo da rejeição é obrigatório",
			"details": err.Error(),
		})
		return
	}

	resp, err := services.RejeitarCustomer(uint(id), req.Motivo)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// =====================================================
// ENDPOINTS CUSTOMER
// =====================================================

// CreateCustomerHandler godoc
// @Summary      Criar usuário customer
// @Description  Cria um novo usuário do tipo customer. O customer começa com status pendente e precisa ser aprovado por um administrativo. Opcionalmente pode ser vinculado a um executivo.
// @Tags         Customers
// @Accept       json
// @Produce      json
// @Param        request body json.CustomerRequest true "Dados do customer"
// @Success      201  {object}  json.CustomerResponse "Customer criado com sucesso (pendente aprovação)"
// @Failure      400  {object}  map[string]interface{} "Dados inválidos"
// @Failure      500  {object}  map[string]interface{} "Erro interno do servidor"
// @Router       /users/customer [post]
func CreateCustomerHandler(c *gin.Context) {
	var req json.CustomerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Dados inválidos",
			"details": err.Error(),
		})
		return
	}

	resp, err := services.CreateCustomer(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, resp)
}

// =====================================================
// ENDPOINTS SOLICITAÇÃO DE EXECUTIVO
// =====================================================

// SolicitarExecutivoHandler godoc
// @Summary      Solicitar virar executivo
// @Description  Permite que um usuário mobile solicite se tornar executivo. A solicitação fica pendente até ser aprovada por um administrativo.
// @Tags         Usuários
// @Accept       json
// @Produce      json
// @Param        id path int true "ID do usuário mobile"
// @Param        request body json.SolicitarExecutivoRequest true "Motivo/justificativa da solicitação"
// @Success      200 {object} json.SolicitacaoExecutivoResponse "Solicitação enviada com sucesso"
// @Failure      400 {object} map[string]interface{} "ID inválido, usuário não é mobile ou já tem solicitação pendente"
// @Failure      404 {object} map[string]interface{} "Usuário não encontrado"
// @Failure      500 {object} map[string]interface{} "Erro interno do servidor"
// @Router       /users/{id}/solicitar-executivo [post]
func SolicitarExecutivoHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID inválido",
		})
		return
	}

	var req json.SolicitarExecutivoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Motivo da solicitação é obrigatório",
			"details": err.Error(),
		})
		return
	}

	resp, err := services.SolicitarExecutivo(uint(id), req.Motivo)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// GetSolicitacoesExecutivoPendentesHandler godoc
// @Summary      Lista solicitações de executivo pendentes
// @Description  Retorna uma lista com todas as solicitações de usuários mobile que querem virar executivo
// @Tags         Administrativo
// @Accept       json
// @Produce      json
// @Success      200 {object} json.SolicitacoesExecutivoListResponse "Lista de solicitações pendentes"
// @Failure      500 {object} map[string]interface{} "Erro interno do servidor"
// @Router       /users/solicitacoes-executivo [get]
func GetSolicitacoesExecutivoPendentesHandler(c *gin.Context) {
	resp, err := services.GetSolicitacoesExecutivoPendentes()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// AprovarSolicitacaoExecutivoHandler godoc
// @Summary      Aprova solicitação de executivo
// @Description  Aprova a solicitação de um usuário mobile para virar executivo. O usuário passa a ter tipo 'executivo'.
// @Tags         Administrativo
// @Accept       json
// @Produce      json
// @Param        id path int true "ID do usuário"
// @Param        request body json.AprovarSolicitacaoExecutivoRequest false "Dados da aprovação (motivo opcional)"
// @Success      200 {object} json.SolicitacaoExecutivoResponse "Solicitação aprovada"
// @Failure      400 {object} map[string]interface{} "ID inválido ou usuário não tem solicitação pendente"
// @Failure      404 {object} map[string]interface{} "Usuário não encontrado"
// @Failure      500 {object} map[string]interface{} "Erro interno do servidor"
// @Router       /users/{id}/aprovar-executivo [post]
func AprovarSolicitacaoExecutivoHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID inválido",
		})
		return
	}

	var req json.AprovarSolicitacaoExecutivoRequest
	// Ignora erro se body estiver vazio (motivo é opcional)
	c.ShouldBindJSON(&req)

	resp, err := services.AprovarSolicitacaoExecutivo(uint(id), req.Motivo)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// RejeitarSolicitacaoExecutivoHandler godoc
// @Summary      Rejeita solicitação de executivo
// @Description  Rejeita a solicitação de um usuário mobile para virar executivo. O usuário continua como mobile.
// @Tags         Administrativo
// @Accept       json
// @Produce      json
// @Param        id path int true "ID do usuário"
// @Param        request body json.RejeitarSolicitacaoExecutivoRequest true "Dados da rejeição (motivo obrigatório)"
// @Success      200 {object} json.SolicitacaoExecutivoResponse "Solicitação rejeitada"
// @Failure      400 {object} map[string]interface{} "ID inválido, motivo não informado ou usuário não tem solicitação pendente"
// @Failure      404 {object} map[string]interface{} "Usuário não encontrado"
// @Failure      500 {object} map[string]interface{} "Erro interno do servidor"
// @Router       /users/{id}/rejeitar-executivo [post]
func RejeitarSolicitacaoExecutivoHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID inválido",
		})
		return
	}

	var req json.RejeitarSolicitacaoExecutivoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Motivo da rejeição é obrigatório",
			"details": err.Error(),
		})
		return
	}

	resp, err := services.RejeitarSolicitacaoExecutivo(uint(id), req.Motivo)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}