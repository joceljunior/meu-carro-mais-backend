# Endpoint de Registro de Interesse em Veículos

## Criar Registro de Interesse

### POST /registro-interesse

Registra o interesse de um usuário em um anúncio de veículo.

#### Request

**Método**: POST  
**URL**: `/registro-interesse`  
**Headers**: 
- `Content-Type: application/json`

**Body**:
```json
{
  "id_anuncio": 1,
  "nome": "João Silva",
  "email": "joao@example.com",
  "telefone": "11999999999",
  "mensagem": "Tenho interesse neste veículo. Poderia me passar mais informações?"
}
```

**Campos obrigatórios**:
- `id_anuncio` (integer): ID do anúncio de veículo
- `nome` (string, 2-255 caracteres): Nome completo do interessado
- `email` (string, formato email válido): Email do interessado
- `telefone` (string, 10-20 caracteres): Telefone de contato
- `mensagem` (string, opcional, máximo 1000 caracteres): Mensagem adicional do interessado

#### Response

**Status: 201 Created**

```json
{
  "id": 1,
  "id_anuncio": 1,
  "nome": "João Silva",
  "email": "joao@example.com",
  "telefone": "11999999999",
  "mensagem": "Tenho interesse neste veículo. Poderia me passar mais informações?",
  "data_cadastro": "2024-01-15T10:30:00Z",
  "data_atualizacao": "2024-01-15T10:30:00Z",
  "anuncio": {
    "id": 1,
    "titulo": "Honda Civic 2020",
    "descricao": "Veículo em excelente estado",
    "preco": 85000.00,
    "imagem": "https://exemplo.com/civic.jpg",
    "destaque": true,
    "id_loja": 1,
    "id_categoria": 1,
    "categoria": "Veículos",
    "loja": {
      "id": 1,
      "nome": "Concessionária ABC",
      "cnpj": "12.345.678/0001-90",
      "imagem": "https://exemplo.com/loja.jpg",
      "latitude": -23.5505,
      "longitude": -46.6333,
      "id_categoria": 1,
      "categoria": "Concessionária"
    }
  }
}
```

**Status: 400 Bad Request** (dados inválidos)
```json
{
  "error": "Dados inválidos",
  "details": "campo 'email' é obrigatório"
}
```

**Status: 500 Internal Server Error**
```json
{
  "error": "anúncio não encontrado"
}
```

---

## Listar Todos os Registros de Interesse

### GET /registro-interesse

Retorna todos os registros de interesse ativos do sistema.

#### Request

**Método**: GET  
**URL**: `/registro-interesse`  
**Headers**: Nenhum obrigatório  
**Parâmetros**: Nenhum

#### Response

**Status: 200 OK**

```json
[
  {
    "id": 1,
    "id_anuncio": 1,
    "nome": "João Silva",
    "email": "joao@example.com",
    "telefone": "11999999999",
    "mensagem": "Tenho interesse neste veículo",
    "data_cadastro": "2024-01-15T10:30:00Z",
    "data_atualizacao": "2024-01-15T10:30:00Z",
    "anuncio": { ... }
  },
  {
    "id": 2,
    "id_anuncio": 2,
    "nome": "Maria Santos",
    "email": "maria@example.com",
    "telefone": "11988888888",
    "mensagem": "",
    "data_cadastro": "2024-01-15T11:00:00Z",
    "data_atualizacao": "2024-01-15T11:00:00Z",
    "anuncio": { ... }
  }
]
```

---

## Buscar Registro de Interesse por ID

### GET /registro-interesse/:id

Retorna um registro de interesse específico pelo ID.

#### Request

**Método**: GET  
**URL**: `/registro-interesse/:id`  
**Headers**: Nenhum obrigatório  
**Parâmetros de URL**:
- `id` (integer): ID do registro de interesse

#### Response

**Status: 200 OK**

```json
{
  "id": 1,
  "id_anuncio": 1,
  "nome": "João Silva",
  "email": "joao@example.com",
  "telefone": "11999999999",
  "mensagem": "Tenho interesse neste veículo",
  "data_cadastro": "2024-01-15T10:30:00Z",
  "data_atualizacao": "2024-01-15T10:30:00Z",
  "anuncio": { ... }
}
```

**Status: 400 Bad Request** (ID inválido)
```json
{
  "error": "ID inválido"
}
```

**Status: 404 Not Found**
```json
{
  "error": "Registro de interesse não encontrado"
}
```

---

## Listar Registros de Interesse por Anúncio

### GET /registro-interesse/anuncio/:anuncio_id

Retorna todos os registros de interesse de um anúncio específico.

#### Request

**Método**: GET  
**URL**: `/registro-interesse/anuncio/:anuncio_id`  
**Headers**: Nenhum obrigatório  
**Parâmetros de URL**:
- `anuncio_id` (integer): ID do anúncio

#### Response

**Status: 200 OK**

```json
[
  {
    "id": 1,
    "id_anuncio": 1,
    "nome": "João Silva",
    "email": "joao@example.com",
    "telefone": "11999999999",
    "mensagem": "Tenho interesse neste veículo",
    "data_cadastro": "2024-01-15T10:30:00Z",
    "data_atualizacao": "2024-01-15T10:30:00Z",
    "anuncio": { ... }
  },
  {
    "id": 3,
    "id_anuncio": 1,
    "nome": "Pedro Costa",
    "email": "pedro@example.com",
    "telefone": "11977777777",
    "mensagem": "Gostaria de agendar uma visita",
    "data_cadastro": "2024-01-15T14:00:00Z",
    "data_atualizacao": "2024-01-15T14:00:00Z",
    "anuncio": { ... }
  }
]
```

**Status: 400 Bad Request** (ID inválido)
```json
{
  "error": "ID do anúncio inválido"
}
```

---

## Excluir Registro de Interesse (Soft Delete)

### DELETE /registro-interesse/:id

Realiza soft delete do registro de interesse, marcando como excluído sem remover do banco.

#### Request

**Método**: DELETE  
**URL**: `/registro-interesse/:id`  
**Headers**: Nenhum obrigatório  
**Parâmetros de URL**:
- `id` (integer): ID do registro de interesse

#### Response

**Status: 200 OK**

```json
{
  "message": "Registro de interesse excluído com sucesso"
}
```

**Status: 400 Bad Request** (ID inválido)
```json
{
  "error": "ID inválido"
}
```

**Status: 404 Not Found**
```json
{
  "error": "Registro de interesse não encontrado"
}
```

---

## Restaurar Registro de Interesse

### POST /registro-interesse/:id/restore

Restaura um registro de interesse que foi soft deleted.

#### Request

**Método**: POST  
**URL**: `/registro-interesse/:id/restore`  
**Headers**: Nenhum obrigatório  
**Parâmetros de URL**:
- `id` (integer): ID do registro de interesse

#### Response

**Status: 200 OK**

```json
{
  "message": "Registro de interesse restaurado com sucesso"
}
```

**Status: 400 Bad Request** (ID inválido)
```json
{
  "error": "ID inválido"
}
```

**Status: 404 Not Found**
```json
{
  "error": "Registro de interesse não encontrado ou não foi excluído"
}
```

---

## Validações

### Campos Obrigatórios
- `id_anuncio`: Deve ser um ID válido de um anúncio existente e não excluído
- `nome`: Mínimo 2 caracteres, máximo 255 caracteres
- `email`: Deve ser um email válido
- `telefone`: Mínimo 10 caracteres, máximo 20 caracteres

### Campos Opcionais
- `mensagem`: Máximo 1000 caracteres

### Regras de Negócio
- O anúncio referenciado deve existir e não estar excluído (soft delete)
- Todos os registros são ordenados por data de cadastro (mais recentes primeiro)
- Soft delete não remove o registro do banco, apenas marca como excluído

---

## Exemplos de Uso

### Exemplo 1: Registrar interesse básico
```bash
curl -X POST http://localhost:8080/registro-interesse \
  -H "Content-Type: application/json" \
  -d '{
    "id_anuncio": 1,
    "nome": "João Silva",
    "email": "joao@example.com",
    "telefone": "11999999999"
  }'
```

### Exemplo 2: Registrar interesse com mensagem
```bash
curl -X POST http://localhost:8080/registro-interesse \
  -H "Content-Type: application/json" \
  -d '{
    "id_anuncio": 1,
    "nome": "Maria Santos",
    "email": "maria@example.com",
    "telefone": "11988888888",
    "mensagem": "Gostaria de agendar uma visita para ver o veículo pessoalmente."
  }'
```

### Exemplo 3: Listar interesses de um anúncio específico
```bash
curl -X GET http://localhost:8080/registro-interesse/anuncio/1
```

### Exemplo 4: Buscar registro específico
```bash
curl -X GET http://localhost:8080/registro-interesse/1
```

### Exemplo 5: Excluir registro (soft delete)
```bash
curl -X DELETE http://localhost:8080/registro-interesse/1
```

### Exemplo 6: Restaurar registro excluído
```bash
curl -X POST http://localhost:8080/registro-interesse/1/restore
```

