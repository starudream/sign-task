package api

import (
	"bytes"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"strings"
	"time"

	"github.com/starudream/go-lib/core/v2/codec/json"
	"github.com/starudream/go-lib/resty/v2"

	"github.com/starudream/sign-task/pkg/aliyun/config"
	"github.com/starudream/sign-task/util"
)

func addSign(r *resty.Request, method, addr, path, action, version string, account config.Account) {
	host := strings.TrimSuffix(addr, "/")
	host = strings.TrimPrefix(host, "https://")
	host = strings.TrimPrefix(host, "http://")

	bodyHex := genBody(r.Body)

	r.SetHeader("Host", host)
	r.SetHeader("x-acs-action", action)
	r.SetHeader("x-acs-version", version)
	r.SetHeader("x-acs-date", time.Now().UTC().Format(time.RFC3339))
	r.SetHeader("x-acs-signature-nonce", util.UUID())
	r.SetHeader("x-acs-content-sha256", bodyHex)

	queryStr := genQuery(r.QueryParam)
	headerStr, headerKeys := genHeader(r.Header)

	signStr := strings.Join([]string{strings.ToUpper(method), path, queryStr, headerStr, headerKeys, bodyHex}, "\n")
	signature := strings.ToLower(util.HMAC256Hex(account.Secret, Algorithm+"\n"+util.SHA256Hex(signStr)))

	r.SetHeader("Authorization", fmt.Sprintf("%s Credential=%s,SignedHeaders=%s,Signature=%s", Algorithm, account.Id, headerKeys, signature))
}

func genBody(body any) string {
	s := ""
	if body != nil {
		s = json.MustMarshalString(body)
	}
	return util.SHA256Hex(s)
}

func genQuery(query url.Values) string {
	keys := make([]string, 0)
	for k := range query {
		keys = append(keys, k)
	}
	slices.Sort(keys)

	buf := &bytes.Buffer{}
	for i := 0; i < len(keys); i++ {
		key := keys[i]
		if buf.Len() > 0 {
			buf.WriteByte('&')
		}
		buf.WriteString(escape(key))
		buf.WriteByte('=')
		buf.WriteString(escape(query.Get(key)))
	}
	return buf.String()
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
		buf.WriteString(header.Get(key))
		buf.WriteByte('\n')
		lks[i] = key
	}
	return buf.String(), strings.Join(lks, ";")
}

var replacer = strings.NewReplacer("+", "%20", "*", "%2A", "%7E", "~")

func escape(s string) string {
	return replacer.Replace(url.QueryEscape(s))
}
