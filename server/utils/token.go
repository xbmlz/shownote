package utils

import (
	"encoding/json"
	"fmt"
	"shownote/global"

	"github.com/go-resty/resty/v2"
)

func GetToken(code, clientId, redirectUri, clientSecret string) (global.Token, error) {
	// https://gitee.com/oauth/token?grant_type=authorization_code&code={code}&client_id={client_id}&redirect_uri={redirect_uri}&client_secret={client_secret}
	client := resty.New()
	var token global.Token
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(fmt.Sprintf(`{"grant_type":"authorization_code", "code":"%s", "client_id":"%s", "redirect_uri":"%s", "client_secret":"%s"}`,
			code, clientId, redirectUri, clientSecret)).
		Post("https://gitee.com/oauth/token")
	if err != nil {
		return token, err
	}
	fmt.Println(resp.String())
	json.Unmarshal(resp.Body(), &token)
	return token, nil
}
