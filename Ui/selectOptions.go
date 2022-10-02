package selectOptions

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/RicardoRomeroMedina/pruebaTecnicaTredicom/handlers"
)

func Ui(options int) {
	handlers.Manejadores()
	switch options {
	case 1:
		fmt.Println("Se descargaran los archivos de github")

	case 2:
		response, err := http.Get(
			"http://localhost:8080/verArchivos")

		if err != nil {
			fmt.Println("Hubo un error", err)
		}

		responseData, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(responseData))
	case 3:
		fmt.Println("Eliminando los archivos de la base de datos")

	}
}
