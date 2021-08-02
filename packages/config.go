package packages

import (
	"context"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	dbclient *mongo.Client
)

/*
	@created at 25 May 2021
	Method for processing Mongo URI returning Mongo Client
*/
func Mongo(url string) *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI(url))
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	err = client.Connect(ctx)
	if err != nil {
		panic(err)
	}

	defer cancel()

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		panic(err)
	}

	return client
}

func LoadMongo() {
	dbclient = Mongo(os.Getenv("MONGO_URI"))
}

func LoadDatabase() *mongo.Database {
	return dbclient.Database(os.Getenv("MONGO_DB"))
}
