package api

import (
	"testing"

	"github.com/starudream/go-lib/core/v2/utils/testutil"
)

func TestClient_TBS(t *testing.T) {
	data, err := C.TBS()
	testutil.LogNoErr(t, err, data)
}
