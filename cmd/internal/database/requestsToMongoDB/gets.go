package requestsToMongoDB

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/net/context"
	"log"
	"minecraft/cmd/internal/config"
	"minecraft/cmd/internal/models"
)

func GetLastID() {
	collection, err := getCollection()
	if err != nil {
		log.Fatal(err)
	}

	findOptions := options.FindOne()
	findOptions.SetSort(bson.M{"id": -1})
	// выбираем первый документ
	var result models.Mod
	err = collection.FindOne(context.Background(), bson.M{}, findOptions).Decode(&result)
	if err != nil {
		log.Printf("Cannot find lastID from dataBase, LastID = 0")
		config.LastID = 0
		return
	}

	// извлекаем значение поля _id
	lastId := result.ID
	log.Printf("getLastID:", lastId)
	config.LastID = lastId
}
