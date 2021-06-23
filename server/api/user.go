package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"shownote/config"
	"shownote/model"
	"shownote/model/response"
	"shownote/service"

	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
)

// @Tags 用户
// @Summary 授权登录
// @Description gitee auth login
// @Param	repo	query	string	true	"仓库类型 gitee or github"
// @Param	code	query	string	true	"{redirect_uri}?code=abc"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /user/auth/:repo [get]
func UserAuthAction(c *gin.Context) {
	code := c.Query("code")
	repo := c.Param("repo")
	// state := c.Query("state")
	// func GetToken(code, clientId, redirectUri, clientSecret string) string
	token, err := service.GetToken(code, repo)
	if err != nil {
		c.Redirect(http.StatusMovedPermanently, config.AppConfig.WebConfig.ErrorUrl)
	} else {
		c.Redirect(http.StatusMovedPermanently, config.AppConfig.WebConfig.IndexUrl+"?token="+token.AccessToken)
	}
}

// @Tags 用户
// @Summary 获取用户信息
// @Description 获取用户的基本信息
// @Param	token	query	string	true	"accesstoken"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /user/info [get]
func UserInfoAction(c *gin.Context) {
	token := c.Query("token")
	if token == "" {
		response.FailWithMessage("接口未授权", c)
	}
	client := resty.New()
	resp, err := client.R().
		SetQueryParam("access_token", token).
		Get("https://gitee.com/api/v5/user")

	if err != nil {
		response.FailWithMessage("获取用户信息失败", c)
	}
	var userInfo model.User
	json.Unmarshal(resp.Body(), &userInfo)
	fmt.Println(userInfo)
	response.OkWithData(userInfo, c)
}
