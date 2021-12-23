package bd

import (
	"github.com/Jopperlord/twittor/models"
	"golang.org/x/crypto/bcrypt"
)

/*IntentoLogin realiza el chequeo del login en la BD*/
func IntentoLogin(emal string, password string) (models.Usuario, bool) {

	usu, encontrado, _ := ChequeoYaExisteUsuario(emal)
	if !encontrado {
		return usu, false
	}

	passwordBytes := []byte(password)
	passwordBD := []byte(usu.Password)

	err := bcrypt.CompareHashAndPassword(passwordBD, passwordBytes)
	if err != nil {
		return usu, false
	}
	return usu, true
}
