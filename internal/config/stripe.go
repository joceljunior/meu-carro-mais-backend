package config

import (
	"os"
	"strconv"

	"github.com/stripe/stripe-go/v76"
	"github.com/stripe/stripe-go/v76/checkout/session"
)

type StripeConfig struct {
	SecretKey      string
	PublishableKey string
	WebhookSecret  string
	MonthlyPriceID string
	YearlyPriceID  string
}

var Stripe *StripeConfig

func InitStripe() {
	Stripe = &StripeConfig{
		SecretKey:      getEnv("STRIPE_SECRET_KEY", ""),
		PublishableKey: getEnv("STRIPE_PUBLISHABLE_KEY", ""),
		WebhookSecret:  getEnv("STRIPE_WEBHOOK_SECRET", ""),
		MonthlyPriceID: getEnv("STRIPE_MONTHLY_PRICE_ID", ""),
		YearlyPriceID:  getEnv("STRIPE_YEARLY_PRICE_ID", ""),
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

// CreateCheckoutSession cria uma sessão de checkout no Stripe
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

	return session.New(params)
}

// GetSession busca uma sessão do Stripe por ID
func (s *StripeConfig) GetSession(sessionID string) (*stripe.CheckoutSession, error) {
	return session.Get(sessionID, nil)
}
