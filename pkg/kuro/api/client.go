package api

import (
	"fmt"
	"time"

	"github.com/starudream/go-lib/resty/v2"

	"github.com/starudream/sign-task/pkg/kuro/config"
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
		SetHeader("User-Agent", UserAgent).
		SetHeader("X-Requested-With", AndroidName).
		SetHeader("Source", SourceAndroid).
		SetHeader("Version", Version).
		SetHeader("DevCode", account.DevCode).
		SetHeader("Token", account.Token)
	return c
}

func (c *Client) R() *resty.Request {
	return c.client.R()
}

type baseResp[T any] struct {
	Code Code   `json:"code"`
	Msg  string `json:"msg"`
	Data T      `json:"data,omitempty"`
}

func (t *baseResp[T]) GetCode() Code {
	if t == nil {
		return 999999
	}
	return t.Code
}

func (t *baseResp[T]) IsSuccess() bool {
	return t != nil && t.Code == 200
}

func (t *baseResp[T]) String() string {
	return fmt.Sprintf("code: %d, message: %s, data: %v", t.Code, t.Msg, t.Data)
}

func Exec[T any](r *resty.Request, method, path string) (t T, _ error) {
	res, err := resty.ParseResp[*baseResp[any], *baseResp[T]](
		r.SetError(&baseResp[any]{}).SetResult(&baseResp[T]{}).Execute(method, Addr+path),
	)
	if err != nil {
		return t, fmt.Errorf("[kuro] %w", err)
	}
	return res.Data, nil
}

func IsCode(err error, code Code) bool {
	if err == nil {
		return false
	}
	v1, ok1 := resty.AsRespErr(err)
	if ok1 {
		v2, ok2 := v1.Result().(interface{ GetCode() Code })
		return ok2 && v2.GetCode() == code
	}
	return false
}
