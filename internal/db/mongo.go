package db

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/sandeep-jaiswar/core-exchange-engine/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
    clientInstance *mongo.Client
    clientOnce     sync.Once
    mongoURI       string
    databaseName   string
)

func InitMongoDB() {
    mongoURI = config.GetMongoURI()
    databaseName = config.GetMongoDatabase()
}

func GetClient() (*mongo.Client, error) {
    var err error
    clientOnce.Do(func() {
        clientOptions := options.Client().ApplyURI(mongoURI)
        ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
        defer cancel()

        client, connectErr := mongo.Connect(ctx, clientOptions)
        if connectErr != nil {
            err = connectErr
            return
        }

        // Ping the database to verify the connection
        pingErr := client.Ping(ctx, nil)
        if pingErr != nil {
            err = pingErr
            return
        }

        clientInstance = client
        fmt.Println("Connected to MongoDB!")
    })

    return clientInstance, err
}

func GetCollection(collectionName string) (*mongo.Collection, error) {
    client, err := GetClient()
    if err != nil {
        return nil, err
    }
    return client.Database(databaseName).Collection(collectionName), nil
}