package api

import (
	"encoding/base64"
	"shownote/global"
	"shownote/model/request"
	"shownote/model/response"
	"shownote/service"

	"github.com/gin-gonic/gin"
)

// @Tags 仓库
// @Summary 获取仓库信息
// @Description 获取仓库信息，没有则初始化仓库，并返回.shownote/workspace.json的内容
// @Param	token	query	string	true	"access_token"
// @Param	login	query	string	true	"用户名"
// @Success 200 object response.FileInfo
// @Router /repo/info [get]
func GetRepoInfoAction(c *gin.Context) {
	token := c.Query("token")
	login := c.Query("login")
	var (
		isExist  bool
		err      error
		fileInfo response.FileInfo
	)
	if token == "" {
		response.FailWithMessage("接口未授权", c)
	}
	// 0.  判断是否存在shownote
	// 1.  存在
	// 1.1 返回workspace.json
	isExist, err = service.ExistRepo(token)
	if err != nil {
		response.FailWithMessage("获取仓库配置失败", c)
	}
	if !isExist {
		// 2.  不存在
		// 2.1 创建一个shownote的仓库
		if err = service.CreateRepo(token, global.REPO_NAME); err != nil {
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
		if fileInfo, err = service.CreateFile(token, login, filePath, encodeStr); err != nil {
			response.FailWithMessage("仓库配置文件失败", c)
		}
	}
	fileInfo, err = service.GetContent(token, login, global.WORKSPACE_PATH)
	if err != nil {
		response.FailWithMessage("获取文件内容失败", c)
	}
	if err != nil {
		response.FailWithMessage("文件解密失败", c)
	}
	response.OkWithDetailed(fileInfo, "账号初始化成功", c)
}

// @Tags 仓库
// @Summary 获取仓库具体路径下的内容
// @Description 获取仓库具体路径下的内容
// @Param	token	query	string	true	"access_token"
// @Param	login	query	string	true	"用户名"
// @Param	path	query	string	true	"文件路径"
// @Success 200 object response.FileInfo    "成功后返回值"
// @Router /repo/content [get]
func GetRepoContentAction(c *gin.Context) {
	token := c.Query("token")
	login := c.Query("login")
	path := c.Query("path")
	var (
		err      error
		fileInfo response.FileInfo
	)
	if token == "" {
		response.FailWithMessage("接口未授权", c)
	}
	fileInfo, err = service.GetContent(token, login, path)
	if err != nil {
		response.FailWithMessage("获取文件内容失败", c)
	}
	if err != nil {
		response.FailWithMessage("文件解密失败", c)
	}
	response.OkWithData(fileInfo, c)
}

// @Tags 仓库
// @Summary 新建文件
// @Description 新建文件
// @Produce json
// @Param	token	body  	request.FileInfo	true	"FileInfo"
// @Success 200 object response.FileInfo "成功后返回值"
// @Router /repo/file [post]
func CreateFileAction(c *gin.Context) {
	var (
		reqFileInfo request.FileInfo
		err         error
		fileInfo    response.FileInfo
	)
	if err := c.ShouldBindJSON(&reqFileInfo); err != nil {
		response.FailWithMessage("参数解析失败", c)
		return
	}
	if fileInfo, err = service.CreateFile(reqFileInfo.Token, reqFileInfo.Login, reqFileInfo.Path, reqFileInfo.Content); err != nil {
		response.FailWithMessage("创建文件失败", c)
	}
	response.OkWithData(fileInfo, c)
}

// @Tags 仓库
// @Summary 更新文件
// @Description 更新文件
// @Produce json
// @Param	token	body  	request.FileInfo	true	"FileInfo"
// @Success 200 object response.FileInfo "成功后返回值"
// @Router /repo/file [put]
func UpdateFileAction(c *gin.Context) {
	var (
		reqFileInfo request.FileInfo
		err         error
		fileInfo    response.FileInfo
	)
	if err := c.ShouldBindJSON(&reqFileInfo); err != nil {
		response.FailWithMessage("参数解析失败", c)
		return
	}
	// encodeStr := base64.StdEncoding.EncodeToString([]byte(content))
	if fileInfo, err = service.UpdateFile(reqFileInfo.Token, reqFileInfo.Login, reqFileInfo.Path, reqFileInfo.Content, reqFileInfo.Sha); err != nil {
		response.FailWithMessage("创建文件失败", c)
	}
	response.OkWithData(fileInfo, c)
}
