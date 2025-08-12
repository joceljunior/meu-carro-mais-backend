# Endpoint de Categorias de Lojista

## Listagem de Categorias

### GET /lojas/categorias

Retorna todas as categorias de lojista disponíveis no sistema.

#### Request

**Método**: GET  
**URL**: `/lojas/categorias`  
**Headers**: Nenhum obrigatório  
**Parâmetros**: Nenhum  

#### Response

**Status: 200 OK**

```json
{
  "categorias": [
    {
      "id": 1,
      "nome": "Oficina Mecânica"
    },
    {
      "id": 2,
      "nome": "Auto Elétrica"
    },
    {
      "id": 3,
      "nome": "Funilaria e Pintura"
    },
    {
      "id": 4,
      "nome": "Vidraçaria"
    },
    {
      "id": 5,
      "nome": "Lavagem e Detalhamento"
    }
  ],
  "total": 5
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
- Pode ser usado em conjunto com o endpoint de lojas por proximidade

#### Exemplo de Uso

```bash
curl -X GET "http://localhost:8080/lojas/categorias" \
  -H "Content-Type: application/json"
```

#### Casos de Uso

1. **Formulários de cadastro**: Para selecionar a categoria ao criar uma loja
2. **Filtros de busca**: Para filtrar lojas por categoria
3. **Navegação**: Para mostrar categorias disponíveis no menu
4. **Relatórios**: Para agrupar lojas por categoria
5. **Dashboard**: Para mostrar distribuição de lojas por categoria

#### Diferença das Categorias de Serviço

- **Categorias de Lojista**: Definem o tipo de estabelecimento (ex: Oficina Mecânica)
- **Categorias de Serviço**: Definem o tipo de serviço oferecido (ex: Troca de Óleo)

Uma loja pode oferecer múltiplos serviços de diferentes categorias, mas pertence a uma única categoria de lojista.
