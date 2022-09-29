package selectOptions

import (
	"fmt"
)

func selectOptions() {
	var options int

	fmt.Scanln(options)

	switch options {
	case 1:
		fmt.Println("Se descargaran los archivos de github")
		break

	case 2:
		fmt.Println("Leyendo los archivos de la base  de datos")
		break
	case 3:
		fmt.Println("Eliminando los archivos de la base de datos")
		break
	}
}
