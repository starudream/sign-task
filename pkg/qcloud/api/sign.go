package api

import (
	"bytes"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"strconv"
	"strings"
	"time"

	"github.com/starudream/go-lib/core/v2/codec/json"
	"github.com/starudream/go-lib/resty/v2"

	"github.com/starudream/sign-task/pkg/qcloud/config"
	"github.com/starudream/sign-task/util"
)

func addSign(r *resty.Request, method, addr, path, action, version, region, service string, account config.Account) {
	host := strings.TrimSuffix(addr, "/")
	host = strings.TrimPrefix(host, "https://")
	host = strings.TrimPrefix(host, "http://")

	now := time.Now()
	date := now.Format("2006-01-02")
	timestamp := strconv.FormatInt(now.Unix(), 10)

	bodyHex := genBody(r.Body)

	r.SetHeader("Host", host)
	r.SetHeader("Content-Type", "application/json")
	r.SetHeader("x-tc-action", action)
	r.SetHeader("x-tc-region", region)
	r.SetHeader("x-tc-version", version)
	r.SetHeader("x-tc-timestamp", timestamp)

	queryStr := genQuery(r.QueryParam)
	headerStr, headerKeys := genHeader(r.Header)

	reqStr := strings.Join([]string{strings.ToUpper(method), path, queryStr, headerStr, headerKeys, bodyHex}, "\n")
	scopeStr := strings.Join([]string{date, service, "tc3_request"}, "/")
	signStr := strings.Join([]string{Algorithm, timestamp, scopeStr, util.SHA256Hex(reqStr)}, "\n")
	keyBs := util.HMAC256(util.HMAC256(util.HMAC256("TC3"+account.Key, date), service), "tc3_request")
	signature := strings.ToLower(util.HMAC256Hex(keyBs, signStr))

	r.SetHeader("Authorization", fmt.Sprintf("%s Credential=%s/%s, SignedHeaders=%s, Signature=%s", Algorithm, account.Id, scopeStr, headerKeys, signature))
}

func genBody(body any) string {
	s := ""
	if body != nil {
		s = json.MustMarshalString(body)
	}
	return util.SHA256Hex(s)
}

func genQuery(query url.Values) string {
	return query.Encode()
}

func genHeader(header http.Header) (string, string) {
	keys := make([]string, 0)
	for k := range header {
		keys = append(keys, k)
	}
	slices.Sort(keys)

	buf := &bytes.Buffer{}
	lks := make([]string, len(keys))
	for i := 0; i < len(keys); i++ {
		key := strings.ToLower(keys[i])
		buf.WriteString(key)
		buf.WriteByte(':')
		buf.WriteString(strings.ToLower(header.Get(key)))
		buf.WriteByte('\n')
		lks[i] = key
	}
	return buf.String(), strings.Join(lks, ";")
}
