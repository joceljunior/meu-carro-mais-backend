package datasource

import (
	"errors"

	"meu-carro-mais/internal/database"
	"meu-carro-mais/internal/database/models"

	"gorm.io/gorm"
)

// AdicionarMoedasLojaUsuarioTx incrementa moedas por loja (uso em transação).
func AdicionarMoedasLojaUsuarioTx(tx *gorm.DB, usuarioID, lojaID uint, delta int) error {
	if delta <= 0 {
		return nil
	}
	var row models.UsuarioMoedasLoja
	err := tx.Where("usuario_id = ? AND loja_id = ?", usuarioID, lojaID).First(&row).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		row = models.UsuarioMoedasLoja{
			UsuarioID: usuarioID,
			LojaID:    lojaID,
			Saldo:     delta,
		}
		return tx.Create(&row).Error
	}
	if err != nil {
		return err
	}
	row.Saldo += delta
	return tx.Save(&row).Error
}

// GetUsuarioMoedasLojaByUsuarioID lista saldos por loja do usuário.
func GetUsuarioMoedasLojaByUsuarioID(usuarioID uint) ([]models.UsuarioMoedasLoja, error) {
	var rows []models.UsuarioMoedasLoja
	err := database.DB.
		Preload("Loja").
		Where("usuario_id = ?", usuarioID).
		Order("loja_id ASC").
		Find(&rows).Error
	if err != nil {
		return nil, err
	}
	return rows, nil
}
