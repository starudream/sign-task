package api

import (
	"testing"

	"github.com/starudream/go-lib/core/v2/utils/testutil"
)

func TestClient_SendPhoneCode(t *testing.T) {
	err := C.SendPhoneCode()
	testutil.LogNoErr(t, err)
}

func TestClient_LoginByPhoneCode(t *testing.T) {
	data, err := C.LoginByPhoneCode("")
	testutil.LogNoErr(t, err, data)
}

func TestClient_GrantApp(t *testing.T) {
	data, err := C.GrantApp("", AppCodeSkland)
	testutil.LogNoErr(t, err, data)
}

func TestClient_AuthLoginByCode(t *testing.T) {
	data, err := C.AuthLoginByCode("")
	testutil.LogNoErr(t, err, data)
}

func TestClient_AuthRefresh(t *testing.T) {
	data, err := C.AuthRefresh("")
	testutil.LogNoErr(t, err, data)
}
