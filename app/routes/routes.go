package routes

import (
	"myproject/app/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	handlers.SetUserHandler(r)
}
