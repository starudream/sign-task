package api

import (
	"testing"

	"github.com/starudream/go-lib/core/v2/utils/testutil"
)

func TestClient_ListBadges(t *testing.T) {
	data, err := C.ListBadges()
	testutil.LogNoErr(t, err, data)
}
