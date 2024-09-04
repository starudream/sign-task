package api

type GetUserData struct {
	UserInfo *UserInfo `json:"user_info"`
}

type UserInfo struct {
	Uid      string `json:"uid"`
	Nickname string `json:"nickname"`
}

func (c *Client) GetUser(uid string) (*GetUserData, error) {
	req := c.R().SetCookies(c.sToken()).SetQueryParam("uid", uid)
	return Exec[*GetUserData](req, "GET", AddrBBS+"/user/api/getUserFullInfo")
}
