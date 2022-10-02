package routers

import (
	"encoding/json"
	// "fmt"
	"net/http"

	"github.com/RicardoRomeroMedina/pruebaTecnicaTredicom/api"
	"github.com/RicardoRomeroMedina/pruebaTecnicaTredicom/db"
	"github.com/RicardoRomeroMedina/pruebaTecnicaTredicom/models"
)

func AgregarArchivo(w http.ResponseWriter, r *http.Request) {
	var docsMany []interface{}

	results := api.ConectarAPIGithub()

	for _, data := range results {
		archivo := models.ModeloArchivo{
			IssueNum: data.Number,
			IssueNom: data.Nombre,
			IssueUrl: data.HtmlUrl,
			Autor:    data.User.Login,
			Labels:   data.Labels,
			Milstone: struct {
				Nombre      string "bson:\"nombre\" json:\"nombre,omitempty\""
				Descripcion string "bson:\"descripcion\" json:\"descripcion,omitempty\""
			}(data.Milestone),
		}
		docsMany = append(docsMany, archivo)
	}

	_, status, err := db.NuevoArchivo(docsMany)

	if !status {
		http.Error(w, "No se pudo registrar el archivo"+err.Error(), 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("Se agrego con exito el archivo")
}
