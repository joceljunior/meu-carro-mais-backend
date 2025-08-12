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
- Filtro por loja
- Filtro por faixa de preço
- Filtro por localização (proximidade)
- Ordenação por preço, destaque, data, etc.
