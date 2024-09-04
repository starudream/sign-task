package api

import (
	"net/http"
)

func (c *Client) cookie() []*http.Cookie {
	return []*http.Cookie{
		{Name: "BDUSS", Value: c.account.BDUSS},
	}
}
