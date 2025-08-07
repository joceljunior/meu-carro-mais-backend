package migrations

import (
	"fmt"
	"log"
	"sort"

	"meu-carro-mais/internal/database/models"

	"gorm.io/gorm"
)

// Migration representa uma migration individual
type Migration struct {
	Version string
	Name    string
	Up      func(*gorm.DB) error
	Down    func(*gorm.DB) error
}

// Migrator gerencia as migrations
type Migrator struct {
	db         *gorm.DB
	migrations []Migration
}

// NewMigrator cria uma nova instância do migrator
func NewMigrator(db *gorm.DB) *Migrator {
	migrator := &Migrator{
		db:         db,
		migrations: []Migration{},
	}

	// Registra todas as migrations
	migrator.registerMigrations()

	return migrator
}

// Helper functions para migrations mais declarativas

// AddColumn adiciona uma coluna a uma tabela
func (m *Migrator) AddColumn(table interface{}, columnName string, columnType string) func(*gorm.DB) error {
	return func(db *gorm.DB) error {
		if !db.Migrator().HasColumn(table, columnName) {
			return db.Migrator().AddColumn(table, columnName)
		}
		return nil
	}
}

// DropColumn remove uma coluna de uma tabela
func (m *Migrator) DropColumn(table interface{}, columnName string) func(*gorm.DB) error {
	return func(db *gorm.DB) error {
		if db.Migrator().HasColumn(table, columnName) {
			return db.Migrator().DropColumn(table, columnName)
		}
		return nil
	}
}

// AddIndex adiciona um índice a uma tabela
func (m *Migrator) AddIndex(tableName, indexName, columns string) func(*gorm.DB) error {
	return func(db *gorm.DB) error {
		return db.Exec(fmt.Sprintf("CREATE INDEX IF NOT EXISTS %s ON %s (%s)",
			indexName, tableName, columns)).Error
	}
}

// DropIndex remove um índice de uma tabela
func (m *Migrator) DropIndex(tableName, indexName string) func(*gorm.DB) error {
	return func(db *gorm.DB) error {
		return db.Exec(fmt.Sprintf("DROP INDEX IF EXISTS %s", indexName)).Error
	}
}

// ExecuteSQL executa uma query SQL diretamente
func (m *Migrator) ExecuteSQL(sql string) func(*gorm.DB) error {
	return func(db *gorm.DB) error {
		return db.Exec(sql).Error
	}
}

// AddForeignKey adiciona uma foreign key
func (m *Migrator) AddForeignKey(tableName, constraintName, columnName, referencedTable, referencedColumn string) func(*gorm.DB) error {
	return func(db *gorm.DB) error {
		return db.Exec(fmt.Sprintf(
			"ALTER TABLE %s ADD CONSTRAINT %s FOREIGN KEY (%s) REFERENCES %s(%s)",
			tableName, constraintName, columnName, referencedTable, referencedColumn,
		)).Error
	}
}

// DropForeignKey remove uma foreign key
func (m *Migrator) DropForeignKey(tableName, constraintName string) func(*gorm.DB) error {
	return func(db *gorm.DB) error {
		return db.Exec(fmt.Sprintf("ALTER TABLE %s DROP CONSTRAINT %s", tableName, constraintName)).Error
	}
}

// MigrationBuilder facilita a criação de migrations declarativas
type MigrationBuilder struct {
	version string
	name    string
	upSQL   []string
	downSQL []string
}

// NewMigration cria uma nova migration builder
func (m *Migrator) NewMigration(version, name string) *MigrationBuilder {
	return &MigrationBuilder{
		version: version,
		name:    name,
		upSQL:   []string{},
		downSQL: []string{},
	}
}

// AddColumnSQL adiciona uma coluna via SQL
func (mb *MigrationBuilder) AddColumnSQL(tableName, columnName, columnType string) *MigrationBuilder {
	mb.upSQL = append(mb.upSQL, fmt.Sprintf("ALTER TABLE %s ADD COLUMN IF NOT EXISTS %s %s", tableName, columnName, columnType))
	mb.downSQL = append([]string{fmt.Sprintf("ALTER TABLE %s DROP COLUMN IF EXISTS %s", tableName, columnName)}, mb.downSQL...)
	return mb
}

// DropColumnSQL remove uma coluna via SQL
func (mb *MigrationBuilder) DropColumnSQL(tableName, columnName string) *MigrationBuilder {
	mb.upSQL = append(mb.upSQL, fmt.Sprintf("ALTER TABLE %s DROP COLUMN IF EXISTS %s", tableName, columnName))
	mb.downSQL = append([]string{fmt.Sprintf("ALTER TABLE %s ADD COLUMN %s", tableName, columnName)}, mb.downSQL...)
	return mb
}

// AddIndexSQL adiciona um índice via SQL
func (mb *MigrationBuilder) AddIndexSQL(tableName, indexName, columns string) *MigrationBuilder {
	mb.upSQL = append(mb.upSQL, fmt.Sprintf("CREATE INDEX IF NOT EXISTS %s ON %s (%s)", indexName, tableName, columns))
	mb.downSQL = append([]string{fmt.Sprintf("DROP INDEX IF EXISTS %s", indexName)}, mb.downSQL...)
	return mb
}

// DropIndexSQL remove um índice via SQL
func (mb *MigrationBuilder) DropIndexSQL(indexName string) *MigrationBuilder {
	mb.upSQL = append(mb.upSQL, fmt.Sprintf("DROP INDEX IF EXISTS %s", indexName))
	// Note: Para rollback de drop index, você precisaria recriar o índice
	return mb
}

// ExecuteSQL adiciona uma query SQL customizada
func (mb *MigrationBuilder) ExecuteSQL(upSQL, downSQL string) *MigrationBuilder {
	mb.upSQL = append(mb.upSQL, upSQL)
	mb.downSQL = append([]string{downSQL}, mb.downSQL...)
	return mb
}

// Build constrói a migration final
func (mb *MigrationBuilder) Build() Migration {
	return Migration{
		Version: mb.version,
		Name:    mb.name,
		Up: func(db *gorm.DB) error {
			for _, sql := range mb.upSQL {
				if err := db.Exec(sql).Error; err != nil {
					return fmt.Errorf("erro ao executar SQL '%s': %v", sql, err)
				}
			}
			return nil
		},
		Down: func(db *gorm.DB) error {
			for _, sql := range mb.downSQL {
				if err := db.Exec(sql).Error; err != nil {
					return fmt.Errorf("erro ao executar rollback SQL '%s': %v", sql, err)
				}
			}
			return nil
		},
	}
}

// registerMigrations registra todas as migrations disponíveis
func (m *Migrator) registerMigrations() {
	// Migration 001: Criar tabela de migrations
	m.migrations = append(m.migrations, Migration{
		Version: "001",
		Name:    "create_migrations_table",
		Up:      m.createMigrationsTable,
		Down:    m.dropMigrationsTable,
	})

	// Migration 002: Criar tabelas iniciais
	m.migrations = append(m.migrations, Migration{
		Version: "002",
		Name:    "create_initial_tables",
		Up:      m.createInitialTables,
		Down:    m.dropInitialTables,
	})

	// Migration 003: Adicionar campo telefone ao usuário (usando abordagem declarativa)
	m.migrations = append(m.migrations, Migration{
		Version: "003",
		Name:    "add_telefone_to_usuario",
		Up:      m.AddColumn(&models.Usuario{}, "telefone", "VARCHAR(20)"),
		Down:    m.DropColumn(&models.Usuario{}, "telefone"),
	})

	// Migration 004: Adicionar campo endereco ao usuário (usando abordagem declarativa)
	m.migrations = append(m.migrations, Migration{
		Version: "004",
		Name:    "add_endereco_to_usuario",
		Up:      m.AddColumn(&models.Usuario{}, "endereco", "VARCHAR(500)"),
		Down:    m.DropColumn(&models.Usuario{}, "endereco"),
	})

	// Migration 005: Exemplo usando MigrationBuilder (mais declarativo)
	migration005 := m.NewMigration("005", "add_data_nascimento_to_usuario").
		AddColumnSQL("usuarios", "data_nascimento", "DATE").
		AddIndexSQL("usuarios", "idx_usuario_data_nascimento", "data_nascimento").
		Build()
	m.migrations = append(m.migrations, migration005)

	// Migration 006: Exemplo usando SQL direto
	migration006 := m.NewMigration("006", "add_cpf_index_to_usuario").
		AddIndexSQL("usuarios", "idx_usuario_cpf", "cpf").
		Build()
	m.migrations = append(m.migrations, migration006)

	// Migration 007: Exemplo de múltiplas operações
	migration007 := m.NewMigration("007", "add_usuario_fields").
		AddColumnSQL("usuarios", "data_cadastro", "TIMESTAMP DEFAULT CURRENT_TIMESTAMP").
		AddColumnSQL("usuarios", "ativo", "BOOLEAN DEFAULT TRUE").
		AddIndexSQL("usuarios", "idx_usuario_ativo", "ativo").
		Build()
	m.migrations = append(m.migrations, migration007)

	// Migration 008: Adicionar campos latitude e longitude ao usuário
	migration008 := m.NewMigration("008", "add_latitude_longitude_to_usuario").
		AddColumnSQL("usuarios", "latitude", "DECIMAL(10,8)").
		AddColumnSQL("usuarios", "longitude", "DECIMAL(11,8)").
		AddIndexSQL("usuarios", "idx_usuario_latitude", "latitude").
		AddIndexSQL("usuarios", "idx_usuario_longitude", "longitude").
		Build()
	m.migrations = append(m.migrations, migration008)

	// Ordena as migrations por versão
	sort.Slice(m.migrations, func(i, j int) bool {
		return m.migrations[i].Version < m.migrations[j].Version
	})
}

// Run executa todas as migrations pendentes
func (m *Migrator) Run() error {
	log.Println("Iniciando execução das migrations...")

	// Verifica se a tabela de migrations existe
	if !m.db.Migrator().HasTable(&models.Migration{}) {
		log.Println("Criando tabela de migrations...")
		if err := m.db.AutoMigrate(&models.Migration{}); err != nil {
			return fmt.Errorf("erro ao criar tabela de migrations: %v", err)
		}
	}

	// Obtém migrations já executadas
	var executedMigrations []models.Migration
	if err := m.db.Find(&executedMigrations).Error; err != nil {
		return fmt.Errorf("erro ao buscar migrations executadas: %v", err)
	}

	executedVersions := make(map[string]bool)
	for _, migration := range executedMigrations {
		executedVersions[migration.Version] = true
	}

	// Executa migrations pendentes
	for _, migration := range m.migrations {
		if !executedVersions[migration.Version] {
			log.Printf("Executando migration %s: %s", migration.Version, migration.Name)

			if err := migration.Up(m.db); err != nil {
				return fmt.Errorf("erro ao executar migration %s (%s): %v", migration.Version, migration.Name, err)
			}

			// Registra a migration como executada
			executedMigration := models.Migration{
				Version: migration.Version,
				Name:    migration.Name,
			}

			if err := m.db.Create(&executedMigration).Error; err != nil {
				return fmt.Errorf("erro ao registrar migration %s: %v", migration.Version, err)
			}

			log.Printf("Migration %s executada com sucesso", migration.Version)
		} else {
			log.Printf("Migration %s já foi executada, pulando...", migration.Version)
		}
	}

	log.Println("Todas as migrations foram executadas com sucesso!")
	return nil
}

// Rollback executa o rollback da última migration
func (m *Migrator) Rollback() error {
	var lastMigration models.Migration
	if err := m.db.Order("version DESC").First(&lastMigration).Error; err != nil {
		return fmt.Errorf("nenhuma migration encontrada para rollback: %v", err)
	}

	// Encontra a migration correspondente
	var targetMigration *Migration
	for _, migration := range m.migrations {
		if migration.Version == lastMigration.Version {
			targetMigration = &migration
			break
		}
	}

	if targetMigration == nil {
		return fmt.Errorf("migration %s não encontrada", lastMigration.Version)
	}

	log.Printf("Executando rollback da migration %s: %s", targetMigration.Version, targetMigration.Name)

	if targetMigration.Down != nil {
		if err := targetMigration.Down(m.db); err != nil {
			return fmt.Errorf("erro ao executar rollback da migration %s: %v", targetMigration.Version, err)
		}
	}

	// Remove o registro da migration
	if err := m.db.Delete(&lastMigration).Error; err != nil {
		return fmt.Errorf("erro ao remover registro da migration %s: %v", targetMigration.Version, err)
	}

	log.Printf("Rollback da migration %s executado com sucesso", targetMigration.Version)
	return nil
}

// Status mostra o status das migrations
func (m *Migrator) Status() error {
	var executedMigrations []models.Migration
	if err := m.db.Order("version").Find(&executedMigrations).Error; err != nil {
		return fmt.Errorf("erro ao buscar migrations: %v", err)
	}

	executedVersions := make(map[string]bool)
	for _, migration := range executedMigrations {
		executedVersions[migration.Version] = true
	}

	log.Println("Status das migrations:")
	log.Println("=====================")

	for _, migration := range m.migrations {
		status := "PENDENTE"
		if executedVersions[migration.Version] {
			status = "EXECUTADA"
		}
		log.Printf("[%s] %s - %s", status, migration.Version, migration.Name)
	}

	return nil
}

// createMigrationsTable cria a tabela de migrations
func (m *Migrator) createMigrationsTable(db *gorm.DB) error {
	return db.AutoMigrate(&models.Migration{})
}

// dropMigrationsTable remove a tabela de migrations
func (m *Migrator) dropMigrationsTable(db *gorm.DB) error {
	return db.Migrator().DropTable(&models.Migration{})
}

// createInitialTables cria todas as tabelas iniciais
func (m *Migrator) createInitialTables(db *gorm.DB) error {
	return db.AutoMigrate(
		&models.TipoPlano{},
		&models.CategoriaLojista{},
		&models.Usuario{},
		&models.Loja{},
		&models.HistoricoPlanoUsuario{},
		&models.Carteira{},
		&models.LogCarteira{},
		&models.Anuncio{},
	)
}

// dropInitialTables remove todas as tabelas iniciais
func (m *Migrator) dropInitialTables(db *gorm.DB) error {
	tables := []interface{}{
		&models.Anuncio{},
		&models.LogCarteira{},
		&models.Carteira{},
		&models.HistoricoPlanoUsuario{},
		&models.Loja{},
		&models.Usuario{},
		&models.CategoriaLojista{},
		&models.TipoPlano{},
	}

	for _, table := range tables {
		if err := db.Migrator().DropTable(table); err != nil {
			return fmt.Errorf("erro ao remover tabela: %v", err)
		}
	}

	return nil
}
