package models

import "time"

// TipoUsuario representa os tipos de usuário do sistema
type TipoUsuario string

const (
	TipoUsuarioMobile        TipoUsuario = "mobile"
	TipoUsuarioAdministrativo TipoUsuario = "administrativo"
	TipoUsuarioCustomer      TipoUsuario = "customer"
	TipoUsuarioExecutivo     TipoUsuario = "executivo"
)

// StatusUsuario representa o status de aprovação do usuário (usado para customers)
type StatusUsuario string

const (
	StatusUsuarioPendente  StatusUsuario = "pendente"
	StatusUsuarioAprovado  StatusUsuario = "aprovado"
	StatusUsuarioRejeitado StatusUsuario = "rejeitado"
)

// StatusSolicitacaoExecutivo representa o status da solicitação para virar executivo
type StatusSolicitacaoExecutivo string

const (
	StatusSolicitacaoNenhuma   StatusSolicitacaoExecutivo = ""
	StatusSolicitacaoPendente  StatusSolicitacaoExecutivo = "pendente"
	StatusSolicitacaoAprovada  StatusSolicitacaoExecutivo = "aprovada"
	StatusSolicitacaoRejeitada StatusSolicitacaoExecutivo = "rejeitada"
)

type Usuario struct {
	ID              uint          `gorm:"primaryKey"`
	Nome            string        `gorm:"size:255"`
	Email           string        `gorm:"size:255;unique"`
	Senha           string        `gorm:"size:255"`
	CPF             string        `gorm:"size:255;unique"`
	Imagem          string        `gorm:"size:255"`
	Telefone        string        `gorm:"size:20"`
	Endereco        string        `gorm:"size:500"`
	DataNascimento  *time.Time    `gorm:"type:date"`
	DataCadastro    time.Time     `gorm:"autoCreateTime"`
	DataAtualizacao time.Time     `gorm:"autoUpdateTime"`
	DataExclusao    *time.Time    `gorm:"index"`
	Ativo           bool          `gorm:"default:true"`
	Latitude        *float64      `gorm:"type:decimal(10,8)"`
	Longitude       *float64      `gorm:"type:decimal(11,8)"`
	IDPlano         uint
	IDLoja          *uint
	Tipo            TipoUsuario   `gorm:"size:20;default:'mobile'"`          // Tipo do usuário: mobile, administrativo, customer, executivo
	Status          StatusUsuario `gorm:"size:20;default:'aprovado'"`        // Status de aprovação (usado para customers)
	IDExecutivo     *uint         `gorm:"index"`                              // ID do executivo que criou o customer (para bonificação)
	
	// Campos para solicitação de virar executivo (usado por usuários mobile)
	SolicitacaoExecutivo       StatusSolicitacaoExecutivo `gorm:"size:20;default:''"` // Status da solicitação: "", pendente, aprovada, rejeitada
	DataSolicitacaoExecutivo   *time.Time                 `gorm:"type:datetime"`      // Data da solicitação
	MotivoSolicitacaoExecutivo string                     `gorm:"size:500"`           // Motivo/justificativa da solicitação
	
	// Campos para vínculo de indicação (opcional)
	// Um usuário pode ter sido indicado por uma loja
	IDLojaIndicadora  *uint      `gorm:"index"`        // ID da loja que indicou este usuário (opcional)
	DataVinculoLoja   *time.Time `gorm:"type:datetime"` // Data do vínculo com a loja indicadora
	
	Plano           TipoPlano     `gorm:"foreignKey:IDPlano"`
	Loja            Loja          `gorm:"foreignKey:IDLoja"`
	Veiculos        []Veiculo     `gorm:"foreignKey:IDUsuario"`
	Executivo       *Usuario      `gorm:"foreignKey:IDExecutivo"`             // Referência ao executivo que criou este customer
	LojaIndicadora  *Loja         `gorm:"foreignKey:IDLojaIndicadora"`        // Referência à loja que indicou este usuário
}
