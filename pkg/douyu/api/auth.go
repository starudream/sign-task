package api

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/starudream/go-lib/core/v2/gh"
)

func (c *Client) Refresh() error {
	t := strconv.FormatInt(time.Now().UnixMilli(), 10)
	query := gh.MS{"client_id": "1", "t": t, "_": t, "callback": "axiosJsonpCallback"}
	resp, err := c.R().SetCookies(c.xCookie()).SetQueryParams(query).Get("https://passport.douyu.com/lapi/passport/iframe/safeAuth")
	if err != nil {
		return fmt.Errorf("[douyu] refresh error: %w", err)
	}

	if code := resp.StatusCode(); code != http.StatusOK {
		return fmt.Errorf("[douyu] http status code %d not 200", code)
	}

	for _, cookie := range resp.Cookies() {
		if cookie.Name == "acf_uid" {
			c.Uid = cookie.Value
		}
		if cookie.Name == "acf_auth" {
			c.Auth = cookie.Value
		}
		if cookie.Name == "acf_stk" {
			c.Stk = cookie.Value
		}
		if cookie.Name == "acf_ltkid" {
			c.Ltkid = cookie.Value
		}
		if cookie.Name == "acf_username" {
			c.Username = cookie.Value
		}
	}

	if c.Uid == "" || c.Auth == "" || c.Stk == "" || c.Ltkid == "" || c.Username == "" {
		return fmt.Errorf("[douyu] cookies not found")
	}

	return nil
}
