# Endpoint de Histórico de Resgate

O histórico de resgate registra quando um usuário resgata um anúncio (produto, serviço ou veículo). O resgate inicia com status "pendente" e precisa ser aprovado pela loja.

## Listar Históricos de Resgate

### GET /historicos-resgate

Retorna todos os históricos de resgate ativos do sistema.

#### Request

**Método**: GET  
**URL**: `/historicos-resgate`  
**Headers**: Nenhum obrigatório  

#### Response

**Status: 200 OK**

```json
{
  "historicos": [
    {
      "id": 1,
      "id_usuario": 1,
      "id_produto": null,
      "id_servico": 5,
      "id_veiculo": null,
      "id_loja": 1,
      "tipo_resgate": "servico",
      "valor": 89.90,
      "status": "pendente",
      "data_resgate": "2024-01-15T14:30:00Z",
      "data_atualizacao": "2024-01-15T14:30:00Z",
      "usuario": {
        "id": 1,
        "nome": "João Silva",
        "email": "joao@example.com"
      },
      "loja": {
        "id": 1,
        "nome": "Oficina São Paulo"
      },
      "servico": {
        "id": 5,
        "titulo": "Troca de Óleo Completa",
        "preco": 89.90
      }
    }
  ],
  "total": 1
}
```

---

## Buscar Histórico por ID

### GET /historicos-resgate/{id}

Retorna um histórico de resgate específico pelo ID.

#### Request

**Método**: GET  
**URL**: `/historicos-resgate/{id}`  
**Parâmetros de Path**: 
- `id` (obrigatório): ID do histórico de resgate

#### Response

**Status: 200 OK**

```json
{
  "id": 1,
  "id_usuario": 1,
  "id_servico": 5,
  "id_loja": 1,
  "tipo_resgate": "servico",
  "valor": 89.90,
  "status": "pendente",
  "data_resgate": "2024-01-15T14:30:00Z",
  "data_atualizacao": "2024-01-15T14:30:00Z",
  "usuario": { ... },
  "loja": { ... },
  "servico": { ... }
}
```

---

## Históricos de Resgate por Usuário

### GET /users/{id}/historicos-resgate

Retorna todos os históricos de resgate de um usuário específico.

#### Request

**Método**: GET  
**URL**: `/users/{id}/historicos-resgate`  
**Parâmetros de Path**: 
- `id` (obrigatório): ID do usuário

#### Response

**Status: 200 OK**

```json
{
  "historicos": [
    {
      "id": 1,
      "id_usuario": 1,
      "tipo_resgate": "servico",
      "valor": 89.90,
      "status": "pendente",
      "data_resgate": "2024-01-15T14:30:00Z",
      "loja": { ... },
      "servico": { ... }
    }
  ],
  "total": 1
}
```

#### Exemplo de Uso

```bash
curl -X GET "http://localhost:8080/users/1/historicos-resgate" \
  -H "Content-Type: application/json"
```

---

## Históricos de Resgate por Loja

### GET /lojas/{id}/historicos-resgate

Retorna todos os históricos de resgate de uma loja específica. Útil para a loja visualizar todos os resgates pendentes de aprovação.

#### Request

**Método**: GET  
**URL**: `/lojas/{id}/historicos-resgate`  
**Parâmetros de Path**: 
- `id` (obrigatório): ID da loja

#### Response

**Status: 200 OK**

```json
{
  "historicos": [
    {
      "id": 1,
      "id_usuario": 1,
      "tipo_resgate": "servico",
      "valor": 89.90,
      "status": "pendente",
      "data_resgate": "2024-01-15T14:30:00Z",
      "usuario": {
        "id": 1,
        "nome": "João Silva",
        "email": "joao@example.com"
      },
      "servico": {
        "id": 5,
        "titulo": "Troca de Óleo Completa",
        "preco": 89.90
      }
    }
  ],
  "total": 1
}
```

#### Exemplo de Uso

```bash
curl -X GET "http://localhost:8080/lojas/1/historicos-resgate" \
  -H "Content-Type: application/json"
```

---

## Aprovar Resgate

### PUT /historicos-resgate/{id}/aprovar

Aprova um resgate pendente, alterando o status para "confirmado". Apenas resgates com status "pendente" podem ser aprovados.

#### Request

**Método**: PUT  
**URL**: `/historicos-resgate/{id}/aprovar`  
**Parâmetros de Path**: 
- `id` (obrigatório): ID do histórico de resgate

#### Response

**Status: 200 OK**

```json
{
  "id": 1,
  "id_usuario": 1,
  "tipo_resgate": "servico",
  "valor": 89.90,
  "status": "confirmado",
  "data_resgate": "2024-01-15T14:30:00Z",
  "data_atualizacao": "2024-01-15T15:00:00Z",
  "usuario": { ... },
  "loja": { ... },
  "servico": { ... }
}
```

#### Erros

**Status: 400 Bad Request** - Resgate não está pendente
```json
{
  "error": "Apenas resgates com status 'pendente' podem ser aprovados"
}
```

**Status: 404 Not Found** - Histórico não encontrado
```json
{
  "error": "Histórico não encontrado"
}
```

#### Exemplo de Uso

```bash
curl -X PUT "http://localhost:8080/historicos-resgate/1/aprovar" \
  -H "Content-Type: application/json"
```

---

## Rejeitar Resgate

### PUT /historicos-resgate/{id}/rejeitar

Rejeita um resgate pendente, alterando o status para "cancelado". Apenas resgates com status "pendente" podem ser rejeitados.

#### Request

**Método**: PUT  
**URL**: `/historicos-resgate/{id}/rejeitar`  
**Parâmetros de Path**: 
- `id` (obrigatório): ID do histórico de resgate

#### Response

**Status: 200 OK**

```json
{
  "id": 1,
  "id_usuario": 1,
  "tipo_resgate": "servico",
  "valor": 89.90,
  "status": "cancelado",
  "data_resgate": "2024-01-15T14:30:00Z",
  "data_atualizacao": "2024-01-15T15:00:00Z",
  "usuario": { ... },
  "loja": { ... },
  "servico": { ... }
}
```

#### Erros

**Status: 400 Bad Request** - Resgate não está pendente
```json
{
  "error": "Apenas resgates com status 'pendente' podem ser rejeitados"
}
```

**Status: 404 Not Found** - Histórico não encontrado
```json
{
  "error": "Histórico não encontrado"
}
```

#### Exemplo de Uso

```bash
curl -X PUT "http://localhost:8080/historicos-resgate/1/rejeitar" \
  -H "Content-Type: application/json"
```

---

## Atualizar Status do Resgate

### PUT /historicos-resgate/{id}/status

Atualiza apenas o status de um histórico de resgate. Permite alterar para qualquer status válido: "pendente", "confirmado" ou "cancelado".

#### Request

**Método**: PUT  
**URL**: `/historicos-resgate/{id}/status`  
**Parâmetros de Path**: 
- `id` (obrigatório): ID do histórico de resgate

**Body**:
```json
{
  "status": "confirmado"
}
```

#### Response

**Status: 200 OK**

```json
{
  "message": "Status atualizado com sucesso",
  "status": "confirmado"
}
```

#### Erros

**Status: 400 Bad Request** - Status inválido
```json
{
  "error": "Status inválido. Valores aceitos: pendente, confirmado, cancelado"
}
```

#### Exemplo de Uso

```bash
curl -X PUT "http://localhost:8080/historicos-resgate/1/status" \
  -H "Content-Type: application/json" \
  -d '{
    "status": "confirmado"
  }'
```

---

## Criar Histórico de Resgate Manualmente

### POST /historicos-resgate

Cria um histórico de resgate manualmente. Normalmente não é necessário usar este endpoint, pois o resgate é criado automaticamente ao resgatar um anúncio.

#### Request

**Método**: POST  
**URL**: `/historicos-resgate`  
**Headers**: 
- `Content-Type: application/json`

**Body**:
```json
{
  "id_usuario": 1,
  "id_produto": null,
  "id_servico": 5,
  "id_veiculo": null,
  "id_loja": 1,
  "tipo_resgate": "servico",
  "valor": 89.90,
  "status": "pendente"
}
```

#### Response

**Status: 201 Created**

```json
{
  "id": 1,
  "id_usuario": 1,
  "tipo_resgate": "servico",
  "valor": 89.90,
  "status": "pendente",
  "data_resgate": "2024-01-15T14:30:00Z",
  "usuario": { ... },
  "loja": { ... },
  "servico": { ... }
}
```

---

## Status do Resgate

Os históricos de resgate podem ter os seguintes status:

- **pendente**: Resgate criado, aguardando aprovação da loja
- **confirmado**: Resgate aprovado pela loja
- **cancelado**: Resgate rejeitado pela loja

### Fluxo de Status

1. **Criação**: Quando um usuário resgata um anúncio, o histórico é criado com status `pendente`
2. **Aprovação**: A loja pode aprovar o resgate, alterando o status para `confirmado`
3. **Rejeição**: A loja pode rejeitar o resgate, alterando o status para `cancelado`

---

## Observações Importantes

- Todos os resgates são criados com status "pendente" por padrão
- Apenas a loja pode aprovar ou rejeitar resgates
- O histórico inclui informações completas do usuário, loja e produto/serviço/veículo
- Resgates excluídos (soft delete) não aparecem nas listagens
- Os endpoints de aprovar/rejeitar validam se o resgate está pendente antes de alterar o status

---

## Relacionamentos

- **Histórico → Usuário**: Cada histórico pertence a um usuário
- **Histórico → Loja**: Cada histórico pertence a uma loja
- **Histórico → Produto/Serviço/Veículo**: Cada histórico está associado a um produto, serviço ou veículo (dependendo do tipo)

---

## Casos de Uso

1. **Usuário visualiza seus resgates**: `GET /users/{id}/historicos-resgate`
2. **Loja visualiza resgates pendentes**: `GET /lojas/{id}/historicos-resgate`
3. **Loja aprova resgate**: `PUT /historicos-resgate/{id}/aprovar`
4. **Loja rejeita resgate**: `PUT /historicos-resgate/{id}/rejeitar`
5. **Sistema cria resgate automaticamente**: `POST /anuncios/{id}/resgatar`

