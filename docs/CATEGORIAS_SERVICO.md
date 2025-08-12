# Endpoint de Categorias de Serviço

## Listagem de Categorias

### GET /servicos/categorias

Retorna todas as categorias de serviço disponíveis no sistema.

#### Request

**Método**: GET  
**URL**: `/servicos/categorias`  
**Headers**: Nenhum obrigatório  
**Parâmetros**: Nenhum  

#### Response

**Status: 200 OK**

```json
{
  "categorias": [
    {
      "id": 1,
      "nome": "Troca de Óleo"
    },
    {
      "id": 2,
      "nome": "Alinhamento"
    },
    {
      "id": 3,
      "nome": "Balanceamento"
    },
    {
      "id": 4,
      "nome": "Freios"
    },
    {
      "id": 5,
      "nome": "Suspensão"
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
- Pode ser usado em conjunto com o endpoint de serviços por proximidade

#### Exemplo de Uso

```bash
curl -X GET "http://localhost:8080/servicos/categorias" \
  -H "Content-Type: application/json"
```

#### Casos de Uso

1. **Formulários de cadastro**: Para selecionar a categoria ao criar um serviço
2. **Filtros de busca**: Para filtrar serviços por categoria
3. **Navegação**: Para mostrar categorias disponíveis no menu
4. **Relatórios**: Para agrupar serviços por categoria
