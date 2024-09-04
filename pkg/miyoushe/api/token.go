package api

type GetCTokenBySTokenData struct {
	Uid         string `json:"uid"`
	CookieToken string `json:"cookie_token"`
}

// GetCTokenBySToken get cookie token by stoken v2
// https://github.com/UIGF-org/mihoyo-api-collect/blob/3a9116ea538941cfead749572df1f364cb9f9c8d/hoyolab/user/token.md#%E9%80%9A%E8%BF%87stoken%E8%8E%B7%E5%8F%96cookie-token
func (c *Client) GetCTokenBySToken() (*GetCTokenBySTokenData, error) {
	req := c.R().SetCookies(c.sToken())
	return Exec[*GetCTokenBySTokenData](req, "GET", AddrTakumi+"/auth/api/getCookieAccountInfoBySToken")
}

type GetLTokenBySTokenData struct {
	LToken string `json:"ltoken"`
}

// GetLTokenBySToken get ltoken v1 by stoken v2
// https://github.com/UIGF-org/mihoyo-api-collect/blob/3a9116ea538941cfead749572df1f364cb9f9c8d/hoyolab/user/token.md#%E9%80%9A%E8%BF%87stoken%E8%8E%B7%E5%8F%96ltokenv1
func (c *Client) GetLTokenBySToken() (*GetLTokenBySTokenData, error) {
	req := c.R().SetCookies(c.sToken())
	return Exec[*GetLTokenBySTokenData](req, "GET", AddrTakumi+"/account/auth/api/getLTokenBySToken")
}
