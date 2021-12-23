package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Jopperlord/twittor/bd"
	"github.com/Jopperlord/twittor/jwt"
	"github.com/Jopperlord/twittor/models"
)

func Login(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Add("content-type", "application/json")

	var t models.Usuario

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(rw, "Usuario y/o contraseña inválidos "+err.Error(), 400)
		return
	}
	if len(t.Email) == 0 {
		http.Error(rw, "El email del usuario es requerido", 400)
		return
	}
	documento, existe := bd.IntentoLogin(t.Email, t.Password)
	if !existe {
		http.Error(rw, "Usuario y/o contraseña inválidos", 400)
		return
	}

	jwtKey, err := jwt.GeneroJWT(documento)
	if err != nil {
		http.Error(rw, "Ocurrió un error al intentar generar el Token correspondiente "+err.Error(), 400)
		return
	}

	resp := models.RespuestaLogin{
		Token: jwtKey,
	}

	rw.Header().Set("content-type", "application/json")
	rw.WriteHeader(http.StatusCreated)
	json.NewEncoder(rw).Encode(resp)

	expirationTime := time.Now().Add(24 * time.Hour)
	http.SetCookie(rw, &http.Cookie{
		Name:    "token",
		Value:   jwtKey,
		Expires: expirationTime,
	})

}
