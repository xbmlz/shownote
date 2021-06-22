package utils

import (
	"encoding/json"
	"fmt"
	"shownote/config"
	"shownote/global"

	"github.com/go-resty/resty/v2"
)

func GetToken(code, repo string) (global.Token, error) {
	// https://gitee.com/oauth/token?grant_type=authorization_code&code={code}&client_id={client_id}&redirect_uri={redirect_uri}&client_secret={client_secret}
	client := resty.New()
	var (
		token        global.Token
		tokenUrl     string
		clientId     string
		redirectUri  string
		clientSecret string
	)
	if repo == "gitee" {
		tokenUrl = config.AppConfig.GiteeConfig.TokenUrl
		clientId = config.AppConfig.GiteeConfig.ClientId
		redirectUri = config.AppConfig.GiteeConfig.RedirectUri
		clientSecret = config.AppConfig.GiteeConfig.ClientSecret
	} else if repo == "github" {
		tokenUrl = config.AppConfig.GiteeConfig.TokenUrl
		clientId = config.AppConfig.GithubConfig.ClientId
		redirectUri = config.AppConfig.GithubConfig.RedirectUri
		clientSecret = config.AppConfig.GithubConfig.ClientSecret
	}
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(fmt.Sprintf(`{"grant_type":"authorization_code", "code":"%s", "client_id":"%s", "redirect_uri":"%s", "client_secret":"%s"}`,
			code, clientId, redirectUri, clientSecret)).
		Post(tokenUrl)
	if err != nil {
		return token, err
	}
	fmt.Println(resp.String())
	json.Unmarshal(resp.Body(), &token)
	return token, nil
}
