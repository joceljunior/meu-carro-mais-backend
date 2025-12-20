# Endpoint de Usuários

## Tipos de Usuário

O sistema possui 4 tipos de usuários:

| Tipo | Descrição | Pode logar no Mobile | Status Inicial |
|------|-----------|---------------------|----------------|
| `mobile` | Usuário padrão do aplicativo mobile | ✅ Sim | `aprovado` |
| `administrativo` | Donos do sistema, todos os poderes | ✅ Sim | `aprovado` |
| `executivo` | Pode criar customers e recebe bonificação | ✅ Sim | `aprovado` |
| `customer` | Pode criar lojas e produtos (web) | ❌ Não | `pendente` |

## Status de Usuário

| Status | Descrição |
|--------|-----------|
| `pendente` | Aguardando aprovação (usado para customers) |
| `aprovado` | Usuário aprovado e ativo |
| `rejeitado` | Usuário rejeitado |

---

## Criação de Usuário Mobile (Padrão)

### POST /users

Cria um novo usuário do tipo **mobile** com todos os dados fornecidos.

#### Request Body

```json
{
  "nome": "João Silva",
  "email": "joao@email.com",
  "senha": "123456",
  "cpf": "123.456.789-00",
  "imagem": "https://exemplo.com/foto.jpg",
  "telefone": "(11) 99999-9999",
  "endereco": "Rua das Flores, 123",
  "data_nascimento": "1990-01-01T00:00:00Z",
  "latitude": -23.5505,
  "longitude": -46.6333
}
```

#### Response

**Status: 201 Created**

```json
{
  "id": 1,
  "nome": "João Silva",
  "email": "joao@email.com",
  "cpf": "123.456.789-00",
  "imagem": "https://exemplo.com/foto.jpg",
  "telefone": "(11) 99999-9999",
  "endereco": "Rua das Flores, 123",
  "data_nascimento": "1990-01-01T00:00:00Z",
  "data_cadastro": "2024-01-15T10:30:00Z",
  "ativo": true,
  "latitude": -23.5505,
  "longitude": -46.6333,
  "id_plano": 1,
  "id_loja": null,
  "tipo": "mobile",
  "status": "aprovado",
  "mensagem": "Usuário criado com sucesso"
}
```

---

## Endpoints Administrativos

### POST /users/administrativo

Cria um novo usuário do tipo **administrativo** (dono do sistema).

#### Request Body

```json
{
  "nome": "Admin Master",
  "email": "admin@meucarro.com",
  "senha": "senha123",
  "cpf": "111.111.111-11",
  "telefone": "(11) 99999-0000"
}
```

#### Response

**Status: 201 Created**

```json
{
  "id": 1,
  "nome": "Admin Master",
  "email": "admin@meucarro.com",
  "tipo": "administrativo",
  "status": "aprovado",
  "mensagem": "Usuário administrativo criado com sucesso"
}
```

---

### POST /users/executivo

Cria um novo usuário do tipo **executivo** (pode criar customers).

#### Request Body

```json
{
  "nome": "Executivo Vendas",
  "email": "executivo@meucarro.com",
  "senha": "senha123",
  "cpf": "222.222.222-22",
  "telefone": "(11) 99999-1111"
}
```

#### Response

**Status: 201 Created**

```json
{
  "id": 2,
  "nome": "Executivo Vendas",
  "email": "executivo@meucarro.com",
  "tipo": "executivo",
  "status": "aprovado",
  "mensagem": "Usuário executivo criado com sucesso"
}
```

---

## Endpoints Customer

### POST /users/customer

Cria um novo usuário do tipo **customer** (sempre inicia com status `pendente`).

#### Request Body

```json
{
  "nome": "Loja do João",
  "email": "loja@joao.com",
  "senha": "senha123",
  "cpf": "333.333.333-33",
  "telefone": "(11) 99999-2222",
  "id_executivo": 2
}
```

> **Nota**: O campo `id_executivo` é opcional. Se informado, vincula o customer a um executivo que receberá bonificação quando o customer for aprovado.

#### Response

**Status: 201 Created**

```json
{
  "id": 3,
  "nome": "Loja do João",
  "email": "loja@joao.com",
  "tipo": "customer",
  "status": "pendente",
  "id_executivo": 2,
  "executivo": {
    "id": 2,
    "nome": "Executivo Vendas",
    "email": "executivo@meucarro.com"
  },
  "mensagem": "Customer criado com sucesso. Aguardando aprovação."
}
```

---

### GET /users/customers

Lista todos os customers com filtro opcional por status.

#### Query Parameters

| Parâmetro | Tipo | Descrição |
|-----------|------|-----------|
| `status` | string | Filtrar por: `pendente`, `aprovado`, `rejeitado` |

#### Exemplos

- `GET /users/customers` - Lista todos
- `GET /users/customers?status=pendente` - Lista apenas pendentes
- `GET /users/customers?status=aprovado` - Lista apenas aprovados

#### Response

**Status: 200 OK**

```json
{
  "customers": [
    {
      "id": 3,
      "nome": "Loja do João",
      "email": "loja@joao.com",
      "tipo": "customer",
      "status": "pendente",
      "id_executivo": 2,
      "executivo": {
        "id": 2,
        "nome": "Executivo Vendas",
        "email": "executivo@meucarro.com"
      }
    }
  ],
  "total": 1,
  "mensagem": "Customers listados com sucesso"
}
```

---

### GET /users/customers/pendentes

Lista apenas customers com status `pendente`.

#### Response

**Status: 200 OK**

```json
{
  "customers": [...],
  "total": 5,
  "mensagem": "Customers pendentes listados com sucesso"
}
```

---

### POST /users/customers/{id}/aprovar

Aprova um customer pendente. Se o customer foi criado por um executivo, o executivo recebe **100 moedas** de bonificação.

#### Request Body (opcional)

```json
{
  "motivo": "Documentação verificada"
}
```

#### Response

**Status: 200 OK**

```json
{
  "id": 3,
  "status": "aprovado",
  "motivo": "Documentação verificada",
  "bonificacao_executivo": 100,
  "mensagem": "Customer aprovado com sucesso. Executivo bonificado com 100 moedas"
}
```

---

### POST /users/customers/{id}/rejeitar

Rejeita um customer pendente (motivo obrigatório).

#### Request Body

```json
{
  "motivo": "Documentação inválida"
}
```

#### Response

**Status: 200 OK**

```json
{
  "id": 3,
  "status": "rejeitado",
  "motivo": "Documentação inválida",
  "mensagem": "Customer rejeitado"
}
```

---

## Validações Importantes

1. **Email** e **CPF** devem ser únicos no sistema
2. **Senha** deve ter pelo menos 6 caracteres
3. **Customers** não podem fazer login no aplicativo mobile
4. Apenas customers com status `pendente` podem ser aprovados/rejeitados
5. A bonificação ao executivo só ocorre na aprovação

## Fluxo de Customer

```
1. Customer é criado (status: pendente)
   ↓
2. Administrativo visualiza lista de pendentes
   ↓
3. Administrativo aprova ou rejeita
   ↓
   Se aprovado + tem executivo vinculado:
   → Executivo recebe 100 moedas na carteira
```
