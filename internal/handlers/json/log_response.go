package json

import "time"

type LogResponse struct {
	ID              uint      `json:"id"`
	IDUsuario       *uint     `json:"id_usuario,omitempty"`
	TipoAcao        string    `json:"tipo_acao"`
	Entidade        string    `json:"entidade"`
	IDEntidade      *uint     `json:"id_entidade,omitempty"`
	Descricao       string    `json:"descricao"`
	DadosAntigos    interface{} `json:"dados_antigos,omitempty"`
	DadosNovos      interface{} `json:"dados_novos,omitempty"`
	IP              string    `json:"ip,omitempty"`
	UserAgent       string    `json:"user_agent,omitempty"`
	MetodoHTTP      string    `json:"metodo_http"`
	Endpoint        string    `json:"endpoint"`
	StatusHTTP      int       `json:"status_http"`
	DataAcao        time.Time `json:"data_acao"`
	DataCadastro    time.Time `json:"data_cadastro"`
	DataAtualizacao time.Time `json:"data_atualizacao"`
	
	// Dados relacionados
	Usuario *UserResponse `json:"usuario,omitempty"`
}

type LogsResponse struct {
	Logs  []LogResponse `json:"logs"`
	Total int64         `json:"total"`
	Limit int           `json:"limit"`
	Offset int          `json:"offset"`
}

