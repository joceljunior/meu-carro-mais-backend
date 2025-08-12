# Endpoint de Veículos

## Listagem de Veículos por Usuário

### GET /usuarios/{id_usuario}/veiculos

Retorna todos os veículos ativos de um usuário específico.

#### Request

**Método**: GET  
**URL**: `/usuarios/{id_usuario}/veiculos`  
**Headers**: Nenhum obrigatório  
**Parâmetros de Path**: 
- `id_usuario` (obrigatório): ID do usuário

#### Response

**Status: 200 OK**

```json
{
  "veiculos": [
    {
      "id": 1,
      "modelo": "Honda Civic",
      "ano": 2020,
      "cor": "Prata",
      "placa": "ABC-1234",
      "id_usuario": 1,
      "data_cadastro": "2024-01-15T10:30:00Z",
      "ativo": true
    },
    {
      "id": 2,
      "modelo": "Toyota Corolla",
      "ano": 2019,
      "cor": "Branco",
      "placa": "XYZ-5678",
      "id_usuario": 1,
      "data_cadastro": "2024-01-20T14:15:00Z",
      "ativo": true
    }
  ],
  "total": 2
}
```

#### Erros

**Status: 400 Bad Request**
```json
{
  "error": "ID de usuário inválido"
}
```

**Status: 500 Internal Server Error**
```json
{
  "error": "Erro interno do servidor"
}
```

#### Estrutura da Resposta

- **veiculos**: Array com todos os veículos ativos do usuário
  - **id**: Identificador único do veículo
  - **modelo**: Modelo do veículo
  - **ano**: Ano de fabricação
  - **cor**: Cor do veículo
  - **placa**: Placa do veículo (única)
  - **id_usuario**: ID do usuário proprietário
  - **data_cadastro**: Data de cadastro do veículo
  - **ativo**: Se o veículo está ativo no sistema
- **total**: Número total de veículos

#### Observações

- Este endpoint não requer autenticação
- Retorna apenas veículos ativos (`ativo = true`)
- A placa é única no sistema
- Útil para mostrar a frota de um usuário
- Pode ser usado em conjunto com histórico de manutenções

#### Exemplo de Uso

```bash
curl -X GET "http://localhost:8080/usuarios/1/veiculos" \
  -H "Content-Type: application/json"
```

#### Casos de Uso

1. **Dashboard do usuário**: Para mostrar todos os veículos do usuário
2. **Seleção de veículo**: Para permitir escolher um veículo para operações
3. **Gestão de frota**: Para administrar múltiplos veículos
4. **Relatórios**: Para gerar relatórios por usuário
5. **Integração**: Para sincronizar com sistemas externos

#### Relacionamentos

- **Veículo → Usuário**: Cada veículo pertence a um usuário
- **Veículo → Histórico**: Cada veículo pode ter múltiplos registros de histórico
- **Histórico → Anúncio**: Cada registro de histórico está vinculado a um anúncio

#### Validações

- **ID do usuário**: Deve ser um número inteiro válido
- **Placa**: Deve ser única no sistema
- **Ano**: Deve ser um ano válido
- **Modelo e Cor**: Não podem estar vazios
