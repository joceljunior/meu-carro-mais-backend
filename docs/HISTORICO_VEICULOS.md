# Endpoint de Histórico de Veículos

## Histórico de Todos os Veículos de um Usuário

### GET /usuarios/{id_usuario}/historico

Retorna o histórico de todos os veículos de um usuário específico, ordenado por data (mais recente primeiro).

#### Request

**Método**: GET  
**URL**: `/usuarios/{id_usuario}/historico`  
**Headers**: Nenhum obrigatório  
**Parâmetros de Path**: 
- `id_usuario` (obrigatório): ID do usuário

#### Response

**Status: 200 OK**

```json
{
  "historicos": [
    {
      "id": 1,
      "id_veiculo": 1,
      "id_anuncio": 5,
      "descricao": "Troca de óleo realizada com sucesso",
      "data": "2024-01-15T14:30:00Z",
      "data_cadastro": "2024-01-15T14:35:00Z"
    },
    {
      "id": 2,
      "id_veiculo": 1,
      "id_anuncio": 8,
      "descricao": "Alinhamento e balanceamento das 4 rodas",
      "data": "2024-01-10T09:15:00Z",
      "data_cadastro": "2024-01-10T09:20:00Z"
    },
    {
      "id": 3,
      "id_veiculo": 2,
      "id_anuncio": 12,
      "descricao": "Troca de pastilhas de freio dianteiras",
      "data": "2024-01-05T16:45:00Z",
      "data_cadastro": "2024-01-05T16:50:00Z"
    }
  ],
  "total": 3
}
```

## Histórico de um Veículo Específico

### GET /veiculos/{id_veiculo}/historico

Retorna o histórico completo de um veículo específico, ordenado por data (mais recente primeiro).

#### Request

**Método**: GET  
**URL**: `/veiculos/{id_veiculo}/historico`  
**Headers**: Nenhum obrigatório  
**Parâmetros de Path**: 
- `id_veiculo` (obrigatório): ID do veículo

#### Response

**Status: 200 OK**

```json
{
  "historicos": [
    {
      "id": 1,
      "id_veiculo": 1,
      "id_anuncio": 5,
      "descricao": "Troca de óleo realizada com sucesso",
      "data": "2024-01-15T14:30:00Z",
      "data_cadastro": "2024-01-15T14:35:00Z"
    },
    {
      "id": 2,
      "id_veiculo": 1,
      "id_anuncio": 8,
      "descricao": "Alinhamento e balanceamento das 4 rodas",
      "data": "2024-01-10T09:15:00Z",
      "data_cadastro": "2024-01-10T09:20:00Z"
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

- **historicos**: Array com todos os registros de histórico
  - **id**: Identificador único do registro
  - **id_veiculo**: ID do veículo
  - **id_anuncio**: ID do anúncio relacionado
  - **descricao**: Descrição detalhada do serviço realizado
  - **data**: Data em que o serviço foi realizado
  - **data_cadastro**: Data de cadastro do registro
- **total**: Número total de registros

#### Observações

- Este endpoint não requer autenticação
- Os registros são ordenados por data (mais recente primeiro)
- Cada registro está vinculado a um anúncio específico
- Útil para acompanhar a manutenção dos veículos
- Pode ser usado para gerar relatórios de manutenção

#### Exemplo de Uso

```bash
# Histórico de todos os veículos de um usuário
curl -X GET "http://localhost:8080/usuarios/1/historico" \
  -H "Content-Type: application/json"

# Histórico de um veículo específico
curl -X GET "http://localhost:8080/veiculos/1/historico" \
  -H "Content-Type: application/json"
```

#### Casos de Uso

1. **Controle de manutenção**: Para acompanhar serviços realizados
2. **Relatórios**: Para gerar relatórios de manutenção por usuário/veículo
3. **Garantia**: Para verificar serviços realizados dentro da garantia
3. **Revenda**: Para mostrar histórico completo ao vender o veículo
4. **Seguro**: Para comprovar manutenções realizadas
5. **Dashboard**: Para mostrar estatísticas de manutenção

#### Relacionamentos

- **Histórico → Veículo**: Cada registro pertence a um veículo
- **Histórico → Anúncio**: Cada registro está vinculado a um anúncio
- **Anúncio → Loja**: O anúncio pertence a uma loja específica

#### Fluxo de Dados

```
Usuário → Veículo → Histórico → Anúncio → Loja
```

1. **Usuário** cadastra **Veículos**
2. **Veículos** recebem **Históricos** de manutenção
3. **Históricos** são vinculados a **Anúncios** de serviços
4. **Anúncios** pertencem a **Lojas** específicas

#### Validações

- **ID do usuário**: Deve ser um número inteiro válido
- **ID do veículo**: Deve ser um número inteiro válido
- **Data**: Deve ser uma data válida
- **Descrição**: Não pode estar vazia
