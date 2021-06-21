package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func IndexAction(context *gin.Context) {
	context.HTML(http.StatusOK, "index.html", nil)
}
