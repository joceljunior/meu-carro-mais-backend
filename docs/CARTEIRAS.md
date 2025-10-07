# Carteiras - Documentação da API

## Visão Geral

O módulo de Carteiras permite gerenciar carteiras digitais dos usuários, incluindo operações de saldo, consultas e controle financeiro.

## Modelo de Dados

### Carteira
```json
{
  "id": 1,
  "usuario_id": 1,
  "saldo": 1000.00,
  "data_criacao": "2023-10-07T13:30:00Z",
  "data_atualizacao": "2023-10-07T13:30:00Z"
}
```

### Relacionamentos
- **Carteira** pertence a **Usuário** (1:1)
- Cada usuário pode ter apenas uma carteira

## Endpoints da API

### 1. Criar Carteira
**POST** `/carteiras`

Cria uma nova carteira para um usuário.

**Request Body:**
```json
{
  "usuario_id": 1,
  "saldo": 1000.00
}
```

**Response (201):**
```json
{
  "id": 1,
  "usuario_id": 1,
  "saldo": 1000.00,
  "data_criacao": "2023-10-07T13:30:00Z",
  "data_atualizacao": "2023-10-07T13:30:00Z",
  "mensagem": "Carteira criada com sucesso"
}
```

### 2. Listar Todas as Carteiras
**GET** `/carteiras`

Retorna todas as carteiras do sistema.

**Response (200):**
```json
{
  "carteiras": [
    {
      "id": 1,
      "usuario_id": 1,
      "saldo": 1000.00,
      "data_criacao": "2023-10-07T13:30:00Z",
      "data_atualizacao": "2023-10-07T13:30:00Z",
      "usuario": {
        "id": 1,
        "nome": "João Silva",
        "email": "joao@email.com"
      }
    }
  ],
  "total": 1,
  "mensagem": "Carteiras listadas com sucesso"
}
```

### 3. Buscar Carteira por ID
**GET** `/carteiras/{id}`

Retorna uma carteira específica pelo ID.

**Response (200):**
```json
{
  "id": 1,
  "usuario_id": 1,
  "saldo": 1000.00,
  "data_criacao": "2023-10-07T13:30:00Z",
  "data_atualizacao": "2023-10-07T13:30:00Z",
  "usuario": {
    "id": 1,
    "nome": "João Silva",
    "email": "joao@email.com"
  },
  "mensagem": "Carteira encontrada com sucesso"
}
```

### 4. Buscar Carteira por Usuário
**GET** `/carteiras/usuario/{usuario_id}`

Retorna a carteira de um usuário específico.

**Response (200):**
```json
{
  "id": 1,
  "usuario_id": 1,
  "saldo": 1000.00,
  "data_criacao": "2023-10-07T13:30:00Z",
  "data_atualizacao": "2023-10-07T13:30:00Z",
  "usuario": {
    "id": 1,
    "nome": "João Silva",
    "email": "joao@email.com"
  },
  "mensagem": "Carteira encontrada com sucesso"
}
```

### 5. Atualizar Carteira
**PUT** `/carteiras/{id}`

Atualiza os dados de uma carteira existente.

**Request Body:**
```json
{
  "usuario_id": 1,
  "saldo": 1500.00
}
```

**Response (200):**
```json
{
  "id": 1,
  "usuario_id": 1,
  "saldo": 1500.00,
  "data_criacao": "2023-10-07T13:30:00Z",
  "data_atualizacao": "2023-10-07T13:35:00Z",
  "mensagem": "Carteira atualizada com sucesso"
}
```

### 6. Atualizar Apenas Saldo
**PUT** `/carteiras/{id}/saldo`

Atualiza apenas o saldo de uma carteira.

**Request Body:**
```json
{
  "saldo": 2000.00
}
```

**Response (200):**
```json
{
  "id": 1,
  "usuario_id": 1,
  "saldo": 2000.00,
  "data_criacao": "2023-10-07T13:30:00Z",
  "data_atualizacao": "2023-10-07T13:40:00Z",
  "mensagem": "Saldo atualizado com sucesso"
}
```

### 7. Adicionar Saldo
**POST** `/carteiras/{id}/adicionar`

Adiciona um valor ao saldo atual da carteira.

**Request Body:**
```json
{
  "valor": 100.00
}
```

**Response (200):**
```json
{
  "id": 1,
  "usuario_id": 1,
  "saldo_anterior": 1000.00,
  "saldo_atual": 1100.00,
  "valor_operacao": 100.00,
  "tipo_operacao": "adicao",
  "data_atualizacao": "2023-10-07T13:45:00Z",
  "mensagem": "Saldo adicionado com sucesso"
}
```

### 8. Subtrair Saldo
**POST** `/carteiras/{id}/subtrair`

Subtrai um valor do saldo atual da carteira.

**Request Body:**
```json
{
  "valor": 50.00
}
```

**Response (200):**
```json
{
  "id": 1,
  "usuario_id": 1,
  "saldo_anterior": 1100.00,
  "saldo_atual": 1050.00,
  "valor_operacao": 50.00,
  "tipo_operacao": "subtracao",
  "data_atualizacao": "2023-10-07T13:50:00Z",
  "mensagem": "Saldo subtraído com sucesso"
}
```

### 9. Buscar Carteiras por Range de Saldo
**GET** `/carteiras/range?saldo_min={min}&saldo_max={max}`

Retorna carteiras com saldo dentro de um range específico.

**Query Parameters:**
- `saldo_min`: Saldo mínimo (obrigatório)
- `saldo_max`: Saldo máximo (obrigatório)

**Exemplo:** `GET /carteiras/range?saldo_min=500&saldo_max=2000`

**Response (200):**
```json
{
  "carteiras": [
    {
      "id": 1,
      "usuario_id": 1,
      "saldo": 1000.00,
      "data_criacao": "2023-10-07T13:30:00Z",
      "data_atualizacao": "2023-10-07T13:30:00Z",
      "usuario": {
        "id": 1,
        "nome": "João Silva",
        "email": "joao@email.com"
      }
    }
  ],
  "total": 1,
  "mensagem": "Carteiras encontradas com sucesso"
}
```

### 10. Buscar Carteiras com Saldo Maior
**GET** `/carteiras/saldo-maior?valor={valor}`

Retorna carteiras com saldo maior que um valor específico.

**Query Parameters:**
- `valor`: Valor mínimo de saldo (obrigatório)

**Exemplo:** `GET /carteiras/saldo-maior?valor=1000`

**Response (200):**
```json
{
  "carteiras": [
    {
      "id": 1,
      "usuario_id": 1,
      "saldo": 1500.00,
      "data_criacao": "2023-10-07T13:30:00Z",
      "data_atualizacao": "2023-10-07T13:30:00Z",
      "usuario": {
        "id": 1,
        "nome": "João Silva",
        "email": "joao@email.com"
      }
    }
  ],
  "total": 1,
  "mensagem": "Carteiras encontradas com sucesso"
}
```

### 11. Excluir Carteira
**DELETE** `/carteiras/{id}`

Remove uma carteira do sistema.

**Response (200):**
```json
{
  "message": "Carteira excluída com sucesso"
}
```

## Códigos de Erro

### 400 - Bad Request
```json
{
  "error": "Dados inválidos",
  "details": "Campo obrigatório não informado"
}
```

### 404 - Not Found
```json
{
  "error": "Carteira não encontrada"
}
```

### 400 - Saldo Insuficiente (apenas para subtração)
```json
{
  "error": "Saldo insuficiente"
}
```

### 500 - Internal Server Error
```json
{
  "error": "Erro interno do servidor"
}
```

## Regras de Negócio

1. **Carteira Única por Usuário**: Cada usuário pode ter apenas uma carteira
2. **Saldo Mínimo**: O saldo não pode ser negativo
3. **Validação de Usuário**: Só é possível criar carteira para usuários existentes
4. **Operações de Saldo**: 
   - Adição: Sempre permitida
   - Subtração: Só permitida se houver saldo suficiente
5. **Auditoria**: Todas as operações são registradas com timestamps

## Exemplos de Uso

### Cenário 1: Criar carteira para novo usuário
```bash
curl -X POST http://localhost:8080/carteiras \
  -H "Content-Type: application/json" \
  -d '{"usuario_id": 1, "saldo": 1000.00}'
```

### Cenário 2: Consultar saldo de um usuário
```bash
curl -X GET http://localhost:8080/carteiras/usuario/1
```

### Cenário 3: Adicionar crédito à carteira
```bash
curl -X POST http://localhost:8080/carteiras/1/adicionar \
  -H "Content-Type: application/json" \
  -d '{"valor": 100.00}'
```

### Cenário 4: Realizar compra (subtrair saldo)
```bash
curl -X POST http://localhost:8080/carteiras/1/subtrair \
  -H "Content-Type: application/json" \
  -d '{"valor": 50.00}'
```

## Integração com Outros Módulos

- **Usuários**: Carteira está vinculada a um usuário específico
- **Pagamentos**: Pode ser usada para processar pagamentos
- **Histórico de Resgates**: Pode ser usada para resgates de produtos/serviços

## Seeds

O sistema inclui seeds que criam carteiras com saldos variados para todos os usuários:

- João Silva: R$ 1.000,00
- Maria Santos: R$ 1.500,00
- Pedro Costa: R$ 800,00
- Carlos Porto Alegre: R$ 1.000,00
- Ana Silva Premium: R$ 5.000,00

Para executar os seeds:
```bash
go run cmd/seed/main.go run
```
