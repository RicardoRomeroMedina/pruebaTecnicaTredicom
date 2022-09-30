package selectOptions

import (
	"fmt"
)

func Ui(options int) {
	switch options {
	case 1:
		fmt.Println("Se descargaran los archivos de github")

	case 2:
		fmt.Println("Leyendo los archivos de la base  de datos")

	case 3:
		fmt.Println("Eliminando los archivos de la base de datos")

	}
}
