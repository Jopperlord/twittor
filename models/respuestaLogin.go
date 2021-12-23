package models

/*RespuestaLogin tiene le token que se devuelve con el login*/
type RespuestaLogin struct {
	Token string `json:"token,omitempty"`
}
