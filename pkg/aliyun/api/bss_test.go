package api

import (
	"testing"

	"github.com/starudream/go-lib/core/v2/utils/testutil"
)

func TestClient_QueryAccountBalance(t *testing.T) {
	data, err := C.QueryAccountBalance()
	testutil.LogNoErr(t, err, data)
}

func TestClient_QueryAccountBill(t *testing.T) {
	data, err := C.QueryAccountBill(&QueryAccountBillReq{
		BillingCycle:     "2025-05",
		IsGroupByProduct: true,
		Granularity:      "DAILY",
		BillingDate:      "2025-05-30",
	})
	testutil.LogNoErr(t, err, data)
}
