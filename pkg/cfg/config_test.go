package cfg

import (
	"testing"

	"github.com/starudream/go-lib/core/v2/utils/testutil"
)

func TestSave(t *testing.T) {
	testutil.LogNoErr(t, Save())
}
