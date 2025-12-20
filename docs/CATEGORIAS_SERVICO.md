# Categorias de Serviço e Produto

## Visão Geral

As categorias de serviços e produtos são gerenciadas através de um campo de texto simples (`categoria`) diretamente nas tabelas de `servicos` e `produtos`. Isso simplifica a estrutura do banco de dados e torna o sistema mais flexível.

## Estrutura

### Serviço

O modelo `Servico` possui um campo `categoria` do tipo string:

```go
type Servico struct {
    ID        uint   `gorm:"primaryKey"`
    Titulo    string `gorm:"size:255"`
    Descricao string `gorm:"size:255"`
    Preco     float64
    Imagem    string
    Destaque  bool
    Categoria string `gorm:"size:100"` // Categoria como string fixa
    IDLoja    uint
    // ... outros campos
}
```

### Produto

O modelo `Produto` também possui um campo `categoria` do tipo string:

```go
type Produto struct {
    ID        uint   `gorm:"primaryKey"`
    Nome      string `gorm:"size:255;not null"`
    Descricao string `gorm:"size:500"`
    Preco     float64
    Imagem    string
    Estoque   int
    Categoria string `gorm:"size:100"` // Categoria como string fixa
    IDLoja    uint
    // ... outros campos
}
```

## Categorias Sugeridas

### Para Serviços

- Manutenção
- Revisão
- Troca de Óleo
- Alinhamento
- Balanceamento
- Ajuste de Freios
- Suspensão
- Elétrica
- Ar Condicionado
- Funilaria

### Para Produtos

- Óleos e Lubrificantes
- Filtros
- Freios
- Pneus
- Suspensão
- Elétrica
- Iluminação
- Motor
- Acessórios
- Limpeza

## Exemplos de Uso

### Criando um Serviço

```http
POST /servicos
Content-Type: application/json

{
  "titulo": "Troca de Óleo",
  "descricao": "Troca completa de óleo do motor com filtro",
  "preco": 89.90,
  "imagem": "https://example.com/troca-oleo.jpg",
  "destaque": true,
  "categoria": "Manutenção",
  "id_loja": 1
}
```

### Criando um Produto

```http
POST /produtos
Content-Type: application/json

{
  "nome": "Óleo de Motor 5W30",
  "descricao": "Óleo sintético para motor, 1 litro",
  "preco": 45.90,
  "imagem": "https://exemplo.com/oleo-motor.jpg",
  "estoque": 50,
  "categoria": "Óleos e Lubrificantes",
  "id_loja": 1
}
```

## Response de Serviço

```json
{
  "id": 1,
  "titulo": "Troca de Óleo",
  "descricao": "Troca completa de óleo do motor com filtro",
  "preco": 89.90,
  "imagem": "https://example.com/troca-oleo.jpg",
  "destaque": true,
  "categoria": "Manutenção",
  "loja": {
    "id": 1,
    "nome": "Auto Center São Paulo",
    // ... outros campos
  }
}
```

## Response de Produto

```json
{
  "id": 1,
  "nome": "Óleo de Motor 5W30",
  "descricao": "Óleo sintético para motor, 1 litro",
  "preco": 45.90,
  "imagem": "https://exemplo.com/oleo-motor.jpg",
  "estoque": 50,
  "ativo": true,
  "categoria": "Óleos e Lubrificantes",
  "id_loja": 1,
  "data_cadastro": "2024-01-15T10:30:00Z",
  "loja": {
    "id": 1,
    "nome": "Auto Center São Paulo",
    // ... outros campos
  }
}
```

## Observações

1. O campo `categoria` é opcional e pode ser deixado vazio
2. A categoria é uma string livre, permitindo flexibilidade na nomenclatura
3. Recomenda-se manter consistência nas nomenclaturas usadas
4. Para filtrar por categoria, use consultas SQL/GORM com `WHERE categoria = ?`
