package api

import (
	"testing"

	"github.com/starudream/go-lib/core/v2/utils/testutil"
)

func TestClient_GetHome(t *testing.T) {
	data, err := C.GetHome(GameIdYS)
	testutil.LogNoErr(t, err, data, data.GetSignActId())
}

func TestClient_GetBusinesses(t *testing.T) {
	data, err := C.GetBusinesses()
	testutil.LogNoErr(t, err, data)
}
