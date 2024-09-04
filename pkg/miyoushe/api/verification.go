package api

import (
	"github.com/starudream/go-lib/core/v2/gh"

	"github.com/starudream/sign-task/pkg/geetest"
)

func (c *Client) CreateVerification() (*geetest.V3Param, error) {
	req := c.R().SetCookies(c.sToken()).SetQueryParam("is_high", "true")
	return Exec[*geetest.V3Param](req, "GET", AddrBBS+"/misc/api/createVerification")
}

type VerifyVerificationData struct {
	Challenge string `json:"challenge"`
}

func (c *Client) VerifyVerification(challenge, validate string) (*VerifyVerificationData, error) {
	body := gh.M{"geetest_challenge": challenge, "geetest_seccode": validate + "|jordan", "geetest_validate": validate}
	req := c.R().SetCookies(c.sToken()).SetBody(body)
	return Exec[*VerifyVerificationData](req, "POST", AddrBBS+"/misc/api/verifyVerification")
}
