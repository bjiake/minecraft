package requestsToMongoDB

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/net/context"
	"log"
	"minecraft/cmd/internal/models"
)

func InsertOne(mod models.Mod) {
	collection, err := getCollection()
	if err != nil {
		log.Fatal(err)
	}

	var result bson.M
	err = collection.FindOne(context.TODO(), bson.D{{"title", mod.Title}}).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			insertResult, err := collection.InsertOne(context.Background(), mod)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("Inserted a single document: ", insertResult.InsertedID)
			log.Println(mod)
			return
		}
		log.Fatal(err)
	}
	log.Println("\n\nAlready inserted document: \n", result, "\n\n")
}

func InsertMany(modList []models.Mod) {
	//Создание документов для БД из ModList
	var mods []interface{}
	for _, mod := range modList {
		mods = append(mods, mod)
	}

	for _, mod := range modList {
		InsertOne(mod)
	}

	log.Println("Inserted multiple documents: ", modList)
	log.Println(mods)
}
