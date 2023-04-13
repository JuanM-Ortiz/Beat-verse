package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Mongo = ConnectDB()

var clientOptions = options.Client().ApplyURI("mongodb+srv://juanmaortiz3:kFqF1q5HL1sp4lgk@cluster0.athz34b.mongodb.net/beat-verse?retryWrites=true&w=majority")

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
