package api

import (
	"testing"

	"github.com/starudream/go-lib/core/v2/utils/testutil"
)

func TestClient_ListPlayer(t *testing.T) {
	data, err := C.ListPlayer()
	testutil.LogNoErr(t, err, data)
}

func TestClient_SignGame(t *testing.T) {
	gid, player := GetPlayer(t)
	data, err := C.SignGame(gid, player.Uid)
	if IsCode(err, CodeGameHasSigned) {
		t.Log("game has signed")
		return
	}
	testutil.LogNoErr(t, err, data)
}

func TestClient_ListSignGame(t *testing.T) {
	gid, player := GetPlayer(t)
	data, err := C.ListSignGame(gid, player.Uid)
	testutil.LogNoErr(t, err, data)
}

func GetPlayer(t *testing.T) (string, *Player) {
	data, err := C.ListPlayer()
	testutil.LogNoErr(t, err, data)
	testutil.MustNotEqual(t, 0, len(data.List))
	testutil.MustNotEqual(t, 0, len(data.List[0].BindingList))
	testutil.MustNotEqual(t, "", GameIdByCode[data.List[0].AppCode])
	return GameIdByCode[data.List[0].AppCode], data.List[0].BindingList[0]
}
