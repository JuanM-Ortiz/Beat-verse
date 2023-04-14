package db

import (
	"context"
	"fmt"
	"time"

	"github.com/JuanM-Ortiz/beat-verse/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func SearchProfile(ID string) (models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := Mongo.Database("beat-verse")
	collection := db.Collection("users")

	var profile models.User

	objID, _ := primitive.ObjectIDFromHex(ID)

	condition := bson.M{
		"_id": objID,
	}

	err := collection.FindOne(ctx, condition).Decode(&profile)
	profile.Password = ""
	if err != nil {
		fmt.Println("No se ha encontrado el Registro. " + err.Error())
		return profile, err
	}

	return profile, nil
}
