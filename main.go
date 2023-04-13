package main

import (
	"log"

	"github.com/JuanM-Ortiz/beat-verse/db"
	"github.com/JuanM-Ortiz/beat-verse/handlers"
)

func main() {
	if db.CheckConnection() == 0 {
		log.Fatal("La base de datos no esta conectada")
		return
	}
	handlers.Handlers()
}
