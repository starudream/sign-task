package api

import (
	"net/http"
)

func (c *Client) xCookie() []*http.Cookie {
	return []*http.Cookie{
		{Name: "dy_did", Value: c.account.Did},
		{Name: "LTP0", Value: c.account.Ltp0},
	}
}

func (c *Client) yCookie() []*http.Cookie {
	return []*http.Cookie{
		{Name: "dy_did", Value: c.account.Did},
		{Name: "acf_uid", Value: c.Uid},
		{Name: "acf_auth", Value: c.Auth},
	}
}
