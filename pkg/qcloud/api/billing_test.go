package api

import (
	"testing"

	"github.com/starudream/go-lib/core/v2/utils/testutil"
)

func TestClient_DescribeAccountBalance(t *testing.T) {
	data, err := C.DescribeAccountBalance()
	testutil.LogNoErr(t, err, data)
}
