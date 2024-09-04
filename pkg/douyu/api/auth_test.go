package api

import (
	"testing"

	"github.com/starudream/go-lib/core/v2/utils/testutil"
)

func TestClient_Refresh(t *testing.T) {
	err := C.Refresh()
	testutil.LogNoErr(t, err)
}
