# Endpoint de Anúncios

## Listagem de Anúncios

### GET /anuncios

Retorna todos os anúncios disponíveis no sistema com informações completas da loja e categoria.

#### Request

**Método**: GET  
**URL**: `/anuncios`  
**Headers**: Nenhum obrigatório  
**Parâmetros**: Nenhum  

#### Response

**Status: 200 OK**

```json
{
  "anuncios": [
    {
      "id": 1,
      "titulo": "Troca de Óleo Completa",
      "descricao": "Troca de óleo com filtro incluído",
      "preco": 89.90,
      "imagem": "https://exemplo.com/troca-oleo.jpg",
      "destaque": true,
      "id_loja": 1,
      "id_categoria": 1,
      "categoria": "Manutenção Preventiva",
      "loja": {
        "id": 1,
        "nome": "Oficina São Paulo",
        "cnpj": "12.345.678/0001-90",
        "imagem": "https://exemplo.com/oficina.jpg",
        "latitude": -23.5505,
        "longitude": -46.6333,
        "id_categoria": 1,
        "categoria": "Oficina Mecânica"
      }
    },
    {
      "id": 2,
      "titulo": "Alinhamento e Balanceamento",
      "descricao": "Alinhamento e balanceamento das 4 rodas",
      "preco": 120.00,
      "imagem": "https://exemplo.com/alinhamento.jpg",
      "destaque": false,
      "id_loja": 2,
      "id_categoria": 2,
      "categoria": "Suspensão",
      "loja": {
        "id": 2,
        "nome": "Auto Center Rio",
        "cnpj": "98.765.432/0001-10",
        "imagem": "https://exemplo.com/autocenter.jpg",
        "latitude": -22.9068,
        "longitude": -43.1729,
        "id_categoria": 1,
        "categoria": "Oficina Mecânica"
      }
    }
  ],
  "total": 2
}
```

#### Erros

**Status: 500 Internal Server Error**
```json
{
  "error": "Erro interno do servidor"
}
```

#### Estrutura da Resposta

- **anuncios**: Array com todos os anúncios disponíveis
  - **id**: Identificador único do anúncio
  - **titulo**: Título do anúncio
  - **descricao**: Descrição detalhada do serviço/produto
  - **preco**: Preço em reais (float64)
  - **imagem**: URL da imagem do anúncio
  - **destaque**: Se o anúncio está em destaque
  - **id_loja**: ID da loja que oferece o serviço
  - **id_categoria**: ID da categoria do anúncio
  - **categoria**: Nome da categoria do anúncio
  - **loja**: Objeto com informações completas da loja
- **total**: Número total de anúncios

#### Observações

- Este endpoint não requer autenticação
- Retorna todos os anúncios ativos do sistema
- Os anúncios incluem informações completas da loja e categoria
- Útil para listar todos os serviços/produtos disponíveis
- Pode ser usado em conjunto com filtros por categoria ou loja

#### Exemplo de Uso

```bash
curl -X GET "http://localhost:8080/anuncios" \
  -H "Content-Type: application/json"
```

#### Casos de Uso

1. **Vitrine de serviços**: Para mostrar todos os serviços disponíveis
2. **Busca geral**: Para permitir busca em todos os anúncios
3. **Comparação de preços**: Para comparar preços entre diferentes lojas
4. **Dashboard**: Para mostrar estatísticas gerais de anúncios
5. **Feed principal**: Para exibir anúncios em destaque primeiro

#### Relacionamentos

- **Anúncio → Loja**: Cada anúncio pertence a uma loja
- **Anúncio → Categoria**: Cada anúncio tem uma categoria específica
- **Loja → Categoria**: Cada loja tem uma categoria de estabelecimento

#### Filtros Futuros

Este endpoint pode ser expandido no futuro para incluir:
- Filtro por categoria de anúncio
- Filtro por faixa de preço
- Filtro por localização (proximidade)
- Ordenação por preço, destaque, data, etc.

---

## Anúncios por Loja

### GET /anuncios/loja/{loja_id}

Retorna todos os anúncios ativos de uma loja específica, ordenados por destaque (primeiro) e data de cadastro (mais recente primeiro).

#### Request

**Método**: GET  
**URL**: `/anuncios/loja/{loja_id}`  
**Headers**: Nenhum obrigatório  
**Parâmetros de Path**: 
- `loja_id` (obrigatório): ID da loja

#### Response

**Status: 200 OK**

```json
{
  "anuncios": [
    {
      "id": 5,
      "titulo": "Troca de Óleo Completa",
      "descricao": "Troca de óleo com filtro incluído",
      "preco": 89.90,
      "imagem": "https://exemplo.com/troca-oleo.jpg",
      "destaque": true,
      "id_loja": 1,
      "id_categoria": 1,
      "categoria": "Manutenção Preventiva",
      "loja": {
        "id": 1,
        "nome": "Oficina São Paulo",
        "cnpj": "12.345.678/0001-90",
        "imagem": "https://exemplo.com/oficina.jpg",
        "latitude": -23.5505,
        "longitude": -46.6333,
        "id_categoria": 1,
        "categoria": "Oficina Mecânica"
      }
    },
    {
      "id": 8,
      "titulo": "Alinhamento e Balanceamento",
      "descricao": "Alinhamento e balanceamento das 4 rodas",
      "preco": 120.00,
      "imagem": "https://exemplo.com/alinhamento.jpg",
      "destaque": false,
      "id_loja": 1,
      "id_categoria": 2,
      "categoria": "Suspensão",
      "loja": {
        "id": 1,
        "nome": "Oficina São Paulo",
        "cnpj": "12.345.678/0001-90",
        "imagem": "https://exemplo.com/oficina.jpg",
        "latitude": -23.5505,
        "longitude": -46.6333,
        "id_categoria": 1,
        "categoria": "Oficina Mecânica"
      }
    }
  ],
  "total": 2
}
```

#### Erros

**Status: 400 Bad Request** - ID de loja inválido
```json
{
  "error": "ID de loja inválido"
}
```

**Status: 500 Internal Server Error**
```json
{
  "error": "Erro interno do servidor"
}
```

#### Observações

- Retorna apenas anúncios ativos (não excluídos)
- Anúncios em destaque aparecem primeiro
- Depois são ordenados por data de cadastro (mais recente primeiro)
- Inclui informações completas da loja e categoria
- Útil para exibir todos os anúncios de uma loja específica

#### Exemplo de Uso

```bash
curl -X GET "http://localhost:8080/anuncios/loja/1" \
  -H "Content-Type: application/json"
```

#### Casos de Uso

1. **Vitrine da loja**: Para mostrar todos os anúncios de uma loja específica
2. **Perfil da loja**: Para exibir produtos/serviços oferecidos pela loja
3. **Filtro por loja**: Para permitir que usuários vejam apenas anúncios de uma loja
4. **Dashboard da loja**: Para lojistas visualizarem seus próprios anúncios

---

## Resgatar Anúncio

### POST /anuncios/{id}/resgatar

Cria um histórico de resgate automaticamente quando um usuário resgata um anúncio. O histórico é criado com status "pendente", aguardando aprovação da loja.

#### Request

**Método**: POST  
**URL**: `/anuncios/{id}/resgatar`  
**Headers**: 
- `Content-Type: application/json`

**Parâmetros de Path**: 
- `id` (obrigatório): ID do anúncio a ser resgatado

**Body**:
```json
{
  "id_usuario": 1
}
```

#### Response

**Status: 201 Created**

```json
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
    "nome": "Oficina São Paulo",
    "cnpj": "12.345.678/0001-90"
  },
  "servico": {
    "id": 5,
    "titulo": "Troca de Óleo Completa",
    "descricao": "Troca de óleo com filtro incluído",
    "preco": 89.90
  }
}
```

#### Erros

**Status: 400 Bad Request** - Dados inválidos
```json
{
  "error": "Dados inválidos",
  "details": "Campo 'id_usuario' é obrigatório"
}
```

**Status: 404 Not Found** - Anúncio não encontrado
```json
{
  "error": "anúncio não encontrado"
}
```

**Status: 500 Internal Server Error**
```json
{
  "error": "anúncio não está disponível"
}
```

#### Observações

- O histórico de resgate é criado automaticamente com status "pendente"
- A loja precisa aprovar ou rejeitar o resgate posteriormente
- O sistema extrai automaticamente: tipo do anúncio, preço, loja e produto/serviço/veículo associado
- O anúncio deve existir e não estar excluído (soft delete)

#### Exemplo de Uso

```bash
curl -X POST "http://localhost:8080/anuncios/5/resgatar" \
  -H "Content-Type: application/json" \
  -d '{
    "id_usuario": 1
  }'
```

#### Fluxo Completo

1. Usuário resgata anúncio → `POST /anuncios/{id}/resgatar`
2. Sistema cria histórico com status "pendente"
3. Loja visualiza resgates pendentes → `GET /lojas/{id}/historicos-resgate`
4. Loja aprova → `PUT /historicos-resgate/{id}/aprovar`
5. Loja rejeita → `PUT /historicos-resgate/{id}/rejeitar`