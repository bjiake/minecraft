package requestsToMongoDB

import (
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/net/context"
	"log"
	"sync"
)

var (
	once   sync.Once
	client *mongo.Client
	err    error
)

const (
	dbName         = "minecraft"
	collectionName = "mods"
	mongoURI       = "mongodb://localhost:27017"
)

func getClient() (*mongo.Client, error) {
	once.Do(func() {
		credential := options.Credential{
			AuthMechanism: "SCRAM-SHA-1",
			Username:      "root",
		}
		clientOpts := options.Client().ApplyURI(mongoURI).SetAuth(credential)
		client, err = mongo.NewClient(clientOpts)
		if err != nil {
			log.Fatal(err)
		}

		// Create connect
		err = client.Connect(context.TODO())
		if err != nil {
			log.Fatal(err)
		}

		// Check the connection
		err = client.Ping(context.TODO(), nil)
		if err != nil {
			log.Fatal(err)
		}

		log.Println("Connected to MongoDB")
	})
	return client, err
}

func getCollection() (*mongo.Collection, error) {
	client, err := getClient()
	if err != nil {
		return nil, err
	}
	collection := client.Database(dbName).Collection(collectionName)

	return collection, nil
}
