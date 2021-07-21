package service

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"shownote/global"
	"shownote/model/response"

	"github.com/bitly/go-simplejson"
	"github.com/go-resty/resty/v2"
)

func ExistRepo(token string) (bool, error) {
	client := resty.New()
	resp, err := client.R().
		SetQueryParam("access_token", token).
		SetQueryParam("q", global.REPO_NAME).
		Get("https://gitee.com/api/v5/user/repos")
	if err != nil {
		return false, err
	}
	var repoArr []interface{}
	json.Unmarshal(resp.Body(), &repoArr)
	if len(repoArr) == 0 {
		return false, nil
	} else {
		return true, nil
	}
}

func CreateRepo(token, name string) error {
	client := resty.New()
	_, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(fmt.Sprintf(`{"access_token":"%s", "name":"%s", "auto_init":true}`, token, name)).
		Post("https://gitee.com/api/v5/user/repos")
	if err != nil {
		return err
	}
	return nil
}

func CreateFile(token, login, path, content string) (response.FileInfo, error) {
	// https://gitee.com/api/v5/repos/{owner}/{repo}/contents/{path}
	var (
		info response.FileInfo
	)
	content = base64.StdEncoding.EncodeToString([]byte(content))
	client := resty.New()
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(fmt.Sprintf(`{"access_token":"%s", "owner":"%s", "repo":"%s", "path":"%s", "content":"%s", "message":"%s"}`,
			token, login, global.REPO_NAME, path, content, global.COMMIT_MSG)).
		Post(fmt.Sprintf("https://gitee.com/api/v5/repos/%s/%s/contents/%s", login, global.REPO_NAME, path))
	if err != nil {
		return info, err
	}
	jsonObj, err := simplejson.NewJson([]byte(resp.Body()))
	if err != nil {
		return info, err
	}
	fmt.Println(jsonObj.Get("message").MustString())
	if resp.RawResponse.StatusCode > 201 {
		return info, errors.New(jsonObj.Get("message").MustString())
	}
	decStr, err := base64.StdEncoding.DecodeString(jsonObj.Get("content").MustString())
	if err != nil {
		return info, err
	}
	info.Content = string(decStr)
	info.Name = jsonObj.Get("content").Get("name").MustString()
	info.Path = jsonObj.Get("content").Get("path").MustString()
	info.Content = jsonObj.Get("content").Get("content").MustString()
	info.Sha = jsonObj.Get("content").Get("sha").MustString()
	info.Size = jsonObj.Get("content").Get("size").MustInt()
	info.Url = jsonObj.Get("content").Get("url").MustString()
	info.FileType = jsonObj.Get("content").Get("type").MustString()
	info.HtmlUrl = jsonObj.Get("content").Get("html_url").MustString()
	info.DownloadUrl = jsonObj.Get("content").Get("download_url").MustString()
	return info, err
}

func UpdateFile(token, login, path, content, sha string) (response.FileInfo, error) {
	// https://gitee.com/api/v5/repos/{owner}/{repo}/contents/{path}
	var (
		info response.FileInfo
	)
	content = base64.StdEncoding.EncodeToString([]byte(content))
	client := resty.New()
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(fmt.Sprintf(`{"access_token":"%s", "owner":"%s", "repo":"%s", "path":"%s", "content":"%s", "sha":"%s", "message":"%s"}`,
			token, login, global.REPO_NAME, path, content, sha, global.COMMIT_MSG)).
		Put(fmt.Sprintf("https://gitee.com/api/v5/repos/%s/%s/contents/%s", login, global.REPO_NAME, path))
	if err != nil {
		return info, err
	}
	jsonObj, err := simplejson.NewJson([]byte(resp.Body()))
	if err != nil {
		return info, err
	}
	if resp.RawResponse.StatusCode > 201 {
		return info, errors.New(jsonObj.Get("message").MustString())
	}
	decStr, err := base64.StdEncoding.DecodeString(jsonObj.Get("content").MustString())
	if err != nil {
		return info, err
	}
	info.Content = string(decStr)
	info.Name = jsonObj.Get("content").Get("name").MustString()
	info.Path = jsonObj.Get("content").Get("path").MustString()
	info.Content = jsonObj.Get("content").Get("content").MustString()
	info.Sha = jsonObj.Get("content").Get("sha").MustString()
	info.Size = jsonObj.Get("content").Get("size").MustInt()
	info.Url = jsonObj.Get("content").Get("url").MustString()
	info.FileType = jsonObj.Get("content").Get("type").MustString()
	info.HtmlUrl = jsonObj.Get("content").Get("html_url").MustString()
	info.DownloadUrl = jsonObj.Get("content").Get("download_url").MustString()
	return info, err
}

func GetContent(token, login, path string) (response.FileInfo, error) {
	// https://gitee.com/api/v5/repos/{owner}/{repo}/contents(/{path})
	var (
		info response.FileInfo
	)
	client := resty.New()
	resp, err := client.R().
		SetQueryParam("access_token", token).
		Get(fmt.Sprintf("https://gitee.com/api/v5/repos/%s/%s/contents/%s", login, global.REPO_NAME, path))
	if err != nil {
		return info, err
	}
	jsonObj, err := simplejson.NewJson([]byte(resp.Body()))
	if err != nil {
		return info, err
	}
	if resp.RawResponse.StatusCode > 201 {
		return info, errors.New(jsonObj.Get("message").MustString())
	}
	decStr, err := base64.StdEncoding.DecodeString(jsonObj.Get("content").MustString())
	if err != nil {
		return info, err
	}
	info.Content = string(decStr)
	info.Name = jsonObj.Get("name").MustString()
	info.Path = jsonObj.Get("path").MustString()
	info.Sha = jsonObj.Get("sha").MustString()
	info.Size = jsonObj.Get("size").MustInt()
	info.Url = jsonObj.Get("url").MustString()
	info.FileType = jsonObj.Get("type").MustString()
	info.HtmlUrl = jsonObj.Get("html_url").MustString()
	info.DownloadUrl = jsonObj.Get("download_url").MustString()
	return info, err
}
