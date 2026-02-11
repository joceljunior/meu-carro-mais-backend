package json

type RegistroInteresseRequest struct {
	IDCupom  uint   `json:"id_cupom" binding:"required"`
	Nome     string `json:"nome" binding:"required,min=2,max=255"`
	Email    string `json:"email" binding:"required,email"`
	Telefone string `json:"telefone" binding:"required,min=10,max=20"`
	Mensagem string `json:"mensagem,omitempty" binding:"max=1000"`
}
