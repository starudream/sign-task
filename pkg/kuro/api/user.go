package api

import (
	"github.com/starudream/go-lib/core/v2/gh"
)

type User struct {
	Mine Mine `json:"mine"`
}

type Mine struct {
	UserId   string `json:"userId"`
	UserName string `json:"userName"`
}

func (c *Client) GetUser() (*User, error) {
	req := c.R().SetFormData(gh.MS{"type": "1", "searchType": "2"})
	return Exec[*User](req, "POST", "/user/mine")
}
