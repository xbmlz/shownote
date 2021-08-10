package api

import (
	"encoding/base64"
	"io/ioutil"
	"net/http"
	"path"
	"shownote/global"
	"shownote/model/request"
	"shownote/model/response"
	"shownote/service"
	"strings"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

// @Tags 仓库
// @Summary 获取仓库信息
// @Description 获取仓库信息，没有则初始化仓库，并返回.shownote/workspace.json的内容
// @Param	token	query	string	true	"access_token"
// @Param	login	query	string	true	"用户名"
// @Success 200 object response.FileInfo
// @Router /repo/info [get]
func GetRepoInfoAction(c *gin.Context) {
	token := c.GetString("token")
	login := c.Query("login")
	var (
		isExist  bool
		err      error
		fileInfo response.FileInfo
	)
	if token == "" {
		response.FailWithMessage("接口未授权", c)
		return
	}
	// 0.  判断是否存在shownote
	// 1.  存在
	// 1.1 返回workspace.json
	isExist, err = service.ExistRepo(token)
	if err != nil {
		response.FailWithMessage("获取仓库配置失败", c)
		return
	}
	if !isExist {
		// 2.  不存在
		// 2.1 创建一个shownote的仓库
		if err = service.CreateRepo(token, global.REPO_NAME); err != nil {
			response.FailWithMessage("仓库初始化失败, "+err.Error(), c)
			return
		}
		// 2.2 创建一个.shownote文件夹
		// 2.3 在.shownote下创建一个workspace.json的配置文件
		const workspaceJson = `{
	"notes": [],
	"trash": [],
	"share": []
}`
		encodeStr := base64.StdEncoding.EncodeToString([]byte(workspaceJson))
		if fileInfo, err = service.CreateFile(token, login, global.WORKSPACE_PATH, encodeStr); err != nil {
			response.FailWithMessage("仓库初始化失败, "+err.Error(), c)
			return
		}
	}
	fileInfo, err = service.GetContent(token, login, global.WORKSPACE_PATH)
	if err != nil {
		response.FailWithMessage("仓库信息获取失败, "+err.Error(), c)
		return
	}
	response.OkWithDetailed(fileInfo, "账号初始化成功", c)
}

// @Tags 仓库
// @Summary 获取workspace内容
// @Description 获取workspace内容
// @Param	token	query	string	true	"access_token"
// @Param	login	query	string	true	"用户名"
// @Success 200 object response.FileInfo    "成功后返回值"
// @Router /repo/workspace [get]
func GetWorkspaceAction(c *gin.Context) {
	token := c.GetString("token")
	login := c.Query("login")
	var (
		err      error
		fileInfo response.FileInfo
	)
	fileInfo, err = service.GetContent(token, login, global.WORKSPACE_PATH)
	if err != nil {
		response.FailWithMessage("仓库信息获取失败, "+err.Error(), c)
		return
	}
	response.OkWithDetailed(fileInfo, "账号初始化成功", c)
}

// @Tags 仓库
// @Summary 更新workspace
// @Description 更新workspace
// @Produce json
// @Param	token	body  	request.FileInfo	true	"FileInfo"
// @Success 200 object response.FileInfo "成功后返回值"
// @Router /repo/workspace [put]
func UpdateWorkspaceAction(c *gin.Context) {
	var (
		reqFileInfo request.FileInfo
		err         error
		fileInfo    response.FileInfo
	)
	if err := c.ShouldBindJSON(&reqFileInfo); err != nil {
		response.FailWithMessage("参数解析失败", c)
		return
	}
	if fileInfo, err = service.UpdateFile(reqFileInfo.Token, reqFileInfo.Login, global.WORKSPACE_PATH, reqFileInfo.Content, reqFileInfo.Sha); err != nil {
		response.FailWithMessage("更新工作空间失败, "+err.Error(), c)
		return
	}
	fileInfo.Content = reqFileInfo.Content
	response.OkWithData(fileInfo, c)
}

// @Tags 仓库
// @Summary 获取仓库具体路径下的内容
// @Description 获取仓库具体路径下的内容
// @Param	token	query	string	true	"access_token"
// @Param	login	query	string	true	"用户名"
// @Param	uid	query	string	true	"文件UID"
// @Success 200 object response.FileInfo    "成功后返回值"
// @Router /repo/file [get]
func GetFileAction(c *gin.Context) {
	token := c.GetString("token")
	login := c.Query("login")
	uid := c.Query("uid")
	var (
		err      error
		fileInfo response.FileInfo
	)
	if token == "" {
		response.FailWithMessage("接口未授权", c)
		return
	}
	fileInfo, err = service.GetContent(token, login, "notes/"+uid+".md")
	if err != nil {
		response.FailWithMessage("获取文件内容失败, "+err.Error(), c)
		return
	}
	// fileInfo.Html = utils.MarkdownToHtml(fileInfo.Content)
	fileInfo.Uid = uid
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
	if err = c.ShouldBindJSON(&reqFileInfo); err != nil {
		response.FailWithMessage("参数解析失败", c)
		return
	}
	uid := strings.ReplaceAll(uuid.NewV4().String(), "-", "")
	reqFileInfo.Content = base64.StdEncoding.EncodeToString([]byte(reqFileInfo.Content))
	if fileInfo, err = service.CreateFile(reqFileInfo.Token, reqFileInfo.Login, "notes/"+uid+".md", reqFileInfo.Content); err != nil {
		response.FailWithMessage("创建文件失败, "+err.Error(), c)
		return
	}
	fileInfo.Uid = uid
	response.OkWithData(fileInfo, c)
}

// @Tags 仓库
// @Summary 上传文件
// @Description 上传文件
// @Produce json
// @Param	type	query	string	true	"文件类型,image|viodo|audio"
// @Success 200 object response.FileInfo "成功后返回值"
// @Router /repo/upload [post]
func UploadFileAction(c *gin.Context) {
	var (
		fileType string
		err      error
		fileInfo response.FileInfo
	)
	token := c.PostForm("token")
	login := c.PostForm("login")
	fileContent, _ := c.FormFile("file")
	o, _ := fileContent.Open()
	io, _ := ioutil.ReadAll(o)
	content := base64.StdEncoding.EncodeToString(io)
	uid := strings.ReplaceAll(uuid.NewV4().String(), "-", "")
	fileSuffix := strings.ToLower(path.Ext(fileContent.Filename))
	if fileSuffix == ".jpeg" || fileSuffix == ".jpg" || fileSuffix == ".png" || fileSuffix == ".bmp" || fileSuffix == ".gif" {
		fileType = "image"
	} else if fileSuffix == ".mp3" || fileSuffix == ".wav" || fileSuffix == ".acc" {
		fileType = "audio"
	} else if fileSuffix == ".mp4" || fileSuffix == ".ogg" || fileSuffix == ".webm" {
		fileType = "video"
	} else {
		fileType = "other"
	}
	fileName := "assets/" + fileType + "/" + uid + fileSuffix
	if fileInfo, err = service.CreateFile(token, login, fileName, content); err != nil {
		response.FailWithMessage("创建文件失败, "+err.Error(), c)
		return
	}
	result := make(map[string]interface{})
	succMap := make(map[string]string)
	succMap[fileInfo.Name] = fileInfo.DownloadUrl
	result["errFiles"] = [...]string{}
	result["succMap"] = succMap
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "",
		"data": result,
	})
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
	if fileInfo, err = service.UpdateFile(reqFileInfo.Token, reqFileInfo.Login, "notes/"+reqFileInfo.Uid+".md", reqFileInfo.Content, reqFileInfo.Sha); err != nil {
		response.FailWithMessage("更新文件失败, "+err.Error(), c)
		return
	}
	fileInfo.Uid = reqFileInfo.Uid
	response.OkWithData(fileInfo, c)
}

// @Tags 仓库
// @Summary 删除文件
// @Description 删除文件
// @Param	token	query	string	true	"access_token"
// @Param	login	query	string	true	"用户名"
// @Param	uid		query	string	true	"文件UID"
// @Param	sha		query	string	true	"文件sha"
// @Success 200 object response.FileInfo    "成功后返回值"
// @Router /repo/file [delete]
func DeleteFileAction(c *gin.Context) {
	token := c.Query("token")
	login := c.Query("login")
	uid := c.Query("uid")
	sha := c.Query("sha")
	var (
		err      error
		fileInfo response.FileInfo
	)
	if token == "" {
		response.FailWithMessage("接口未授权", c)
		return
	}
	fileInfo, err = service.DeleteFile(token, login, "notes/"+uid+".md", sha)
	if err != nil {
		response.FailWithMessage("删除文件失败, "+err.Error(), c)
		return
	}
	response.OkWithData(fileInfo, c)
}
