package routers

import (
	"encoding/json"
	"net/http"

	"github.com/RicardoRomeroMedina/pruebaTecnicaTredicom/api"
	"github.com/RicardoRomeroMedina/pruebaTecnicaTredicom/db"
	"github.com/RicardoRomeroMedina/pruebaTecnicaTredicom/models"
)

// Esta funcion sirve para llamar a la api de github y para agregar la informacion necesaria a la base de datos
func AgregarArchivo(w http.ResponseWriter, r *http.Request) {
	var docsMany []interface{}

	//Esta es la llamada a la api de github
	results := api.ConectarAPIGithub()

	/*
		Este ciclo sirve para poder darle el formato correcto a la informacion en la variable results y para agregar esa informacion
		ordenada en la variable docsMany que es una lista de interface ya que es el formato de la informacion que se necesita para poder
		en la funcion de db.NuevoArchivo ya que el InsertMany de mongo es el formato que requiere para agregar a la base de datos la
		informacio
	*/
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
	//En esta linea se llama a la funcion de agregar a la base de datos
	_, status, err := db.NuevoArchivo(docsMany)

	if !status {
		http.Error(w, "No se pudo registrar el archivo"+err.Error(), 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("Se agrego con exito el archivo")
}
