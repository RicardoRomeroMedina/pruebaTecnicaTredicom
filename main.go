package main

import (
	"log"

	"github.com/RicardoRomeroMedina/pruebaTecnicaTredicom/db"
	handlers "github.com/RicardoRomeroMedina/pruebaTecnicaTredicom/handlers"
)

func main() {
	if db.CheckConnection() == 0 {
		log.Fatal("Sin conexion a la base de datos")
		return
	}
	handlers.Manejadores()

}
