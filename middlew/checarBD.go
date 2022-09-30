package middlew

import (
	"net/http"

	"github.com/RicardoRomeroMedina/pruebaTecnicaTredicom/db"
)

func ChecarDB(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if db.CheckConnection() == 0 {
			http.Error(w, "Conexion perdida con la BD", 500)
			return
		}
		next.ServeHTTP(w, r)
	}
}
