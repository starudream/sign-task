package api

import (
	"fmt"
	"time"

	"github.com/starudream/go-lib/core/v2/codec/json"
	"github.com/starudream/go-lib/resty/v2"

	"github.com/starudream/sign-task/pkg/geetest"
	"github.com/starudream/sign-task/pkg/miyoushe/config"
)

type Client struct {
	client  *resty.Client
	account config.Account
}

func NewClient(account config.Account) *Client {
	c := &Client{
		account: account,
	}
	// https://github.com/UIGF-org/mihoyo-api-collect/blob/3a9116ea538941cfead749572df1f364cb9f9c8d/other/authentication.md#%E8%AF%B7%E6%B1%82%E5%A4%B4
	c.client = resty.New().
		SetTimeout(30*time.Second).
		SetHeader("Accept-Encoding", "gzip").
		SetHeader("User-Agent", UserAgent).
		SetHeader("Referer", RefererApp).
		SetHeader("x-rpc-app_version", AppVersion).
		SetHeader("x-rpc-app_id", AppIdMiyoushe).
		SetHeader("x-rpc-verify_key", AppIdMiyoushe).
		SetHeader("x-rpc-device_id", account.Device.Id).
		SetHeader("x-rpc-client_type", account.Device.Type).
		SetHeader("x-rpc-device_name", account.Device.Name).
		SetHeader("x-rpc-device_model", account.Device.Model).
		SetHeader("x-rpc-sys_version", account.Device.Version).
		SetHeader("x-rpc-channel", account.Device.Channel)
	return c
}

func (c *Client) R() *resty.Request {
	return c.client.R()
}

type baseResp[T any] struct {
	RetCode *int   `json:"retcode"`
	Message string `json:"message"`
	Data    T      `json:"data"`
}

func (t *baseResp[T]) GetRetCode() int {
	if t == nil || t.RetCode == nil {
		return 999999
	}
	return *t.RetCode
}

func (t *baseResp[T]) IsSuccess() bool {
	return t != nil && t.RetCode != nil && *t.RetCode == 0
}

func (t *baseResp[T]) String() string {
	if t == nil || t.RetCode == nil {
		return "<nil>"
	}
	return fmt.Sprintf("retcode: %d, message: %s, data: %v", *t.RetCode, t.Message, t.Data)
}

func Exec[T any](r *resty.Request, method, url string, args ...any) (t T, _ error) {
	ds := SignDS1
	for i := 0; i < len(args); i++ {
		switch arg := args[i].(type) {
		case SignDS:
			ds = arg
		case *geetest.V3Data:
			if arg == nil {
				continue
			}
			r.SetHeader("x-rpc-challenge", arg.Challenge)
			r.SetHeader("x-rpc-validate", arg.Validate)
		}
	}
	if ds == SignDS2 {
		r.SetHeader("DS", DS2(r))
	} else {
		r.SetHeader("DS", DS1())
	}
	res, err := resty.ParseResp[*baseResp[any], *baseResp[T]](
		r.SetError(&baseResp[any]{}).SetResult(&baseResp[T]{}).Execute(method, url),
	)
	if err != nil {
		return t, fmt.Errorf("[miyoushe] %w", err)
	}
	return res.Data, nil
}

func IsRetCode(err error, rc int) bool {
	if err == nil {
		return false
	}
	v1, ok1 := resty.AsRespErr(err)
	if ok1 {
		v2, ok2 := v1.Result().(interface{ GetRetCode() int })
		return ok2 && v2.GetRetCode() == rc
	}
	return false
}

type Aigis struct {
	SessionId string `json:"session_id"`
	MmtType   int    `json:"mmt_type"`
	Data      string `json:"data"`
}

func GetAigisData(err error) (string, *geetest.V3Param) {
	if err != nil {
		v1, ok1 := resty.AsRespErr(err)
		if ok1 {
			str := v1.Response.Header().Get(xRpcAigis)
			if str != "" {
				aigis, e1 := json.UnmarshalTo[*Aigis](str)
				if e1 == nil {
					param, e2 := json.UnmarshalTo[*geetest.V3Param](aigis.Data)
					if e2 == nil {
						return aigis.SessionId, param
					}
				}
			}
		}
	}
	return "", nil
}
