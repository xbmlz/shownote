package api

import (
	"shownote/model/response"
	"shownote/service"

	"github.com/gin-gonic/gin"
)

// @Summary 初始化
// @Description初始化仓库信息
// @Param	token	query	string	true	"accesstoken"
// @Param	login	query	string	true	"用户名"
// @Success 200 {string string "{"success":true,"data":{},"msg":"账号初始化成功"}"
// @Router /init [get]
func InitAction(c *gin.Context) {
	token := c.Query("token")
	login := c.Query("login")
	if token == "" {
		response.FailWithMessage("接口未授权", c)
	}
	// 1. 创建一个shownote的仓库
	if err := service.CreateRepo(token, "shownote"); err != nil {
		response.FailWithMessage("仓库初始化失败", c)
	}
	// 2. 创建一个.shownote文件夹
	// 3. 在.shownote下创建一个workspace.json的配置文件
	const filePath = ".shownote/workspace.json"
	const workspaceJson = `{
		"note": []
	}`
	if err := service.CreateFile(token, login, filePath, workspaceJson); err != nil {
		response.FailWithMessage("仓库配置文件失败", c)
	}
	response.OkWithMessage("账号初始化成功", c)
}
