# Endpoint de Usuários

## Tipos de Usuário

O sistema possui 4 tipos de usuários:

| Tipo | Descrição | Login Mobile | Login Web | Status Inicial |
|------|-----------|--------------|-----------|----------------|
| `mobile` | Usuário padrão do aplicativo mobile | ✅ Sim | ❌ Não | `aprovado` |
| `administrativo` | Donos do sistema, todos os poderes | ❌ Não | ✅ Sim | `aprovado` |
| `executivo` | Pode criar customers e recebe bonificação | ❌ Não | ✅ Sim | `aprovado` |
| `customer` | Pode criar lojas e produtos (web) | ❌ Não | ✅ Sim* | `pendente` |

> *Customers precisam estar **aprovados** para fazer login na web.

## Status de Usuário

| Status | Descrição |
|--------|-----------|
| `pendente` | Aguardando aprovação (usado para customers) |
| `aprovado` | Usuário aprovado e ativo |
| `rejeitado` | Usuário rejeitado |

---

## Endpoints de Autenticação

### POST /login (Mobile)

Login para o aplicativo mobile. Se o usuário não existir, cria automaticamente como tipo `mobile`.

**Tipos permitidos:** `mobile` (apenas)

**Tipos bloqueados:** `customer`, `administrativo`, `executivo`

#### Request Body

```json
{
  "email": "usuario@email.com",
  "senha": "senha123"
}
```

#### Response

**Status: 200 OK**

```json
{
  "id": 1,
  "nome": "João Silva",
  "email": "usuario@email.com",
  "tipo": "mobile",
  "status": "aprovado",
  "nome_plano": "Gratuito",
  "loja": {...}
}
```

---

### POST /login/web

Login para a plataforma web. Não cria usuário automaticamente.

**Tipos permitidos:** `administrativo`, `executivo`, `customer` (aprovados)

**Tipos bloqueados:** `mobile`

#### Request Body

```json
{
  "email": "admin@meucarro.com",
  "senha": "senha123"
}
```

#### Response

**Status: 200 OK**

```json
{
  "id": 1,
  "nome": "Admin Master",
  "email": "admin@meucarro.com",
  "tipo": "administrativo",
  "status": "aprovado",
  "nome_plano": "Gratuito",
  "loja": {...}
}
```

#### Erros Possíveis

| Status | Mensagem | Causa |
|--------|----------|-------|
| 404 | "usuário não encontrado" | Email não existe no sistema |
| 401 | "usuários do tipo mobile não podem fazer login na plataforma web" | Tipo não permitido |
| 401 | "sua conta está pendente de aprovação" | Customer ainda não aprovado |
| 401 | "sua conta foi rejeitada" | Customer foi rejeitado |

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

## Solicitação para Virar Executivo

Um usuário **mobile** pode solicitar se tornar **executivo** para poder indicar customers e receber bonificações.

### POST /users/{id}/solicitar-executivo

Usuário mobile solicita virar executivo.

#### Request Body

```json
{
  "motivo": "Quero indicar lojas e ganhar bonificações. Tenho contatos no ramo automotivo."
}
```

#### Response

**Status: 200 OK**

```json
{
  "id": 5,
  "nome": "João Mobile",
  "email": "joao@email.com",
  "tipo": "mobile",
  "solicitacao_executivo": "pendente",
  "data_solicitacao_executivo": "2024-01-15T10:30:00Z",
  "motivo_solicitacao_executivo": "Quero indicar lojas e ganhar bonificações.",
  "mensagem": "Solicitação para virar executivo enviada com sucesso. Aguardando aprovação."
}
```

---

### GET /users/solicitacoes-executivo

Lista todas as solicitações de executivo pendentes (para administrativo).

#### Response

**Status: 200 OK**

```json
{
  "solicitacoes": [
    {
      "id": 5,
      "nome": "João Mobile",
      "email": "joao@email.com",
      "tipo": "mobile",
      "solicitacao_executivo": "pendente",
      "data_solicitacao_executivo": "2024-01-15T10:30:00Z",
      "motivo_solicitacao_executivo": "Quero indicar lojas..."
    }
  ],
  "total": 1,
  "mensagem": "Solicitações pendentes listadas com sucesso"
}
```

---

### POST /users/{id}/aprovar-executivo

Aprova a solicitação. O usuário passa de **mobile** para **executivo**.

#### Request Body (opcional)

```json
{
  "motivo": "Perfil aprovado"
}
```

#### Response

**Status: 200 OK**

```json
{
  "id": 5,
  "nome": "João Mobile",
  "email": "joao@email.com",
  "tipo": "executivo",
  "solicitacao_executivo": "aprovada",
  "mensagem": "Solicitação aprovada! Usuário agora é executivo."
}
```

---

### POST /users/{id}/rejeitar-executivo

Rejeita a solicitação (motivo obrigatório). O usuário continua como **mobile**.

#### Request Body

```json
{
  "motivo": "Perfil não atende aos requisitos"
}
```

#### Response

**Status: 200 OK**

```json
{
  "id": 5,
  "nome": "João Mobile",
  "email": "joao@email.com",
  "tipo": "mobile",
  "solicitacao_executivo": "rejeitada",
  "mensagem": "Solicitação rejeitada."
}
```

---

## Validações Importantes

1. **Email** e **CPF** devem ser únicos no sistema
2. **Senha** deve ter pelo menos 6 caracteres
3. **Customers** não podem fazer login no aplicativo mobile
4. Apenas customers com status `pendente` podem ser aprovados/rejeitados
5. A bonificação ao executivo só ocorre na aprovação
6. Apenas usuários **mobile** podem solicitar virar executivo
7. Usuário não pode ter mais de uma solicitação pendente

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

## Fluxo de Solicitação de Executivo

```
1. Usuário mobile solicita virar executivo
   ↓
2. Solicitação fica pendente (solicitacao_executivo: "pendente")
   ↓
3. Administrativo visualiza lista de solicitações
   ↓
4. Administrativo aprova ou rejeita
   ↓
   Se APROVADO:
   → Tipo muda de "mobile" para "executivo"
   → Usuário pode criar customers e receber bonificações
   
   Se REJEITADO:
   → Usuário continua como "mobile"
   → Pode solicitar novamente no futuro
```
