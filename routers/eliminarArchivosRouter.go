package routers

import (
	"encoding/json"
	"net/http"

	"github.com/RicardoRomeroMedina/pruebaTecnicaTredicom/db"
)

// En esta funcion se llama a la funcion que elimina los archivos de la base de datos
func EliminarArchivosRouters(w http.ResponseWriter, r *http.Request) {
	err := db.EliminarArchivosDB()

	if err != nil {
		http.Error(w, "Ocurrio un error al intentar borrar los archivos "+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-type", "application/json")

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("Se eliminaron con exito los archivos")
}
