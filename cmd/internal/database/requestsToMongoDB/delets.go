package requestsToMongoDB

import (
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/net/context"
	"log"
	"minecraft/cmd/internal/models"
)

func DeleteOne(mod models.Mod) {
	collection, err := getCollection()
	if err != nil {
		log.Fatal(err)
	}

	filter := bson.D{{"title", mod.Title}}
	deleteResult, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Deleted: %v document in the mods collection\n", deleteResult)
}
func DeleteMany(modList []models.Mod) {
	for _, mod := range modList {
		DeleteOne(mod)
	}
}
