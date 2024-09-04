package api

import (
	"testing"

	"github.com/starudream/go-lib/core/v2/utils/testutil"
)

func TestClient_SendGift(t *testing.T) {
	data, err := C.SendGift(RoomYYF, GiftGlowSticks, 1)
	testutil.LogNoErr(t, err, data)
}

func TestClient_ListGift(t *testing.T) {
	data, err := C.ListGift()
	testutil.LogNoErr(t, err, data)
}
