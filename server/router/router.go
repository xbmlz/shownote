package router

import (
	"shownote/api"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/", api.IndexAction)
	router.GET("/auth", api.AuthAction)
	return router
}
