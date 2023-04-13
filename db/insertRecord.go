package db

import (
	"context"
	"fmt"
	"time"

	"github.com/JuanM-Ortiz/beat-verse/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func InsertRecord(u models.User) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := Mongo.Database("beat-verse")
	collection := db.Collection("users")

	u.Password, _ = EncryptPassword(u.Password)

	result, err := collection.InsertOne(ctx, u)
	if err != nil {
		return "", false, err
	}

	insertedID, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return "", false, fmt.Errorf("no se pudo convertir el InsertedID a primitive.ObjectID")
	}

	return insertedID.String(), true, nil
}
