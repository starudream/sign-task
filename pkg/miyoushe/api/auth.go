package api

import (
	"github.com/starudream/go-lib/core/v2/gh"
)

const phoneAreaCodeCN = "+86"

type SendPhoneCodeData struct {
	SentNew    bool   `json:"sent_new"`
	Countdown  int    `json:"countdown"`
	ActionType string `json:"action_type"`
}

func (c *Client) SendPhoneCode(aigis string) (*SendPhoneCodeData, error) {
	req := c.R().SetBody(gh.M{
		"area_code": RSAEncrypt(phoneAreaCodeCN),
		"mobile":    RSAEncrypt(c.account.Phone),
	})
	if aigis != "" {
		req.SetHeader(xRpcAigis, aigis)
	}
	return Exec[*SendPhoneCodeData](req, "POST", AddrPassport+"/account/ma-cn-verifier/verifier/createLoginCaptcha")
}

type LoginByPhoneCodeData struct {
	Token    *TokenInfo     `json:"token"`
	UserInfo *LoginUserInfo `json:"user_info"`
}

type TokenInfo struct {
	TokenType int    `json:"token_type"` // 1
	Token     string `json:"token"`      // stoken_v2
}

type LoginUserInfo struct {
	Aid string `json:"aid"` // uid
	Mid string `json:"mid"`
}

const actionTypeLoginByMobileCaptcha = "login_by_mobile_captcha"

func (c *Client) LoginByPhoneCode(code string) (*LoginByPhoneCodeData, error) {
	req := c.R().SetBody(gh.M{
		"area_code":   RSAEncrypt(phoneAreaCodeCN),
		"mobile":      RSAEncrypt(c.account.Phone),
		"action_type": actionTypeLoginByMobileCaptcha,
		"captcha":     code,
	})
	return Exec[*LoginByPhoneCodeData](req, "POST", AddrPassport+"/account/ma-cn-passport/app/loginByMobileCaptcha")
}
