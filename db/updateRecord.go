package db

import (
	"context"
	"time"

	"github.com/JuanM-Ortiz/beat-verse/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func UpdateRecord(u models.User, ID string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := Mongo.Database("beat-verse")
	collection := db.Collection("users")

	record := make(map[string]interface{})
	if len(u.Name) > 0 {
		record["name"] = u.Name
	}

	if len(u.LastName) > 0 {
		record["lastname"] = u.LastName
	}

	record["dateofbirth"] = u.DateOfBirth

	if len(u.Avatar) > 0 {
		record["avatar"] = u.Avatar
	}

	if len(u.Banner) > 0 {
		record["banner"] = u.Banner
	}

	if len(u.Biography) > 0 {
		record["biography"] = u.Biography
	}

	if len(u.Location) > 0 {
		record["location"] = u.Location
	}

	if len(u.WebSite) > 0 {
		record["website"] = u.WebSite
	}

	updateString := bson.M{
		"$set": record,
	}

	objId, _ := primitive.ObjectIDFromHex(ID)
	filter := bson.M{"_id": bson.M{"$eq": objId}}

	_, err := collection.UpdateOne(ctx, filter, updateString)
	if err != nil {
		return false, err
	}

	return true, nil
}
