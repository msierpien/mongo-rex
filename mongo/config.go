package database

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
    DatabaseHost    string
    DatabaseName    string
    DatabaseUser    string
	DatabasePass    string
	DatabasePort    string
}

func LoadConfig(env string) (*Config, error) {
    
    if env == "dev" {
        if err := godotenv.Load(".env"); err != nil {
            return nil, err
        }
    }
    fmt.Println("MONGO_HOST", os.Getenv("MONGO_HOST"))
    return &Config{
        DatabaseHost: os.Getenv("MONGO_HOST"),
        DatabaseName: os.Getenv("MONGO_DBNAME"),
        DatabaseUser: os.Getenv("MONGO_USERNAME"),
        DatabasePass: os.Getenv("MONGO_PASSWORD"),
        DatabasePort: os.Getenv("MONGO_PORT"),
    }, nil
}
func (c *Config) GetConfigStringMongo() string {
    return fmt.Sprintf("mongodb://%s:%s@%s:%s",
        c.DatabaseUser,
        c.DatabasePass,
        c.DatabaseHost,
        c.DatabasePort,
    )
}
func (n *Config) GetDatabaseName() string {
    return n.DatabaseName
}