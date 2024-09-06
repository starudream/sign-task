package geetest_test

import (
	"testing"

	"github.com/starudream/go-lib/core/v2/utils/testutil"

	"github.com/starudream/sign-task/pkg/geetest"
)

func TestTTPoint(t *testing.T) {
	resp, err := geetest.TTPoint(&geetest.V3Param{})
	testutil.LogNoErr(t, err, resp)
}

func TestTT(t *testing.T) {
	resp, err := geetest.TT(&geetest.V3Param{
		GT:        "",
		Challenge: "",
	})
	testutil.LogNoErr(t, err, resp)
}
