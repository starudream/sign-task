package api

import (
	"github.com/starudream/go-lib/core/v2/gh"
)

func (c *Client) SendPhoneCode() error {
	req := c.R().SetBody(gh.M{"type": 2, "phone": c.account.Phone})
	_, err := Exec[any](req, "POST", AddrHypergryph+"/general/v1/send_phone_code")
	return err
}

type LoginByPhoneCodeData struct {
	Token string `json:"token"`
}

func (c *Client) LoginByPhoneCode(code string) (*LoginByPhoneCodeData, error) {
	req := c.R().SetBody(gh.M{"phone": c.account.Phone, "code": code})
	return Exec[*LoginByPhoneCodeData](req, "POST", AddrHypergryph+"/user/auth/v2/token_by_phone_code")
}

type GrantAppData struct {
	Uid  string `json:"uid"`
	Code string `json:"code"`
}

func (c *Client) GrantApp(token string, code string) (*GrantAppData, error) {
	req := c.R().SetBody(gh.M{"type": 0, "token": token, "appCode": code})
	return Exec[*GrantAppData](req, "POST", AddrHypergryph+"/user/oauth2/v2/grant")
}

type GenCredByCodeData struct {
	UserId string `json:"userId"`
	Cred   string `json:"cred"`
	Token  string `json:"token"`
}

func (c *Client) AuthLoginByCode(code string) (*GenCredByCodeData, error) {
	req := c.R().SetBody(gh.M{"kind": 1, "code": code})
	return Exec[*GenCredByCodeData](req, "POST", AddrZonai+"/api/v1/user/auth/generate_cred_by_code")
}

type AuthRefreshData struct {
	Token string `json:"token"`
}

func (c *Client) AuthRefresh(cred string) (*AuthRefreshData, error) {
	req := c.R().SetHeader("cred", cred)
	return Exec[*AuthRefreshData](req, "GET", AddrZonai+"/api/v1/auth/refresh")
}
