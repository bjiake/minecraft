package requestsToMongoDB

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/net/context"
	"log"
	"minecraft/cmd/internal/models"
)

func FindOne(mod models.Mod) models.Mod {
	collection, err := getCollection()
	if err != nil {
		log.Fatal(err)
	}

	var result models.Mod

	filter := bson.D{{"title", mod.Title}}
	err = collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Found a single document: %+v\n", result)
	return result
}
func FindByID(id int) models.Mod {
	collection, err := getCollection()
	if err != nil {
		log.Fatal(err)
	}

	var result models.Mod

	filter := bson.D{{"id", id}}
	err = collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Found a single document: %+v\n", result)
	return result
}
func FindByPage(page int) []models.Mod {

	startID := (page - 1) * 10 // вычисляем стартовый ID для этой страницы
	endID := startID + 10      // вычисляем конечный ID для этой страницы

	collection, err := getCollection()
	if err != nil {
		log.Fatal(err)
	}

	filter := bson.M{"id": bson.M{"$gte": startID, "$lte": endID}}
	cur, err := collection.Find(context.Background(), filter)

	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer cur.Close(context.Background())

	var result []models.Mod

	for cur.Next(context.Background()) {
		var mod models.Mod
		err := cur.Decode(&mod)
		if err != nil {
			log.Fatal(err)
		}
		result = append(result, mod)
	}

	if err := cur.Err(); err != nil {
		fmt.Println(err)
		return nil
	}

	fmt.Printf("FindByPage(%d): %v", page, result)

	return result
}
func FindAll() []models.Mod {
	collection, err := getCollection()
	if err != nil {
		log.Fatal(err)
	}

	filter := bson.M{}
	var results []models.Mod

	// Finding multiple documents returns a cursor
	// Iterating through the cursor allows us to decode documents one at a time
	cur, err := collection.Find(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}

	for cur.Next(context.TODO()) {
		var mod models.Mod
		err := cur.Decode(&mod)
		if err != nil {
			log.Fatal(err)
		}
		results = append(results, mod)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	// Close the cursor once finished
	cur.Close(context.TODO())

	fmt.Printf("Found multiple documents (array of pointers): %+v\n", results)
	return results
}
