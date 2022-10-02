package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"

	"github.com/RicardoRomeroMedina/pruebaTecnicaTredicom/middlew"
	"github.com/RicardoRomeroMedina/pruebaTecnicaTredicom/routers"
)

// manejadores setea el purto y el handler, tambien pone a escuchar al servidor
func Manejadores() {
	router := mux.NewRouter()

	//Esta es la ruta para agregar los archivos a la base de datos
	router.HandleFunc("/agregarArchivo", middlew.ChecarDB(routers.AgregarArchivo)).Methods("POST")
	//Esta es la ruta para poder ver los archivos que hay en la base de datos
	router.HandleFunc("/verArchivos", middlew.ChecarDB(routers.VerArchiVosRouter)).Methods("GET")
	//Esta es la ruta par eliminar todos los archivos de la base de datos
	router.HandleFunc("/eliminarArchivos", middlew.ChecarDB(routers.EliminarArchivosRouters)).Methods("DELETE")

	PORT := os.Getenv("PORT")

	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
