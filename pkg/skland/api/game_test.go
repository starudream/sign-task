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
	gid, player := GetPlayer(t, 0)
	data, err := C.SignGame(gid, player.Uid, player.GetDefaultRole().RoleId, player.GetDefaultRole().ServerId)
	if IsCode(err, CodeGameHasSigned) {
		t.Log("game has signed")
		return
	}
	testutil.LogNoErr(t, err, data)
}

func TestClient_ListSignGame(t *testing.T) {
	gid, player := GetPlayer(t, 0)
	data, err := C.ListSignGame(gid, player.Uid, player.GetDefaultRole().RoleId, player.GetDefaultRole().ServerId)
	testutil.LogNoErr(t, err, data)
}

func GetPlayer(t *testing.T, i int) (string, *Player) {
	data, err := C.ListPlayer()
	testutil.LogNoErr(t, err, data)
	testutil.MustNotEqual(t, 0, len(data.List))
	testutil.MustNotEqual(t, 0, len(data.List[i].BindingList))
	testutil.MustNotEqual(t, "", GameIdByCode[data.List[i].AppCode])
	return GameIdByCode[data.List[i].AppCode], data.List[i].BindingList[0]
}
