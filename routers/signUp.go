package routers

import (
	"encoding/json"
	"net/http"

	"github.com/JuanM-Ortiz/beat-verse/db"
	"github.com/JuanM-Ortiz/beat-verse/models"
)

func SignUp(rw http.ResponseWriter, r *http.Request) {
	var t models.User
	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		http.Error(rw, "Error en los datos recibidos: "+err.Error(), http.StatusBadRequest)
		return
	}

	if len(t.Email) == 0 {
		http.Error(rw, "El email de usuario es requerido", http.StatusBadRequest)
		return
	}

	if len(t.Password) < 6 {
		http.Error(rw, "La contraseña debe tener al menos 6 caracteres.", http.StatusBadRequest)
		return
	}

	_, found, _ := db.CheckUser(t.Email)
	if found {
		http.Error(rw, "Ya existe un usuario registrado con ese Email", http.StatusBadRequest)
		return
	}

	_, status, err := db.InsertRecord(t)
	if err != nil {
		http.Error(rw, "Ha ocurrido un error al realizar el registro: "+err.Error(), http.StatusBadRequest)
		return
	}

	if !status {
		http.Error(rw, "No se ha logrado insertar el registro del usuario.", http.StatusBadRequest)
		return
	}

	rw.WriteHeader(http.StatusCreated)
}
