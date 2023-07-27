package app

import (
	"BestCleanTemplate/api/handler"
	"BestCleanTemplate/pkg/repository"
	"BestCleanTemplate/pkg/services"
	"BestCleanTemplate/scripts/database"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

func main() {
	gin.SetMode(gin.DebugMode)
	r := gin.New()
	err := r.SetTrustedProxies([]string{"127.0.0.1"})
	if err != nil {
		println("error:%s", err)
	}

	db := database.DatabaseSetup()

	repositories := repository.InitRepositories(db)

	apiServices := services.InitServices(repositories)

	handler.SetHandlers(r, apiServices)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	println("port:", port)

	if err := r.Run(":" + port); err != nil {
		log.Panicf("error: %s", err)
	}
}
