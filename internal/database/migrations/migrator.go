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

	// Migration 019: Criar tabela de descontos para lojas
	migration019 := m.NewMigration("019", "create_descontos_table").
		ExecuteSQL(`
			CREATE TABLE IF NOT EXISTS descontos (
				id SERIAL PRIMARY KEY,
				id_loja INTEGER NOT NULL,
				porcentagem DECIMAL(5,2) NOT NULL,
				ativo BOOLEAN DEFAULT TRUE,
				data_validade TIMESTAMP NOT NULL,
				data_cadastro TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
				data_atualizacao TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
				data_exclusao TIMESTAMP,
				CONSTRAINT fk_desconto_loja FOREIGN KEY (id_loja) REFERENCES lojas(id) ON DELETE CASCADE
			)
		`, `
			DROP TABLE IF EXISTS descontos CASCADE
		`).
		AddIndexSQL("descontos", "idx_desconto_id_loja", "id_loja").
		AddIndexSQL("descontos", "idx_desconto_ativo", "ativo").
		AddIndexSQL("descontos", "idx_desconto_data_exclusao", "data_exclusao").
		AddIndexSQL("descontos", "idx_desconto_data_validade", "data_validade").
		AddIndexSQL("descontos", "idx_desconto_loja_ativo", "id_loja, ativo").
		Build()
	m.migrations = append(m.migrations, migration019)

	// Migration 020: Criar tabela de ofertas Auto Mais
	migration020 := m.NewMigration("020", "create_ofertas_auto_mais_table").
		ExecuteSQL(`
			CREATE TABLE IF NOT EXISTS oferta_auto_mais (
				id SERIAL PRIMARY KEY,
				id_loja INTEGER NOT NULL,
				nome VARCHAR(255) NOT NULL,
				descricao VARCHAR(500),
				moedas INTEGER NOT NULL,
				porcentagem DECIMAL(5,2) NOT NULL,
				ativo BOOLEAN DEFAULT TRUE,
				data_validade TIMESTAMP,
				data_cadastro TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
				data_atualizacao TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
				data_exclusao TIMESTAMP,
				CONSTRAINT fk_oferta_auto_mais_loja FOREIGN KEY (id_loja) REFERENCES lojas(id) ON DELETE CASCADE
			)
		`, `
			DROP TABLE IF EXISTS oferta_auto_mais CASCADE
		`).
		AddIndexSQL("oferta_auto_mais", "idx_oferta_auto_mais_id_loja", "id_loja").
		AddIndexSQL("oferta_auto_mais", "idx_oferta_auto_mais_ativo", "ativo").
		AddIndexSQL("oferta_auto_mais", "idx_oferta_auto_mais_data_exclusao", "data_exclusao").
		AddIndexSQL("oferta_auto_mais", "idx_oferta_auto_mais_data_validade", "data_validade").
		AddIndexSQL("oferta_auto_mais", "idx_oferta_auto_mais_loja_ativo", "id_loja, ativo").
		Build()
	m.migrations = append(m.migrations, migration020)

	// Migration 021: Adicionar campo id_oferta_auto_mais à tabela anuncios
	migration021 := m.NewMigration("021", "add_oferta_auto_mais_to_anuncios").
		AddColumnSQL("anuncios", "id_oferta_auto_mais", "INTEGER").
		ExecuteSQL(`
			ALTER TABLE anuncios ADD CONSTRAINT fk_anuncio_oferta_auto_mais 
			FOREIGN KEY (id_oferta_auto_mais) REFERENCES oferta_auto_mais(id) ON DELETE SET NULL
		`, `
			ALTER TABLE anuncios DROP CONSTRAINT IF EXISTS fk_anuncio_oferta_auto_mais
		`).
		AddIndexSQL("anuncios", "idx_anuncio_oferta_auto_mais", "id_oferta_auto_mais").
		Build()
	m.migrations = append(m.migrations, migration021)

	// Migration 022: Adicionar campos de tipo de usuário e atualizar todos para mobile
	migration022 := m.NewMigration("022", "add_user_type_fields_and_update_to_mobile").
		// Adiciona novos campos
		AddColumnSQL("usuarios", "tipo", "VARCHAR(20) DEFAULT 'mobile'").
		AddColumnSQL("usuarios", "status", "VARCHAR(20) DEFAULT 'aprovado'").
		AddColumnSQL("usuarios", "id_executivo", "INTEGER").
		AddColumnSQL("usuarios", "solicitacao_executivo", "VARCHAR(20) DEFAULT ''").
		AddColumnSQL("usuarios", "data_solicitacao_executivo", "TIMESTAMP").
		AddColumnSQL("usuarios", "motivo_solicitacao_executivo", "VARCHAR(500)").
		// Adiciona índices
		AddIndexSQL("usuarios", "idx_usuario_tipo", "tipo").
		AddIndexSQL("usuarios", "idx_usuario_status", "status").
		AddIndexSQL("usuarios", "idx_usuario_id_executivo", "id_executivo").
		AddIndexSQL("usuarios", "idx_usuario_solicitacao_executivo", "solicitacao_executivo").
		// Atualiza TODOS os usuários existentes para tipo mobile e status aprovado
		ExecuteSQL(`
			UPDATE usuarios 
			SET tipo = 'mobile', 
				status = 'aprovado',
				solicitacao_executivo = ''
			WHERE tipo IS NULL OR tipo = ''
		`, `
			-- Rollback: não faz nada, pois não podemos saber o tipo original
			SELECT 'Rollback não implementado para atualização de tipo' as message
		`).
		Build()
	m.migrations = append(m.migrations, migration022)

	// Migration 023: Adicionar campos categoria e id_usuario à tabela lojas
	migration023 := m.NewMigration("023", "add_categoria_and_id_usuario_to_lojas").
		AddColumnSQL("lojas", "categoria", "VARCHAR(255)").
		AddColumnSQL("lojas", "id_usuario", "INTEGER").
		AddIndexSQL("lojas", "idx_loja_categoria", "categoria").
		AddIndexSQL("lojas", "idx_loja_id_usuario", "id_usuario").
		ExecuteSQL(`
			-- Adiciona foreign key para id_usuario se não existir
			DO $$
			BEGIN
				IF NOT EXISTS (
					SELECT 1 FROM pg_constraint WHERE conname = 'fk_loja_usuario'
				) THEN
					ALTER TABLE lojas ADD CONSTRAINT fk_loja_usuario 
					FOREIGN KEY (id_usuario) REFERENCES usuarios(id) ON DELETE SET NULL;
				END IF;
			END $$;
		`, `
			-- Rollback: remove foreign key
			ALTER TABLE lojas DROP CONSTRAINT IF EXISTS fk_loja_usuario;
		`).
		Build()
	m.migrations = append(m.migrations, migration023)

	// Migration 024: Adicionar campos de desconto à tabela anuncios
	migration024 := m.NewMigration("024", "add_desconto_fields_to_anuncios").
		AddColumnSQL("anuncios", "porcentagem_desconto", "DECIMAL(5,2) DEFAULT 0").
		AddColumnSQL("anuncios", "preco_com_desconto", "DECIMAL(10,2)").
		Build()
	m.migrations = append(m.migrations, migration024)

	// Migration 025: Adicionar campos completos à tabela veiculos
	migration025 := m.NewMigration("025", "add_complete_fields_to_veiculos").
		// Renomear campo 'ano' para 'ano_fabricacao' e adicionar 'ano_modelo'
		ExecuteSQL(`
			-- Renomeia 'ano' para 'ano_fabricacao' se ainda não foi renomeado
			DO $$
			BEGIN
				IF EXISTS (SELECT 1 FROM information_schema.columns 
					WHERE table_name = 'veiculos' AND column_name = 'ano' 
					AND NOT EXISTS (SELECT 1 FROM information_schema.columns 
						WHERE table_name = 'veiculos' AND column_name = 'ano_fabricacao')) THEN
					ALTER TABLE veiculos RENAME COLUMN ano TO ano_fabricacao;
				END IF;
			END $$;
		`, `
			-- Rollback: renomeia de volta
			DO $$
			BEGIN
				IF EXISTS (SELECT 1 FROM information_schema.columns 
					WHERE table_name = 'veiculos' AND column_name = 'ano_fabricacao') THEN
					ALTER TABLE veiculos RENAME COLUMN ano_fabricacao TO ano;
				END IF;
			END $$;
		`).
		// Adiciona novos campos
		AddColumnSQL("veiculos", "marca", "VARCHAR(100) NOT NULL DEFAULT ''").
		AddColumnSQL("veiculos", "ano_modelo", "INTEGER NOT NULL DEFAULT 0").
		AddColumnSQL("veiculos", "renavam", "VARCHAR(20)").
		AddColumnSQL("veiculos", "chassi", "VARCHAR(50)").
		AddColumnSQL("veiculos", "tipo_veiculo", "VARCHAR(50)").
		AddColumnSQL("veiculos", "combustivel", "VARCHAR(50)").
		AddColumnSQL("veiculos", "preco", "DECIMAL(10,2)").
		AddColumnSQL("veiculos", "licenciamento", "VARCHAR(50)").
		AddColumnSQL("veiculos", "ipva_pago", "BOOLEAN DEFAULT FALSE").
		AddColumnSQL("veiculos", "possui_financiamento", "BOOLEAN DEFAULT FALSE").
		AddColumnSQL("veiculos", "possui_multas", "BOOLEAN DEFAULT FALSE").
		// Atualiza ano_modelo para ser igual ao ano_fabricacao se não foi definido
		ExecuteSQL(`
			UPDATE veiculos 
			SET ano_modelo = ano_fabricacao 
			WHERE ano_modelo = 0 OR ano_modelo IS NULL
		`, `
			-- Rollback: não faz nada
			SELECT 'Rollback não implementado para atualização de ano_modelo' as message
		`).
		Build()
	m.migrations = append(m.migrations, migration025)

	// Migration 026: Adicionar campos faltantes à tabela veiculos
	migration026 := m.NewMigration("026", "add_missing_fields_to_veiculos").
		AddColumnSQL("veiculos", "quilometragem", "INTEGER").
		AddColumnSQL("veiculos", "observacoes", "TEXT").
		AddColumnSQL("veiculos", "data_atualizacao", "TIMESTAMP DEFAULT CURRENT_TIMESTAMP").
		AddColumnSQL("veiculos", "data_exclusao", "TIMESTAMP").
		AddIndexSQL("veiculos", "idx_veiculo_data_exclusao", "data_exclusao").
		Build()
	m.migrations = append(m.migrations, migration026)

	// Migration 027: Adicionar campos de quantidade e desconto ao histórico de resgates
	migration027 := m.NewMigration("027", "add_quantidade_desconto_to_historico_resgates").
		AddColumnSQL("historico_resgates", "quantidade", "INTEGER DEFAULT 1").
		AddColumnSQL("historico_resgates", "valor_unitario", "DECIMAL(10,2) DEFAULT 0").
		AddColumnSQL("historico_resgates", "valor_original", "DECIMAL(10,2) DEFAULT 0").
		AddColumnSQL("historico_resgates", "desconto_aplicado", "DECIMAL(10,2) DEFAULT 0").
		Build()
	m.migrations = append(m.migrations, migration027)

	// Migration 028: Adicionar campos de indicação para usuarios e lojas
	migration028 := m.NewMigration("028", "add_indicacao_fields_to_usuarios_and_lojas").
		// Campos de indicação no Usuario
		AddColumnSQL("usuarios", "id_loja_indicadora", "INTEGER").
		AddColumnSQL("usuarios", "data_vinculo_loja", "TIMESTAMP").
		AddColumnSQL("usuarios", "id_usuario_indicador", "INTEGER").
		AddColumnSQL("usuarios", "data_vinculo_usuario", "TIMESTAMP").
		// Campos de indicação na Loja
		AddColumnSQL("lojas", "id_usuario_indicador", "INTEGER").
		AddColumnSQL("lojas", "data_vinculo_usuario", "TIMESTAMP").
		AddColumnSQL("lojas", "endereco", "VARCHAR(500)").
		// Índices
		AddIndexSQL("usuarios", "idx_usuario_loja_indicadora", "id_loja_indicadora").
		AddIndexSQL("usuarios", "idx_usuario_usuario_indicador", "id_usuario_indicador").
		AddIndexSQL("lojas", "idx_loja_usuario_indicador", "id_usuario_indicador").
		// Foreign keys
		ExecuteSQL(`
			-- FK: Usuario -> Loja indicadora
			DO $$
			BEGIN
				IF NOT EXISTS (
					SELECT 1 FROM pg_constraint WHERE conname = 'fk_usuario_loja_indicadora'
				) THEN
					ALTER TABLE usuarios ADD CONSTRAINT fk_usuario_loja_indicadora 
					FOREIGN KEY (id_loja_indicadora) REFERENCES lojas(id) ON DELETE SET NULL;
				END IF;
			END $$;
			
			-- FK: Usuario -> Usuario indicador
			DO $$
			BEGIN
				IF NOT EXISTS (
					SELECT 1 FROM pg_constraint WHERE conname = 'fk_usuario_usuario_indicador'
				) THEN
					ALTER TABLE usuarios ADD CONSTRAINT fk_usuario_usuario_indicador 
					FOREIGN KEY (id_usuario_indicador) REFERENCES usuarios(id) ON DELETE SET NULL;
				END IF;
			END $$;
			
			-- FK: Loja -> Usuario indicador
			DO $$
			BEGIN
				IF NOT EXISTS (
					SELECT 1 FROM pg_constraint WHERE conname = 'fk_loja_usuario_indicador'
				) THEN
					ALTER TABLE lojas ADD CONSTRAINT fk_loja_usuario_indicador 
					FOREIGN KEY (id_usuario_indicador) REFERENCES usuarios(id) ON DELETE SET NULL;
				END IF;
			END $$;
		`, `
			-- Rollback: remove foreign keys
			ALTER TABLE usuarios DROP CONSTRAINT IF EXISTS fk_usuario_loja_indicadora;
			ALTER TABLE usuarios DROP CONSTRAINT IF EXISTS fk_usuario_usuario_indicador;
			ALTER TABLE lojas DROP CONSTRAINT IF EXISTS fk_loja_usuario_indicador;
		`).
		Build()
	m.migrations = append(m.migrations, migration028)

	// Migration 029: Adicionar campo porcentagem_desconto ao histórico de resgates
	migration029 := m.NewMigration("029", "add_porcentagem_desconto_to_historico_resgates").
		AddColumnSQL("historico_resgates", "porcentagem_desconto", "DECIMAL(5,2) DEFAULT 0").
		Build()
	m.migrations = append(m.migrations, migration029)

	// Migration 030: Adicionar campos opcionais às avaliações (servico, produto, anuncio)
	// NOTA: GORM pluraliza Avaliacao como avaliacaos (não avaliacoes)
	migration030 := m.NewMigration("030", "add_optional_fields_to_avaliacoes").
		// Torna id_loja opcional (permite NULL)
		ExecuteSQL(`
			ALTER TABLE avaliacaos ALTER COLUMN id_loja DROP NOT NULL;
		`, `
			-- Rollback: restaura NOT NULL (pode falhar se houver dados NULL)
			ALTER TABLE avaliacaos ALTER COLUMN id_loja SET NOT NULL;
		`).
		AddColumnSQL("avaliacaos", "id_servico", "INTEGER").
		AddColumnSQL("avaliacaos", "id_produto", "INTEGER").
		AddColumnSQL("avaliacaos", "id_anuncio", "INTEGER").
		// Adiciona foreign keys
		ExecuteSQL(`
			DO $$
			BEGIN
				IF NOT EXISTS (
					SELECT 1 FROM pg_constraint WHERE conname = 'fk_avaliacao_servico'
				) THEN
					ALTER TABLE avaliacaos ADD CONSTRAINT fk_avaliacao_servico 
					FOREIGN KEY (id_servico) REFERENCES servicos(id) ON DELETE SET NULL;
				END IF;
				IF NOT EXISTS (
					SELECT 1 FROM pg_constraint WHERE conname = 'fk_avaliacao_produto'
				) THEN
					ALTER TABLE avaliacaos ADD CONSTRAINT fk_avaliacao_produto 
					FOREIGN KEY (id_produto) REFERENCES produtos(id) ON DELETE SET NULL;
				END IF;
				IF NOT EXISTS (
					SELECT 1 FROM pg_constraint WHERE conname = 'fk_avaliacao_anuncio'
				) THEN
					ALTER TABLE avaliacaos ADD CONSTRAINT fk_avaliacao_anuncio 
					FOREIGN KEY (id_anuncio) REFERENCES anuncios(id) ON DELETE SET NULL;
				END IF;
			END $$;
		`, `
			-- Rollback: remove foreign keys
			ALTER TABLE avaliacaos DROP CONSTRAINT IF EXISTS fk_avaliacao_servico;
			ALTER TABLE avaliacaos DROP CONSTRAINT IF EXISTS fk_avaliacao_produto;
			ALTER TABLE avaliacaos DROP CONSTRAINT IF EXISTS fk_avaliacao_anuncio;
		`).
		Build()
	m.migrations = append(m.migrations, migration030)

	// Migration 031: Criar tabela de transferências de veículos
	migration031 := m.NewMigration("031", "create_transferencia_veiculos_table").
		ExecuteSQL(`
			CREATE TABLE IF NOT EXISTS transferencia_veiculos (
				id SERIAL PRIMARY KEY,
				id_veiculo INTEGER NOT NULL,
				id_usuario_origem INTEGER NOT NULL,
				id_usuario_destino INTEGER NOT NULL,
				id_loja_venda INTEGER,
				id_historico_resgate INTEGER,
				tipo_transferencia VARCHAR(20) NOT NULL DEFAULT 'manual',
				status VARCHAR(20) NOT NULL DEFAULT 'confirmada',
				observacoes TEXT,
				data_transferencia TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
				data_atualizacao TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
				data_exclusao TIMESTAMP,
				CONSTRAINT fk_transferencia_veiculo FOREIGN KEY (id_veiculo) REFERENCES veiculos(id),
				CONSTRAINT fk_transferencia_usuario_origem FOREIGN KEY (id_usuario_origem) REFERENCES usuarios(id),
				CONSTRAINT fk_transferencia_usuario_destino FOREIGN KEY (id_usuario_destino) REFERENCES usuarios(id),
				CONSTRAINT fk_transferencia_loja_venda FOREIGN KEY (id_loja_venda) REFERENCES lojas(id) ON DELETE SET NULL,
				CONSTRAINT fk_transferencia_historico_resgate FOREIGN KEY (id_historico_resgate) REFERENCES historico_resgates(id) ON DELETE SET NULL
			)
		`, `
			DROP TABLE IF EXISTS transferencia_veiculos CASCADE
		`).
		AddIndexSQL("transferencia_veiculos", "idx_transferencia_veiculo", "id_veiculo").
		AddIndexSQL("transferencia_veiculos", "idx_transferencia_usuario_origem", "id_usuario_origem").
		AddIndexSQL("transferencia_veiculos", "idx_transferencia_usuario_destino", "id_usuario_destino").
		AddIndexSQL("transferencia_veiculos", "idx_transferencia_loja_venda", "id_loja_venda").
		AddIndexSQL("transferencia_veiculos", "idx_transferencia_data_exclusao", "data_exclusao").
		AddIndexSQL("transferencia_veiculos", "idx_transferencia_tipo", "tipo_transferencia").
		AddIndexSQL("transferencia_veiculos", "idx_transferencia_status", "status").
		Build()
	m.migrations = append(m.migrations, migration031)

	// Migration 032: Adicionar campos id_anuncio e id_veiculo_usuario ao histórico de resgates
	migration032 := m.NewMigration("032", "add_id_anuncio_and_id_veiculo_usuario_to_historico_resgates").
		AddColumnSQL("historico_resgates", "id_anuncio", "INTEGER").
		AddColumnSQL("historico_resgates", "id_veiculo_usuario", "INTEGER").
		ExecuteSQL(`
			DO $$
			BEGIN
				IF NOT EXISTS (
					SELECT 1 FROM pg_constraint WHERE conname = 'fk_historico_resgate_anuncio'
				) THEN
					ALTER TABLE historico_resgates ADD CONSTRAINT fk_historico_resgate_anuncio 
					FOREIGN KEY (id_anuncio) REFERENCES anuncios(id) ON DELETE SET NULL;
				END IF;
				IF NOT EXISTS (
					SELECT 1 FROM pg_constraint WHERE conname = 'fk_historico_resgate_veiculo_usuario'
				) THEN
					ALTER TABLE historico_resgates ADD CONSTRAINT fk_historico_resgate_veiculo_usuario 
					FOREIGN KEY (id_veiculo_usuario) REFERENCES veiculos(id) ON DELETE SET NULL;
				END IF;
			END $$;
		`, `
			-- Rollback: remove foreign keys
			ALTER TABLE historico_resgates DROP CONSTRAINT IF EXISTS fk_historico_resgate_anuncio;
			ALTER TABLE historico_resgates DROP CONSTRAINT IF EXISTS fk_historico_resgate_veiculo_usuario;
		`).
		AddIndexSQL("historico_resgates", "idx_historico_resgate_anuncio", "id_anuncio").
		AddIndexSQL("historico_resgates", "idx_historico_resgate_veiculo_usuario", "id_veiculo_usuario").
		Build()
	m.migrations = append(m.migrations, migration032)

	// Migration 033: Renomear anuncios para cupons e atualizar referências
	migration033 := m.NewMigration("033", "rename_anuncios_to_cupons").
		ExecuteSQL(`
			-- 1. Renomear tabela anuncios para cupons
			ALTER TABLE IF EXISTS anuncios RENAME TO cupons;

			-- 2. Renomear coluna tipo_anuncio para tipo_cupom na tabela cupons
			DO $$
			BEGIN
				IF EXISTS (SELECT 1 FROM information_schema.columns 
					WHERE table_name = 'cupons' AND column_name = 'tipo_anuncio') THEN
					ALTER TABLE cupons RENAME COLUMN tipo_anuncio TO tipo_cupom;
				END IF;
			END $$;

			-- 3. Renomear colunas id_anuncio para id_cupom nas tabelas referenciadas
			DO $$
			BEGIN
				-- historico_veiculos
				IF EXISTS (SELECT 1 FROM information_schema.columns 
					WHERE table_name = 'historico_veiculos' AND column_name = 'id_anuncio') THEN
					ALTER TABLE historico_veiculos RENAME COLUMN id_anuncio TO id_cupom;
				END IF;
				-- registro_interesses
				IF EXISTS (SELECT 1 FROM information_schema.columns 
					WHERE table_name = 'registro_interesses' AND column_name = 'id_anuncio') THEN
					ALTER TABLE registro_interesses RENAME COLUMN id_anuncio TO id_cupom;
				END IF;
				-- avaliacaos
				IF EXISTS (SELECT 1 FROM information_schema.columns 
					WHERE table_name = 'avaliacaos' AND column_name = 'id_anuncio') THEN
					ALTER TABLE avaliacaos RENAME COLUMN id_anuncio TO id_cupom;
				END IF;
				-- historico_resgates
				IF EXISTS (SELECT 1 FROM information_schema.columns 
					WHERE table_name = 'historico_resgates' AND column_name = 'id_anuncio') THEN
					ALTER TABLE historico_resgates RENAME COLUMN id_anuncio TO id_cupom;
				END IF;
			END $$;

			-- 4. Atualizar FK constraints (drop old, add new)
			ALTER TABLE historico_veiculos DROP CONSTRAINT IF EXISTS fk_historico_anuncio;
			ALTER TABLE registro_interesses DROP CONSTRAINT IF EXISTS fk_registro_interesse_anuncio;
			ALTER TABLE avaliacaos DROP CONSTRAINT IF EXISTS fk_avaliacao_anuncio;
			ALTER TABLE historico_resgates DROP CONSTRAINT IF EXISTS fk_historico_resgate_anuncio;
			ALTER TABLE cupons DROP CONSTRAINT IF EXISTS fk_anuncio_oferta_auto_mais;

			DO $$
			BEGIN
				IF NOT EXISTS (SELECT 1 FROM pg_constraint WHERE conname = 'fk_historico_cupom') THEN
					ALTER TABLE historico_veiculos ADD CONSTRAINT fk_historico_cupom 
					FOREIGN KEY (id_cupom) REFERENCES cupons(id);
				END IF;
				IF NOT EXISTS (SELECT 1 FROM pg_constraint WHERE conname = 'fk_registro_interesse_cupom') THEN
					ALTER TABLE registro_interesses ADD CONSTRAINT fk_registro_interesse_cupom 
					FOREIGN KEY (id_cupom) REFERENCES cupons(id);
				END IF;
				IF NOT EXISTS (SELECT 1 FROM pg_constraint WHERE conname = 'fk_avaliacao_cupom') THEN
					ALTER TABLE avaliacaos ADD CONSTRAINT fk_avaliacao_cupom 
					FOREIGN KEY (id_cupom) REFERENCES cupons(id) ON DELETE SET NULL;
				END IF;
				IF NOT EXISTS (SELECT 1 FROM pg_constraint WHERE conname = 'fk_historico_resgate_cupom') THEN
					ALTER TABLE historico_resgates ADD CONSTRAINT fk_historico_resgate_cupom 
					FOREIGN KEY (id_cupom) REFERENCES cupons(id) ON DELETE SET NULL;
				END IF;
				IF NOT EXISTS (SELECT 1 FROM pg_constraint WHERE conname = 'fk_cupom_oferta_auto_mais') THEN
					ALTER TABLE cupons ADD CONSTRAINT fk_cupom_oferta_auto_mais 
					FOREIGN KEY (id_oferta_auto_mais) REFERENCES oferta_auto_mais(id) ON DELETE SET NULL;
				END IF;
			END $$;

			-- 5. Atualizar índices (drop old, create new)
			DROP INDEX IF EXISTS idx_anuncio_data_exclusao;
			DROP INDEX IF EXISTS idx_anuncio_tipo;
			DROP INDEX IF EXISTS idx_anuncio_oferta_auto_mais;
			DROP INDEX IF EXISTS idx_historico_anuncio;
			DROP INDEX IF EXISTS idx_registro_interesse_anuncio;
			DROP INDEX IF EXISTS idx_historico_resgate_anuncio;

			CREATE INDEX IF NOT EXISTS idx_cupom_data_exclusao ON cupons(data_exclusao);
			CREATE INDEX IF NOT EXISTS idx_cupom_tipo ON cupons(tipo_cupom);
			CREATE INDEX IF NOT EXISTS idx_cupom_oferta_auto_mais ON cupons(id_oferta_auto_mais);
			CREATE INDEX IF NOT EXISTS idx_historico_cupom ON historico_veiculos(id_cupom);
			CREATE INDEX IF NOT EXISTS idx_registro_interesse_cupom ON registro_interesses(id_cupom);
			CREATE INDEX IF NOT EXISTS idx_historico_resgate_cupom ON historico_resgates(id_cupom);
		`, `
			-- Rollback: renomear cupons de volta para anuncios
			ALTER TABLE IF EXISTS cupons RENAME TO anuncios;
			
			DO $$
			BEGIN
				IF EXISTS (SELECT 1 FROM information_schema.columns 
					WHERE table_name = 'anuncios' AND column_name = 'tipo_cupom') THEN
					ALTER TABLE anuncios RENAME COLUMN tipo_cupom TO tipo_anuncio;
				END IF;
			END $$;
			
			DO $$
			BEGIN
				IF EXISTS (SELECT 1 FROM information_schema.columns 
					WHERE table_name = 'historico_veiculos' AND column_name = 'id_cupom') THEN
					ALTER TABLE historico_veiculos RENAME COLUMN id_cupom TO id_anuncio;
				END IF;
				IF EXISTS (SELECT 1 FROM information_schema.columns 
					WHERE table_name = 'registro_interesses' AND column_name = 'id_cupom') THEN
					ALTER TABLE registro_interesses RENAME COLUMN id_cupom TO id_anuncio;
				END IF;
				IF EXISTS (SELECT 1 FROM information_schema.columns 
					WHERE table_name = 'avaliacaos' AND column_name = 'id_cupom') THEN
					ALTER TABLE avaliacaos RENAME COLUMN id_cupom TO id_anuncio;
				END IF;
				IF EXISTS (SELECT 1 FROM information_schema.columns 
					WHERE table_name = 'historico_resgates' AND column_name = 'id_cupom') THEN
					ALTER TABLE historico_resgates RENAME COLUMN id_cupom TO id_anuncio;
				END IF;
			END $$;
		`).
		Build()
	m.migrations = append(m.migrations, migration033)

	// Migration 034: Simplificar tabela historico_resgates e atualizar status
	migration034 := m.NewMigration("034", "simplify_historico_resgates_and_update_status").
		ExecuteSQL(`
			-- 1. Atualizar status existentes: aprovado/confirmado -> efetivado
			UPDATE historico_resgates SET status = 'efetivado' WHERE status IN ('aprovado', 'confirmado');

			-- 2. Remover colunas que não são mais utilizadas
			ALTER TABLE historico_resgates DROP COLUMN IF EXISTS quantidade;
			ALTER TABLE historico_resgates DROP COLUMN IF EXISTS valor_unitario;
			ALTER TABLE historico_resgates DROP COLUMN IF EXISTS valor_original;
			ALTER TABLE historico_resgates DROP COLUMN IF EXISTS desconto_aplicado;
			ALTER TABLE historico_resgates DROP COLUMN IF EXISTS porcentagem_desconto;
			ALTER TABLE historico_resgates DROP COLUMN IF EXISTS id_veiculo_usuario;

			-- 3. Remover foreign keys e índices das colunas removidas
			ALTER TABLE historico_resgates DROP CONSTRAINT IF EXISTS fk_historico_resgate_veiculo_usuario;
			DROP INDEX IF EXISTS idx_historico_resgate_veiculo_usuario;
		`, `
			-- Rollback: recria as colunas removidas
			ALTER TABLE historico_resgates ADD COLUMN IF NOT EXISTS quantidade INTEGER DEFAULT 1;
			ALTER TABLE historico_resgates ADD COLUMN IF NOT EXISTS valor_unitario DECIMAL(10,2) DEFAULT 0;
			ALTER TABLE historico_resgates ADD COLUMN IF NOT EXISTS valor_original DECIMAL(10,2) DEFAULT 0;
			ALTER TABLE historico_resgates ADD COLUMN IF NOT EXISTS desconto_aplicado DECIMAL(10,2) DEFAULT 0;
			ALTER TABLE historico_resgates ADD COLUMN IF NOT EXISTS porcentagem_desconto DECIMAL(5,2) DEFAULT 0;
			ALTER TABLE historico_resgates ADD COLUMN IF NOT EXISTS id_veiculo_usuario INTEGER;

			-- Rollback: reverte status efetivado para aprovado
			UPDATE historico_resgates SET status = 'aprovado' WHERE status = 'efetivado';
		`).
		Build()
	m.migrations = append(m.migrations, migration034)

	// Migration 035: Remover colunas remanescentes da tabela historico_resgates
	migration035 := m.NewMigration("035", "remove_remaining_columns_from_historico_resgates").
		ExecuteSQL(`
			-- Remover foreign keys das colunas que serão removidas
			ALTER TABLE historico_resgates DROP CONSTRAINT IF EXISTS fk_historico_resgate_produto;
			ALTER TABLE historico_resgates DROP CONSTRAINT IF EXISTS fk_historico_resgate_servico;
			ALTER TABLE historico_resgates DROP CONSTRAINT IF EXISTS fk_historico_resgate_veiculo;
			ALTER TABLE historico_resgates DROP CONSTRAINT IF EXISTS fk_historico_resgate_loja;
			ALTER TABLE historico_resgates DROP CONSTRAINT IF EXISTS fk_historico_resgates_produto;
			ALTER TABLE historico_resgates DROP CONSTRAINT IF EXISTS fk_historico_resgates_servico;
			ALTER TABLE historico_resgates DROP CONSTRAINT IF EXISTS fk_historico_resgates_veiculo;
			ALTER TABLE historico_resgates DROP CONSTRAINT IF EXISTS fk_historico_resgates_loja;

			-- Remover colunas que não fazem mais parte do modelo
			ALTER TABLE historico_resgates DROP COLUMN IF EXISTS id_produto;
			ALTER TABLE historico_resgates DROP COLUMN IF EXISTS id_servico;
			ALTER TABLE historico_resgates DROP COLUMN IF EXISTS id_veiculo;
			ALTER TABLE historico_resgates DROP COLUMN IF EXISTS id_loja;
			ALTER TABLE historico_resgates DROP COLUMN IF EXISTS tipo_resgate;
			ALTER TABLE historico_resgates DROP COLUMN IF EXISTS valor;
			ALTER TABLE historico_resgates DROP COLUMN IF EXISTS data_exclusao;
		`, `
			-- Rollback: recria as colunas removidas
			ALTER TABLE historico_resgates ADD COLUMN IF NOT EXISTS id_produto INTEGER;
			ALTER TABLE historico_resgates ADD COLUMN IF NOT EXISTS id_servico INTEGER;
			ALTER TABLE historico_resgates ADD COLUMN IF NOT EXISTS id_veiculo INTEGER;
			ALTER TABLE historico_resgates ADD COLUMN IF NOT EXISTS id_loja INTEGER;
			ALTER TABLE historico_resgates ADD COLUMN IF NOT EXISTS tipo_resgate VARCHAR(20);
			ALTER TABLE historico_resgates ADD COLUMN IF NOT EXISTS valor NUMERIC(10,2);
			ALTER TABLE historico_resgates ADD COLUMN IF NOT EXISTS data_exclusao TIMESTAMPTZ;
		`).
		Build()
	m.migrations = append(m.migrations, migration035)

	migration036 := m.NewMigration("036", "add_id_usuario_to_cupons").
		ExecuteSQL(`
			-- Adicionar coluna id_usuario na tabela cupons (opcional, pode ser null)
			ALTER TABLE cupons ADD COLUMN IF NOT EXISTS id_usuario INTEGER;

			-- Criar índice para a coluna id_usuario
			CREATE INDEX IF NOT EXISTS idx_cupons_id_usuario ON cupons(id_usuario);

			-- Adicionar foreign key para a tabela usuarios
			ALTER TABLE cupons ADD CONSTRAINT fk_cupons_usuario FOREIGN KEY (id_usuario) REFERENCES usuarios(id) ON DELETE SET NULL;
		`, `
			-- Rollback: remover a coluna id_usuario
			ALTER TABLE cupons DROP CONSTRAINT IF EXISTS fk_cupons_usuario;
			DROP INDEX IF EXISTS idx_cupons_id_usuario;
			ALTER TABLE cupons DROP COLUMN IF EXISTS id_usuario;
		`).
		Build()
	m.migrations = append(m.migrations, migration036)

	migration037 := m.NewMigration("037", "add_desconto_geral_porcentagem_to_lojas").
		ExecuteSQL(`
			ALTER TABLE lojas ADD COLUMN IF NOT EXISTS desconto_geral_porcentagem DECIMAL(5,2) NOT NULL DEFAULT 0;
		`, `
			ALTER TABLE lojas DROP COLUMN IF EXISTS desconto_geral_porcentagem;
		`).
		Build()
	m.migrations = append(m.migrations, migration037)

	migration038 := m.NewMigration("038", "add_moedas_utilizadas_to_historico_resgates").
		ExecuteSQL(`
			ALTER TABLE historico_resgates ADD COLUMN IF NOT EXISTS moedas_utilizadas INTEGER NOT NULL DEFAULT 0;
		`, `
			ALTER TABLE historico_resgates DROP COLUMN IF EXISTS moedas_utilizadas;
		`).
		Build()
	m.migrations = append(m.migrations, migration038)

	migration039 := m.NewMigration("039", "create_vendas_produto_avulso_table").
		ExecuteSQL(`
			CREATE TABLE IF NOT EXISTS vendas_produto_avulso (
				id SERIAL PRIMARY KEY,
				id_usuario INTEGER NOT NULL,
				id_loja INTEGER NOT NULL,
				valor DECIMAL(10,2) NOT NULL,
				descricao_produto VARCHAR(500) NOT NULL,
				data_venda TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
				CONSTRAINT fk_vendas_produto_avulso_usuario FOREIGN KEY (id_usuario) REFERENCES usuarios(id) ON DELETE CASCADE,
				CONSTRAINT fk_vendas_produto_avulso_loja FOREIGN KEY (id_loja) REFERENCES lojas(id) ON DELETE CASCADE
			);
			CREATE INDEX IF NOT EXISTS idx_vendas_produto_avulso_usuario ON vendas_produto_avulso(id_usuario);
			CREATE INDEX IF NOT EXISTS idx_vendas_produto_avulso_loja ON vendas_produto_avulso(id_loja);
			CREATE INDEX IF NOT EXISTS idx_vendas_produto_avulso_data ON vendas_produto_avulso(data_venda);
		`, `
			DROP TABLE IF EXISTS vendas_produto_avulso CASCADE;
		`).
		Build()
	m.migrations = append(m.migrations, migration039)

	migration040 := m.NewMigration("040", "moedas_gerais_por_loja_e_credito_resgate").
		ExecuteSQL(`
			ALTER TABLE carteiras RENAME COLUMN saldo TO saldo_geral;

			ALTER TABLE historico_resgates ADD COLUMN IF NOT EXISTS moedas_loja_ja_creditadas BOOLEAN NOT NULL DEFAULT FALSE;

			CREATE TABLE IF NOT EXISTS usuario_moedas_loja (
				id SERIAL PRIMARY KEY,
				usuario_id INTEGER NOT NULL,
				loja_id INTEGER NOT NULL,
				saldo INTEGER NOT NULL DEFAULT 0,
				CONSTRAINT uq_usuario_moedas_loja UNIQUE (usuario_id, loja_id),
				CONSTRAINT fk_umj_usuario FOREIGN KEY (usuario_id) REFERENCES usuarios(id) ON DELETE CASCADE,
				CONSTRAINT fk_umj_loja FOREIGN KEY (loja_id) REFERENCES lojas(id) ON DELETE CASCADE
			);
			CREATE INDEX IF NOT EXISTS idx_usuario_moedas_loja_usuario ON usuario_moedas_loja(usuario_id);
			CREATE INDEX IF NOT EXISTS idx_usuario_moedas_loja_loja ON usuario_moedas_loja(loja_id);
		`, `
			DROP TABLE IF EXISTS usuario_moedas_loja CASCADE;
			ALTER TABLE historico_resgates DROP COLUMN IF EXISTS moedas_loja_ja_creditadas;
			ALTER TABLE carteiras RENAME COLUMN saldo_geral TO saldo;
		`).
		Build()
	m.migrations = append(m.migrations, migration040)

	migration041 := m.NewMigration("041", "add_loja_redes_sociais_e_horario").
		ExecuteSQL(`
			ALTER TABLE lojas ADD COLUMN IF NOT EXISTS link_instagram VARCHAR(500) NOT NULL DEFAULT '';
			ALTER TABLE lojas ADD COLUMN IF NOT EXISTS link_facebook VARCHAR(500) NOT NULL DEFAULT '';
			ALTER TABLE lojas ADD COLUMN IF NOT EXISTS horario_funcionamento TEXT NOT NULL DEFAULT '';
		`, `
			ALTER TABLE lojas DROP COLUMN IF EXISTS horario_funcionamento;
			ALTER TABLE lojas DROP COLUMN IF EXISTS link_facebook;
			ALTER TABLE lojas DROP COLUMN IF EXISTS link_instagram;
		`).
		Build()
	m.migrations = append(m.migrations, migration041)

	migration042 := m.NewMigration("042", "add_link_site_to_lojas").
		ExecuteSQL(`
			ALTER TABLE lojas ADD COLUMN IF NOT EXISTS link_site VARCHAR(500) NOT NULL DEFAULT '';
		`, `
			ALTER TABLE lojas DROP COLUMN IF EXISTS link_site;
		`).
		Build()
	m.migrations = append(m.migrations, migration042)

	migration043 := m.NewMigration("043", "add_telefone_to_lojas").
		ExecuteSQL(`
			ALTER TABLE lojas ADD COLUMN IF NOT EXISTS telefone VARCHAR(30) NOT NULL DEFAULT '';
		`, `
			ALTER TABLE lojas DROP COLUMN IF EXISTS telefone;
		`).
		Build()
	m.migrations = append(m.migrations, migration043)

	migration044 := m.NewMigration("044", "loja_rating_default_zero").
		ExecuteSQL(`
			ALTER TABLE lojas ALTER COLUMN rating SET DEFAULT 0;
		`, `
			ALTER TABLE lojas ALTER COLUMN rating SET DEFAULT 5;
		`).
		Build()
	m.migrations = append(m.migrations, migration044)

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
		&models.Cupom{},
		&models.Veiculo{},
		&models.HistoricoVeiculo{},
	)
}

// dropInitialTables remove todas as tabelas iniciais
func (m *Migrator) dropInitialTables(db *gorm.DB) error {
	tables := []interface{}{
		&models.HistoricoVeiculo{},
		&models.Veiculo{},
		&models.Cupom{},
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
		&models.VendaProdutoAvulso{},
	)
}

// dropComplementaryTables remove as tabelas complementares
func (m *Migrator) dropComplementaryTables(db *gorm.DB) error {
	tables := []interface{}{
		&models.VendaProdutoAvulso{},
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
