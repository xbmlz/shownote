package router

import (
	"shownote/api"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/", api.LoginAction)
	router.GET("/auth/:repo", api.AuthAction)
	router.GET("/user", api.UserAction)
	return router
}
