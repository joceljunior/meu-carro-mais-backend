# Sistema de Migrations

Este documento descreve como usar o sistema de migrations manual implementado na API.

## Visão Geral

O sistema de migrations foi implementado para substituir o auto-migrate do GORM, oferecendo:

- ✅ Controle total sobre as mudanças no banco
- ✅ Versionamento de schema
- ✅ Rollback de migrations
- ✅ Rastreamento de migrations executadas
- ✅ Segurança para ambiente de produção
- ✅ **NOVO**: Abordagem declarativa e simplificada

## Estrutura

```
internal/database/
├── migrations/
│   └── migrator.go          # Sistema principal de migrations
├── models/
│   ├── migration.go         # Model para controlar migrations
│   └── usuario.go           # Models da aplicação
└── db.go                    # Configuração do banco
```

## Como Funciona

### 1. Tabela de Controle

O sistema cria uma tabela `migrations` para rastrear quais migrations já foram executadas:

```sql
CREATE TABLE migrations (
    id SERIAL PRIMARY KEY,
    version VARCHAR(255) UNIQUE NOT NULL,
    name VARCHAR(255) NOT NULL,
    executed_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

### 2. Execução de Migrations

As migrations são executadas automaticamente quando a aplicação inicia, mas também podem ser executadas manualmente.

### 3. Versionamento

Cada migration tem:
- **Version**: Identificador único (ex: "001", "002")
- **Name**: Nome descritivo da migration
- **Up**: Função que executa a migration
- **Down**: Função que faz o rollback

## Comandos Disponíveis

### Executar Migrations

```bash
go run cmd/migrate/main.go run
```

### Verificar Status

```bash
go run cmd/migrate/main.go status
```

### Fazer Rollback

```bash
go run cmd/migrate/main.go rollback
```

## Criando Novas Migrations

### 🆕 Abordagem Declarativa (Recomendada)

Agora você pode criar migrations de forma muito mais simples e declarativa:

#### 1. Usando MigrationBuilder (Mais Simples)

```go
// Migration 008: Adicionar campos ao usuário
migration008 := m.NewMigration("008", "add_usuario_profile_fields").
    AddColumnSQL("usuarios", "bio", "TEXT").
    AddColumnSQL("usuarios", "website", "VARCHAR(255)").
    AddIndexSQL("usuarios", "idx_usuario_website", "website").
    Build()
m.migrations = append(m.migrations, migration008)
```

#### 2. Usando SQL Direto

```go
// Migration 009: Adicionar constraint
migration009 := m.NewMigration("009", "add_usuario_email_constraint").
    ExecuteSQL(
        "ALTER TABLE usuarios ADD CONSTRAINT chk_email_format CHECK (email ~* '^[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\\.[A-Za-z]{2,}$')",
        "ALTER TABLE usuarios DROP CONSTRAINT IF EXISTS chk_email_format",
    ).
    Build()
m.migrations = append(m.migrations, migration009)
```

#### 3. Múltiplas Operações

```go
// Migration 010: Múltiplas operações
migration010 := m.NewMigration("010", "add_usuario_complete_profile").
    AddColumnSQL("usuarios", "genero", "VARCHAR(20)").
    AddColumnSQL("usuarios", "data_ultimo_login", "TIMESTAMP").
    AddColumnSQL("usuarios", "tentativas_login", "INTEGER DEFAULT 0").
    AddIndexSQL("usuarios", "idx_usuario_ultimo_login", "data_ultimo_login").
    AddIndexSQL("usuarios", "idx_usuario_genero", "genero").
    Build()
m.migrations = append(m.migrations, migration010)
```

### Métodos Disponíveis no MigrationBuilder

| Método | Descrição | Exemplo |
|--------|-----------|---------|
| `AddColumnSQL` | Adiciona uma coluna | `AddColumnSQL("usuarios", "idade", "INTEGER")` |
| `DropColumnSQL` | Remove uma coluna | `DropColumnSQL("usuarios", "idade")` |
| `AddIndexSQL` | Adiciona um índice | `AddIndexSQL("usuarios", "idx_nome", "nome")` |
| `DropIndexSQL` | Remove um índice | `DropIndexSQL("idx_nome")` |
| `ExecuteSQL` | Executa SQL customizado | `ExecuteSQL("UPDATE usuarios SET ativo = true", "UPDATE usuarios SET ativo = false")` |

### Abordagem Tradicional (Ainda Disponível)

Se preferir, você ainda pode usar a abordagem tradicional:

```go
// Migration 011: Abordagem tradicional
m.migrations = append(m.migrations, Migration{
    Version: "011",
    Name:    "add_campo_tradicional",
    Up: func(db *gorm.DB) error {
        return db.Exec("ALTER TABLE usuarios ADD COLUMN IF NOT EXISTS campo_tradicional VARCHAR(100)").Error
    },
    Down: func(db *gorm.DB) error {
        return db.Exec("ALTER TABLE usuarios DROP COLUMN IF EXISTS campo_tradicional").Error
    },
})
```

## Migrations Existentes

### 001 - create_migrations_table
- **Objetivo**: Criar a tabela de controle de migrations
- **Status**: Sempre executada primeiro

### 002 - create_initial_tables
- **Objetivo**: Criar todas as tabelas iniciais da aplicação
- **Tabelas**: TipoPlano, CategoriaLojista, Usuario, Loja, HistoricoPlanoUsuario, Carteira, LogCarteira, Anuncio

### 003 - add_telefone_to_usuario
- **Objetivo**: Adicionar campo telefone ao usuário
- **Campo**: `telefone VARCHAR(20)`

### 004 - add_endereco_to_usuario
- **Objetivo**: Adicionar campo endereco ao usuário
- **Campo**: `endereco VARCHAR(500)`

### 005 - add_data_nascimento_to_usuario
- **Objetivo**: Adicionar campo data_nascimento e índice
- **Campos**: `data_nascimento DATE`, `idx_usuario_data_nascimento`

### 006 - add_cpf_index_to_usuario
- **Objetivo**: Adicionar índice no CPF
- **Índice**: `idx_usuario_cpf`

### 007 - add_usuario_fields
- **Objetivo**: Adicionar campos de controle
- **Campos**: `data_cadastro TIMESTAMP`, `ativo BOOLEAN`
- **Índice**: `idx_usuario_ativo`

## Exemplos Práticos

### Exemplo 1: Adicionar Campo Simples

```go
// Adicionar campo idade ao usuário
migration := m.NewMigration("012", "add_idade_to_usuario").
    AddColumnSQL("usuarios", "idade", "INTEGER").
    Build()
```

### Exemplo 2: Adicionar Campo com Índice

```go
// Adicionar campo cidade com índice
migration := m.NewMigration("013", "add_cidade_to_usuario").
    AddColumnSQL("usuarios", "cidade", "VARCHAR(100)").
    AddIndexSQL("usuarios", "idx_usuario_cidade", "cidade").
    Build()
```

### Exemplo 3: Múltiplas Operações

```go
// Adicionar vários campos de uma vez
migration := m.NewMigration("014", "add_usuario_profile").
    AddColumnSQL("usuarios", "bio", "TEXT").
    AddColumnSQL("usuarios", "website", "VARCHAR(255)").
    AddColumnSQL("usuarios", "redes_sociais", "JSONB").
    AddIndexSQL("usuarios", "idx_usuario_website", "website").
    AddIndexSQL("usuarios", "idx_usuario_redes_sociais", "redes_sociais").
    Build()
```

### Exemplo 4: SQL Customizado

```go
// Executar SQL customizado
migration := m.NewMigration("015", "update_existing_data").
    ExecuteSQL(
        "UPDATE usuarios SET ativo = true WHERE ativo IS NULL",
        "UPDATE usuarios SET ativo = false WHERE ativo IS NULL",
    ).
    Build()
```

## Boas Práticas

### 1. Sempre Teste em Desenvolvimento
- Execute as migrations em um ambiente de desenvolvimento primeiro
- Verifique se os dados existentes não são afetados

### 2. Use Transações
- Para migrations complexas, considere usar transações
- Isso garante que a migration seja executada completamente ou revertida

### 3. Documente Mudanças
- Sempre documente o que cada migration faz
- Mantenha um changelog das mudanças

### 4. Backup Antes de Migrations
- Sempre faça backup antes de executar migrations em produção
- Teste o rollback em ambiente de desenvolvimento

### 5. Versionamento
- Use números sequenciais para versões (001, 002, 003...)
- Mantenha as migrations em ordem cronológica

### 6. Nomenclatura
- Use nomes descritivos para as migrations
- Siga um padrão: `add_<campo>_to_<tabela>`, `remove_<campo>_from_<tabela>`, etc.

## Troubleshooting

### Migration Já Executada
Se uma migration já foi executada e você precisa reexecutá-la:
1. Remova o registro da tabela `migrations`
2. Execute a migration novamente

### Erro Durante Migration
Se uma migration falhar:
1. Verifique os logs para identificar o erro
2. Corrija o problema no código
3. Execute a migration novamente

### Rollback Falhou
Se o rollback falhar:
1. Verifique se a migration tem um método `Down` implementado
2. Execute o rollback manualmente se necessário
3. Remova o registro da tabela `migrations` 