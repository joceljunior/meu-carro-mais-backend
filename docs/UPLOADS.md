# Endpoint de Uploads

O sistema de uploads permite salvar imagens e documentos associados a diferentes entidades (usuários, veículos, produtos, serviços, lojas).

## Criar Upload

### POST /uploads

Cria um novo upload (imagem ou documento) no sistema.

#### Request

**Método**: POST  
**URL**: `/uploads`  
**Headers**: 
- `Content-Type: application/json`

**Body**:
```json
{
  "id_usuario": 1,
  "tipo_entidade": "usuario",
  "tipo": "Imagem",
  "url": "https://exemplo.com/foto.jpg",
  "nome_arquivo": "foto_perfil.jpg",
  "tamanho": 102400,
  "tipo_mime": "image/jpeg",
  "principal": true,
  "ordem": 0
}
```

**Campos**:
- `id_usuario`, `id_veiculo`, `id_produto`, `id_servico`, `id_loja` (opcional): Apenas um deve ser informado
- `tipo_entidade` (obrigatório): "usuario", "veiculo", "veiculo_loja", "produto", "servico" ou "loja"
- `tipo` (obrigatório): "Imagem" ou "Documento"
- `url` (obrigatório): URL do arquivo
- `nome_arquivo` (obrigatório): Nome do arquivo
- `tamanho` (obrigatório): Tamanho em bytes
- `tipo_mime` (obrigatório): Tipo MIME (image/jpeg, image/png, application/pdf, etc.)
- `principal` (opcional): Se é a imagem principal (apenas para imagens)
- `ordem` (opcional): Ordem de exibição

#### Response

**Status: 201 Created**

```json
{
  "id": 1,
  "id_usuario": 1,
  "tipo_entidade": "usuario",
  "tipo": "Imagem",
  "url": "https://exemplo.com/foto.jpg",
  "nome_arquivo": "foto_perfil.jpg",
  "tamanho": 102400,
  "tipo_mime": "image/jpeg",
  "principal": true,
  "ordem": 0,
  "data_upload": "2024-01-15T14:30:00Z",
  "data_atualizacao": "2024-01-15T14:30:00Z",
  "usuario": {
    "id": 1,
    "nome": "João Silva",
    "email": "joao@example.com"
  }
}
```

#### Exemplo de Uso

```bash
# Upload de imagem de usuário
curl -X POST "http://localhost:8080/uploads" \
  -H "Content-Type: application/json" \
  -d '{
    "id_usuario": 1,
    "tipo_entidade": "usuario",
    "tipo": "Imagem",
    "url": "https://exemplo.com/foto.jpg",
    "nome_arquivo": "foto_perfil.jpg",
    "tamanho": 102400,
    "tipo_mime": "image/jpeg",
    "principal": true
  }'

# Upload de documento de veículo
curl -X POST "http://localhost:8080/uploads" \
  -H "Content-Type: application/json" \
  -d '{
    "id_veiculo": 5,
    "tipo_entidade": "veiculo",
    "tipo": "Documento",
    "url": "https://exemplo.com/doc.pdf",
    "nome_arquivo": "crlv.pdf",
    "tamanho": 204800,
    "tipo_mime": "application/pdf"
  }'
```

---

## Listar Todos os Uploads

### GET /uploads

Retorna todos os uploads ativos do sistema. Pode filtrar por tipo usando o query parameter `tipo`.

#### Request

**Método**: GET  
**URL**: `/uploads`  
**Query Parameters**:
- `tipo` (opcional): Filtrar por tipo - "Imagem" ou "Documento"  

#### Response

**Status: 200 OK**

```json
{
  "uploads": [
    {
      "id": 1,
      "id_usuario": 1,
      "tipo_entidade": "usuario",
      "tipo": "Imagem",
      "url": "https://exemplo.com/foto.jpg",
      "nome_arquivo": "foto_perfil.jpg",
      "tamanho": 102400,
      "tipo_mime": "image/jpeg",
      "principal": true,
      "data_upload": "2024-01-15T14:30:00Z"
    }
  ],
  "total": 1
}
```

#### Exemplos de Uso

```bash
# Listar todos os uploads
curl -X GET "http://localhost:8080/uploads"

# Listar apenas imagens
curl -X GET "http://localhost:8080/uploads?tipo=Imagem"

# Listar apenas documentos
curl -X GET "http://localhost:8080/uploads?tipo=Documento"
```

#### Erros

**Status: 400 Bad Request** - Tipo inválido
```json
{
  "error": "Tipo inválido. Use 'Imagem' ou 'Documento'"
}
```

---

## Buscar Upload por ID

### GET /uploads/{id}

Retorna um upload específico pelo ID.

#### Request

**Método**: GET  
**URL**: `/uploads/{id}`  
**Parâmetros de Path**: 
- `id` (obrigatório): ID do upload

#### Response

**Status: 200 OK**

```json
{
  "id": 1,
  "id_usuario": 1,
  "tipo_entidade": "usuario",
  "tipo": "Imagem",
  "url": "https://exemplo.com/foto.jpg",
  "nome_arquivo": "foto_perfil.jpg",
  "tamanho": 102400,
  "tipo_mime": "image/jpeg",
  "principal": true,
  "ordem": 0,
  "data_upload": "2024-01-15T14:30:00Z",
  "usuario": { ... }
}
```

---

## Uploads por Usuário

### GET /usuarios/{id_usuario}/uploads

Retorna todos os uploads de um usuário específico.

#### Request

**Método**: GET  
**URL**: `/usuarios/{id_usuario}/uploads`  
**Parâmetros de Path**: 
- `id_usuario` (obrigatório): ID do usuário

#### Response

**Status: 200 OK**

```json
{
  "uploads": [
    {
      "id": 1,
      "tipo": "Imagem",
      "url": "https://exemplo.com/foto.jpg",
      "nome_arquivo": "foto_perfil.jpg",
      "principal": true
    },
    {
      "id": 2,
      "tipo": "Documento",
      "url": "https://exemplo.com/doc.pdf",
      "nome_arquivo": "cpf.pdf"
    }
  ],
  "total": 2
}
```

---

## Uploads por Veículo

### GET /veiculos/{id}/uploads

Retorna todos os uploads de um veículo específico.

#### Request

**Método**: GET  
**URL**: `/veiculos/{id}/uploads`  
**Parâmetros de Path**: 
- `id` (obrigatório): ID do veículo

#### Response

**Status: 200 OK**

```json
{
  "uploads": [
    {
      "id": 3,
      "tipo": "Imagem",
      "url": "https://exemplo.com/carro.jpg",
      "nome_arquivo": "carro_frente.jpg",
      "principal": true
    },
    {
      "id": 4,
      "tipo": "Documento",
      "url": "https://exemplo.com/crlv.pdf",
      "nome_arquivo": "crlv.pdf"
    }
  ],
  "total": 2
}
```

---

## Uploads por Produto

### GET /produtos/{id}/uploads

Retorna todos os uploads de um produto específico.

#### Request

**Método**: GET  
**URL**: `/produtos/{id}/uploads`  
**Parâmetros de Path**: 
- `id` (obrigatório): ID do produto

---

## Uploads por Serviço

### GET /servicos/{id}/uploads

Retorna todos os uploads de um serviço específico.

#### Request

**Método**: GET  
**URL**: `/servicos/{id}/uploads`  
**Parâmetros de Path**: 
- `id` (obrigatório): ID do serviço

---

## Uploads por Loja

### GET /lojas/{id}/uploads

Retorna todos os uploads de uma loja específica.

#### Request

**Método**: GET  
**URL**: `/lojas/{id}/uploads`  
**Parâmetros de Path**: 
- `id` (obrigatório): ID da loja

---

## Upload Principal de uma Entidade

### GET /uploads/principal/{tipo}/{id}

Retorna o upload principal (imagem) de uma entidade específica.

#### Request

**Método**: GET  
**URL**: `/uploads/principal/{tipo}/{id}`  
**Parâmetros de Path**: 
- `tipo` (obrigatório): Tipo da entidade (usuario, veiculo, veiculo_loja, produto, servico, loja)
- `id` (obrigatório): ID da entidade

#### Response

**Status: 200 OK**

```json
{
  "id": 1,
  "tipo": "Imagem",
  "url": "https://exemplo.com/foto.jpg",
  "nome_arquivo": "foto_perfil.jpg",
  "principal": true
}
```

---

## Atualizar Upload

### PUT /uploads/{id}

Atualiza os dados de um upload existente.

#### Request

**Método**: PUT  
**URL**: `/uploads/{id}`  
**Body**: Mesmo formato do POST

---

## Definir Upload como Principal

### PUT /uploads/{id}/principal

Define um upload como principal da entidade. **Apenas imagens podem ser principais**.

#### Request

**Método**: PUT  
**URL**: `/uploads/{id}/principal`  

#### Response

**Status: 200 OK**

```json
{
  "message": "Upload definido como principal com sucesso"
}
```

#### Erros

**Status: 400 Bad Request** - Upload não é imagem
```json
{
  "error": "apenas imagens podem ser definidas como principais"
}
```

---

## Excluir Upload (Soft Delete)

### DELETE /uploads/{id}

Realiza soft delete de um upload.

#### Request

**Método**: DELETE  
**URL**: `/uploads/{id}`  

---

## Restaurar Upload

### POST /uploads/{id}/restore

Restaura um upload que foi soft deleted.

#### Request

**Método**: POST  
**URL**: `/uploads/{id}/restore`  

---

## Tipos de Upload

O sistema suporta dois tipos de upload:

- **Imagem**: Para fotos e imagens (pode ser definida como principal)
- **Documento**: Para documentos PDF, documentos de texto, etc. (não pode ser principal)

---

## Entidades Suportadas

Uploads podem ser associados a:

- **usuario**: Uploads de usuários (fotos de perfil, documentos)
- **veiculo**: Uploads de veículos (fotos, documentos do veículo)
- **veiculo_loja**: Uploads de veículos de loja
- **produto**: Uploads de produtos (fotos, documentos)
- **servico**: Uploads de serviços (fotos, documentos)
- **loja**: Uploads de lojas (fotos, documentos)

---

## Observações Importantes

- Apenas **imagens** podem ser definidas como principais
- Um usuário pode ter **N uploads** (imagens e documentos)
- Um veículo pode ter **N uploads** (imagens e documentos)
- Um produto pode ter **N uploads** (imagens e documentos)
- Um serviço pode ter **N uploads** (imagens e documentos)
- Uploads excluídos (soft delete) não aparecem nas consultas
- A ordenação padrão é: tipo (Imagem primeiro), principal (principal primeiro), ordem, data de upload

---

## Casos de Uso

1. **Foto de perfil do usuário**: Upload de imagem principal do usuário
2. **Fotos de veículo**: Múltiplas fotos de um veículo para venda
3. **Documentos do veículo**: CRLV, documentos de transferência, etc.
4. **Fotos de produtos**: Galeria de imagens de produtos
5. **Documentos de produtos**: Manuais, garantias, etc.
6. **Fotos de serviços**: Imagens ilustrativas de serviços oferecidos

