package main

import (
	"fmt"
	"log"

	selectOptions "github.com/RicardoRomeroMedina/pruebaTecnicaTredicom/Ui"
	"github.com/RicardoRomeroMedina/pruebaTecnicaTredicom/db"
	"github.com/RicardoRomeroMedina/pruebaTecnicaTredicom/handlers"
)

func main() {
	if db.CheckConnection() == 0 {
		log.Fatal("Sin conexion a la base de datos")
		return
	}
	handlers.Manejadores()

	fmt.Println("Eliga la opcion que desea ejecutar")
	fmt.Println("1.- Descargar información")
	fmt.Println("2.- Ver la informacion en la base de datos")
	fmt.Println("3.- Eliminar informacion en la base de datos")
	var opt int
	fmt.Scanln(&opt)
	selectOptions.Ui(opt)
}
