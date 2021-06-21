package api

import (
	"shownote/model/response"
	"shownote/utils"

	"github.com/gin-gonic/gin"
)

// @Summary gitee授权登录
// @Description gitee auth login
// @Param	code	query	string	true	"{redirect_uri}?code=abc"
// @Accept  json
// @Produce json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /auth/ [get]
func AuthAction(c *gin.Context) {
	code := c.Query("code")
	// state := c.Query("state")
	// func GetToken(code, clientId, redirectUri, clientSecret string) string
	token, err := utils.GetToken(code, "f5763537e579c4f97c56c69b80489a17e250d8186e48efbe3e3fba4c4b6c9558", "http://127.0.0.1:8000/auth", "43896c34a76ebf45b408cd05436dd5e8cabea1e2c20c6e7ba9dfb6fde908c5ba")
	if err != nil {
		response.Fail(c)
	} else {
		response.OkWithData(token, c)
	}
}
