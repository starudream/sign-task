package api

import (
	"testing"

	"github.com/starudream/go-lib/core/v2/utils/testutil"
)

func TestClient_GetCTokenBySToken(t *testing.T) {
	data, err := C.GetCTokenBySToken()
	testutil.LogNoErr(t, err, data)
}

func TestClient_GetLTokenBySToken(t *testing.T) {
	data, err := C.GetLTokenBySToken()
	testutil.LogNoErr(t, err, data)
}
