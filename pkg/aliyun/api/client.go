package api

import (
	"fmt"
	"time"

	"github.com/starudream/go-lib/core/v2/utils/reflectutil"
	"github.com/starudream/go-lib/resty/v2"

	"github.com/starudream/sign-task/pkg/aliyun/config"
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
		SetHeader("User-Agent", resty.UAWindowsChrome)
	return c
}

func (c *Client) R() *resty.Request {
	return c.client.R()
}

type baseResp[T any] struct {
	RequestId string `json:"RequestId"`
	Success   bool   `json:"Success"`
	Code      string `json:"Code"`
	Message   string `json:"Message"`
	Recommend string `json:"Recommend,omitempty"`
	Data      T      `json:"Data,omitempty"`
	Result    T      `json:"Result,omitempty"`
}

func (t *baseResp[T]) IsSuccess() bool {
	return t != nil
}

func (t *baseResp[T]) String() string {
	if t == nil {
		return "<nil>"
	}
	return fmt.Sprintf("code: %s, message: %s, recommend: %s", t.Code, t.Message, t.Recommend)
}

func Exec[T any](r *resty.Request, method, addr, path, action, version string, args ...any) (t T, _ error) {
	for i := 0; i < len(args); i++ {
		switch arg := args[i].(type) {
		case config.Account:
			addSign(r, method, addr, path, action, version, arg)
		}
	}
	res, err := resty.ParseResp[*baseResp[any], *baseResp[T]](
		r.SetError(&baseResp[any]{}).SetResult(&baseResp[T]{}).Execute(method, addr+path),
	)
	if err != nil {
		return t, fmt.Errorf("[aliyun] %w", err)
	}
	if !reflectutil.IsNil(res.Result) {
		return res.Result, nil
	}
	return res.Data, nil
}
