package app

import (
	"minecraft/cmd/internal/config"
	"minecraft/cmd/internal/database/requestsToMongoDB"
	"minecraft/cmd/internal/services"
	"minecraft/cmd/internal/transport"
)

func Start() {
	requestsToMongoDB.GetLastID()
	services.GetLastPageNumber()

	if config.LastID == 0 {
		//modeList := services.GetAllPages()
		modeList := services.GetPages("1,6,7")
		requestsToMongoDB.InsertMany(modeList)
	}
	transport.Transport()
	requestsToMongoDB.CloseConnection()
}
