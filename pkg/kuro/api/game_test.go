package api

import (
	"testing"

	"github.com/starudream/go-lib/core/v2/utils/testutil"
)

func TestClient_ListRole(t *testing.T) {
	data, err := C.ListRole(GameIdMC)
	testutil.LogNoErr(t, err, data)
}

func TestClient_SignGame(t *testing.T) {
	role := GetRole(t)
	data, err := C.SignGame(role.GameId, role.ServerId, role.RoleId, role.UserId)
	if IsCode(err, CodeHasSigned) {
		t.Log("game has signed")
		return
	}
	testutil.LogNoErr(t, err, data)
}

func TestClient_ListSignGame(t *testing.T) {
	role := GetRole(t)
	data, err := C.ListSignGame(role.GameId, role.ServerId, role.RoleId, role.UserId)
	testutil.LogNoErr(t, err, data, data.GoodsMap())
}

func TestClient_ListSignGameRecord(t *testing.T) {
	role := GetRole(t)
	data, err := C.ListSignGameRecord(role.GameId, role.ServerId, role.RoleId, role.UserId)
	testutil.LogNoErr(t, err, data, len(data), data.Today(), data.Today().ShortString())
}

func GetRole(t *testing.T) *Role {
	data, err := C.ListRole(GameIdMC)
	testutil.LogNoErr(t, err, data)
	testutil.MustNotEqual(t, 0, len(data))
	return data[0]
}
