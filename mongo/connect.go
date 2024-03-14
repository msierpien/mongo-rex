package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)


func Connect() *DB{
	// Wczytaj konfigurację
	

	fmt.Printf("ENV: %s\n", os.Getenv("ENV"))
	cfg, err := LoadConfig(os.Getenv("ENV"))
	if err != nil {
		log.Fatal("Failed to load config", err)
	}
	mongoURL := cfg.GetConfigStringMongo()
	fmt.Println(mongoURL)
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoURL))
	if err != nil {
		log.Fatal(err)
	}
	clientDbName := client.Database(cfg.GetDatabaseName())
	// Opcjonalnie: Ustaw limit czasu dla próby połączenia
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Sprawdź połączenie
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	db := &DB{
        client: clientDbName,
    }
	// TODO Watro to przenieć do osobnego pliku indexes.go
	indexConfigs := []IndexConfig{
        {
            CollectionName: "users",
            IndexModel: mongo.IndexModel{
                Keys: bson.M{"email": 1},
                Options: options.Index().SetUnique(true),
            },

        },
		{
			CollectionName: "pages",
			IndexModel: mongo.IndexModel{
				Keys: bson.D{{Key: "project_id", Value: 1}, {Key: "slug", Value: 1}},
				Options: options.Index().SetUnique(false),
			},
		},

        // Możesz dodać więcej konfiguracji indeksów tutaj
    }

    db.CreateIndexes(indexConfigs);
	


	return db
}
