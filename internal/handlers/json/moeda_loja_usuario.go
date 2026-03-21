package json

// MoedaLojaUsuarioItem saldo de moedas restritas a uma loja.
type MoedaLojaUsuarioItem struct {
	IDLoja   uint   `json:"id_loja"`
	NomeLoja string `json:"nome_loja,omitempty"`
	Saldo    int    `json:"saldo"`
}
