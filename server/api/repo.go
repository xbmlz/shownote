package api

import (
	"encoding/base64"
	"shownote/global"
	"shownote/model/response"
	"shownote/service"

	"github.com/gin-gonic/gin"
)

// @Tags 仓库
// @Summary 获取仓库信息
// @Description 获取仓库信息，没有则初始化仓库
// @Param	token	query	string	true	"accesstoken"
// @Param	login	query	string	true	"用户名"
// @Success 200 {string string "{"success":true,"data":{},"msg":"账号初始化成功"}"
// @Router /repo/info [get]
func RepoInfoAction(c *gin.Context) {
	token := c.Query("token")
	login := c.Query("login")
	if token == "" {
		response.FailWithMessage("接口未授权", c)
	}
	// 0.  判断是否存在shownote
	// 1.  存在
	// 1.1 返回workspace.json
	isExist, err := service.ExistRepo(token)
	if err != nil {
		response.FailWithMessage("获取仓库配置失败", c)
	}
	if !isExist {
		// 2.  不存在
		// 2.1 创建一个shownote的仓库
		if err := service.CreateRepo(token, global.REPO_NAME); err != nil {
			response.FailWithMessage("仓库初始化失败", c)
		}
		// 2.2 创建一个.shownote文件夹
		// 2.3 在.shownote下创建一个workspace.json的配置文件
		const filePath = ".shownote/workspace.json"
		const workspaceJson = `{
	"note": [],
	"trash": [],
	"share": []
}`
		encodeStr := base64.StdEncoding.EncodeToString([]byte(workspaceJson))
		if err := service.CreateFile(token, login, filePath, encodeStr); err != nil {
			response.FailWithMessage("仓库配置文件失败", c)
		}
	}
	content, err := service.GetContent(token, login, global.WORKSPACE_PATH)
	if err != nil {
		response.FailWithMessage("获取文件内容失败", c)
	}
	decodeStr, err := base64.StdEncoding.DecodeString(content)
	if err != nil {
		response.FailWithMessage("文件解密失败", c)
	}
	response.OkWithDetailed(string(decodeStr), "账号初始化成功", c)
}

// @Tags 仓库
// @Summary 获取文件信息
// @Description 获取文件信息
// @Param	token	query	string	true	"accesstoken"
// @Param	login	query	string	true	"用户名"
// @Param	path	query	string	true	"文件路径"
// @Success 200 {string string "{"success":true,"data":{},"msg":"账号初始化成功"}"
// @Router /repo/content [get]
func RepoContentAction(c *gin.Context) {
	token := c.Query("token")
	login := c.Query("login")
	path := c.Query("path")
	if token == "" {
		response.FailWithMessage("接口未授权", c)
	}
	content, err := service.GetContent(token, login, path)
	if err != nil {
		response.FailWithMessage("获取文件内容失败", c)
	}
	decodeStr, err := base64.StdEncoding.DecodeString(content)
	if err != nil {
		response.FailWithMessage("文件解密失败", c)
	}
	response.OkWithData(string(decodeStr), c)
}
