# Deploy no Railway

## Configuração de Variáveis de Ambiente

### 1. Variáveis Obrigatórias

No painel do Railway, configure as seguintes variáveis:

```env
# Swagger Configuration
SWAGGER_HOST=meu-carro-mais-production.up.railway.app

# Server Configuration
PORT=8080
GIN_MODE=release

# Stripe Configuration
STRIPE_SECRET_KEY=sk_live_...
STRIPE_PUBLISHABLE_KEY=pk_live_...
STRIPE_WEBHOOK_SECRET=whsec_...

# URLs de Redirecionamento
SUCCESS_URL=https://seu-frontend.com/success
CANCEL_URL=https://seu-frontend.com/cancel
```

### 2. Como Adicionar Variáveis no Railway

1. Acesse seu projeto no Railway
2. Clique na aba **"Variables"**
3. Adicione cada variável listada acima
4. Clique em **"Deploy"** para aplicar as mudanças

### 3. Verificar Deployment

Após o deploy, acesse:

- **API**: https://meu-carro-mais-production.up.railway.app/
- **Swagger UI**: https://meu-carro-mais-production.up.railway.app/swagger/index.html

### 4. Configurar Webhooks do Stripe

No painel do Stripe:

1. Vá em **Developers** → **Webhooks**
2. Adicione um novo endpoint: `https://meu-carro-mais-production.up.railway.app/pagamentos/webhook`
3. Selecione os eventos:
   - `checkout.session.completed`
   - `customer.subscription.created`
   - `customer.subscription.updated`
   - `customer.subscription.deleted`
4. Copie o **Signing Secret** e adicione como `STRIPE_WEBHOOK_SECRET` no Railway

### 5. Executar Migrations em Produção

As migrations são executadas automaticamente quando a API inicia. Se precisar executar manualmente:

```bash
# Via Railway CLI
railway run go run cmd/migrate/main.go run

# Verificar status
railway run go run cmd/migrate/main.go status
```

### 6. Monitoramento

- **Logs**: Visualize no painel do Railway em tempo real
- **Métricas**: CPU, memória e uso de rede disponíveis no dashboard
- **Health Check**: Configure um endpoint `/health` se necessário

### 7. Troubleshooting

**Problema: Swagger não aparece**
- Verifique se a variável `SWAGGER_HOST` está configurada corretamente
- Certifique-se que não tem `http://` ou `https://` no valor

**Problema: Erros de banco de dados**
- Verifique se as migrations foram executadas
- Confirme a connection string do PostgreSQL

**Problema: Webhooks do Stripe não funcionam**
- Verifique o `STRIPE_WEBHOOK_SECRET`
- Certifique-se que a URL do webhook está correta no Stripe
- Teste com a Stripe CLI: `stripe listen --forward-to https://meu-carro-mais-production.up.railway.app/pagamentos/webhook`

### 8. Atualizações

Para fazer deploy de novas mudanças:

1. Faça push para a branch `main` no GitHub
2. O Railway fará deploy automaticamente
3. Aguarde a conclusão (geralmente 2-3 minutos)
4. Verifique os logs para confirmar sucesso

### 9. Rollback

Se algo der errado:

1. Vá em **Deployments** no Railway
2. Encontre o último deployment estável
3. Clique em **"Redeploy"**

## Comandos Úteis

```bash
# Instalar Railway CLI
npm install -g @railway/cli

# Login
railway login

# Linkar ao projeto
railway link

# Ver logs em tempo real
railway logs

# Executar comandos no ambiente Railway
railway run <comando>

# Ver variáveis de ambiente
railway variables
```

## Custos

Railway oferece:
- **$5 de crédito gratuito por mês** para o hobby plan
- Cobrança por uso além do crédito gratuito
- Monitoramento de custos no dashboard

## Segurança

✅ **Recomendações:**
- Nunca commite secrets no código
- Use variáveis de ambiente para configurações sensíveis
- Configure CORS adequadamente
- Use HTTPS sempre que possível
- Mantenha dependências atualizadas

## Suporte

- **Railway Docs**: https://docs.railway.app/
- **Discord Railway**: https://discord.gg/railway
- **Stripe Docs**: https://stripe.com/docs

