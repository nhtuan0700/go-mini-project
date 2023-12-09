package api

import (
	"fmt"
	"myproject/app/routes"
	"myproject/app/util/db"
	"myproject/app/util/env"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func RunServer() {
	server := SetServer()

	serverPort := fmt.Sprintf(":%s", env.ServerPortNo())
	err := server.Run(serverPort)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Server is running: %s", serverPort)
}

func SetServer() *gin.Engine {
	server := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = env.AllowOrigins()
	server.Use(cors.New(config))
	db.Open(env.DBName())

	healthCheck(server)

	routes.SetupRoutes(server)

	return server
}

func healthCheck(server *gin.Engine) {
	server.GET("health", func(context *gin.Context) {
		context.JSON(http.StatusOK, map[string]string{"message": "ok"})
	})
}
