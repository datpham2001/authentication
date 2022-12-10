package db

import (
	"authentication/config"
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	once   *sync.Once
	Client *mongo.Client
	dbErr  error
)

func ConnectDB() {
	once.Do(func() {
		if Client, dbErr = mongo.NewClient(options.Client().ApplyURI(config.AppConfig.MongoURI)); dbErr != nil {
			log.Fatal(dbErr)
		}

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		// connect db
		dbErr = Client.Connect(ctx)
		if dbErr != nil {
			log.Fatal(dbErr)
		}

		// ping db
		dbErr = Client.Ping(ctx, nil)
		if dbErr != nil {
			log.Fatal(dbErr)
		}

		fmt.Println("Connected to MongoDB successfully")
	})
}
