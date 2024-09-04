package api

import (
	"github.com/starudream/go-lib/core/v2/codec/json"
	"github.com/starudream/go-lib/core/v2/gh"

	"github.com/starudream/sign-task/pkg/geetest"
)

type SendPhoneCodeData struct {
	GeeTest bool `json:"geeTest"`
}

func (c *Client) SendPhoneCode(gt *geetest.V4Data) (*SendPhoneCodeData, error) {
	data := "{}"
	if gt != nil {
		data = json.MustMarshalString(gt)
	}
	req := c.R().SetFormData(gh.MS{"mobile": c.account.Phone, "geeTestData": data})
	return Exec[*SendPhoneCodeData](req, "POST", "/user/getSmsCode")
}

type LoginByPhoneCodeData struct {
	UserId   string `json:"userId"`
	UserName string `json:"userName"`
	Token    string `json:"token"`
}

func (c *Client) LoginByPhoneCode(code string) (*LoginByPhoneCodeData, error) {
	req := c.R().SetFormData(gh.MS{"mobile": c.account.Phone, "code": code, "devCode": c.account.DevCode})
	return Exec[*LoginByPhoneCodeData](req, "POST", "/user/sdkLogin")
}
