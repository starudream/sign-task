package api

import (
	"bytes"
	"slices"
	"strings"

	"github.com/starudream/go-lib/resty/v2"

	"github.com/starudream/sign-task/util"
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

	r.FormData.Set("sign", strings.ToUpper(util.MD5Hex(buf.String())))
	return r
}
