package routers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/RicardoRomeroMedina/pruebaTecnicaTredicom/db"
)

func VerArchiVosRouter(w http.ResponseWriter, r *http.Request) {
	respuesta, correcto := db.VerArchivos()

	if !correcto {
		http.Error(w, "Error al leer los archivos", http.StatusBadRequest)
		return
	}

	for i := 0; i < len(respuesta); i++ {
		fmt.Println("Issue #" + strconv.Itoa(respuesta[i].IssueNum))
		fmt.Println("	-URL: " + respuesta[i].IssueUrl)
		fmt.Println("	-Nombre: " + respuesta[i].IssueNom)
		fmt.Println("	-Autor: " + respuesta[i].Autor)
		fmt.Println("	-Tags: ", respuesta[i].Labels)
		fmt.Println("	-Milestone")
		fmt.Println("		-Nombre: " + respuesta[i].Milstone.Nombre)
		fmt.Println("		-Descripción: " + respuesta[i].Milstone.Descripcion)
	}

	w.Header().Set("Content-type", "application/json")

	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(respuesta)
}
