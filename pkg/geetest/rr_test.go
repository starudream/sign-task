package geetest_test

import (
	"testing"

	"github.com/starudream/go-lib/core/v2/utils/testutil"

	"github.com/starudream/sign-task/pkg/geetest"
)

func TestRRPoint(t *testing.T) {
	resp, err := geetest.RRPoint(&geetest.V3Param{})
	testutil.LogNoErr(t, err, resp)
}

func TestRR(t *testing.T) {
	resp, err := geetest.RR(&geetest.V3Param{
		GT:        "",
		Challenge: "",
		Referer:   "https://act.mihoyo.com",
	})
	testutil.LogNoErr(t, err, resp)
}
