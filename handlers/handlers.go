package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/JuanM-Ortiz/beat-verse/middlew"
	"github.com/JuanM-Ortiz/beat-verse/routers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func Handlers() {
	router := mux.NewRouter()

	router.HandleFunc("/signUp", middlew.CheckDB(routers.SignUp)).Methods("POST")
	router.HandleFunc("/login", middlew.CheckDB(routers.Login)).Methods(("POST"))
	router.HandleFunc("/viewProfile", middlew.CheckDB(routers.ViewProfile)).Methods("GET")
	router.HandleFunc("/updateProfile", middlew.CheckDB(middlew.ValidateJwt(routers.UpdateProfile))).Methods("PUT")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
