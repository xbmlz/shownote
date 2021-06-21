package main

import (
	_ "shownote/docs"
	"shownote/router"

	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod download

// @title ShowNote API
// @version 0.0.1
// @description This is a sample server Petstore server.
// @host localhost:8000
// @BasePath /
func main() {
	router := router.InitRouter()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.LoadHTMLGlob("templates/*")
	router.Run(":8000") // listen and serve on 0.0.0.0:8000
}
