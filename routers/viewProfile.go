package routers

import (
	"encoding/json"
	"net/http"

	"github.com/JuanM-Ortiz/beat-verse/db"
)

func ViewProfile(rw http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(rw, "Debe enviar el parametro ID", http.StatusBadRequest)
	}

	profile, err := db.SearchProfile(ID)
	if err != nil {
		http.Error(rw, "Error al buscar el registro."+err.Error(), http.StatusBadRequest)
		return
	}

	rw.Header().Set("context-type", "application/json")
	rw.WriteHeader(http.StatusOK)
	json.NewEncoder(rw).Encode(profile)
}
