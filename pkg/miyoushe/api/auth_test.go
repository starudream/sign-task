package api

import (
	"testing"

	"github.com/starudream/go-lib/core/v2/utils/testutil"
)

func TestClient_SendPhoneCode(t *testing.T) {
	data, err := C.SendPhoneCode("")
	testutil.LogNoErr(t, err, data)
}

func TestClient_LoginByPhoneCode(t *testing.T) {
	data, err := C.LoginByPhoneCode("")
	testutil.LogNoErr(t, err, data)
}
