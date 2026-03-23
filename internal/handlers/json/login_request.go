package json

type LoginRequest struct {
	Email              string `json:"email"`
	Senha              string `json:"senha"`
	IDLojaIndicadora   *uint  `json:"id_loja_indicadora,omitempty"` // ao criar conta pelo app: vincula loja + cupom destaque
} 