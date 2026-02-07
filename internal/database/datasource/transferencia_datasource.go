package datasource

import (
	"errors"
	"meu-carro-mais/internal/database"
	"meu-carro-mais/internal/database/models"
	"time"
)

// TransferirVeiculo realiza a transferência de um veículo de um usuário para outro
func TransferirVeiculo(idVeiculo, idUsuarioOrigem, idUsuarioDestino uint, tipoTransferencia models.TipoTransferencia, observacoes string, idLojaVenda, idHistoricoResgate *uint) (*models.TransferenciaVeiculo, error) {
	// Verifica se o veículo existe e pertence ao usuário origem
	veiculo, err := GetVeiculoByID(idVeiculo)
	if err != nil {
		return nil, errors.New("veículo não encontrado")
	}

	if veiculo.IDUsuario != idUsuarioOrigem {
		return nil, errors.New("veículo não pertence ao usuário de origem")
	}

	// Verifica se o usuário destino existe
	_, err = GetUserByID(idUsuarioDestino)
	if err != nil {
		return nil, errors.New("usuário destino não encontrado")
	}

	// Verifica se o usuário destino é diferente do origem
	if idUsuarioOrigem == idUsuarioDestino {
		return nil, errors.New("usuário de origem e destino não podem ser o mesmo")
	}

	// Inicia transação
	tx := database.DB.Begin()

	// Atualiza o veículo com o novo dono
	err = tx.Model(&models.Veiculo{}).
		Where("id = ?", idVeiculo).
		Update("id_usuario", idUsuarioDestino).Error
	if err != nil {
		tx.Rollback()
		return nil, errors.New("erro ao atualizar proprietário do veículo")
	}

	// Cria o registro de transferência
	transferencia := models.TransferenciaVeiculo{
		IDVeiculo:          idVeiculo,
		IDUsuarioOrigem:    idUsuarioOrigem,
		IDUsuarioDestino:   idUsuarioDestino,
		IDLojaVenda:        idLojaVenda,
		IDHistoricoResgate: idHistoricoResgate,
		TipoTransferencia:  tipoTransferencia,
		Status:             models.StatusTransferenciaConfirmada,
		Observacoes:        observacoes,
	}

	err = tx.Create(&transferencia).Error
	if err != nil {
		tx.Rollback()
		return nil, errors.New("erro ao registrar transferência")
	}

	// Commit da transação
	tx.Commit()

	// Recarrega a transferência com os relacionamentos
	return GetTransferenciaByID(transferencia.ID)
}

// GetTransferenciaByID busca uma transferência por ID
func GetTransferenciaByID(id uint) (*models.TransferenciaVeiculo, error) {
	var transferencia models.TransferenciaVeiculo

	err := database.DB.
		Preload("Veiculo").
		Preload("UsuarioOrigem").
		Preload("UsuarioDestino").
		Preload("LojaVenda").
		Where("id = ? AND data_exclusao IS NULL", id).
		First(&transferencia).Error

	if err != nil {
		return nil, err
	}

	return &transferencia, nil
}

// GetTransferenciasByVeiculo retorna todas as transferências de um veículo
func GetTransferenciasByVeiculo(idVeiculo uint) ([]models.TransferenciaVeiculo, error) {
	var transferencias []models.TransferenciaVeiculo

	err := database.DB.
		Preload("Veiculo").
		Preload("UsuarioOrigem").
		Preload("UsuarioDestino").
		Preload("LojaVenda").
		Where("id_veiculo = ? AND data_exclusao IS NULL", idVeiculo).
		Order("data_transferencia DESC").
		Find(&transferencias).Error

	if err != nil {
		return nil, err
	}

	return transferencias, nil
}

// GetTransferenciasByUsuario retorna todas as transferências envolvendo um usuário (origem ou destino)
func GetTransferenciasByUsuario(idUsuario uint) ([]models.TransferenciaVeiculo, error) {
	var transferencias []models.TransferenciaVeiculo

	err := database.DB.
		Preload("Veiculo").
		Preload("UsuarioOrigem").
		Preload("UsuarioDestino").
		Preload("LojaVenda").
		Where("(id_usuario_origem = ? OR id_usuario_destino = ?) AND data_exclusao IS NULL", idUsuario, idUsuario).
		Order("data_transferencia DESC").
		Find(&transferencias).Error

	if err != nil {
		return nil, err
	}

	return transferencias, nil
}

// GetAllTransferencias retorna todas as transferências
func GetAllTransferencias() ([]models.TransferenciaVeiculo, error) {
	var transferencias []models.TransferenciaVeiculo

	err := database.DB.
		Preload("Veiculo").
		Preload("UsuarioOrigem").
		Preload("UsuarioDestino").
		Preload("LojaVenda").
		Where("data_exclusao IS NULL").
		Order("data_transferencia DESC").
		Find(&transferencias).Error

	if err != nil {
		return nil, err
	}

	return transferencias, nil
}

// SoftDeleteTransferencia realiza soft delete de uma transferência
func SoftDeleteTransferencia(id uint) error {
	_, err := GetTransferenciaByID(id)
	if err != nil {
		return errors.New("transferência não encontrada")
	}

	now := time.Now()
	err = database.DB.Model(&models.TransferenciaVeiculo{}).
		Where("id = ?", id).
		Update("data_exclusao", now).Error

	return err
}

// BuscarUsuariosParaTransferencia busca usuários ativos para selecionar como novo dono
// Permite buscar por nome, email ou CPF
func BuscarUsuariosParaTransferencia(termo string) ([]models.Usuario, error) {
	var usuarios []models.Usuario

	query := database.DB.
		Preload("Plano").
		Where("data_exclusao IS NULL AND ativo = ?", true)

	if termo != "" {
		termoBusca := "%" + termo + "%"
		query = query.Where(
			"nome ILIKE ? OR email ILIKE ? OR cpf ILIKE ?",
			termoBusca, termoBusca, termoBusca,
		)
	}

	err := query.
		Order("nome ASC").
		Limit(50).
		Find(&usuarios).Error

	if err != nil {
		return nil, err
	}

	return usuarios, nil
}
