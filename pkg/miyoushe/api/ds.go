package api

import (
	"fmt"
	"math/rand/v2"
	"time"

	"github.com/starudream/go-lib/core/v2/codec/json"
	"github.com/starudream/go-lib/resty/v2"

	"github.com/starudream/sign-task/util"
)

type SignDS int

const (
	SignDS1 SignDS = iota
	SignDS2
)

const (
	// AppSaltK2 https://blog.starudream.cn/2023/11/09/miyoushe-salt-2.62.2/
	AppSaltK2 = "WGtruoQrwczmsjLOPXzJLnaAYycsLavx" // 2.73.1
	// AppSalt6X https://github.com/UIGF-org/mihoyo-api-collect/issues/1
	AppSalt6X = "t0qEgfub6cvueAPgR5m9aQWWVciEer7v"
)

// DS1 https://github.com/UIGF-org/mihoyo-api-collect/blob/3a9116ea538941cfead749572df1f364cb9f9c8d/other/authentication.md#ds
func DS1() string {
	_, s := ds1(time.Now().Unix(), util.RandAlphanum(6))
	return s
}

func ds1(t int64, r string) (string, string) {
	s := fmt.Sprintf("salt=%s&t=%d&r=%s", AppSaltK2, t, r)
	return s, fmt.Sprintf("%d,%s,%s", t, r, util.MD5Hex(s))
}

// DS2 https://github.com/UIGF-org/mihoyo-api-collect/blob/3a9116ea538941cfead749572df1f364cb9f9c8d/other/authentication.md#ds
func DS2(q *resty.Request) string {
	_, s := ds2(time.Now().Unix(), 100000+rand.IntN(100000), json.MustMarshalString(q.Body), q.QueryParam.Encode())
	return s
}

func ds2(t int64, r int, body, query string) (string, string) {
	if r == 100000 {
		r += 542367
	}
	s := fmt.Sprintf("salt=%s&t=%d&r=%d&b=%s&q=%s", AppSalt6X, t, r, body, query)
	return s, fmt.Sprintf("%d,%d,%s", t, r, util.MD5Hex(s))
}
