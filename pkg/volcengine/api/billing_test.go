package api

import (
	"testing"

	"github.com/starudream/go-lib/core/v2/utils/testutil"
)

func TestClient_QueryBalanceAcct(t *testing.T) {
	data, err := C.QueryBalanceAcct()
	testutil.LogNoErr(t, err, data)
}
