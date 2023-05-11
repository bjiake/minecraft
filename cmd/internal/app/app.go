package app

import (
	"encoding/json"
	"log"
	"minecraft/cmd/internal/models"
	"minecraft/cmd/internal/services"
	"minecraft/cmd/internal/transport"
)

func Start() {
	//requestsToMongoDB.GetLastID()
	services.GetLastPageNumber()
	//modeList := services.GetAllPages()
	//modeList := services.GetPages("1,6,7,8,9")
	//convertModListJson(modeList)
	//requestsToMongoDB.InsertMany(modeList)
	transport.Transport()
	//requestsToMongoDB.CloseConnection()
}

func convertModListJson(modeList []models.Mod) {
	_, err := json.Marshal(modeList)
	if err != nil {
		log.Fatal(err)
	}
	//log.Println(string(result))
}
