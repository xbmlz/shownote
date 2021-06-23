package router

import (
	"shownote/api"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/", api.LoginAction)

	user := router.Group("user")
	{
		user.GET("/auth/:repo", api.UserAuthAction)
		user.GET("/info", api.UserInfoAction)
	}

	repo := router.Group("repo")
	{
		repo.GET("/info", api.RepoInfoAction)
		repo.GET("/content", api.RepoContentAction)
	}
	return router
}
