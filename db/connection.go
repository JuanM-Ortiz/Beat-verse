package db

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Mongo = ConnectDB()

var URL = os.Getenv("MONGO_URL")

var clientOptions = options.Client().ApplyURI(URL)

func ConnectDB() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	log.Println("Conexion exitosa con la base de datos.")
	return client
}

func CheckConnection() int {
	err := Mongo.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}
	return 1
}
