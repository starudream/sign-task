package api

import (
	"testing"

	"github.com/starudream/go-lib/core/v2/utils/testutil"
)

func TestClient_QueryBalanceAcct(t *testing.T) {
	data, err := C.QueryBalanceAcct()
	testutil.LogNoErr(t, err, data)
}

func TestClient_ListBillDetail(t *testing.T) {
	data, err := C.ListBillDetail(&ListBillDetailReq{
		BillPeriod:    "2025-05",
		ExpenseDate:   "2025-05-16",
		GroupPeriod:   1,
		GroupTerm:     0,
		IgnoreZero:    1,
		NeedRecordNum: 1,
	})
	testutil.LogNoErr(t, err, data)
}
