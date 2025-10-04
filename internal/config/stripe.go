package config

import (
	"errors"
	"os"
	"strconv"

	"github.com/stripe/stripe-go/v83"
	"github.com/stripe/stripe-go/v83/billingportal/session"
	checkoutsession "github.com/stripe/stripe-go/v83/checkout/session"
	"github.com/stripe/stripe-go/v83/customer"
	"github.com/stripe/stripe-go/v83/price"
	"github.com/stripe/stripe-go/v83/webhook"
)

type StripeConfig struct {
	SecretKey      string
	PublishableKey string
	WebhookSecret  string
	MonthlyPriceID string
	YearlyPriceID  string
	Domain         string
}

var Stripe *StripeConfig

func InitStripe() {
	Stripe = &StripeConfig{
		SecretKey:      getEnv("STRIPE_SECRET_KEY", "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"),
		PublishableKey: getEnv("STRIPE_PUBLISHABLE_KEY", ""),
		WebhookSecret:  getEnv("STRIPE_WEBHOOK_SECRET", ""),
		MonthlyPriceID: getEnv("STRIPE_MONTHLY_PRICE_ID", ""),
		YearlyPriceID:  getEnv("STRIPE_YEARLY_PRICE_ID", ""),
		Domain:         getEnv("STRIPE_DOMAIN", "http://localhost:8080"),
	}

	// Configura a chave secreta do Stripe
	stripe.Key = Stripe.SecretKey
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// CreateCheckoutSession cria uma sessão de checkout no Stripe para pagamento único
func (s *StripeConfig) CreateCheckoutSession(userID uint, userEmail, tipoPlano, successURL, cancelURL string) (*stripe.CheckoutSession, error) {
	var priceID string
	var amount int64

	// Define o preço baseado no tipo de plano
	switch tipoPlano {
	case "monthly":
		priceID = s.MonthlyPriceID
		amount = 2990 // R$ 29,90 em centavos
	case "yearly":
		priceID = s.YearlyPriceID
		amount = 29900 // R$ 299,00 em centavos
	default:
		// Fallback para preço customizado se não tiver price_id configurado
		priceID = ""
	}

	params := &stripe.CheckoutSessionParams{
		PaymentMethodTypes: stripe.StringSlice([]string{
			"card",
		}),
		LineItems: []*stripe.CheckoutSessionLineItemParams{
			{
				Price:    stripe.String(priceID),
				Quantity: stripe.Int64(1),
			},
		},
		Mode:          stripe.String(string(stripe.CheckoutSessionModePayment)),
		SuccessURL:    stripe.String(successURL),
		CancelURL:     stripe.String(cancelURL),
		CustomerEmail: stripe.String(userEmail),
		Metadata: map[string]string{
			"user_id":    strconv.FormatUint(uint64(userID), 10),
			"tipo_plano": tipoPlano,
		},
	}

	// Se não tiver price_id configurado, usa preço customizado
	if priceID == "" {
		params.LineItems = []*stripe.CheckoutSessionLineItemParams{
			{
				PriceData: &stripe.CheckoutSessionLineItemPriceDataParams{
					Currency: stripe.String("brl"),
					ProductData: &stripe.CheckoutSessionLineItemPriceDataProductDataParams{
						Name: stripe.String("Plano Premium " + tipoPlano),
					},
					UnitAmount: stripe.Int64(amount),
				},
				Quantity: stripe.Int64(1),
			},
		}
	}

	return checkoutsession.New(params)
}

// CreateSubscriptionCheckoutSession cria uma sessão de checkout para assinatura recorrente
func (s *StripeConfig) CreateSubscriptionCheckoutSession(userID uint, userEmail, lookupKey, successURL, cancelURL string) (*stripe.CheckoutSession, error) {
	// Busca o preço pelo lookup_key
	params := &stripe.PriceListParams{
		LookupKeys: stripe.StringSlice([]string{lookupKey}),
	}

	i := price.List(params)
	if !i.Next() {
		return nil, errors.New("preço não encontrado para a chave: " + lookupKey)
	}

	priceObj := i.Price()

	checkoutParams := &stripe.CheckoutSessionParams{
		Mode: stripe.String(string(stripe.CheckoutSessionModeSubscription)),
		LineItems: []*stripe.CheckoutSessionLineItemParams{
			{
				Price:    stripe.String(priceObj.ID),
				Quantity: stripe.Int64(1),
			},
		},
		SuccessURL:    stripe.String(successURL),
		CancelURL:     stripe.String(cancelURL),
		CustomerEmail: stripe.String(userEmail),
		Metadata: map[string]string{
			"user_id": strconv.FormatUint(uint64(userID), 10),
		},
	}

	return checkoutsession.New(checkoutParams)
}

// CreateCustomerPortalSession cria uma sessão do portal de cobrança do cliente
func (s *StripeConfig) CreateCustomerPortalSession(customerID, returnURL string) (*stripe.BillingPortalSession, error) {
	params := &stripe.BillingPortalSessionParams{
		Customer:  stripe.String(customerID),
		ReturnURL: stripe.String(returnURL),
	}

	return session.New(params)
}

// GetCustomer busca um cliente pelo ID
func (s *StripeConfig) GetCustomer(customerID string) (*stripe.Customer, error) {
	return customer.Get(customerID, nil)
}

// GetSession busca uma sessão do Stripe por ID
func (s *StripeConfig) GetSession(sessionID string) (*stripe.CheckoutSession, error) {
	return checkoutsession.Get(sessionID, nil)
}

// ConstructWebhookEvent constrói um evento de webhook verificando a assinatura
func (s *StripeConfig) ConstructWebhookEvent(payload []byte, signature string) (stripe.Event, error) {
	return webhook.ConstructEvent(payload, signature, s.WebhookSecret)
}
