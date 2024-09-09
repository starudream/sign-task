package api

import (
	"testing"

	"github.com/starudream/go-lib/core/v2/utils/testutil"
)

func TestClient_QueryAccountBalance(t *testing.T) {
	data, err := C.QueryAccountBalance()
	testutil.LogNoErr(t, err, data)
}
