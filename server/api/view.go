package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func LoginAction(context *gin.Context) {
	context.HTML(http.StatusOK, "login.html", nil)
}
