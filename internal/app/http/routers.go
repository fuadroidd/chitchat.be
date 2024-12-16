// routers
package http

import (
	"github.com/fuadroidd/chitchat.be/internal/app/http/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/", handlers.HomeHandler)
	return router

}
