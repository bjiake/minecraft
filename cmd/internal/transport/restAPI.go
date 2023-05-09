package transport

import (
	"github.com/gin-gonic/gin"
	"minecraft/cmd/internal/database/requestsToMongoDB"
	"net/http"
	"strconv"
	"strings"
)

func Transport() {
	router := gin.Default()
	router.GET("/mods", getMods)
	router.GET("/mods/:page", getModsByPage)

	router.Run("25.10.209.53:8000")
}

// * &
func getMods(contex *gin.Context) {
	contex.IndentedJSON(http.StatusOK, requestsToMongoDB.FindAll())
}

func getModsByPage(context *gin.Context) {
	pageStr := context.Param("page")
	parts := strings.Split(pageStr, "=")

	page, err := strconv.Atoi(parts[1]) // конвертируем строку в число
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "invalid page number"})
		return
	}
	modList := requestsToMongoDB.FindByPage(page)

	context.IndentedJSON(http.StatusOK, modList)
	return

	context.IndentedJSON(http.StatusNotFound, gin.H{"message": "mods not found"})
}
