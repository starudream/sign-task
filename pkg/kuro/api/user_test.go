package api

import (
	"testing"

	"github.com/starudream/go-lib/core/v2/utils/testutil"
)

func TestClient_GetUser(t *testing.T) {
	data, err := C.GetUser()
	testutil.LogNoErr(t, err, data)
}
