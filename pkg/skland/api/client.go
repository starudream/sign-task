package api

import (
	"fmt"
	"strings"
	"time"

	"github.com/starudream/go-lib/resty/v2"

	"github.com/starudream/sign-task/pkg/skland/config"
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
		SetHeader("User-Agent", UserAgent)
	return c
}

func (c *Client) R() *resty.Request {
	return c.client.R()
}

type baseResp[T any] struct {
	// hypergryph
	Status *int   `json:"status"`
	Msg    string `json:"msg"`

	// skland
	Code    *int   `json:"code"`
	Message string `json:"message"`

	Data T `json:"data,omitempty"`
}

func (t *baseResp[T]) GetCode() int {
	if t == nil || t.Code == nil {
		return 999999
	}
	return *t.Code
}

func (t *baseResp[T]) IsSuccess() bool {
	return t != nil && ((t.Status != nil && *t.Status == 0) || (t.Code != nil && *t.Code == 0))
}

func (t *baseResp[T]) String() string {
	if t == nil || (t.Status == nil && t.Code == nil) {
		return "<nil>"
	}
	if t.Status != nil {
		return fmt.Sprintf("status: %d, msg: %s, data: %v", *t.Status, t.Msg, t.Data)
	}
	return fmt.Sprintf("code: %d, message: %s, data: %v", *t.Code, t.Message, t.Data)
}

func Exec[T any](r *resty.Request, method, url string, args ...any) (t T, _ error) {
	for i := 0; i < len(args); i++ {
		switch arg := args[i].(type) {
		case config.Account:
			addSign(r, method, strings.TrimPrefix(url, AddrZonai), arg)
		}
	}
	res, err := resty.ParseResp[*baseResp[any], *baseResp[T]](
		r.SetError(&baseResp[any]{}).SetResult(&baseResp[T]{}).Execute(method, url),
	)
	if err != nil {
		return t, fmt.Errorf("[skland] %w", err)
	}
	return res.Data, nil
}

func IsCode(err error, code int) bool {
	if err == nil {
		return false
	}
	v1, ok1 := resty.AsRespErr(err)
	if ok1 {
		v2, ok2 := v1.Result().(interface{ GetCode() int })
		v3, ok3 := v1.Response.Error().(interface{ GetCode() int })
		return (ok2 && v2.GetCode() == code) || (ok3 && v3.GetCode() == code)
	}
	return false
}
