package middlew

import (
	"net/http"

	"github.com/JuanM-Ortiz/beat-verse/db"
)

func CheckDB(next http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		if db.CheckConnection() == 0 {
			http.Error(rw, "Se ha perdido la conexion con la base de datos.", http.StatusInternalServerError)
			return
		}

		next.ServeHTTP(rw, r)
	}
}
