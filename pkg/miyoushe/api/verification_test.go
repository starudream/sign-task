package api

import (
	"testing"

	"github.com/starudream/go-lib/core/v2/utils/testutil"
)

func TestClient_CreateVerification(t *testing.T) {
	data, err := C.CreateVerification()
	testutil.LogNoErr(t, err, data)
}

func TestClient_VerifyVerification(t *testing.T) {
	data, err := C.VerifyVerification("", "")
	testutil.LogNoErr(t, err, data)
}
