package api

import (
	"testing"

	"github.com/starudream/go-lib/core/v2/utils/testutil"

	"github.com/starudream/sign-task/util"
)

func TestClient_ListArkEndpoints(t *testing.T) {
	data, err := C.ListArkEndpoints(&ListArkEndpointsReq{
		PageNumber: 1,
		PageSize:   100,
	})
	testutil.LogNoErr(t, err, data)
}

func TestClient_GetArkUsage(t *testing.T) {
	data, err := C.GetArkUsage(&GetArkUsageReq{
		StartTime: int(util.GetToday(-1).Unix()),
		EndTime:   int(util.GetToday().Unix() - 1),
		Interval:  86400,
	})
	testutil.LogNoErr(t, err, data)
}
