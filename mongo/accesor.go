package database

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type DB struct {
	client *mongo.Database
}

type DataAccessor interface {
	GetCollection(name string) *mongo.Collection
	WithCollection(collectionName string, action func(*mongo.Collection, context.Context) error) error
	CreateIndexes(indexConfigs []IndexConfig) error

}


