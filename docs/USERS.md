# Endpoint de Usuários

## Criação de Usuário Completo

### POST /users

Cria um novo usuário com todos os dados fornecidos.

#### Request Body

```json
{
  "nome": "João Silva",
  "email": "joao@email.com",
  "senha": "123456",
  "cpf": "123.456.789-00",
  "imagem": "https://exemplo.com/foto.jpg",
  "telefone": "(11) 99999-9999",
  "endereco": "Rua das Flores, 123",
  "data_nascimento": "1990-01-01T00:00:00Z",
  "latitude": -23.5505,
  "longitude": -46.6333
}
```

#### Campos Obrigatórios

- `nome`: Nome completo do usuário
- `email`: Email válido do usuário (deve ser único)
- `senha`: Senha com mínimo de 6 caracteres
- `cpf`: CPF do usuário (deve ser único)

#### Campos Opcionais

- `imagem`: URL da imagem do usuário
- `telefone`: Telefone do usuário
- `endereco`: Endereço do usuário
- `data_nascimento`: Data de nascimento em formato ISO 8601
- `latitude`: Latitude da localização do usuário
- `longitude`: Longitude da localização do usuário

#### Response

**Status: 201 Created**

```json
{
  "id": 1,
  "nome": "João Silva",
  "email": "joao@email.com",
  "cpf": "123.456.789-00",
  "imagem": "https://exemplo.com/foto.jpg",
  "telefone": "(11) 99999-9999",
  "endereco": "Rua das Flores, 123",
  "data_nascimento": "1990-01-01T00:00:00Z",
  "data_cadastro": "2024-01-15T10:30:00Z",
  "ativo": true,
  "latitude": -23.5505,
  "longitude": -46.6333,
  "id_plano": 1,
  "id_loja": null,
  "mensagem": "Usuário criado com sucesso"
}
```

#### Erros

**Status: 400 Bad Request**
```json
{
  "error": "Dados inválidos",
  "details": "Campo 'email' é obrigatório"
}
```

**Status: 500 Internal Server Error**
```json
{
  "error": "Erro interno do servidor"
}
```

#### Validações

- Email deve ser único no sistema
- CPF deve ser único no sistema
- Senha deve ter pelo menos 6 caracteres
- Email deve estar em formato válido
- Campos obrigatórios não podem estar vazios

#### Observações

- O usuário é criado automaticamente com plano padrão (ID: 1 - Gratuito)
- O usuário é criado como ativo por padrão
- A data de cadastro é preenchida automaticamente
- O ID da loja é inicialmente nulo (pode ser associado posteriormente)
