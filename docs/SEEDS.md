# Sistema de Seeds - Meu Carro Mais

Este documento explica como usar o sistema de seeds para popular as tabelas com dados de exemplo.

## 🌱 O que são Seeds?

Seeds são dados de exemplo que são inseridos nas tabelas do banco de dados para facilitar o desenvolvimento e testes. Eles permitem que você tenha dados realistas para trabalhar sem precisar criar manualmente cada registro.

## 📁 Estrutura dos Seeds

```
internal/database/seeds/
└── seeder.go          # Lógica principal dos seeds

cmd/seed/
└── main.go            # Comando CLI para executar seeds

scripts/
└── seed.sh            # Script shell para facilitar execução
```

## 🚀 Como Usar

### 1. Executar Todos os Seeds

```bash
# Usando o comando Go
go run cmd/seed/main.go run

# Usando o script shell
./scripts/seed.sh run
```

### 2. Ver Ajuda

```bash
# Usando o comando Go
go run cmd/seed/main.go help

# Usando o script shell
./scripts/seed.sh help
```

### 3. Ver Status dos Seeds

```bash
# Usando o script shell
./scripts/seed.sh status
```

## 📊 Dados Inseridos

### Tipos de Plano
- **Gratuito**: Plano básico sem custo
- **Básico**: Plano com funcionalidades básicas
- **Premium**: Plano com funcionalidades avançadas
- **Enterprise**: Plano empresarial

### Categorias de Lojista
- **Concessionária**: Lojas de veículos novos
- **Loja de Peças**: Lojas de peças automotivas
- **Oficina Mecânica**: Oficinas de reparo
- **Lava Jato**: Serviços de limpeza
- **Seguros**: Corretoras de seguros
- **Financiamento**: Instituições financeiras

### Categorias de Anúncio
- **Carros**: Veículos de passeio
- **Motos**: Motocicletas
- **Caminhões**: Veículos comerciais
- **Peças**: Peças automotivas
- **Serviços**: Serviços automotivos
- **Acessórios**: Acessórios para veículos

### Usuários de Exemplo
1. **João Silva**
   - Email: `joao@email.com`
   - Senha: `senha123`
   - Telefone: `(11) 99999-9999`
   - Endereço: `Rua das Flores, 123 - São Paulo/SP`

2. **Maria Santos**
   - Email: `maria@email.com`
   - Senha: `senha123`
   - Telefone: `(11) 88888-8888`
   - Endereço: `Av. Paulista, 1000 - São Paulo/SP`

3. **Pedro Costa**
   - Email: `pedro@email.com`
   - Senha: `senha123`
   - Telefone: `(11) 77777-7777`
   - Endereço: `Rua Augusta, 500 - São Paulo/SP`

4. **Carlos Porto Alegre**
   - Email: `carlos@email.com`
   - Senha: `senha123`
   - Telefone: `(51) 99999-9999`
   - Endereço: `Av. Borges de Medeiros, 500 - Porto Alegre/RS`
   - Latitude: `-29.992596`
   - Longitude: `-51.1592592`

### Lojas de Exemplo
1. **Auto Center São Paulo**
   - CNPJ: `12.345.678/0001-90`
   - Categoria: Concessionária
   - Localização: `-23.5505, -46.6333`

2. **Oficina do João**
   - CNPJ: `98.765.432/0001-10`
   - Categoria: Oficina Mecânica
   - Localização: `-23.5489, -46.6388`

3. **Carros Premium**
   - CNPJ: `11.222.333/0001-44`
   - Categoria: Concessionária
   - Localização: `-23.5520, -46.6310`

### Anúncios de Exemplo
1. **Honda Civic 2020**
   - Preço: R$ 85.000,00
   - Descrição: Honda Civic EXL 2.0, automático, completo, único dono
   - Destaque: Sim

2. **Toyota Corolla 2019**
   - Preço: R$ 75.000,00
   - Descrição: Toyota Corolla XEi 2.0, automático, revisões em dia
   - Destaque: Não

3. **Volkswagen Golf GTI**
   - Preço: R$ 95.000,00
   - Descrição: Golf GTI 2.0 TSI, manual, esportivo, baixa quilometragem
   - Destaque: Sim

### Carteiras de Exemplo
- **João Silva**: R$ 1.000,00

## 🔄 Comportamento dos Seeds

### Verificação de Duplicatas
Os seeds verificam se os dados já existem antes de inserir:

- **Se o dado não existe**: Cria um novo registro
- **Se o dado já existe**: Pula a inserção (não duplica)

### Ordem de Execução
Os seeds são executados em ordem de dependência:

1. `tipo_plano` (sem dependências)
2. `categoria_lojista` (sem dependências)
3. `categoria_anuncio` (sem dependências)
4. `usuario` (depende de `tipo_plano`)
5. `loja` (depende de `categoria_lojista`)
6. `anuncio` (depende de `loja` e `categoria_anuncio`)
7. `carteira` (depende de `usuario`)

## 🛠️ Personalizando Seeds

### Adicionar Novos Dados

Para adicionar novos dados, edite o arquivo `internal/database/seeds/seeder.go`:

```go
// Exemplo: Adicionar novo usuário
usuarios := []models.Usuario{
    // ... usuários existentes ...
    {
        Nome:     "Ana Oliveira",
        Email:    "ana@email.com",
        Senha:    "senha123",
        CPF:      "55566677788",
        Imagem:   "https://via.placeholder.com/150",
        Telefone: "(11) 66666-6666",
        Endereco: "Rua das Palmeiras, 456 - São Paulo/SP",
        Ativo:    true,
        IDPlano:  planoBásico.ID,
    },
}
```

### Criar Novo Seed

Para criar um novo seed:

1. Adicione uma nova função no `seeder.go`:

```go
// seedNovaTabela popula a tabela nova_tabela
func (s *Seeder) seedNovaTabela() error {
    log.Println("📝 Populando tabela nova_tabela...")
    
    dados := []models.NovaTabela{
        // ... dados ...
    }
    
    for _, dado := range dados {
        var existing models.NovaTabela
        if err := s.db.Where("campo = ?", dado.Campo).First(&existing).Error; err != nil {
            if err := s.db.Create(&dado).Error; err != nil {
                return fmt.Errorf("erro ao criar dado: %v", err)
            }
            log.Printf("✅ Dado criado: %s", dado.Campo)
        } else {
            log.Printf("⏭️ Dado já existe: %s", dado.Campo)
        }
    }
    
    return nil
}
```

2. Adicione a chamada no método `Run()`:

```go
func (s *Seeder) Run() error {
    // ... seeds existentes ...
    
    if err := s.seedNovaTabela(); err != nil {
        return fmt.Errorf("erro ao executar seed nova_tabela: %v", err)
    }
    
    return nil
}
```

## 🚨 Considerações Importantes

### Ambiente de Desenvolvimento
- Os seeds são ideais para ambiente de desenvolvimento
- **Nunca execute seeds em produção** sem revisão cuidadosa
- Sempre faça backup antes de executar seeds

### Dados Sensíveis
- Os seeds incluem dados de exemplo (emails, CPFs, etc.)
- **Nunca use dados reais** nos seeds
- Use dados fictícios para testes

### Performance
- Os seeds verificam duplicatas antes de inserir
- Isso pode tornar a execução mais lenta em tabelas grandes
- Para grandes volumes, considere usar `INSERT IGNORE` ou similar

## 📝 Logs

Os seeds geram logs detalhados:

```
🌱 Iniciando seeds...
📝 Populando tabela tipo_plano...
✅ TipoPlano criado: Gratuito
✅ TipoPlano criado: Básico
⏭️ TipoPlano já existe: Premium
📝 Populando tabela categoria_lojista...
✅ CategoriaLojista criada: Concessionária
...
✅ Seeds executados com sucesso!
```

## 🔧 Troubleshooting

### Erro de Conexão
```
Erro ao conectar no banco: connection refused
```
- Verifique se o banco está rodando
- Verifique as credenciais de conexão

### Erro de Dependência
```
Erro ao buscar plano básico: record not found
```
- Execute as migrations primeiro: `go run cmd/migrate/main.go run`
- Verifique se os dados de dependência existem

### Erro de Duplicata
```
Erro ao criar usuario: duplicate key value violates unique constraint
```
- Os seeds já verificam duplicatas automaticamente
- Se ocorrer, pode haver dados inconsistentes no banco

## 📚 Exemplos Práticos

### Executar Seeds em Desenvolvimento
```bash
# 1. Execute as migrations
go run cmd/migrate/main.go run

# 2. Execute os seeds
go run cmd/seed/main.go run

# 3. Verifique os dados
# Use sua aplicação ou ferramentas de banco para verificar
```

### Executar Seeds com Script
```bash
# 1. Torne o script executável (Linux/Mac)
chmod +x scripts/seed.sh

# 2. Execute os seeds
./scripts/seed.sh run

# 3. Ver status
./scripts/seed.sh status
```

### Verificar Dados Inseridos
```sql
-- Verificar usuários
SELECT id, nome, email, telefone FROM usuarios;

-- Verificar lojas
SELECT id, nome, cnpj, latitude, longitude FROM lojas;

-- Verificar anúncios
SELECT id, titulo, preco, destaque FROM anuncios;
``` 