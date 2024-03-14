package main

import (
	"context"
	"fmt"
	"time"

	database "github.com/msierpien/mongo-rex/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Page struct {
	ID               string         `json:"_id" bson:"_id"`
	Title            string         `json:"title" bson:"title"`
	Description      string         `json:"description" bson:"description"`
	ShortDescription string         `json:"short_description" bson:"short_description"`
	Slug             string         `json:"slug" bson:"slug"`
}
func main() {
	var db = database.Connect()

	pageCollection := db.GetCollection("pages")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	_id, _ := primitive.ObjectIDFromHex("65f329034a0079b8b690f1d6")
	filter := bson.M{
		"$and": []interface{}{
			bson.M{"_id": _id},
		},
	}
	var page Page
	err := pageCollection.FindOne(ctx, filter).Decode(&page)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(page)

}
