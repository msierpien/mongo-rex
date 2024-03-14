package database

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type IndexConfig struct {
	CollectionName string
	IndexModel     mongo.IndexModel
}




func (db *DB) CreateIndexes(indexConfigs []IndexConfig) error {
    for _, config := range indexConfigs {
        _, err := db.GetCollection(config.CollectionName).Indexes().CreateOne(context.Background(), config.IndexModel)
        if err != nil {
            return err
        }
    }
    return nil
}


