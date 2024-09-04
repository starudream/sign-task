package api

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"slices"
	"strings"

	"github.com/starudream/go-lib/resty/v2"
)

func addSign(r *resty.Request) *resty.Request {
	if len(r.FormData) == 0 {
		return r
	}

	keys := make([]string, 0)
	for key := range r.FormData {
		keys = append(keys, key)
	}
	slices.Sort(keys)

	buf := &bytes.Buffer{}
	for i := 0; i < len(keys); i++ {
		buf.WriteString(keys[i])
		buf.WriteByte('=')
		buf.WriteString(r.FormData[keys[i]][0])
	}
	buf.WriteString("tiebaclient!!!")

	b := md5.Sum(buf.Bytes())
	r.FormData.Set("sign", strings.ToUpper(hex.EncodeToString(b[:])))
	return r
}
