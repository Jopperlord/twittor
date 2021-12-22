package main

import (
	"log"

	"github.com/Jopperlord/twittor/bd"
	"github.com/Jopperlord/twittor/handlers"
)

func main() {

	if bd.PingConnection() == 0 {
		log.Fatalln("Sin conexion a la BD")
	}
	handlers.Manejadores()

}
