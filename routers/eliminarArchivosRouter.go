package routers

import (
	"net/http"

	"github.com/RicardoRomeroMedina/pruebaTecnicaTredicom/db"
)

func EliminarArchivosRouters(w http.ResponseWriter, r *http.Request) {
	err := db.EliminarArchivosDB()

	if err != nil {
		http.Error(w, "Ocurrio un error al intentar borrar el tweet "+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-type", "application/json")

	w.WriteHeader(http.StatusCreated)
}
