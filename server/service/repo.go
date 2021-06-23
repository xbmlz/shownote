package service

import (
	"encoding/json"
	"fmt"
	"shownote/global"

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

func CreateFile(token, login, path, content string) error {
	// https://gitee.com/api/v5/repos/{owner}/{repo}/contents/{path}
	client := resty.New()
	_, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(fmt.Sprintf(`{"access_token":"%s", "owner":"%s", "repo":"%s", "path":"%s", "content":"%s", "message":"%s"}`,
			token, login, global.REPO_NAME, path, content, "update")).
		Post(fmt.Sprintf("https://gitee.com/api/v5/repos/%s/%s/contents/%s", login, global.REPO_NAME, path))
	if err != nil {
		return err
	}
	return nil
}

func GetContent(token, login, path string) (string, error) {
	// https://gitee.com/api/v5/repos/{owner}/{repo}/contents(/{path})
	client := resty.New()
	resp, err := client.R().
		SetQueryParam("access_token", token).
		Get(fmt.Sprintf("https://gitee.com/api/v5/repos/%s/%s/contents/%s", login, global.REPO_NAME, path))
	if err != nil {
		return "", err
	}
	var respJson map[string]interface{}
	json.Unmarshal(resp.Body(), &respJson)
	if respJson == nil {
		return "", nil
	} else {
		return respJson["content"].(string), nil
	}
}
