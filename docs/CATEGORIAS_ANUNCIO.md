# Endpoint de Categorias de Anúncio

## Listagem de Categorias

### GET /anuncios/categorias

Retorna todas as categorias de anúncio disponíveis no sistema.

#### Request

**Método**: GET  
**URL**: `/anuncios/categorias`  
**Headers**: Nenhum obrigatório  
**Parâmetros**: Nenhum  

#### Response

**Status: 200 OK**

```json
{
  "categorias": [
    {
      "id": 1,
      "nome": "Manutenção Preventiva"
    },
    {
      "id": 2,
      "nome": "Suspensão"
    },
    {
      "id": 3,
      "nome": "Freios"
    },
    {
      "id": 4,
      "nome": "Motor"
    },
    {
      "id": 5,
      "nome": "Elétrica"
    },
    {
      "id": 6,
      "nome": "Funilaria"
    },
    {
      "id": 7,
      "nome": "Pintura"
    },
    {
      "id": 8,
      "nome": "Lavagem"
    }
  ],
  "total": 8
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

- **categorias**: Array com todas as categorias disponíveis
  - **id**: Identificador único da categoria
  - **nome**: Nome da categoria
- **total**: Número total de categorias

#### Observações

- Este endpoint não requer autenticação
- Retorna todas as categorias ativas do sistema
- As categorias são ordenadas por ID (ordem de criação)
- Útil para preencher dropdowns e formulários de filtro
- Pode ser usado em conjunto com o endpoint de anúncios

#### Exemplo de Uso

```bash
curl -X GET "http://localhost:8080/anuncios/categorias" \
  -H "Content-Type: application/json"
```

#### Casos de Uso

1. **Formulários de cadastro**: Para selecionar a categoria ao criar um anúncio
2. **Filtros de busca**: Para filtrar anúncios por categoria
3. **Navegação**: Para mostrar categorias disponíveis no menu
4. **Relatórios**: Para agrupar anúncios por categoria
5. **Dashboard**: Para mostrar distribuição de anúncios por categoria

#### Diferença das Outras Categorias

- **Categorias de Lojista**: Definem o tipo de estabelecimento (ex: Oficina Mecânica)
- **Categorias de Serviço**: Definem o tipo de serviço oferecido (ex: Troca de Óleo)
- **Categorias de Anúncio**: Definem a classificação do anúncio (ex: Manutenção Preventiva)

#### Hierarquia de Categorias

```
Categoria de Lojista (Oficina Mecânica)
├── Categoria de Serviço (Troca de Óleo)
└── Categoria de Anúncio (Manutenção Preventiva)
```

Uma loja pode oferecer múltiplos serviços de diferentes categorias, e cada serviço pode ser anunciado com uma categoria de anúncio específica.

#### Exemplos de Categorias

- **Manutenção Preventiva**: Troca de óleo, filtros, correias
- **Suspensão**: Alinhamento, balanceamento, amortecedores
- **Freios**: Pastilhas, discos, fluido de freio
- **Motor**: Correia dentada, junta do cabeçote, válvulas
- **Elétrica**: Bateria, alternador, sistema de ignição
- **Funilaria**: Reparo de amassados, troca de peças
- **Pintura**: Retoque, pintura completa, polimento
- **Lavagem**: Lavagem simples, detalhamento, cera
