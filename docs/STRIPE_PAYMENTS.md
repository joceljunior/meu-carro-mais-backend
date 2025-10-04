# Fluxo de Pagamento com Stripe

Este documento descreve como usar o sistema de pagamentos integrado com o Stripe na API do Meu Carro Mais.

## Configuração

### Variáveis de Ambiente

Configure as seguintes variáveis de ambiente:

```bash
STRIPE_SECRET_KEY=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
STRIPE_PUBLISHABLE_KEY=pk_live_...
STRIPE_WEBHOOK_SECRET=whsec_...
STRIPE_MONTHLY_PRICE_ID=price_...
STRIPE_YEARLY_PRICE_ID=price_...
STRIPE_DOMAIN=http://localhost:8080
```

### Inicialização

A configuração do Stripe é inicializada automaticamente quando a API é iniciada através da função `InitStripe()`.

## Endpoints Disponíveis

### 1. Criar Sessão de Checkout (Pagamento Único)

**POST** `/api/v1/pagamentos/checkout`

Cria uma sessão de checkout para pagamento único (mensal ou anual).

```json
{
  "id_usuario": 1,
  "tipo_plano": "monthly",
  "success_url": "http://localhost:3000/success",
  "cancel_url": "http://localhost:3000/cancel"
}
```

**Resposta:**
```json
{
  "session_url": "https://checkout.stripe.com/c/pay/cs_...",
  "session_id": "cs_...",
  "mensagem": "Sessão de checkout criada com sucesso"
}
```

### 2. Criar Sessão de Checkout para Assinatura

**POST** `/api/v1/pagamentos/subscription-checkout`

Cria uma sessão de checkout para assinatura recorrente usando lookup_key.

```json
{
  "id_usuario": 1,
  "lookup_key": "premium_monthly",
  "success_url": "http://localhost:3000/success",
  "cancel_url": "http://localhost:3000/cancel"
}
```

**Resposta:**
```json
{
  "session_url": "https://checkout.stripe.com/c/pay/cs_...",
  "session_id": "cs_...",
  "mensagem": "Sessão de checkout para assinatura criada com sucesso"
}
```

### 3. Portal de Cobrança do Cliente

**POST** `/api/v1/pagamentos/customer-portal`

Cria uma sessão do portal de cobrança onde o cliente pode gerenciar sua assinatura.

```json
{
  "session_id": "cs_test_1234567890"
}
```

**Resposta:**
```json
{
  "portal_url": "https://billing.stripe.com/session/...",
  "mensagem": "Sessão do portal de cobrança criada com sucesso"
}
```

### 4. Webhook do Stripe

**POST** `/api/v1/pagamentos/webhook`

Processa webhooks do Stripe para atualizar status de pagamentos e assinaturas.

**Headers necessários:**
- `Stripe-Signature`: Assinatura do webhook para verificação

**Eventos suportados:**
- `checkout.session.completed`
- `checkout.session.expired`
- `customer.subscription.created`
- `customer.subscription.updated`
- `customer.subscription.deleted`
- `customer.subscription.trial_will_end`
- `payment_intent.succeeded`
- `payment_intent.payment_failed`

### 5. Histórico de Pagamentos

**GET** `/api/v1/pagamentos/historicos`
Lista todos os históricos de pagamento.

**GET** `/api/v1/pagamentos/historicos/{id}`
Busca um histórico específico por ID.

**GET** `/api/v1/usuarios/{id_usuario}/historicos-pagamento`
Lista históricos de pagamento de um usuário específico.

## Fluxo de Pagamento

### 1. Pagamento Único

1. Cliente solicita checkout através do endpoint `/checkout`
2. API cria sessão no Stripe e retorna URL de checkout
3. Cliente é redirecionado para o Stripe
4. Após pagamento, Stripe envia webhook `checkout.session.completed`
5. API atualiza status do pagamento e torna usuário premium

### 2. Assinatura Recorrente

1. Cliente solicita assinatura através do endpoint `/subscription-checkout`
2. API busca preço pelo `lookup_key` e cria sessão de assinatura
3. Cliente é redirecionado para o Stripe
4. Após criação da assinatura, Stripe envia webhook `customer.subscription.created`
5. API processa o evento e ativa funcionalidades premium

### 3. Gerenciamento de Assinatura

1. Cliente acessa portal através do endpoint `/customer-portal`
2. API cria sessão do portal usando customer ID da sessão anterior
3. Cliente é redirecionado para portal de cobrança do Stripe
4. Cliente pode cancelar, atualizar ou visualizar sua assinatura

## Modelos de Dados

### HistoricoPagamento

```go
type HistoricoPagamento struct {
    ID              uint       `gorm:"primaryKey"`
    IDUsuario       uint       `gorm:"not null"`
    StripeSessionID string     `gorm:"size:255;unique;not null"`
    StripePaymentID string     `gorm:"size:255;unique"`
    Status          string     `gorm:"size:50;not null;default:'pending'"`
    TipoPlano       string     `gorm:"size:50;not null"`
    Valor           float64    `gorm:"type:decimal(10,2);not null"`
    Moeda           string     `gorm:"size:3;not null;default:'BRL'"`
    DataPagamento   *time.Time `gorm:"null"`
    DataVencimento  *time.Time `gorm:"null"`
    DataCriacao     time.Time  `gorm:"autoCreateTime"`
    DataAtualizacao time.Time  `gorm:"autoUpdateTime"`
    DataExclusao    *time.Time `gorm:"index"`
    
    Usuario Usuario `gorm:"foreignKey:IDUsuario"`
}
```

### Status de Pagamento

- `pending`: Pagamento pendente
- `completed`: Pagamento concluído
- `failed`: Pagamento falhou
- `canceled`: Pagamento cancelado

### Tipos de Plano

- `monthly`: Plano mensal
- `yearly`: Plano anual
- `subscription`: Assinatura recorrente

## Segurança

### Verificação de Webhook

O sistema verifica a assinatura dos webhooks usando o `webhook secret` configurado:

```go
event, err := webhook.ConstructEvent(payload, signature, webhookSecret)
```

### Chaves de API

- **Secret Key**: Usada no backend para operações sensíveis
- **Publishable Key**: Usada no frontend para criar elementos de pagamento
- **Webhook Secret**: Usada para verificar a autenticidade dos webhooks

## Exemplos de Uso

Consulte o arquivo `examples/stripe_payment_examples.http` para exemplos completos de como usar todos os endpoints.

## Monitoramento

### Logs

O sistema registra logs para:
- Criação de sessões de checkout
- Processamento de webhooks
- Atualizações de status de pagamento
- Erros de processamento

### Métricas

Monitore as seguintes métricas:
- Taxa de conversão de checkout
- Taxa de sucesso de pagamentos
- Volume de assinaturas criadas/canceladas
- Tempo de processamento de webhooks

## Troubleshooting

### Problemas Comuns

1. **Webhook não processado**: Verifique se o webhook secret está correto
2. **Sessão expirada**: Sessões do Stripe expiram em 24 horas
3. **Preço não encontrado**: Verifique se o lookup_key existe no Stripe
4. **Usuário não encontrado**: Verifique se o ID do usuário é válido

### Logs de Debug

Para debug, verifique os logs da aplicação que mostram:
- Detalhes dos webhooks recebidos
- Erros de processamento
- Status de atualizações no banco de dados
