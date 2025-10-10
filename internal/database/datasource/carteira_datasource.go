package datasource

import (
	"errors"
	"meu-carro-mais/internal/database"
	"meu-carro-mais/internal/database/models"
	"meu-carro-mais/internal/handlers/json"
)

// GetCarteiraByID busca uma carteira por ID
func GetCarteiraByID(id uint) (*models.Carteira, error) {
	var carteira models.Carteira
	err := database.DB.
		Preload("Usuario").
		Where("id = ?", id).
		First(&carteira).Error
	if err != nil {
		return nil, err
	}
	return &carteira, nil
}

// GetCarteiraByUsuarioID busca a carteira de um usuário específico
func GetCarteiraByUsuarioID(usuarioID uint) (*models.Carteira, error) {
	var carteira models.Carteira
	err := database.DB.
		Preload("Usuario").
		Where("usuario_id = ?", usuarioID).
		First(&carteira).Error
	if err != nil {
		return nil, err
	}
	return &carteira, nil
}

// GetAllCarteiras retorna todas as carteiras
func GetAllCarteiras() ([]models.Carteira, error) {
	var carteiras []models.Carteira
	err := database.DB.
		Preload("Usuario").
		Order("data_criacao DESC").
		Find(&carteiras).Error
	if err != nil {
		return nil, err
	}
	return carteiras, nil
}

// CreateCarteira cria uma nova carteira
func CreateCarteira(req json.CarteiraRequest) (*models.Carteira, error) {
	// Verifica se o usuário existe
	var usuario models.Usuario
	err := database.DB.Where("id = ?", req.UsuarioID).First(&usuario).Error
	if err != nil {
		return nil, errors.New("usuário não encontrado")
	}

	// Verifica se o usuário já possui uma carteira
	var existingCarteira models.Carteira
	err = database.DB.Where("usuario_id = ?", req.UsuarioID).First(&existingCarteira).Error
	if err == nil {
		return nil, errors.New("usuário já possui uma carteira")
	}

	carteira := models.Carteira{
		UsuarioID: req.UsuarioID,
		Saldo:     req.Saldo,
	}

	err = database.DB.Create(&carteira).Error
	if err != nil {
		return nil, err
	}

	// Recarrega a carteira com os relacionamentos
	return GetCarteiraByID(carteira.ID)
}

// UpdateCarteira atualiza uma carteira existente
func UpdateCarteira(id uint, req json.CarteiraRequest) (*models.Carteira, error) {
	// Verifica se a carteira existe
	carteira, err := GetCarteiraByID(id)
	if err != nil {
		return nil, errors.New("carteira não encontrada")
	}

	// Verifica se o usuário existe (se foi alterado)
	if carteira.UsuarioID != req.UsuarioID {
		var usuario models.Usuario
		err = database.DB.Where("id = ?", req.UsuarioID).First(&usuario).Error
		if err != nil {
			return nil, errors.New("usuário não encontrado")
		}
	}

	// Atualiza os campos
	carteira.UsuarioID = req.UsuarioID
	carteira.Saldo = req.Saldo

	err = database.DB.Save(&carteira).Error
	if err != nil {
		return nil, err
	}

	// Recarrega a carteira com os relacionamentos
	return GetCarteiraByID(id)
}

// UpdateCarteiraSaldo atualiza apenas o saldo de uma carteira
func UpdateCarteiraSaldo(id uint, novoSaldo int) (*models.Carteira, error) {
	// Verifica se a carteira existe
	_, err := GetCarteiraByID(id)
	if err != nil {
		return nil, errors.New("carteira não encontrada")
	}

	// Atualiza apenas o saldo
	err = database.DB.Model(&models.Carteira{}).
		Where("id = ?", id).
		Update("saldo", novoSaldo).Error
	if err != nil {
		return nil, err
	}

	// Recarrega a carteira com os relacionamentos
	return GetCarteiraByID(id)
}

// AdicionarSaldo adiciona valor ao saldo da carteira
func AdicionarSaldo(id uint, valor int) (*models.Carteira, error) {
	// Verifica se a carteira existe
	carteira, err := GetCarteiraByID(id)
	if err != nil {
		return nil, errors.New("carteira não encontrada")
	}

	// Adiciona o valor ao saldo atual
	novoSaldo := carteira.Saldo + valor
	return UpdateCarteiraSaldo(id, novoSaldo)
}

// SubtrairSaldo subtrai valor do saldo da carteira
func SubtrairSaldo(id uint, valor int) (*models.Carteira, error) {
	// Verifica se a carteira existe
	carteira, err := GetCarteiraByID(id)
	if err != nil {
		return nil, errors.New("carteira não encontrada")
	}

	// Verifica se há saldo suficiente
	if carteira.Saldo < valor {
		return nil, errors.New("saldo insuficiente")
	}

	// Subtrai o valor do saldo atual
	novoSaldo := carteira.Saldo - valor
	return UpdateCarteiraSaldo(id, novoSaldo)
}

// DeleteCarteira remove uma carteira
func DeleteCarteira(id uint) error {
	// Verifica se a carteira existe
	_, err := GetCarteiraByID(id)
	if err != nil {
		return errors.New("carteira não encontrada")
	}

	// Remove a carteira
	err = database.DB.Delete(&models.Carteira{}, id).Error
	if err != nil {
		return err
	}

	return nil
}

// GetCarteirasBySaldoRange busca carteiras com saldo dentro de um range
func GetCarteirasBySaldoRange(saldoMin, saldoMax int) ([]models.Carteira, error) {
	var carteiras []models.Carteira
	err := database.DB.
		Preload("Usuario").
		Where("saldo >= ? AND saldo <= ?", saldoMin, saldoMax).
		Order("saldo DESC").
		Find(&carteiras).Error
	if err != nil {
		return nil, err
	}
	return carteiras, nil
}

// GetCarteirasComSaldoMaior busca carteiras com saldo maior que um valor
func GetCarteirasComSaldoMaior(valor int) ([]models.Carteira, error) {
	var carteiras []models.Carteira
	err := database.DB.
		Preload("Usuario").
		Where("saldo > ?", valor).
		Order("saldo DESC").
		Find(&carteiras).Error
	if err != nil {
		return nil, err
	}
	return carteiras, nil
}
