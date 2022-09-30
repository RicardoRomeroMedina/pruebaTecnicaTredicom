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

	router.HandleFunc("/agregarArchivo", middlew.ChecarDB(routers.AgregarArchivo)).Methods("POST")
	// router.HandleFunc("/verInventario", middlew.ChecarDB(routers.VerInventarioRouter)).Methods("GET")
	// router.HandleFunc("/agregarVenta", middlew.ChecarDB(routers.AgregarVenta)).Methods("POST")
	// router.HandleFunc("/agregarCredito", middlew.ChecarDB(routers.AgregarCredito)).Methods("POST")

	PORT := os.Getenv("PORT")

	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
