# Endpoint de Logs

O sistema de logs registra todas as ações dos usuários na aplicação, permitindo auditoria completa e rastreamento de atividades.

## Listar Todos os Logs

### GET /logs

Retorna todos os logs do sistema com paginação.

#### Request

**Método**: GET  
**URL**: `/logs`  
**Headers**: Nenhum obrigatório  
**Query Parameters**: 
- `limit` (opcional): Limite de resultados por página (padrão: 50, máximo: 1000)
- `offset` (opcional): Offset para paginação (padrão: 0)

#### Response

**Status: 200 OK**

```json
{
  "logs": [
    {
      "id": 1,
      "id_usuario": 1,
      "tipo_acao": "criar",
      "entidade": "anuncio",
      "id_entidade": 5,
      "descricao": "Anúncio criado: Troca de Óleo Completa",
      "dados_novos": {
        "id": 5,
        "titulo": "Troca de Óleo Completa",
        "preco": 89.90
      },
      "ip": "192.168.1.100",
      "user_agent": "Mozilla/5.0...",
      "metodo_http": "POST",
      "endpoint": "/anuncios",
      "status_http": 201,
      "data_acao": "2024-01-15T14:30:00Z",
      "data_cadastro": "2024-01-15T14:30:00Z",
      "data_atualizacao": "2024-01-15T14:30:00Z",
      "usuario": {
        "id": 1,
        "nome": "João Silva",
        "email": "joao@example.com"
      }
    }
  ],
  "total": 150,
  "limit": 50,
  "offset": 0
}
```

#### Exemplo de Uso

```bash
# Listar primeiros 50 logs
curl -X GET "http://localhost:8080/logs" \
  -H "Content-Type: application/json"

# Listar próximos 50 logs (página 2)
curl -X GET "http://localhost:8080/logs?limit=50&offset=50" \
  -H "Content-Type: application/json"
```

---

## Buscar Log por ID

### GET /logs/{id}

Retorna um log específico pelo ID.

#### Request

**Método**: GET  
**URL**: `/logs/{id}`  
**Parâmetros de Path**: 
- `id` (obrigatório): ID do log

#### Response

**Status: 200 OK**

```json
{
  "id": 1,
  "id_usuario": 1,
  "tipo_acao": "criar",
  "entidade": "anuncio",
  "id_entidade": 5,
  "descricao": "Anúncio criado: Troca de Óleo Completa",
  "dados_novos": { ... },
  "ip": "192.168.1.100",
  "user_agent": "Mozilla/5.0...",
  "metodo_http": "POST",
  "endpoint": "/anuncios",
  "status_http": 201,
  "data_acao": "2024-01-15T14:30:00Z",
  "usuario": { ... }
}
```

#### Exemplo de Uso

```bash
curl -X GET "http://localhost:8080/logs/1" \
  -H "Content-Type: application/json"
```

---

## Logs de um Usuário

### GET /logs/usuario/{id}

Retorna todos os logs de um usuário específico.

#### Request

**Método**: GET  
**URL**: `/logs/usuario/{id}`  
**Parâmetros de Path**: 
- `id` (obrigatório): ID do usuário

#### Response

**Status: 200 OK**

```json
[
  {
    "id": 1,
    "id_usuario": 1,
    "tipo_acao": "criar",
    "entidade": "anuncio",
    "descricao": "Anúncio criado",
    "data_acao": "2024-01-15T14:30:00Z"
  },
  {
    "id": 2,
    "id_usuario": 1,
    "tipo_acao": "resgatar",
    "entidade": "anuncio",
    "descricao": "Anúncio resgatado",
    "data_acao": "2024-01-15T15:00:00Z"
  }
]
```

#### Exemplo de Uso

```bash
curl -X GET "http://localhost:8080/logs/usuario/1" \
  -H "Content-Type: application/json"
```

---

## Logs de uma Entidade

### GET /logs/entidade/{entidade}/{id}

Retorna todos os logs de uma entidade específica (anúncio, produto, serviço, etc.).

#### Request

**Método**: GET  
**URL**: `/logs/entidade/{entidade}/{id}`  
**Parâmetros de Path**: 
- `entidade` (obrigatório): Nome da entidade (anuncio, produto, servico, veiculo, historico_resgate, registro_interesse, usuario, loja)
- `id` (obrigatório): ID da entidade

#### Response

**Status: 200 OK**

```json
[
  {
    "id": 1,
    "id_usuario": 1,
    "tipo_acao": "criar",
    "entidade": "anuncio",
    "id_entidade": 5,
    "descricao": "Anúncio criado",
    "data_acao": "2024-01-15T14:30:00Z"
  },
  {
    "id": 3,
    "id_usuario": 1,
    "tipo_acao": "atualizar",
    "entidade": "anuncio",
    "id_entidade": 5,
    "descricao": "Anúncio atualizado",
    "dados_antigos": { ... },
    "dados_novos": { ... },
    "data_acao": "2024-01-15T16:00:00Z"
  }
]
```

#### Exemplo de Uso

```bash
# Logs de um anúncio específico
curl -X GET "http://localhost:8080/logs/entidade/anuncio/5" \
  -H "Content-Type: application/json"

# Logs de um produto específico
curl -X GET "http://localhost:8080/logs/entidade/produto/10" \
  -H "Content-Type: application/json"
```

---

## Logs por Tipo de Ação

### GET /logs/acao/{tipo}

Retorna todos os logs de um tipo de ação específico.

#### Request

**Método**: GET  
**URL**: `/logs/acao/{tipo}`  
**Parâmetros de Path**: 
- `tipo` (obrigatório): Tipo de ação (criar, atualizar, deletar, resgatar, aprovar, rejeitar, restaurar, visualizar, registrar_interesse)

#### Response

**Status: 200 OK**

```json
[
  {
    "id": 1,
    "id_usuario": 1,
    "tipo_acao": "criar",
    "entidade": "anuncio",
    "descricao": "Anúncio criado",
    "data_acao": "2024-01-15T14:30:00Z"
  },
  {
    "id": 5,
    "id_usuario": 2,
    "tipo_acao": "criar",
    "entidade": "produto",
    "descricao": "Produto criado",
    "data_acao": "2024-01-15T17:00:00Z"
  }
]
```

#### Exemplo de Uso

```bash
# Todos os logs de criação
curl -X GET "http://localhost:8080/logs/acao/criar" \
  -H "Content-Type: application/json"

# Todos os logs de resgate
curl -X GET "http://localhost:8080/logs/acao/resgatar" \
  -H "Content-Type: application/json"

# Todos os logs de exclusão
curl -X GET "http://localhost:8080/logs/acao/deletar" \
  -H "Content-Type: application/json"
```

---

## Tipos de Ação Registrados

O sistema registra os seguintes tipos de ação:

- **criar**: Quando uma entidade é criada
- **atualizar**: Quando uma entidade é atualizada
- **deletar**: Quando uma entidade é excluída (soft delete)
- **restaurar**: Quando uma entidade excluída é restaurada
- **resgatar**: Quando um usuário resgata um anúncio
- **aprovar**: Quando uma loja aprova um resgate
- **rejeitar**: Quando uma loja rejeita um resgate
- **registrar_interesse**: Quando um usuário registra interesse em um veículo
- **visualizar**: Quando uma entidade é visualizada (se implementado)

---

## Entidades Monitoradas

O sistema registra logs para as seguintes entidades:

- **anuncio**: Anúncios de produtos/serviços/veículos
- **produto**: Produtos das lojas
- **servico**: Serviços oferecidos pelas lojas
- **veiculo**: Veículos dos usuários
- **historico_resgate**: Histórico de resgates
- **registro_interesse**: Registros de interesse em veículos
- **usuario**: Usuários do sistema
- **loja**: Lojas cadastradas
- **carteira**: Operações de carteira
- **avaliacao**: Avaliações de lojas/serviços

---

## Estrutura do Log

Cada log contém as seguintes informações:

- **id**: Identificador único do log
- **id_usuario**: ID do usuário que executou a ação (pode ser null)
- **tipo_acao**: Tipo de ação executada
- **entidade**: Nome da entidade afetada
- **id_entidade**: ID da entidade afetada
- **descricao**: Descrição textual da ação
- **dados_antigos**: Dados antes da alteração (para updates)
- **dados_novos**: Dados após a alteração
- **ip**: IP do usuário que executou a ação
- **user_agent**: Navegador/dispositivo usado
- **metodo_http**: Método HTTP usado (GET, POST, PUT, DELETE)
- **endpoint**: Endpoint chamado
- **status_http**: Status HTTP da resposta
- **data_acao**: Timestamp da ação
- **usuario**: Dados do usuário (se disponível)

---

## Casos de Uso

1. **Auditoria**: Rastrear todas as ações dos usuários
2. **Debugging**: Identificar problemas e comportamentos inesperados
3. **Análise**: Entender padrões de uso da aplicação
4. **Segurança**: Detectar atividades suspeitas
5. **Histórico**: Ver histórico completo de alterações em uma entidade
6. **Compliance**: Atender requisitos de auditoria e compliance

---

## Observações Importantes

- Os logs são salvos de forma **assíncrona** (não bloqueiam a resposta ao usuário)
- O sistema tenta extrair automaticamente o ID do usuário do contexto da requisição
- Logs excluídos (soft delete) não aparecem nas consultas
- A paginação é recomendada para grandes volumes de dados
- Os dados antigos e novos são armazenados em formato JSONB para auditoria completa

---

## Exemplos de Consultas Úteis

```bash
# Ver todas as criações de anúncios hoje
curl -X GET "http://localhost:8080/logs/acao/criar?entidade=anuncio"

# Ver histórico completo de um anúncio específico
curl -X GET "http://localhost:8080/logs/entidade/anuncio/5"

# Ver todas as ações de um usuário
curl -X GET "http://localhost:8080/logs/usuario/1"

# Ver todos os resgates aprovados
curl -X GET "http://localhost:8080/logs/acao/aprovar"
```

