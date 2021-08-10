package api

import (
	"net/http"
	"shownote/config"
	"shownote/middleware"
	"shownote/model/response"
	"shownote/service"

	"github.com/gin-gonic/gin"
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
	repoType := c.Param("repo")
	authInfo, err := service.GetToken(repoType, code)
	token := middleware.GenToken(authInfo)
	if err != nil {
		c.Redirect(http.StatusMovedPermanently, config.AppConfig.WebConfig.ErrorUrl)
	} else {
		c.Redirect(http.StatusMovedPermanently, config.AppConfig.WebConfig.IndexUrl+"?token="+token)
	}
}

// @Tags 用户
// @Summary 刷新授权
// @Description 刷新token
// @Param	refresh_token	query	string	true	"refresh_token"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /user/auth/refresh [get]
func UserAuthRefreshAction(c *gin.Context) {
	refreshToken := c.Query("refresh_token")
	token, err := service.RefreshToken(refreshToken)
	if err != nil {
		response.FailWithMessage("刷新token失败", c)
		return
	}
	response.OkWithData(token, c)
}

// @Tags 用户
// @Summary 获取用户信息
// @Description 获取用户的基本信息
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /user/info [get]
func UserInfoAction(c *gin.Context) {
	token := c.GetString("token")
	user, err := service.GetUserInfo(token)
	if err != nil {
		response.FailWithMessage("获取用户信息失败", c)
		return
	}
	response.OkWithData(user, c)
}
