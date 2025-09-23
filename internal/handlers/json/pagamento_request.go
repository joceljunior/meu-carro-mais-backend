package json

type CheckoutRequest struct {
	IDUsuario  uint   `json:"id_usuario" binding:"required"`
	TipoPlano  string `json:"tipo_plano" binding:"required,oneof=monthly yearly"`
	SuccessURL string `json:"success_url" binding:"required"`
	CancelURL  string `json:"cancel_url" binding:"required"`
}

type WebhookRequest struct {
	ID              string              `json:"id"`
	Object          string              `json:"object"`
	Type            string              `json:"type"`
	Data            WebhookData         `json:"data"`
	Created         int64               `json:"created"`
	Livemode        bool                `json:"livemode"`
	PendingWebhooks int                 `json:"pending_webhooks"`
	Request         *WebhookRequestData `json:"request,omitempty"`
}

type WebhookData struct {
	Object WebhookObject `json:"object"`
}

type WebhookObject struct {
	ID            string                 `json:"id"`
	Object        string                 `json:"object"`
	AmountTotal   int64                  `json:"amount_total"`
	Currency      string                 `json:"currency"`
	Customer      string                 `json:"customer"`
	CustomerEmail string                 `json:"customer_email"`
	PaymentStatus string                 `json:"payment_status"`
	Status        string                 `json:"status"`
	URL           string                 `json:"url"`
	Metadata      map[string]interface{} `json:"metadata"`
}

type WebhookRequestData struct {
	ID             string `json:"id"`
	IdempotencyKey string `json:"idempotency_key"`
}
