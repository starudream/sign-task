package api

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/starudream/go-lib/core/v2/codec/json"
	"github.com/starudream/go-lib/resty/v2"

	"github.com/starudream/sign-task/pkg/tieba/config"
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

type baseResp struct {
	ErrorCode string `json:"error_code"`
	ErrorMsg  string `json:"error_msg"`
}

func (t *baseResp) GetCode() string {
	if t == nil {
		return ""
	}
	return t.ErrorCode
}

func (t *baseResp) IsSuccess() bool {
	return t != nil && t.ErrorCode == "0"
}

func (t *baseResp) String() string {
	if t == nil {
		return "<nil>"
	}
	return fmt.Sprintf("code: %s, msg: %s", t.ErrorCode, t.ErrorMsg)
}

type iRespRes interface {
	IsSuccess() bool
	String() string
}

func Exec[T iRespRes](r *resty.Request, method, url string) (t T, _ error) {
	resp, err := r.Execute(method, url)
	if err != nil {
		return t, fmt.Errorf("[tieba] execute error: %w", err)
	}
	if code := resp.StatusCode(); code != http.StatusOK {
		return t, fmt.Errorf("[tieba] http status code %d not 200", code)
	}
	t, err = json.UnmarshalTo[T](resp.Body())
	if err != nil {
		return t, fmt.Errorf("[tieba] unmarshal error: %w", err)
	}
	if !t.IsSuccess() {
		return t, fmt.Errorf("[tieba] %w", &iRespErr{Response: resp, esg: t.String(), data: t})
	}
	return t, nil
}

type iRespErr struct {
	*resty.Response
	esg  string
	data any
}

func (e *iRespErr) String() string {
	return fmt.Sprintf("response status: %s, error: %s", e.Response.Status(), e.esg)
}

func (e *iRespErr) Error() string {
	return e.String()
}

func IsCode(err error, code string) bool {
	if err == nil {
		return false
	}
	if re := new(iRespErr); errors.As(err, &re) {
		v1, ok1 := re.data.(interface{ GetCode() string })
		return ok1 && v1.GetCode() == code
	}
	return false
}
