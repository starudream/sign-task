package api

import (
	"fmt"
)

type TBSData struct {
	TBS     string `json:"tbs"`
	IsLogin int    `json:"is_login"`
}

func (t *TBSData) IsSuccess() bool {
	return t != nil && t.TBS != ""
}

func (t *TBSData) String() string {
	if t == nil {
		return "<nil>"
	}
	return fmt.Sprintf("login: %b, tbs: %s", t.IsLogin, t.TBS)
}

func (t *TBSData) Value() string {
	if t == nil {
		return ""
	}
	return t.TBS
}

func (c *Client) TBS() (string, error) {
	r := c.R().SetCookies(c.cookie())
	data, err := Exec[*TBSData](r, "GET", Addr+"/dc/common/tbs")
	return data.Value(), err
}
