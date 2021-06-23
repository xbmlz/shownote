package service

import (
	"fmt"

	"github.com/go-resty/resty/v2"
)

func CreateRepo(token, name string) error {
	client := resty.New()
	_, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(fmt.Sprintf(`{"access_token":"%s", "name":"%s"}`, token, name)).
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
			token, login, "shownote", path, content, "update")).
		Post(fmt.Sprintf("https://gitee.com/api/v5/repos/%s/%s/contents/%s", login, "shownote", path))
	if err != nil {
		return err
	}
	return nil
}
