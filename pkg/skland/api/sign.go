package api

import (
	"net/url"
	"strconv"
	"time"

	"github.com/starudream/go-lib/core/v2/codec/json"
	"github.com/starudream/go-lib/resty/v2"

	"github.com/starudream/sign-task/pkg/skland/config"
	"github.com/starudream/sign-task/util"
)

type signHeaders struct {
	Platform  string `json:"platform"`
	Timestamp string `json:"timestamp"`
	DId       string `json:"dId"`
	VName     string `json:"vName"`
}

func addSign(r *resty.Request, method, path string, account config.Account) {
	ts := strconv.FormatInt(time.Now().Unix(), 10)

	// use struct to fix the order of headers
	headers := signHeaders{Platform: Platform, Timestamp: ts, DId: DId, VName: VName}

	r.SetHeaders(util.ToMap[string](headers))

	_, signature := sign(headers, method, path, account.Token, r.QueryParam, r.Body)

	r.SetHeader("cred", account.Cred)
	r.SetHeader("sign", signature)
}

func sign(headers signHeaders, method, path, token string, query url.Values, body any) (string, string) {
	str := query.Encode()
	if method != "GET" {
		str = json.MustMarshalString(body)
	}

	content := path + str + headers.Timestamp + json.MustMarshalString(headers)

	return content, util.MD5Hex(util.HMAC256Hex(token, content))
}
