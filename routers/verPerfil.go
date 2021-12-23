package routers

import (
	"encoding/json"
	"net/http"

	"github.com/Jopperlord/twittor/bd"
)

/*VerPerfil permite extraer los valores del perifl*/
func VerPerfil(rw http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(rw, "Debe enviar el parametro ID", http.StatusBadGateway)
		return
	}

	perfil, err := bd.BuscoPerfil(ID)
	if err != nil {
		http.Error(rw, "Ocurrio un error al intentar buscar el registro "+err.Error(), 400)
		return
	}

	rw.Header().Set("context-type", "application/json")
	rw.WriteHeader(http.StatusCreated)
	json.NewEncoder(rw).Encode(perfil)

}
