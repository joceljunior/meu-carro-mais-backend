package services

import (
	"fmt"

	"meu-carro-mais/internal/database/datasource"
	"meu-carro-mais/internal/database/models"
	"meu-carro-mais/internal/handlers/json"
)

// BonificacaoAprovacaoCustomer é o valor de moedas bonificadas ao executivo quando um customer é aprovado
const BonificacaoAprovacaoCustomer = 100

func CreateUser(req json.UserRequest) (*json.UserResponse, error) {
	user, err := datasource.CreateNewUser(req)
	if err != nil {
		return nil, err
	}

	response := &json.UserResponse{
		ID:             user.ID,
		Nome:           user.Nome,
		Email:          user.Email,
		CPF:            user.CPF,
		Imagem:         user.Imagem,
		Telefone:       user.Telefone,
		Endereco:       user.Endereco,
		DataNascimento: user.DataNascimento,
		DataCadastro:   user.DataCadastro,
		Ativo:          user.Ativo,
		Latitude:       user.Latitude,
		Longitude:      user.Longitude,
		IDPlano:        user.IDPlano,
		IDLoja:         user.IDLoja,
		Tipo:           string(user.Tipo),
		Status:         string(user.Status),
		Mensagem:       "Usuário criado com sucesso",
	}

	return response, nil
}

// GetUserByID busca um usuário por ID
func GetUserByID(id uint) (*json.UserResponse, error) {
	user, err := datasource.GetUserByID(id)
	if err != nil {
		return nil, err
	}

	var lojaResponse *json.LojaUsuarioResponse
	if user.Loja.ID != 0 {
		lojaResponse = &json.LojaUsuarioResponse{
			Id:   user.Loja.ID,
			Nome: user.Loja.Nome,
			Logo: user.Loja.Imagem,
		}
	}

	// Monta a resposta da loja indicadora (se existir)
	var lojaIndicadoraResponse *json.LojaUsuarioResponse
	if user.LojaIndicadora != nil && user.LojaIndicadora.ID != 0 {
		lojaIndicadoraResponse = &json.LojaUsuarioResponse{
			Id:   user.LojaIndicadora.ID,
			Nome: user.LojaIndicadora.Nome,
			Logo: user.LojaIndicadora.Imagem,
		}
	}

	response := &json.UserResponse{
		ID:                         user.ID,
		Nome:                       user.Nome,
		Email:                      user.Email,
		CPF:                        user.CPF,
		Imagem:                     user.Imagem,
		Telefone:                   user.Telefone,
		Endereco:                   user.Endereco,
		DataNascimento:             user.DataNascimento,
		DataCadastro:               user.DataCadastro,
		Ativo:                      user.Ativo,
		Latitude:                   user.Latitude,
		Longitude:                  user.Longitude,
		IDPlano:                    user.IDPlano,
		IDLoja:                     user.IDLoja,
		Tipo:                       string(user.Tipo),
		Status:                     string(user.Status),
		SolicitacaoExecutivo:       string(user.SolicitacaoExecutivo),
		DataSolicitacaoExecutivo:   user.DataSolicitacaoExecutivo,
		MotivoSolicitacaoExecutivo: user.MotivoSolicitacaoExecutivo,
		IDLojaIndicadora:           user.IDLojaIndicadora,
		DataVinculoLoja:            user.DataVinculoLoja,
		LojaIndicadora:             lojaIndicadoraResponse,
		Loja:                       lojaResponse,
		Mensagem:                   "Usuário encontrado com sucesso",
	}

	return response, nil
}

// GetAllUsers retorna todos os usuários ativos
func GetAllUsers() ([]json.UserResponse, error) {
	users, err := datasource.GetAllUsers()
	if err != nil {
		return nil, err
	}

	var responses []json.UserResponse
	for _, user := range users {
		var lojaResponse *json.LojaUsuarioResponse
		if user.Loja.ID != 0 {
			lojaResponse = &json.LojaUsuarioResponse{
				Id:   user.Loja.ID,
				Nome: user.Loja.Nome,
				Logo: user.Loja.Imagem,
			}
		}

		response := json.UserResponse{
			ID:             user.ID,
			Nome:           user.Nome,
			Email:          user.Email,
			CPF:            user.CPF,
			Imagem:         user.Imagem,
			Telefone:       user.Telefone,
			Endereco:       user.Endereco,
			DataNascimento: user.DataNascimento,
			DataCadastro:   user.DataCadastro,
			Ativo:          user.Ativo,
			Latitude:       user.Latitude,
			Longitude:      user.Longitude,
			IDPlano:        user.IDPlano,
			IDLoja:         user.IDLoja,
			Tipo:           string(user.Tipo),
			Status:         string(user.Status),
			Loja:           lojaResponse,
		}
		responses = append(responses, response)
	}

	return responses, nil
}

// UpdateUser atualiza um usuário existente
func UpdateUser(id uint, req json.UserRequest) (*json.UserResponse, error) {
	user, err := datasource.UpdateUser(id, req)
	if err != nil {
		return nil, err
	}

	var lojaResponse *json.LojaUsuarioResponse
	if user.Loja.ID != 0 {
		lojaResponse = &json.LojaUsuarioResponse{
			Id:   user.Loja.ID,
			Nome: user.Loja.Nome,
			Logo: user.Loja.Imagem,
		}
	}

	response := &json.UserResponse{
		ID:             user.ID,
		Nome:           user.Nome,
		Email:          user.Email,
		CPF:            user.CPF,
		Imagem:         user.Imagem,
		Telefone:       user.Telefone,
		Endereco:       user.Endereco,
		DataNascimento: user.DataNascimento,
		DataCadastro:   user.DataCadastro,
		Ativo:          user.Ativo,
		Latitude:       user.Latitude,
		Longitude:      user.Longitude,
		IDPlano:        user.IDPlano,
		IDLoja:         user.IDLoja,
		Tipo:           string(user.Tipo),
		Status:         string(user.Status),
		Loja:           lojaResponse,
		Mensagem:       "Usuário atualizado com sucesso",
	}

	return response, nil
}

// SoftDeleteUser realiza soft delete do usuário
func SoftDeleteUser(id uint) error {
	return datasource.SoftDeleteUser(id)
}

// RestoreUser restaura um usuário que foi soft deleted
func RestoreUser(id uint) error {
	return datasource.RestoreUser(id)
}

// GetUserPlanStatus verifica o status do plano de um usuário
func GetUserPlanStatus(id uint) (*json.UserPlanStatusResponse, error) {
	usuario, historico, err := datasource.GetUserPlanStatus(id)
	if err != nil {
		return nil, err
	}

	// Determina se o usuário é premium baseado no ID do plano
	isPremium := usuario.IDPlano > 1 // Assumindo que ID 1 é gratuito e IDs > 1 são premium

	response := &json.UserPlanStatusResponse{
		IDUsuario:    usuario.ID,
		NomeUsuario:  usuario.Nome,
		EmailUsuario: usuario.Email,
		IDPlano:      usuario.IDPlano,
		NomePlano:    usuario.Plano.Nome,
		IsPremium:    isPremium,
		Mensagem:     "Status do plano verificado com sucesso",
	}

	// Se há histórico de pagamento, adiciona informações adicionais
	if historico != nil {
		response.DataVencimento = historico.DataVencimento
		response.StatusPagamento = historico.Status
	}

	return response, nil
}

// CreateAdministrativo cria um novo usuário administrativo
func CreateAdministrativo(req json.AdministrativoRequest) (*json.CustomerResponse, error) {
	user, err := datasource.CreateAdministrativo(req)
	if err != nil {
		return nil, err
	}

	response := &json.CustomerResponse{
		ID:             user.ID,
		Nome:           user.Nome,
		Email:          user.Email,
		CPF:            user.CPF,
		Imagem:         user.Imagem,
		Telefone:       user.Telefone,
		Endereco:       user.Endereco,
		DataNascimento: user.DataNascimento,
		DataCadastro:   user.DataCadastro,
		Ativo:          user.Ativo,
		Latitude:       user.Latitude,
		Longitude:      user.Longitude,
		IDPlano:        user.IDPlano,
		IDLoja:         user.IDLoja,
		Tipo:           string(user.Tipo),
		Status:         string(user.Status),
		Mensagem:       "Usuário administrativo criado com sucesso",
	}

	return response, nil
}

// CreateExecutivo cria um novo usuário executivo
func CreateExecutivo(req json.ExecutivoRequest) (*json.CustomerResponse, error) {
	user, err := datasource.CreateExecutivo(req)
	if err != nil {
		return nil, err
	}

	response := &json.CustomerResponse{
		ID:             user.ID,
		Nome:           user.Nome,
		Email:          user.Email,
		CPF:            user.CPF,
		Imagem:         user.Imagem,
		Telefone:       user.Telefone,
		Endereco:       user.Endereco,
		DataNascimento: user.DataNascimento,
		DataCadastro:   user.DataCadastro,
		Ativo:          user.Ativo,
		Latitude:       user.Latitude,
		Longitude:      user.Longitude,
		IDPlano:        user.IDPlano,
		IDLoja:         user.IDLoja,
		Tipo:           string(user.Tipo),
		Status:         string(user.Status),
		Mensagem:       "Usuário executivo criado com sucesso",
	}

	return response, nil
}

// CreateCustomer cria um novo usuário customer (sempre pendente)
func CreateCustomer(req json.CustomerRequest) (*json.CustomerResponse, error) {
	user, err := datasource.CreateCustomer(req)
	if err != nil {
		return nil, err
	}

	var executivoInfo *json.ExecutivoInfo
	if user.Executivo != nil {
		executivoInfo = &json.ExecutivoInfo{
			ID:    user.Executivo.ID,
			Nome:  user.Executivo.Nome,
			Email: user.Executivo.Email,
		}
	}

	response := &json.CustomerResponse{
		ID:             user.ID,
		Nome:           user.Nome,
		Email:          user.Email,
		CPF:            user.CPF,
		Imagem:         user.Imagem,
		Telefone:       user.Telefone,
		Endereco:       user.Endereco,
		DataNascimento: user.DataNascimento,
		DataCadastro:   user.DataCadastro,
		Ativo:          user.Ativo,
		Latitude:       user.Latitude,
		Longitude:      user.Longitude,
		IDPlano:        user.IDPlano,
		IDLoja:         user.IDLoja,
		Tipo:           string(user.Tipo),
		Status:         string(user.Status),
		IDExecutivo:    user.IDExecutivo,
		Executivo:      executivoInfo,
		Mensagem:       "Customer criado com sucesso. Aguardando aprovação.",
	}

	return response, nil
}

// GetAllCustomers retorna todos os customers
func GetAllCustomers() (*json.CustomersListResponse, error) {
	customers, err := datasource.GetAllCustomers()
	if err != nil {
		return nil, err
	}

	return buildCustomersListResponse(customers, "Customers listados com sucesso")
}

// GetCustomersPendentes retorna todos os customers pendentes
func GetCustomersPendentes() (*json.CustomersListResponse, error) {
	customers, err := datasource.GetCustomersPendentes()
	if err != nil {
		return nil, err
	}

	return buildCustomersListResponse(customers, "Customers pendentes listados com sucesso")
}

// getFotoPerfilUsuario busca a foto de perfil de um usuário
func getFotoPerfilUsuario(idUsuario uint) string {
	upload, err := datasource.GetUploadPrincipalByEntidade("usuario", idUsuario)
	if err != nil {
		return ""
	}
	return upload.URL
}

// gerarLinkCompartilhamentoExecutivo gera o link de compartilhamento para um executivo
func gerarLinkCompartilhamentoExecutivo(idExecutivo uint) string {
	// URL base do app - pode ser configurada via environment variable
	baseURL := "https://app.meucarromais.com.br"
	return baseURL + "/executivo/" + string(rune(idExecutivo))
}

// GetAllExecutivos retorna todos os executivos
func GetAllExecutivos() (*json.ExecutivosListResponse, error) {
	executivos, err := datasource.GetAllExecutivos()
	if err != nil {
		return nil, err
	}

	var responses []json.ExecutivoResponse
	for _, executivo := range executivos {
		var lojaResponse *json.LojaUsuarioResponse
		if executivo.Loja.ID != 0 {
			lojaResponse = &json.LojaUsuarioResponse{
				Id:   executivo.Loja.ID,
				Nome: executivo.Loja.Nome,
				Logo: executivo.Loja.Imagem,
			}
		}

		// Busca a foto de perfil do executivo
		fotoPerfil := getFotoPerfilUsuario(executivo.ID)
		
		// Gera o link de compartilhamento
		linkCompartilhamento := fmt.Sprintf("https://app.meucarromais.com.br/executivo/%d", executivo.ID)

		response := json.ExecutivoResponse{
			ID:                   executivo.ID,
			Nome:                 executivo.Nome,
			Email:                executivo.Email,
			CPF:                  executivo.CPF,
			Imagem:               executivo.Imagem,
			FotoPerfil:           fotoPerfil,
			Telefone:             executivo.Telefone,
			Endereco:             executivo.Endereco,
			DataNascimento:       executivo.DataNascimento,
			DataCadastro:         executivo.DataCadastro,
			Ativo:                executivo.Ativo,
			Latitude:             executivo.Latitude,
			Longitude:            executivo.Longitude,
			IDPlano:              executivo.IDPlano,
			IDLoja:               executivo.IDLoja,
			Tipo:                 string(executivo.Tipo),
			Status:               string(executivo.Status),
			LinkCompartilhamento: linkCompartilhamento,
			Loja:                 lojaResponse,
		}
		responses = append(responses, response)
	}

	return &json.ExecutivosListResponse{
		Executivos: responses,
		Total:      len(responses),
		Mensagem:   "Executivos listados com sucesso",
	}, nil
}

// GetCustomersByStatus retorna customers filtrados por status
func GetCustomersByStatus(status string) (*json.CustomersListResponse, error) {
	var statusModel models.StatusUsuario
	switch status {
	case "pendente":
		statusModel = models.StatusUsuarioPendente
	case "aprovado":
		statusModel = models.StatusUsuarioAprovado
	case "rejeitado":
		statusModel = models.StatusUsuarioRejeitado
	default:
		return nil, nil
	}

	customers, err := datasource.GetCustomersByStatus(statusModel)
	if err != nil {
		return nil, err
	}

	return buildCustomersListResponse(customers, "Customers filtrados com sucesso")
}

// AprovarCustomer aprova um customer e bonifica o executivo se houver
func AprovarCustomer(id uint, motivo string) (*json.AprovacaoResponse, error) {
	// Busca o customer antes de aprovar para pegar o executivo
	customer, err := datasource.GetUserByID(id)
	if err != nil {
		return nil, err
	}

	// Aprova o customer
	_, err = datasource.AprovarCustomer(id)
	if err != nil {
		return nil, err
	}

	response := &json.AprovacaoResponse{
		ID:       id,
		Status:   string(models.StatusUsuarioAprovado),
		Motivo:   motivo,
		Mensagem: "Customer aprovado com sucesso",
	}

	// Se o customer tem um executivo vinculado, bonifica
	if customer.IDExecutivo != nil {
		// Busca a carteira do executivo
		carteira, err := datasource.GetCarteiraByUsuarioID(*customer.IDExecutivo)
		if err == nil && carteira != nil {
			// Adiciona a bonificação
			_, err = datasource.AdicionarSaldo(carteira.ID, BonificacaoAprovacaoCustomer)
			if err == nil {
				bonificacao := BonificacaoAprovacaoCustomer
				response.BonificacaoExecutivo = &bonificacao
				response.Mensagem = "Customer aprovado com sucesso. Executivo bonificado com " + 
					string(rune(BonificacaoAprovacaoCustomer)) + " moedas"
			}
		}
	}

	return response, nil
}

// RejeitarCustomer rejeita um customer
func RejeitarCustomer(id uint, motivo string) (*json.AprovacaoResponse, error) {
	_, err := datasource.RejeitarCustomer(id)
	if err != nil {
		return nil, err
	}

	response := &json.AprovacaoResponse{
		ID:       id,
		Status:   string(models.StatusUsuarioRejeitado),
		Motivo:   motivo,
		Mensagem: "Customer rejeitado",
	}

	return response, nil
}

// buildCustomersListResponse constrói a resposta de lista de customers
func buildCustomersListResponse(customers []models.Usuario, mensagem string) (*json.CustomersListResponse, error) {
	var responses []json.CustomerResponse
	for _, customer := range customers {
		var lojaResponse *json.LojaUsuarioResponse
		if customer.Loja.ID != 0 {
			lojaResponse = &json.LojaUsuarioResponse{
				Id:   customer.Loja.ID,
				Nome: customer.Loja.Nome,
				Logo: customer.Loja.Imagem,
			}
		}

		var executivoInfo *json.ExecutivoInfo
		if customer.Executivo != nil {
			executivoInfo = &json.ExecutivoInfo{
				ID:    customer.Executivo.ID,
				Nome:  customer.Executivo.Nome,
				Email: customer.Executivo.Email,
			}
		}

		response := json.CustomerResponse{
			ID:             customer.ID,
			Nome:           customer.Nome,
			Email:          customer.Email,
			CPF:            customer.CPF,
			Imagem:         customer.Imagem,
			Telefone:       customer.Telefone,
			Endereco:       customer.Endereco,
			DataNascimento: customer.DataNascimento,
			DataCadastro:   customer.DataCadastro,
			Ativo:          customer.Ativo,
			Latitude:       customer.Latitude,
			Longitude:      customer.Longitude,
			IDPlano:        customer.IDPlano,
			IDLoja:         customer.IDLoja,
			Tipo:           string(customer.Tipo),
			Status:         string(customer.Status),
			IDExecutivo:    customer.IDExecutivo,
			Loja:           lojaResponse,
			Executivo:      executivoInfo,
		}
		responses = append(responses, response)
	}

	return &json.CustomersListResponse{
		Customers: responses,
		Total:     len(responses),
		Mensagem:  mensagem,
	}, nil
}

// =====================================================
// FUNÇÕES PARA SOLICITAÇÃO DE EXECUTIVO
// =====================================================

// SolicitarExecutivo registra a solicitação de um usuário mobile para virar executivo
func SolicitarExecutivo(id uint, motivo string) (*json.SolicitacaoExecutivoResponse, error) {
	user, err := datasource.SolicitarExecutivo(id, motivo)
	if err != nil {
		return nil, err
	}

	response := &json.SolicitacaoExecutivoResponse{
		ID:                         user.ID,
		Nome:                       user.Nome,
		Email:                      user.Email,
		Tipo:                       string(user.Tipo),
		SolicitacaoExecutivo:       string(user.SolicitacaoExecutivo),
		DataSolicitacaoExecutivo:   user.DataSolicitacaoExecutivo,
		MotivoSolicitacaoExecutivo: user.MotivoSolicitacaoExecutivo,
		Mensagem:                   "Solicitação para virar executivo enviada com sucesso. Aguardando aprovação.",
	}

	return response, nil
}

// GetSolicitacoesExecutivoPendentes retorna todas as solicitações pendentes
func GetSolicitacoesExecutivoPendentes() (*json.SolicitacoesExecutivoListResponse, error) {
	usuarios, err := datasource.GetSolicitacoesExecutivoPendentes()
	if err != nil {
		return nil, err
	}

	var responses []json.SolicitacaoExecutivoResponse
	for _, user := range usuarios {
		response := json.SolicitacaoExecutivoResponse{
			ID:                         user.ID,
			Nome:                       user.Nome,
			Email:                      user.Email,
			Tipo:                       string(user.Tipo),
			SolicitacaoExecutivo:       string(user.SolicitacaoExecutivo),
			DataSolicitacaoExecutivo:   user.DataSolicitacaoExecutivo,
			MotivoSolicitacaoExecutivo: user.MotivoSolicitacaoExecutivo,
		}
		responses = append(responses, response)
	}

	return &json.SolicitacoesExecutivoListResponse{
		Solicitacoes: responses,
		Total:        len(responses),
		Mensagem:     "Solicitações pendentes listadas com sucesso",
	}, nil
}

// AprovarSolicitacaoExecutivo aprova a solicitação de um usuário para virar executivo
func AprovarSolicitacaoExecutivo(id uint, motivo string) (*json.SolicitacaoExecutivoResponse, error) {
	user, err := datasource.AprovarSolicitacaoExecutivo(id)
	if err != nil {
		return nil, err
	}

	response := &json.SolicitacaoExecutivoResponse{
		ID:                         user.ID,
		Nome:                       user.Nome,
		Email:                      user.Email,
		Tipo:                       string(user.Tipo),
		SolicitacaoExecutivo:       string(user.SolicitacaoExecutivo),
		DataSolicitacaoExecutivo:   user.DataSolicitacaoExecutivo,
		MotivoSolicitacaoExecutivo: motivo,
		Mensagem:                   "Solicitação aprovada! Usuário agora é executivo.",
	}

	return response, nil
}

// RejeitarSolicitacaoExecutivo rejeita a solicitação de um usuário para virar executivo
func RejeitarSolicitacaoExecutivo(id uint, motivo string) (*json.SolicitacaoExecutivoResponse, error) {
	user, err := datasource.RejeitarSolicitacaoExecutivo(id)
	if err != nil {
		return nil, err
	}

	response := &json.SolicitacaoExecutivoResponse{
		ID:                         user.ID,
		Nome:                       user.Nome,
		Email:                      user.Email,
		Tipo:                       string(user.Tipo),
		SolicitacaoExecutivo:       string(user.SolicitacaoExecutivo),
		DataSolicitacaoExecutivo:   user.DataSolicitacaoExecutivo,
		MotivoSolicitacaoExecutivo: motivo,
		Mensagem:                   "Solicitação rejeitada.",
	}

	return response, nil
}

// CancelarExecutivo cancela um executivo aprovado, revertendo para mobile
func CancelarExecutivo(id uint, motivo string) (*json.SolicitacaoExecutivoResponse, error) {
	user, err := datasource.CancelarExecutivo(id)
	if err != nil {
		return nil, err
	}

	response := &json.SolicitacaoExecutivoResponse{
		ID:                         user.ID,
		Nome:                       user.Nome,
		Email:                      user.Email,
		Tipo:                       string(user.Tipo),
		SolicitacaoExecutivo:       string(user.SolicitacaoExecutivo),
		DataSolicitacaoExecutivo:   user.DataSolicitacaoExecutivo,
		MotivoSolicitacaoExecutivo: motivo,
		Mensagem:                   "Executivo cancelado com sucesso. Usuário revertido para mobile.",
	}

	return response, nil
}