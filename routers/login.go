package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/JuanM-Ortiz/beat-verse/db"
	"github.com/JuanM-Ortiz/beat-verse/jwt"
	"github.com/JuanM-Ortiz/beat-verse/models"
)

func Login(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Add("content-type", "application/json")

	var t models.User

	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		http.Error(rw, "Usuario y/o contraseña invalidos."+err.Error(), http.StatusBadRequest)
		return
	}

	if len(t.Email) == 0 {
		http.Error(rw, "El email del usuario es requerido", http.StatusBadRequest)
	}

	doc, exists := db.LoginAttempt(t.Email, t.Password)
	if !exists {
		http.Error(rw, "Usuario y/o contraseña invalidos.", http.StatusBadRequest)
		return
	}

	jwtkey, err := jwt.GenerateJwt(doc)
	if err != nil {
		http.Error(rw, "Ocurrio un error al intentar generar el Token correspondiente: "+err.Error(), http.StatusBadRequest)
	}

	resp := models.ResponseLogin{
		Token: jwtkey,
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusCreated)
	json.NewEncoder(rw).Encode(resp)

	expirationTime := time.Now().Add(24 * time.Hour)
	http.SetCookie(rw, &http.Cookie{
		Name:    "token",
		Value:   jwtkey,
		Expires: expirationTime,
	})
}
