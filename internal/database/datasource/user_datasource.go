package datasource

import (
	"errors"
	"meu-carro-mais/internal/database"
	"meu-carro-mais/internal/database/models"
	"meu-carro-mais/internal/handlers/json"
	"time"
)

func GetUserByEmail(email string, senha string) (*models.Usuario, error) {
	var usuario models.Usuario
	err := database.DB.
		Preload("Loja").
		Preload("Plano").
		Where("email = ? AND senha = ? AND data_exclusao IS NULL", email, senha).
		First(&usuario).Error
	if err != nil {
		return nil, err
	}
	return &usuario, nil
}

// GetUserByEmailOnly busca usuário apenas por email (sem validação de senha)
func GetUserByEmailOnly(email string) (*models.Usuario, error) {
	var usuario models.Usuario
	err := database.DB.
		Preload("Loja").
		Preload("Plano").
		Where("email = ? AND data_exclusao IS NULL", email).
		First(&usuario).Error
	if err != nil {
		return nil, err
	}
	return &usuario, nil
}

func CreateNewUser(json json.UserRequest) (*models.Usuario, error) {
	user := models.Usuario{
		Email:          json.Email,
		Senha:          json.Senha,
		IDLoja:         nil,
		Nome:           json.Nome,
		CPF:            json.CPF,
		Imagem:         json.Imagem,
		Telefone:       json.Telefone,
		Endereco:       json.Endereco,
		DataNascimento: json.DataNascimento,
		Latitude:       json.Latitude,
		Longitude:      json.Longitude,
		IDPlano:        1, // Plano padrão (Gratuito)
		Ativo:          true,
		Tipo:           models.TipoUsuarioMobile, // Tipo padrão é mobile
		Status:         models.StatusUsuarioAprovado,
	}
	err := database.DB.Create(&user).Error
	if err != nil {
		return nil, err
	}

	// Recarrega o usuário com os relacionamentos
	return GetUserByEmailOnly(json.Email)
}

// CreateAdministrativo cria um novo usuário administrativo
func CreateAdministrativo(req json.AdministrativoRequest) (*models.Usuario, error) {
	user := models.Usuario{
		Email:          req.Email,
		Senha:          req.Senha,
		IDLoja:         nil,
		Nome:           req.Nome,
		CPF:            req.CPF,
		Imagem:         req.Imagem,
		Telefone:       req.Telefone,
		Endereco:       req.Endereco,
		DataNascimento: req.DataNascimento,
		Latitude:       req.Latitude,
		Longitude:      req.Longitude,
		IDPlano:        1,
		Ativo:          true,
		Tipo:           models.TipoUsuarioAdministrativo,
		Status:         models.StatusUsuarioAprovado,
	}
	err := database.DB.Create(&user).Error
	if err != nil {
		return nil, err
	}

	return GetUserByEmailOnly(req.Email)
}

// CreateExecutivo cria um novo usuário executivo
func CreateExecutivo(req json.ExecutivoRequest) (*models.Usuario, error) {
	user := models.Usuario{
		Email:          req.Email,
		Senha:          req.Senha,
		IDLoja:         nil,
		Nome:           req.Nome,
		CPF:            req.CPF,
		Imagem:         req.Imagem,
		Telefone:       req.Telefone,
		Endereco:       req.Endereco,
		DataNascimento: req.DataNascimento,
		Latitude:       req.Latitude,
		Longitude:      req.Longitude,
		IDPlano:        1,
		Ativo:          true,
		Tipo:           models.TipoUsuarioExecutivo,
		Status:         models.StatusUsuarioAprovado,
	}
	err := database.DB.Create(&user).Error
	if err != nil {
		return nil, err
	}

	return GetUserByEmailOnly(req.Email)
}

// CreateCustomer cria um novo usuário customer (sempre pendente)
func CreateCustomer(req json.CustomerRequest) (*models.Usuario, error) {
	user := models.Usuario{
		Email:          req.Email,
		Senha:          req.Senha,
		IDLoja:         nil,
		Nome:           req.Nome,
		CPF:            req.CPF,
		Imagem:         req.Imagem,
		Telefone:       req.Telefone,
		Endereco:       req.Endereco,
		DataNascimento: req.DataNascimento,
		Latitude:       req.Latitude,
		Longitude:      req.Longitude,
		IDPlano:        1,
		Ativo:          true,
		Tipo:           models.TipoUsuarioCustomer,
		Status:         models.StatusUsuarioPendente, // Customer sempre começa pendente
		IDExecutivo:    req.IDExecutivo,              // Referência ao executivo que criou (se houver)
	}
	err := database.DB.Create(&user).Error
	if err != nil {
		return nil, err
	}

	return GetUserByID(user.ID)
}

// GetCustomersPendentes retorna todos os customers com status pendente
func GetCustomersPendentes() ([]models.Usuario, error) {
	var usuarios []models.Usuario
	err := database.DB.
		Preload("Loja").
		Preload("Plano").
		Preload("Executivo").
		Where("tipo = ? AND status = ? AND data_exclusao IS NULL", models.TipoUsuarioCustomer, models.StatusUsuarioPendente).
		Order("data_cadastro DESC").
		Find(&usuarios).Error
	if err != nil {
		return nil, err
	}
	return usuarios, nil
}

// GetAllCustomers retorna todos os customers (independente do status)
func GetAllCustomers() ([]models.Usuario, error) {
	var usuarios []models.Usuario
	err := database.DB.
		Preload("Loja").
		Preload("Plano").
		Preload("Executivo").
		Where("tipo = ? AND data_exclusao IS NULL", models.TipoUsuarioCustomer).
		Order("data_cadastro DESC").
		Find(&usuarios).Error
	if err != nil {
		return nil, err
	}
	return usuarios, nil
}

// GetCustomersByStatus retorna customers filtrados por status
func GetCustomersByStatus(status models.StatusUsuario) ([]models.Usuario, error) {
	var usuarios []models.Usuario
	err := database.DB.
		Preload("Loja").
		Preload("Plano").
		Preload("Executivo").
		Where("tipo = ? AND status = ? AND data_exclusao IS NULL", models.TipoUsuarioCustomer, status).
		Order("data_cadastro DESC").
		Find(&usuarios).Error
	if err != nil {
		return nil, err
	}
	return usuarios, nil
}

// AprovarCustomer aprova um customer pendente
func AprovarCustomer(id uint) (*models.Usuario, error) {
	// Busca o usuário
	usuario, err := GetUserByID(id)
	if err != nil {
		return nil, errors.New("customer não encontrado")
	}

	// Verifica se é um customer
	if usuario.Tipo != models.TipoUsuarioCustomer {
		return nil, errors.New("usuário não é um customer")
	}

	// Verifica se está pendente
	if usuario.Status != models.StatusUsuarioPendente {
		return nil, errors.New("customer não está pendente")
	}

	// Atualiza o status para aprovado
	err = database.DB.Model(&models.Usuario{}).
		Where("id = ?", id).
		Update("status", models.StatusUsuarioAprovado).Error
	if err != nil {
		return nil, err
	}

	return GetUserByID(id)
}

// RejeitarCustomer rejeita um customer pendente
func RejeitarCustomer(id uint) (*models.Usuario, error) {
	// Busca o usuário
	usuario, err := GetUserByID(id)
	if err != nil {
		return nil, errors.New("customer não encontrado")
	}

	// Verifica se é um customer
	if usuario.Tipo != models.TipoUsuarioCustomer {
		return nil, errors.New("usuário não é um customer")
	}

	// Verifica se está pendente
	if usuario.Status != models.StatusUsuarioPendente {
		return nil, errors.New("customer não está pendente")
	}

	// Atualiza o status para rejeitado
	err = database.DB.Model(&models.Usuario{}).
		Where("id = ?", id).
		Update("status", models.StatusUsuarioRejeitado).Error
	if err != nil {
		return nil, err
	}

	return GetUserByID(id)
}

// GetExecutivoByCustomerID retorna o executivo que criou um customer
func GetExecutivoByCustomerID(customerID uint) (*models.Usuario, error) {
	// Busca o customer
	customer, err := GetUserByID(customerID)
	if err != nil {
		return nil, err
	}

	// Verifica se tem executivo vinculado
	if customer.IDExecutivo == nil {
		return nil, nil
	}

	// Busca o executivo
	return GetUserByID(*customer.IDExecutivo)
}

// CreateUserFromLogin cria um novo usuário a partir dos dados do login
func CreateUserFromLogin(loginReq json.LoginRequest) (*models.Usuario, error) {
	// Gera um CPF temporário baseado no email para evitar problemas com constraint unique
	tempCPF := "TEMP_" + loginReq.Email

	user := models.Usuario{
		Email:   loginReq.Email,
		Senha:   loginReq.Senha, // Mantém a senha para compatibilidade, mas não será usada para validação
		IDLoja:  nil,
		Nome:    "",      // Será preenchido pelo frontend se necessário
		CPF:     tempCPF, // CPF temporário que pode ser atualizado depois
		Imagem:  "",
		IDPlano: 1,                           // Plano padrão (Gratuito)
		Tipo:    models.TipoUsuarioMobile,    // Tipo padrão é mobile
		Status:  models.StatusUsuarioAprovado,
	}
	err := database.DB.Create(&user).Error
	if err != nil {
		return nil, err
	}

	// Recarrega o usuário com os relacionamentos
	return GetUserByEmailOnly(loginReq.Email)
}

// GetUserByID busca usuário por ID (apenas usuários não excluídos)
func GetUserByID(id uint) (*models.Usuario, error) {
	var usuario models.Usuario
	err := database.DB.
		Preload("Loja").
		Preload("Plano").
		Where("id = ? AND data_exclusao IS NULL", id).
		First(&usuario).Error
	if err != nil {
		return nil, err
	}
	return &usuario, nil
}

// GetAllUsers retorna todos os usuários ativos (não excluídos)
func GetAllUsers() ([]models.Usuario, error) {
	var usuarios []models.Usuario
	err := database.DB.
		Preload("Loja").
		Preload("Plano").
		Where("data_exclusao IS NULL").
		Order("data_cadastro DESC").
		Find(&usuarios).Error
	if err != nil {
		return nil, err
	}
	return usuarios, nil
}

// UpdateUser atualiza um usuário existente
func UpdateUser(id uint, req json.UserRequest) (*models.Usuario, error) {
	// Verifica se o usuário existe e não foi excluído
	usuario, err := GetUserByID(id)
	if err != nil {
		return nil, errors.New("usuário não encontrado")
	}

	// Atualiza os campos
	usuario.Nome = req.Nome
	usuario.Email = req.Email
	usuario.Senha = req.Senha
	usuario.CPF = req.CPF
	usuario.Imagem = req.Imagem
	usuario.Telefone = req.Telefone
	usuario.Endereco = req.Endereco
	usuario.DataNascimento = req.DataNascimento
	usuario.Latitude = req.Latitude
	usuario.Longitude = req.Longitude

	err = database.DB.Save(&usuario).Error
	if err != nil {
		return nil, err
	}

	// Recarrega o usuário com os relacionamentos
	return GetUserByID(id)
}

// SoftDeleteUser realiza soft delete do usuário (marca como excluído)
func SoftDeleteUser(id uint) error {
	// Verifica se o usuário existe e não foi excluído
	_, err := GetUserByID(id)
	if err != nil {
		return errors.New("usuário não encontrado")
	}

	// Atualiza a data de exclusão
	now := time.Now()
	err = database.DB.Model(&models.Usuario{}).
		Where("id = ?", id).
		Update("data_exclusao", now).Error
	if err != nil {
		return err
	}

	return nil
}

// RestoreUser restaura um usuário que foi soft deleted
func RestoreUser(id uint) error {
	var usuario models.Usuario
	err := database.DB.Where("id = ? AND data_exclusao IS NOT NULL", id).First(&usuario).Error
	if err != nil {
		return errors.New("usuário não encontrado ou não foi excluído")
	}

	// Remove a data de exclusão
	err = database.DB.Model(&models.Usuario{}).
		Where("id = ?", id).
		Update("data_exclusao", nil).Error
	if err != nil {
		return err
	}

	return nil
}

// UpdateUserPlano atualiza apenas o plano de um usuário
func UpdateUserPlano(id uint, planoID uint) error {
	// Verifica se o usuário existe
	_, err := GetUserByID(id)
	if err != nil {
		return errors.New("usuário não encontrado")
	}

	// Atualiza apenas o plano
	err = database.DB.Model(&models.Usuario{}).
		Where("id = ?", id).
		Update("id_plano", planoID).Error
	if err != nil {
		return err
	}

	return nil
}

// GetUserPlanStatus busca o status do plano de um usuário
func GetUserPlanStatus(id uint) (*models.Usuario, *models.HistoricoPagamento, error) {
	// Busca o usuário com o plano
	usuario, err := GetUserByID(id)
	if err != nil {
		return nil, nil, errors.New("usuário não encontrado")
	}

	// Busca o último pagamento válido do usuário
	var historico models.HistoricoPagamento
	err = database.DB.
		Where("id_usuario = ? AND status = ? AND data_exclusao IS NULL", id, models.StatusPagamentoCompleted).
		Order("data_criacao DESC").
		First(&historico).Error

	// Se não encontrar histórico de pagamento, retorna apenas o usuário
	if err != nil {
		return usuario, nil, nil
	}

	return usuario, &historico, nil
}

// =====================================================
// FUNÇÕES PARA SOLICITAÇÃO DE EXECUTIVO
// =====================================================

// SolicitarExecutivo registra a solicitação de um usuário mobile para virar executivo
func SolicitarExecutivo(id uint, motivo string) (*models.Usuario, error) {
	// Busca o usuário
	usuario, err := GetUserByID(id)
	if err != nil {
		return nil, errors.New("usuário não encontrado")
	}

	// Verifica se é um usuário mobile
	if usuario.Tipo != models.TipoUsuarioMobile {
		return nil, errors.New("apenas usuários mobile podem solicitar virar executivo")
	}

	// Verifica se já tem uma solicitação pendente
	if usuario.SolicitacaoExecutivo == models.StatusSolicitacaoPendente {
		return nil, errors.New("usuário já possui uma solicitação pendente")
	}

	// Registra a solicitação
	now := time.Now()
	err = database.DB.Model(&models.Usuario{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"solicitacao_executivo":        models.StatusSolicitacaoPendente,
			"data_solicitacao_executivo":   now,
			"motivo_solicitacao_executivo": motivo,
		}).Error
	if err != nil {
		return nil, err
	}

	return GetUserByID(id)
}

// GetSolicitacoesExecutivoPendentes retorna todos os usuários com solicitação de executivo pendente
func GetSolicitacoesExecutivoPendentes() ([]models.Usuario, error) {
	var usuarios []models.Usuario
	err := database.DB.
		Preload("Loja").
		Preload("Plano").
		Where("solicitacao_executivo = ? AND data_exclusao IS NULL", models.StatusSolicitacaoPendente).
		Order("data_solicitacao_executivo DESC").
		Find(&usuarios).Error
	if err != nil {
		return nil, err
	}
	return usuarios, nil
}

// AprovarSolicitacaoExecutivo aprova a solicitação de executivo de um usuário
func AprovarSolicitacaoExecutivo(id uint) (*models.Usuario, error) {
	// Busca o usuário
	usuario, err := GetUserByID(id)
	if err != nil {
		return nil, errors.New("usuário não encontrado")
	}

	// Verifica se tem solicitação pendente
	if usuario.SolicitacaoExecutivo != models.StatusSolicitacaoPendente {
		return nil, errors.New("usuário não possui solicitação pendente")
	}

	// Aprova a solicitação e muda o tipo para executivo
	err = database.DB.Model(&models.Usuario{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"solicitacao_executivo": models.StatusSolicitacaoAprovada,
			"tipo":                  models.TipoUsuarioExecutivo,
		}).Error
	if err != nil {
		return nil, err
	}

	return GetUserByID(id)
}

// RejeitarSolicitacaoExecutivo rejeita a solicitação de executivo de um usuário
func RejeitarSolicitacaoExecutivo(id uint) (*models.Usuario, error) {
	// Busca o usuário
	usuario, err := GetUserByID(id)
	if err != nil {
		return nil, errors.New("usuário não encontrado")
	}

	// Verifica se tem solicitação pendente
	if usuario.SolicitacaoExecutivo != models.StatusSolicitacaoPendente {
		return nil, errors.New("usuário não possui solicitação pendente")
	}

	// Rejeita a solicitação (mantém o tipo mobile)
	err = database.DB.Model(&models.Usuario{}).
		Where("id = ?", id).
		Update("solicitacao_executivo", models.StatusSolicitacaoRejeitada).Error
	if err != nil {
		return nil, err
	}

	return GetUserByID(id)
}
