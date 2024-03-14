package database

import (
	"context"
	"time"
	"go.mongodb.org/mongo-driver/mongo"
)

func (db *DB) GetCollection(name string) *mongo.Collection {

	return db.client.Collection(name)
}


func (db *DB) WithCollection(name string, action func(*mongo.Collection, context.Context) error) error {
    ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
    defer cancel()

    collection := db.client.Collection(name)
    return action(collection, ctx)
}

