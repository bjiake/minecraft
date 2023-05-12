package transport

import (
	"github.com/gin-gonic/gin"
	"log"
	"minecraft/cmd/internal/config"
	"minecraft/cmd/internal/models"
	"minecraft/cmd/internal/services"
	"net/http"
	"strconv"
)

func Transport() {
	router := gin.Default()
	//router.GET("/mods", getMods)
	router.GET("/mods", getModsByQuery)

	router.Run("localhost:8080")
}

func getModsByQuery(context *gin.Context) {
	pageStr := context.Query("page")
	//parts := strings.Split(pageStr, "=")
	//_ = parts

	page, err := strconv.Atoi(pageStr) // конвертируем строку в число
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "invalid page number to convert int"})
		log.Fatal(err)
		return
	}

	var result models.RequestAPI
	newModList := services.GetPage(pageStr)
	//modList := requestsToMongoDB.FindByPage(page)
	if newModList == nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"error": "invalid page number"})
		return
	}
	totalPage := int(config.LastID/10) + 1

	prevPage := page - 1
	var prevPageStr string
	if prevPage <= 0 {
		prevPageStr = "null"
	} else {
		prevPageStr = strconv.Itoa(prevPage)
	}

	nextPage := page + 1
	var nextPageStr string
	if nextPage >= totalPage {
		nextPageStr = "null"
	} else {
		nextPageStr = strconv.Itoa(nextPage)
	}

	result = models.RequestAPI{CurrentPage: page, Data: newModList, PrevPage: prevPageStr, NextPage: nextPageStr, TotalPage: totalPage}

	context.IndentedJSON(http.StatusOK, result)
	return
}
