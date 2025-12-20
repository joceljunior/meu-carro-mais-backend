# Descontos - Sistema de Desconto para Lojas

## Visão Geral

O sistema de descontos permite que uma loja defina uma porcentagem mínima de desconto que será aplicada em seus produtos e serviços. 

**Regra Principal**: Uma loja só pode ter **um desconto ativo por vez**. Para criar um novo desconto, é necessário cancelar o atual primeiro.

## Estrutura da Tabela

```sql
CREATE TABLE descontos (
    id SERIAL PRIMARY KEY,
    id_loja INTEGER NOT NULL,          -- FK para lojas
    porcentagem DECIMAL(5,2) NOT NULL, -- Porcentagem de desconto (0-100)
    ativo BOOLEAN DEFAULT TRUE,        -- Se o desconto está ativo
    data_validade TIMESTAMP NOT NULL,  -- Data de expiração do desconto
    data_cadastro TIMESTAMP,           -- Data de criação
    data_atualizacao TIMESTAMP,        -- Data de última atualização
    data_exclusao TIMESTAMP            -- Soft delete
);
```

## Endpoints

### Descontos

| Método | Endpoint | Descrição |
|--------|----------|-----------|
| POST | `/descontos` | Criar novo desconto |
| GET | `/descontos` | Listar todos os descontos |
| GET | `/descontos/ativos` | Listar descontos ativos |
| GET | `/descontos/:id` | Buscar desconto por ID |
| POST | `/descontos/:id/cancelar` | Cancelar desconto (desativar) |
| DELETE | `/descontos/:id` | Soft delete do desconto |
| POST | `/descontos/:id/restore` | Restaurar desconto excluído |

### Descontos via Loja

| Método | Endpoint | Descrição |
|--------|----------|-----------|
| GET | `/lojas/:id/descontos` | Histórico de descontos da loja |
| GET | `/lojas/:id/desconto-ativo` | Desconto ativo da loja |
| POST | `/lojas/:id/desconto-ativo/cancelar` | Cancelar desconto ativo da loja |

## Request/Response

### Criar Desconto

**Request:**
```json
POST /descontos
{
  "id_loja": 1,
  "porcentagem": 15.5,
  "data_validade": "2025-12-31T23:59:59Z"
}
```

**Response (201 Created):**
```json
{
  "id": 1,
  "id_loja": 1,
  "porcentagem": 15.5,
  "ativo": true,
  "data_validade": "2025-12-31T23:59:59Z",
  "data_cadastro": "2024-01-15T10:30:00Z",
  "data_atualizacao": "2024-01-15T10:30:00Z",
  "loja": {
    "id": 1,
    "nome": "Auto Center XYZ",
    "cnpj": "12.345.678/0001-90",
    ...
  }
}
```

**Erro (400 - Já possui desconto ativo):**
```json
{
  "error": "esta loja já possui um desconto ativo. Cancele o desconto atual antes de criar um novo"
}
```

### Listar Descontos Ativos

**Response:**
```json
{
  "descontos": [
    {
      "id": 1,
      "id_loja": 1,
      "porcentagem": 15.5,
      "ativo": true,
      "data_validade": "2025-12-31T23:59:59Z",
      ...
    }
  ],
  "total": 1
}
```

### Cancelar Desconto

**Request:**
```
POST /descontos/1/cancelar
```

**Response (200):**
```json
{
  "message": "Desconto cancelado com sucesso"
}
```

## Fluxo de Uso

1. **Criar desconto**: A loja cria um desconto com porcentagem e data de validade
2. **Desconto ativo**: O desconto fica ativo automaticamente
3. **Validação**: O sistema impede criar novo desconto enquanto houver um ativo
4. **Cancelar**: Para mudar o desconto, primeiro cancele o atual
5. **Novo desconto**: Após cancelar, pode criar um novo desconto
6. **Expiração automática**: Descontos expirados não aparecem como "ativos"

## Regras de Negócio

1. **Um desconto ativo por loja**: Não é possível ter dois descontos ativos simultaneamente
2. **Porcentagem válida**: Deve estar entre 0 e 100
3. **Data de validade obrigatória**: O desconto deve ter data de expiração
4. **Cancelamento vs Exclusão**: 
   - `Cancelar`: Desativa o desconto (mantém histórico)
   - `Delete`: Soft delete (marca como excluído)
5. **Descontos expirados**: Automaticamente não são considerados ativos

## Exemplo de Uso

```bash
# 1. Criar desconto de 10% para loja 1
curl -X POST http://localhost:8080/descontos \
  -H "Content-Type: application/json" \
  -d '{"id_loja": 1, "porcentagem": 10, "data_validade": "2025-12-31T23:59:59Z"}'

# 2. Verificar desconto ativo
curl http://localhost:8080/lojas/1/desconto-ativo

# 3. Tentar criar outro (vai falhar)
curl -X POST http://localhost:8080/descontos \
  -H "Content-Type: application/json" \
  -d '{"id_loja": 1, "porcentagem": 20, "data_validade": "2025-12-31T23:59:59Z"}'
# Retorna: {"error": "esta loja já possui um desconto ativo..."}

# 4. Cancelar o desconto atual
curl -X POST http://localhost:8080/lojas/1/desconto-ativo/cancelar

# 5. Agora pode criar novo desconto
curl -X POST http://localhost:8080/descontos \
  -H "Content-Type: application/json" \
  -d '{"id_loja": 1, "porcentagem": 20, "data_validade": "2025-12-31T23:59:59Z"}'

# 6. Ver histórico de descontos da loja
curl http://localhost:8080/lojas/1/descontos
```

## Migration

A tabela é criada pela migration `019_create_descontos_table`. Execute:

```bash
go run cmd/migrate/main.go
```

