package api

import (
	"fmt"
	"time"

	"github.com/starudream/go-lib/core/v2/config/version"
	"github.com/starudream/go-lib/resty/v2"

	"github.com/starudream/sign-task/pkg/qcloud/config"
)

type Client struct {
	client  *resty.Client
	account config.Account
}

func NewClient(account config.Account) *Client {
	c := &Client{
		account: account,
	}
	c.client = resty.New().
		SetTimeout(30*time.Second).
		SetHeader("Accept-Encoding", "gzip").
		SetHeader("User-Agent", "sign-task/"+version.GetVersionInfo().GitVersion)
	return c
}

func (c *Client) R() *resty.Request {
	return c.client.R()
}

type baseResp[T respInt] struct {
	Response T `json:"Response"`
}

type respInt interface {
	code() string
	message() string
}

type respError struct {
	RequestId string `json:"RequestId"`
	Error     struct {
		Code    string `json:"Code"`
		Message string `json:"Message"`
	} `json:"Error"`
}

var _ respInt = respError{}

func (e respError) code() string    { return e.Error.Code }
func (e respError) message() string { return e.Error.Message }

func (t *baseResp[T]) IsSuccess() bool {
	return t != nil && t.Response.code() == ""
}

func (t *baseResp[T]) String() string {
	if t == nil {
		return "<nil>"
	}
	return fmt.Sprintf("code: %s, message: %s", t.Response.code(), t.Response.message())
}

func Exec[T respInt](r *resty.Request, method, addr, path, action, version, region, service string, args ...any) (t T, _ error) {
	for i := 0; i < len(args); i++ {
		switch arg := args[i].(type) {
		case config.Account:
			addSign(r, method, addr, path, action, version, region, service, arg)
		}
	}
	res, err := resty.ParseResp[*baseResp[respInt], *baseResp[T]](
		r.SetError(&baseResp[respInt]{}).SetResult(&baseResp[T]{}).Execute(method, addr+path),
	)
	if err != nil {
		return t, fmt.Errorf("[qcloud] %w", err)
	}
	return res.Response, nil
}
