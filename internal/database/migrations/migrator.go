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

	// Migration 009: Criar tabelas de veículos e histórico
	migration009 := m.NewMigration("009", "create_veiculos_tables").
		ExecuteSQL(`
			CREATE TABLE IF NOT EXISTS veiculos (
				id SERIAL PRIMARY KEY,
				modelo VARCHAR(255) NOT NULL,
				ano INTEGER NOT NULL,
				cor VARCHAR(100) NOT NULL,
				placa VARCHAR(10) UNIQUE NOT NULL,
				id_usuario INTEGER NOT NULL,
				data_cadastro TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
				ativo BOOLEAN DEFAULT TRUE,
				CONSTRAINT fk_veiculo_usuario FOREIGN KEY (id_usuario) REFERENCES usuarios(id)
			)
		`, `
			DROP TABLE IF EXISTS veiculos CASCADE
		`).
		ExecuteSQL(`
			CREATE TABLE IF NOT EXISTS historico_veiculos (
				id SERIAL PRIMARY KEY,
				id_veiculo INTEGER NOT NULL,
				id_anuncio INTEGER NOT NULL,
				descricao VARCHAR(500) NOT NULL,
				data TIMESTAMP NOT NULL,
				data_cadastro TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
				CONSTRAINT fk_historico_veiculo FOREIGN KEY (id_veiculo) REFERENCES veiculos(id),
				CONSTRAINT fk_historico_anuncio FOREIGN KEY (id_anuncio) REFERENCES anuncios(id)
			)
		`, `
			DROP TABLE IF EXISTS historico_veiculos CASCADE
		`).
		AddIndexSQL("veiculos", "idx_veiculo_usuario", "id_usuario").
		AddIndexSQL("veiculos", "idx_veiculo_placa", "placa").
		AddIndexSQL("historico_veiculos", "idx_historico_veiculo", "id_veiculo").
		AddIndexSQL("historico_veiculos", "idx_historico_anuncio", "id_anuncio").
		Build()
	m.migrations = append(m.migrations, migration009)

	// Migration 010: Criar tabelas complementares
	m.migrations = append(m.migrations, Migration{
		Version: "010",
		Name:    "create_complementary_tables",
		Up:      m.createComplementaryTables,
		Down:    m.dropComplementaryTables,
	})

	// Migration 011: Remover constraint única do campo stripe_payment_id
	migration011 := m.NewMigration("011", "remove_stripe_payment_id_unique_constraint").
		ExecuteSQL(`
			-- Remove a constraint única do campo stripe_payment_id se ela existir
			DO $$ 
			BEGIN
				-- Verifica se a constraint existe antes de tentar removê-la
				IF EXISTS (
					SELECT 1 FROM information_schema.table_constraints 
					WHERE constraint_name = 'uni_historico_pagamentos_stripe_payment_id'
					AND table_name = 'historico_pagamentos'
				) THEN
					ALTER TABLE historico_pagamentos DROP CONSTRAINT uni_historico_pagamentos_stripe_payment_id;
				END IF;
			END $$;
		`, `
			-- Rollback: Recria a constraint única (não recomendado em produção)
			-- ALTER TABLE historico_pagamentos ADD CONSTRAINT uni_historico_pagamentos_stripe_payment_id UNIQUE (stripe_payment_id);
			SELECT 'Rollback não implementado para esta migration' as message;
		`).
		Build()
	m.migrations = append(m.migrations, migration011)

	// Migration 012: Corrigir estrutura da tabela anuncios
	migration012 := m.NewMigration("012", "fix_anuncios_table_structure").
		AddColumnSQL("anuncios", "data_cadastro", "TIMESTAMP DEFAULT CURRENT_TIMESTAMP").
		AddColumnSQL("anuncios", "data_atualizacao", "TIMESTAMP DEFAULT CURRENT_TIMESTAMP").
		AddColumnSQL("anuncios", "data_exclusao", "TIMESTAMP").
		AddColumnSQL("anuncios", "id_produto", "INTEGER").
		AddColumnSQL("anuncios", "id_servico", "INTEGER").
		AddColumnSQL("anuncios", "id_veiculo", "INTEGER").
		AddColumnSQL("anuncios", "tipo_anuncio", "VARCHAR(20) NOT NULL DEFAULT ''").
		AddIndexSQL("anuncios", "idx_anuncio_data_exclusao", "data_exclusao").
		AddIndexSQL("anuncios", "idx_anuncio_tipo", "tipo_anuncio").
		Build()
	m.migrations = append(m.migrations, migration012)

	// Migration 013: Adicionar campos rating e isMeuCarroMais à tabela lojas
	migration013 := m.NewMigration("013", "add_rating_and_premium_to_lojas").
		AddColumnSQL("lojas", "rating", "INTEGER DEFAULT 5").
		AddColumnSQL("lojas", "is_meu_carro_mais", "BOOLEAN DEFAULT FALSE").
		AddIndexSQL("lojas", "idx_loja_rating", "rating").
		AddIndexSQL("lojas", "idx_loja_premium", "is_meu_carro_mais").
		Build()
	m.migrations = append(m.migrations, migration013)

	// Migration 014: Alterar campo saldo da carteira de decimal para int (moedas do app)
	migration014 := m.NewMigration("014", "change_carteira_saldo_to_int").
		ExecuteSQL(`
			-- Altera o campo saldo de DECIMAL para INTEGER
			-- Converte valores existentes para inteiros (arredonda para baixo)
			ALTER TABLE carteiras ALTER COLUMN saldo TYPE INTEGER USING FLOOR(saldo);
		`, `
			-- Rollback: Converte de volta para DECIMAL
			ALTER TABLE carteiras ALTER COLUMN saldo TYPE DECIMAL(10,2);
		`).
		Build()
	m.migrations = append(m.migrations, migration014)

	// Migration 015: Criar tabela de registro de interesse em veículos
	migration015 := m.NewMigration("015", "create_registro_interesse_table").
		ExecuteSQL(`
			CREATE TABLE IF NOT EXISTS registro_interesses (
				id SERIAL PRIMARY KEY,
				id_anuncio INTEGER NOT NULL,
				nome VARCHAR(255) NOT NULL,
				email VARCHAR(255) NOT NULL,
				telefone VARCHAR(20) NOT NULL,
				mensagem VARCHAR(1000),
				data_cadastro TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
				data_atualizacao TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
				data_exclusao TIMESTAMP,
				CONSTRAINT fk_registro_interesse_anuncio FOREIGN KEY (id_anuncio) REFERENCES anuncios(id)
			)
		`, `
			DROP TABLE IF EXISTS registro_interesses CASCADE
		`).
		AddIndexSQL("registro_interesses", "idx_registro_interesse_anuncio", "id_anuncio").
		AddIndexSQL("registro_interesses", "idx_registro_interesse_data_exclusao", "data_exclusao").
		AddIndexSQL("registro_interesses", "idx_registro_interesse_email", "email").
		Build()
	m.migrations = append(m.migrations, migration015)

	// Migration 016: Criar tabela de logs
	migration016 := m.NewMigration("016", "create_logs_table").
		ExecuteSQL(`
			CREATE TABLE IF NOT EXISTS logs (
				id SERIAL PRIMARY KEY,
				id_usuario INTEGER,
				tipo_acao VARCHAR(50) NOT NULL,
				entidade VARCHAR(50) NOT NULL,
				id_entidade INTEGER,
				descricao VARCHAR(500),
				dados_antigos JSONB,
				dados_novos JSONB,
				ip VARCHAR(45),
				user_agent VARCHAR(500),
				metodo_http VARCHAR(10),
				endpoint VARCHAR(255),
				status_http INTEGER DEFAULT 200,
				data_acao TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
				data_cadastro TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
				data_atualizacao TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
				data_exclusao TIMESTAMP,
				CONSTRAINT fk_log_usuario FOREIGN KEY (id_usuario) REFERENCES usuarios(id) ON DELETE SET NULL
			)
		`, `
			DROP TABLE IF EXISTS logs CASCADE
		`).
		AddIndexSQL("logs", "idx_log_id_usuario", "id_usuario").
		AddIndexSQL("logs", "idx_log_tipo_acao", "tipo_acao").
		AddIndexSQL("logs", "idx_log_entidade", "entidade").
		AddIndexSQL("logs", "idx_log_id_entidade", "id_entidade").
		AddIndexSQL("logs", "idx_log_data_acao", "data_acao").
		AddIndexSQL("logs", "idx_log_data_exclusao", "data_exclusao").
		AddIndexSQL("logs", "idx_log_usuario_entidade", "id_usuario, entidade").
		AddIndexSQL("logs", "idx_log_entidade_acao", "entidade, tipo_acao").
		Build()
	m.migrations = append(m.migrations, migration016)

	// Migration 017: Renomear tabela fotos para uploads e adicionar campos
	migration017 := m.NewMigration("017", "rename_fotos_to_uploads").
		ExecuteSQL(`
			-- Renomeia a tabela
			ALTER TABLE IF EXISTS fotos RENAME TO uploads;
			
			-- Adiciona campo id_usuario se não existir
			ALTER TABLE uploads ADD COLUMN IF NOT EXISTS id_usuario INTEGER;
			
			-- Adiciona campo tipo se não existir
			ALTER TABLE uploads ADD COLUMN IF NOT EXISTS tipo VARCHAR(20) NOT NULL DEFAULT 'Imagem';
			
			-- Adiciona foreign key para usuario
			DO $$
			BEGIN
				IF NOT EXISTS (
					SELECT 1 FROM pg_constraint WHERE conname = 'fk_upload_usuario'
				) THEN
					ALTER TABLE uploads ADD CONSTRAINT fk_upload_usuario 
					FOREIGN KEY (id_usuario) REFERENCES usuarios(id) ON DELETE SET NULL;
				END IF;
			END $$;
			
			-- Adiciona índices
			CREATE INDEX IF NOT EXISTS idx_upload_id_usuario ON uploads(id_usuario);
			CREATE INDEX IF NOT EXISTS idx_upload_tipo ON uploads(tipo);
			
			-- Atualiza tipo_entidade para incluir 'usuario' se necessário
			-- (não precisa fazer nada, apenas documentar que agora suporta)
		`, `
			-- Rollback: renomeia de volta para fotos e remove campos
			ALTER TABLE IF EXISTS uploads RENAME TO fotos;
			ALTER TABLE IF EXISTS fotos DROP COLUMN IF EXISTS id_usuario;
			ALTER TABLE IF EXISTS fotos DROP COLUMN IF EXISTS tipo;
			DROP INDEX IF EXISTS idx_upload_id_usuario;
			DROP INDEX IF EXISTS idx_upload_tipo;
			ALTER TABLE IF EXISTS fotos DROP CONSTRAINT IF EXISTS fk_upload_usuario;
		`).
		Build()
	m.migrations = append(m.migrations, migration017)

	// Migration 018: Simplificar categorias - usar string fixa em vez de tabela
	migration018 := m.NewMigration("018", "simplify_categories_to_string").
		ExecuteSQL(`
			-- Adiciona campo categoria às tabelas produto, servico e anuncio
			ALTER TABLE produtos ADD COLUMN IF NOT EXISTS categoria VARCHAR(100);
			ALTER TABLE servicos ADD COLUMN IF NOT EXISTS categoria VARCHAR(100);
			ALTER TABLE anuncios ADD COLUMN IF NOT EXISTS categoria VARCHAR(100);
			
			-- Migra dados existentes de servicos (se houver relação com categoria_servicos)
			UPDATE servicos s 
			SET categoria = (SELECT nome FROM categoria_servicos WHERE id = s.id_categoria)
			WHERE s.id_categoria IS NOT NULL AND EXISTS (SELECT 1 FROM categoria_servicos WHERE id = s.id_categoria);
			
			-- Migra dados existentes de anuncios (se houver relação com categoria_anuncios)
			UPDATE anuncios a 
			SET categoria = (SELECT nome FROM categoria_anuncios WHERE id = a.id_categoria)
			WHERE a.id_categoria IS NOT NULL AND EXISTS (SELECT 1 FROM categoria_anuncios WHERE id = a.id_categoria);
			
			-- Remove a coluna id_categoria da tabela servicos
			ALTER TABLE servicos DROP CONSTRAINT IF EXISTS fk_servicos_categoria;
			ALTER TABLE servicos DROP COLUMN IF EXISTS id_categoria;
			
			-- Remove a coluna id_categoria da tabela anuncios
			ALTER TABLE anuncios DROP CONSTRAINT IF EXISTS fk_anuncios_categoria;
			ALTER TABLE anuncios DROP COLUMN IF EXISTS id_categoria;
			
			-- Remove as tabelas de categorias
			DROP TABLE IF EXISTS categoria_servicos CASCADE;
			DROP TABLE IF EXISTS categoria_anuncios CASCADE;
		`, `
			-- Rollback: recria a estrutura anterior
			CREATE TABLE IF NOT EXISTS categoria_servicos (
				id SERIAL PRIMARY KEY,
				nome VARCHAR(255)
			);
			CREATE TABLE IF NOT EXISTS categoria_anuncios (
				id SERIAL PRIMARY KEY,
				nome VARCHAR(255)
			);
			ALTER TABLE servicos ADD COLUMN IF NOT EXISTS id_categoria INTEGER;
			ALTER TABLE anuncios ADD COLUMN IF NOT EXISTS id_categoria INTEGER;
			ALTER TABLE produtos DROP COLUMN IF EXISTS categoria;
			ALTER TABLE servicos DROP COLUMN IF EXISTS categoria;
			ALTER TABLE anuncios DROP COLUMN IF EXISTS categoria;
		`).
		Build()
	m.migrations = append(m.migrations, migration018)

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
		&models.Veiculo{},
		&models.HistoricoVeiculo{},
	)
}

// dropInitialTables remove todas as tabelas iniciais
func (m *Migrator) dropInitialTables(db *gorm.DB) error {
	tables := []interface{}{
		&models.HistoricoVeiculo{},
		&models.Veiculo{},
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

// createComplementaryTables cria as tabelas complementares (migration 010)
func (m *Migrator) createComplementaryTables(db *gorm.DB) error {
	return db.AutoMigrate(
		&models.Produto{},
		&models.Servico{},
		&models.VeiculoLoja{},
		&models.Upload{},
		&models.Avaliacao{},
		&models.HistoricoPagamento{},
		&models.HistoricoResgate{},
	)
}

// dropComplementaryTables remove as tabelas complementares
func (m *Migrator) dropComplementaryTables(db *gorm.DB) error {
	tables := []interface{}{
		&models.HistoricoResgate{},
		&models.HistoricoPagamento{},
		&models.Avaliacao{},
		&models.Upload{},
		&models.VeiculoLoja{},
		&models.Servico{},
		&models.Produto{},
	}

	for _, table := range tables {
		if err := db.Migrator().DropTable(table); err != nil {
			return fmt.Errorf("erro ao remover tabela: %v", err)
		}
	}

	return nil
}
