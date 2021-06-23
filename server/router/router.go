package router

import (
	"shownote/api"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/", api.LoginAction)

	userRouter := router.Group("user")
	{
		userRouter.GET("/auth/:repo", api.AuthAction)
		userRouter.GET("/info", api.UserAction)
	}

	repoRouter := router.Group("repo")
	{
		repoRouter.GET("/init")
	}
	return router
}
