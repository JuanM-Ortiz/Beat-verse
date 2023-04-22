package routers

import (
	"encoding/json"
	"net/http"

	"github.com/JuanM-Ortiz/beat-verse/db"
	"github.com/JuanM-Ortiz/beat-verse/models"
)

func UpdateProfile(rw http.ResponseWriter, r *http.Request) {
	var t models.User

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(rw, "Datos incorrectos"+err.Error(), http.StatusBadRequest)
	}

	status, err2 := db.UpdateRecord(t, UserID)
	if err2 != nil {
		http.Error(rw, "Ocurrio un error al intentar modifical el registro. Reintente nuevamente."+err.Error(), http.StatusBadRequest)
		return
	}

	if !status {
		http.Error(rw, "No se ha logrado modificar el registro del usuario", http.StatusBadRequest)
		return
	}

	rw.WriteHeader(http.StatusCreated)
}
