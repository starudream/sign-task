package api

import (
	"fmt"
	"time"

	"github.com/starudream/go-lib/resty/v2"

	"github.com/starudream/sign-task/pkg/douyu/config"
)

type Client struct {
	client  *resty.Client
	account config.Account

	Uid      string // cookie: acf_uid
	Auth     string // cookie: acf_auth
	Stk      string // cookie: acf_stk
	Ltkid    string // cookie: acf_ltkid
	Username string // cookie: acf_username
}

func NewClient(account config.Account) *Client {
	c := &Client{
		account: account,
	}
	c.client = resty.New().
		SetTimeout(30*time.Second).
		SetHeader("Accept-Encoding", "gzip").
		SetHeader("User-Agent", resty.UAWindowsChrome).
		SetHeader("Referer", Addr)
	return c
}

func (c *Client) R() *resty.Request {
	return c.client.R()
}

type baseResp[T any] struct {
	Error *int   `json:"error"`
	Msg   string `json:"msg"`
	Data  T      `json:"data"`
}

func (t *baseResp[T]) IsSuccess() bool {
	return t != nil && t.Error != nil && *t.Error == 0
}

func (t *baseResp[T]) String() string {
	if t == nil || t.Error == nil {
		return "<nil>"
	}
	return fmt.Sprintf("error: %d, msg: %s, data: %v", *t.Error, t.Msg, t.Data)
}

func Exec[T any](r *resty.Request, method, path string) (t T, _ error) {
	res, err := resty.ParseResp[*baseResp[any], *baseResp[T]](
		r.SetError(&baseResp[any]{}).SetResult(&baseResp[T]{}).Execute(method, Addr+path),
	)
	if err != nil {
		return t, fmt.Errorf("[douyu] %w", err)
	}
	return res.Data, nil
}
