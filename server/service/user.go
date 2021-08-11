package service

import (
	"encoding/json"
	"fmt"
	"shownote/config"
	"shownote/global"
	"shownote/model"

	"github.com/go-resty/resty/v2"
)

func RefreshToken(refreshToken string) (auth global.AuthInfo, err error) {
	// https://gitee.com/oauth/token?grant_type=refresh_token&refresh_token={refresh_token}
	client := resty.New()
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(fmt.Sprintf(`{"grant_type":"refresh_token", "refresh_token":"%s"}`, refreshToken)).
		Post(config.AppConfig.GiteeConfig.TokenUrl)
	if err != nil {
		return auth, err
	}
	json.Unmarshal(resp.Body(), &auth)
	return auth, nil
}

func GetUserInfo(token string) (user model.User, err error) {
	client := resty.New()
	resp, err := client.R().
		SetQueryParam("access_token", token).
		Get("https://gitee.com/api/v5/user")
	if err != nil {
		return user, err
	}
	if resp.RawResponse.StatusCode != 200 {
		return user, err
	}
	json.Unmarshal(resp.Body(), &user)
	return user, nil
}

func GetToken(repoType, code string) (auth global.AuthInfo, err error) {
	// https://gitee.com/oauth/token?grant_type=authorization_code&code={code}&client_id={client_id}&redirect_uri={redirect_uri}&client_secret={client_secret}
	client := resty.New()
	var (
		tokenUrl     string
		clientId     string
		redirectUri  string
		clientSecret string
	)
	if repoType == "gitee" {
		tokenUrl = config.AppConfig.GiteeConfig.TokenUrl
		clientId = config.AppConfig.GiteeConfig.ClientId
		redirectUri = config.AppConfig.GiteeConfig.RedirectUri
		clientSecret = config.AppConfig.GiteeConfig.ClientSecret
	} else if repoType == "github" {
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
		return auth, err
	}
	fmt.Println(string(resp.Body()))
	json.Unmarshal(resp.Body(), &auth)
	auth.RepoType = repoType
	return auth, nil
}
