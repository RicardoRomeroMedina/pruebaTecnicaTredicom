package main

import (
	"log"

	"github.com/RicardoRomeroMedina/pruebaTecnicaTredicom/db"
	handlers "github.com/RicardoRomeroMedina/pruebaTecnicaTredicom/handlers"
)

/*
La funcion main es la principal la que se llama para levantar el servidor y poder ver la informacion
ya que en el main se crea la conexion a la base de datos y ademas se crean las rutas de los endpoints a llamar para asi ejecutar las
acciones necesarias
*/
func main() {
	if db.CheckConnection() == 0 {
		log.Fatal("Sin conexion a la base de datos")
		return
	}
	handlers.Manejadores()

}
