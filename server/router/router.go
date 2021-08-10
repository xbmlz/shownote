package router

import (
	"shownote/api"
	"shownote/middleware"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	router.Use(middleware.Cors())
	router.GET("/", api.LoginAction)
	router.GET("/auth/:repo", api.UserAuthAction)
	router.GET("/auth/refresh", api.UserAuthRefreshAction)

	user := router.Group("user", middleware.JwtAuth())
	{
		user.GET("/info", api.UserInfoAction)
	}

	repo := router.Group("repo", middleware.JwtAuth())
	{
		repo.GET("/info", api.GetRepoInfoAction)
		repo.GET("/workspace", api.GetWorkspaceAction)
		repo.PUT("/workspace", api.UpdateWorkspaceAction)
		repo.GET("/file", api.GetFileAction)
		repo.POST("/file", api.CreateFileAction)
		repo.PUT("/file", api.UpdateFileAction)
		repo.DELETE("/file", api.DeleteFileAction)
		repo.POST("/upload", api.UploadFileAction)
	}
	return router
}
