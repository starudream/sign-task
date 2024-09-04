package api

import (
	"net/url"
	"strings"
)

type GetHomeData struct {
	Navigator []*HomeNav `json:"navigator"`
}

type HomeNav struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	AppPath string `json:"app_path"`
}

func (t *GetHomeData) GetSignActId() string {
	for _, nav := range t.Navigator {
		if strings.Contains(nav.Name, "签到") {
			u, err := url.Parse(nav.AppPath)
			if err == nil {
				return u.Query().Get("act_id")
			}
		}
	}
	return ""
}

func (c *Client) GetHome(gameId string) (*GetHomeData, error) {
	req := c.R().SetQueryParam("gids", gameId)
	return Exec[*GetHomeData](req, "GET", AddrBBS+"/apihub/api/home/new")
}

type GetBusinessesData struct {
	Businesses []string `json:"businesses"`
}

func (c *Client) GetBusinesses() (*GetBusinessesData, error) {
	req := c.R().SetCookies(c.sToken()).SetQueryParam("uid", c.account.Uid)
	return Exec[*GetBusinessesData](req, "GET", AddrBBS+"/user/api/getUserBusinesses")
}
