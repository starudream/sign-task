package api

import (
	"fmt"
	"time"

	"github.com/starudream/go-lib/core/v2/config/version"
	"github.com/starudream/go-lib/resty/v2"

	"github.com/starudream/sign-task/pkg/volcengine/config"
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

type baseResp[T any] struct {
	ResponseMetadata respMetadata `json:"ResponseMetadata"`
	Result           T            `json:"Result,omitempty"`
}

type respMetadata struct {
	RequestId string    `json:"RequestId"`
	Action    string    `json:"Action"`
	Version   string    `json:"Version"`
	Service   string    `json:"Service"`
	Region    string    `json:"Region,omitempty"`
	Error     respError `json:"Error,omitempty"`
}

type respError struct {
	CodeN   int    `json:"CodeN"`
	Code    string `json:"Code"`
	Message string `json:"Message"`
}

func (t *baseResp[T]) IsSuccess() bool {
	return t != nil && t.ResponseMetadata.Error.CodeN == 0
}

func (t *baseResp[T]) String() string {
	if t == nil {
		return "<nil>"
	}
	e := t.ResponseMetadata.Error
	return fmt.Sprintf("code: %s, message: %s", e.Code, e.Message)
}

func Exec[T any](r *resty.Request, method, addr, path, action, version, region, service string, args ...any) (t T, _ error) {
	for i := 0; i < len(args); i++ {
		switch arg := args[i].(type) {
		case config.Account:
			addSign(r, method, addr, path, action, version, region, service, arg)
		}
	}
	res, err := resty.ParseResp[*baseResp[any], *baseResp[T]](
		r.SetError(&baseResp[any]{}).SetResult(&baseResp[T]{}).Execute(method, addr+path),
	)
	if err != nil {
		return t, fmt.Errorf("[volcengine] %w", err)
	}
	return res.Result, nil
}
