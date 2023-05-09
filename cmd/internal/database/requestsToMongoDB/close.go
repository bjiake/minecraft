package requestsToMongoDB

import (
	"golang.org/x/net/context"
	"log"
)

func CloseConnection() {
	client, err := getClient()
	if err != nil {
		log.Fatal(err)
	}

	err = client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connection closed")
}
