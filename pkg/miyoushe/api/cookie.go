package api

import (
	"net/http"
)

// sToken generate base stoken auth cookies
func (c *Client) sToken() []*http.Cookie {
	return []*http.Cookie{
		{Name: "mid", Value: c.account.Mid},
		{Name: "stoken", Value: c.account.SToken},
	}
}

// cToken generate base cookie_token auth cookies
func (c *Client) cToken() []*http.Cookie {
	return []*http.Cookie{
		{Name: "account_id", Value: c.account.Uid},
		{Name: "cookie_token", Value: c.account.CToken},
	}
}
