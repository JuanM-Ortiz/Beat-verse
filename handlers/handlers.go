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

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}