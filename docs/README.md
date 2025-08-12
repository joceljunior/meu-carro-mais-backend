# API Meu Carro Mais

API para gerenciamento de usuários, lojas e serviços automotivos.

## Endpoints Disponíveis

### Usuários

- **POST /users** - Criação de usuário completo
  - [Documentação detalhada](USERS.md)

### Autenticação

- **POST /login** - Login/Criação de usuário

### Lojas

- **GET /lojas/proximidade** - Lista lojas por proximidade
- **GET /lojas/categorias** - Lista categorias de lojista
  - [Documentação detalhada](CATEGORIAS_LOJISTA.md)

### Serviços

- **GET /servicos/proximidade** - Lista serviços por proximidade
- **GET /servicos/categorias** - Lista categorias de serviço
  - [Documentação detalhada](CATEGORIAS_SERVICO.md)

### Anúncios

- **GET /anuncios** - Lista todos os anúncios
  - [Documentação detalhada](ANUNCIOS.md)
- **GET /anuncios/categorias** - Lista categorias de anúncio
  - [Documentação detalhada](CATEGORIAS_ANUNCIO.md)

### Veículos

- **GET /usuarios/{id_usuario}/veiculos** - Lista veículos de um usuário
  - [Documentação detalhada](VEICULOS.md)
- **GET /usuarios/{id_usuario}/historico** - Lista histórico de todos os veículos de um usuário
  - [Documentação detalhada](HISTORICO_VEICULOS.md)
- **GET /veiculos/{id_veiculo}/historico** - Lista histórico de um veículo específico
  - [Documentação detalhada](HISTORICO_VEICULOS.md)

## Documentação Swagger

A documentação interativa da API está disponível em:

- **Raiz da API** (`/`) - Swagger UI
- **/swagger** - Swagger UI alternativo

## Estrutura do Projeto

```
meu-carro-mais/
├── cmd/
│   ├── api/           # Servidor principal da API
│   ├── migrate/       # Sistema de migrations
│   └── seed/          # Sistema de seeds
├── internal/
│   ├── config/        # Configurações
│   ├── database/      # Banco de dados e models
│   ├── handlers/      # Handlers HTTP
│   └── services/      # Lógica de negócio
└── docs/              # Documentação
```

## Como Executar

### 1. Executar Migrations

```bash
go run cmd/migrate/main.go run
```

### 2. Executar Seeds (opcional)

```bash
go run cmd/seed/main.go
```

### 3. Iniciar a API

```bash
go run cmd/api/main.go
```

## Validações

### Usuário

- **Campos obrigatórios**: nome, email, senha, CPF
- **Email**: deve ser único e válido
- **CPF**: deve ser único
- **Senha**: mínimo 6 caracteres
- **Plano**: atribuído automaticamente (Gratuito)

## Tecnologias

- **Go** - Linguagem principal
- **Gin** - Framework web
- **GORM** - ORM para banco de dados
- **Swagger** - Documentação da API
- **PostgreSQL** - Banco de dados (recomendado)

## Contribuição

Para adicionar novos endpoints:

1. Crie o handler em `internal/handlers/`
2. Crie o service em `internal/services/`
3. Crie o datasource em `internal/database/datasource/`
4. Crie o router em `cmd/api/router/`
5. Atualize a documentação Swagger
6. Adicione testes se necessário
