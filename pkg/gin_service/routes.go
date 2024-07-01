package gin_service

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func setRoutes(ginGroup *gin.RouterGroup) {
	ginGroup.POST("/comment", CreateComment)
	ginGroup.GET("/comment/:id", GetComment)
	ginGroup.PUT("/comment/:id", UpdateComment)
	ginGroup.GET("/comments/", ListComment)
	ginGroup.DELETE("/comment/:id", DeleteComment)
	ginGroup.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
