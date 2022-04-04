package flyer

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func WxGetOpenid(code string) (string, error) {
	code2sessionURL := "https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code"
	appID := config.Wechat.AppID
	appSecret := config.Wechat.AppSecret
	url := fmt.Sprintf(code2sessionURL, appID, appSecret, code)
	resp, err := http.DefaultClient.Get(url)
	if err != nil {
		return "", err
	}
	var wxMap map[string]string
	err = json.NewDecoder(resp.Body).Decode(&wxMap)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	return wxMap["openid"], nil
}
