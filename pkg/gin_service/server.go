package gin_service

import (
	_ "go_grpc_practice/docs"

	"github.com/gin-gonic/gin"
)

func setRouter() *gin.Engine {
	router := gin.Default()
	v1Group := router.Group("/api/v1")
	setRoutes(v1Group)

	return router
}

func Start(port string) {
	router := setRouter()
	router.Run(":" + port)
}
