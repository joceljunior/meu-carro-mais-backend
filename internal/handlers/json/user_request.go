package json

type UserRequest struct {
	Email    string `json:"email"`
	Senha    string `json:"senha"`
	Nome     string `json:"nome"`
	CPF      string `json:"cpf"`
	Password string `json:"password"`
	Imagem   string `json:"imagem,omitempty"`
}
