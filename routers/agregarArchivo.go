package routers

import (
	"encoding/json"
	"net/http"

	"github.com/RicardoRomeroMedina/pruebaTecnicaTredicom/db"
	"github.com/RicardoRomeroMedina/pruebaTecnicaTredicom/models"
)

func AgregarArchivo(w http.ResponseWriter, r *http.Request) {
	var a models.ModeloArchivo
	err := json.NewDecoder(r.Body).Decode(&a)

	archivo := models.ModeloArchivo{
		IssueNum: a.IssueNum,
		IssueUrl: a.IssueUrl,
		IssueNom: a.IssueNom,
		Autor:    a.Autor,
		Tags:     a.Tags,
		Milstone: a.Milstone,
	}
	if err != nil {
		http.Error(w, "No se pudo registrar el arcgivo"+err.Error(), 400)
		return
	}

	_, status, err := db.NuevoArchivo(archivo)

	if !status {
		http.Error(w, "No se pudo registrar el archivo"+err.Error(), 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("Se agrego con exito el archivo")
}
