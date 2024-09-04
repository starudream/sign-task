package api

import (
	"strings"
	"testing"

	"github.com/starudream/go-lib/core/v2/utils/testutil"

	"github.com/starudream/sign-task/pkg/geetest"
)

func TestClient_ListGameRole(t *testing.T) {
	data, err := C.ListGameRole("")
	testutil.LogNoErr(t, err, data)
}

func TestClient_SignGame(t *testing.T) {
	gt := &geetest.V3Data{
		Challenge: "",
		Validate:  "",
		Seccode:   "",
	}

	gameName, actId, region, uid := GetRole(t, GameBizYSCN)
	data, err := C.SignGame(gameName, actId, region, uid, gt)
	if IsRetCode(err, RetCodeGameHasSigned) {
		t.Log("game has signed")
		return
	}
	testutil.LogNoErr(t, err, data)
}

func TestClient_GetSignGame(t *testing.T) {
	gameName, actId, region, uid := GetRole(t, GameBizYSCN)
	data, err := C.GetSignGame(gameName, actId, region, uid)
	testutil.LogNoErr(t, err, data)
}

func TestClient_ListSignGame(t *testing.T) {
	gameName, actId, _, _ := GetRole(t, GameBizYSCN)
	data, err := C.ListSignGame(gameName, actId)
	testutil.LogNoErr(t, err, data)
}

func TestClient_ListSignGameAward(t *testing.T) {
	gameName, actId, region, uid := GetRole(t, GameBizYSCN)
	data, err := C.ListSignGameAward(gameName, actId, region, uid)
	testutil.LogNoErr(t, err, data, len(data), data.Today(), data.Today().ShortString())
}

func TestClient_ListSignGameAwardPage(t *testing.T) {
	gameName, actId, region, uid := GetRole(t, GameBizYSCN)
	data, err := C.ListSignGameAwardPage(gameName, actId, region, uid, 1, 3)
	testutil.LogNoErr(t, err, data)
}

func GetRole(t *testing.T, gameBiz string) (string, string, string, string) {
	gameName := strings.Split(gameBiz, "_")[0]
	gameId := GameIdByName[gameName]

	data1, err := C.GetHome(gameId)
	testutil.LogNoErr(t, err, data1)
	actId := data1.GetSignActId()
	testutil.MustNotEqual(t, "", actId)

	data2, err := C.ListGameRole(gameBiz)
	testutil.LogNoErr(t, err, data2)
	testutil.MustNotEqual(t, 0, len(data2.List))
	region, uid := data2.List[0].Region, data2.List[0].GameUid

	return gameName, actId, region, uid
}
