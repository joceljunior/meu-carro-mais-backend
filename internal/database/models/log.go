package models

import (
	"database/sql/driver"
	"encoding/json"
	"time"
)

// JSONB é um tipo customizado para armazenar JSON no PostgreSQL
type JSONB map[string]interface{}

// Value implementa o driver.Valuer para JSONB
func (j JSONB) Value() (driver.Value, error) {
	if j == nil {
		return nil, nil
	}
	return json.Marshal(j)
}

// Scan implementa o sql.Scanner para JSONB
func (j *JSONB) Scan(value interface{}) error {
	if value == nil {
		*j = nil
		return nil
	}
	bytes, ok := value.([]byte)
	if !ok {
		return nil
	}
	return json.Unmarshal(bytes, j)
}

type Log struct {
	ID              uint       `gorm:"primaryKey"`
	IDUsuario       *uint      `gorm:"index"` // Pode ser null se for ação do sistema
	TipoAcao        string     `gorm:"size:50;not null;index"` // "criar", "atualizar", "deletar", "resgatar", "aprovar", "rejeitar", "restaurar", "visualizar"
	Entidade        string     `gorm:"size:50;not null;index"` // "cupom", "produto", "servico", "veiculo", "historico_resgate", "registro_interesse", "usuario", "loja", etc.
	IDEntidade      *uint      `gorm:"index"`                  // ID da entidade afetada
	Descricao       string     `gorm:"size:500"`               // Descrição da ação
	DadosAntigos    JSONB      `gorm:"type:jsonb"`             // Dados antes da alteração (para updates)
	DadosNovos      JSONB      `gorm:"type:jsonb"`             // Dados após a alteração
	IP              string     `gorm:"size:45"`                 // IP do usuário (suporta IPv6)
	UserAgent       string     `gorm:"size:500"`                // User-Agent do navegador/dispositivo
	MetodoHTTP      string     `gorm:"size:10"`                 // GET, POST, PUT, DELETE, etc.
	Endpoint        string     `gorm:"size:255"`                // Endpoint chamado
	StatusHTTP      int        `gorm:"default:200"`             // Status HTTP da resposta
	DataAcao        time.Time  `gorm:"autoCreateTime;index"`     // Timestamp da ação
	DataCadastro    time.Time  `gorm:"autoCreateTime"`
	DataAtualizacao time.Time  `gorm:"autoUpdateTime"`
	DataExclusao    *time.Time `gorm:"index"`                   // Soft delete

	// Relacionamentos
	Usuario *Usuario `gorm:"foreignKey:IDUsuario"`
}

