package middlew

import (
	"net/http"

	"github.com/Jopperlord/twittor/bd"
)

func ChequeoBD(next http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		if bd.PingConnection() == 0 {
			http.Error(rw, "Conexión perdida con la BD", 500)
			return
		}
		next.ServeHTTP(rw, r)
	}
}
