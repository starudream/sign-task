package api

import (
	"testing"

	"github.com/starudream/go-lib/core/v2/utils/testutil"
)

func TestClient_ListPost(t *testing.T) {
	data, err := C.ListPost(GameIdMC, ForumIdMC10, 1)
	testutil.LogNoErr(t, err, data)
}

func TestClient_GetPost(t *testing.T) {
	data, err := C.GetPost("1272592437753516032")
	testutil.LogNoErr(t, err, data)
}

func TestClient_LikePost(t *testing.T) {
	err := C.LikePost(GameIdMC, ForumIdMC10, "1272592437753516032", "10381279", true)
	testutil.LogNoErr(t, err)
}

func TestClient_SharePost(t *testing.T) {
	err := C.SharePost(GameIdMC)
	testutil.LogNoErr(t, err)
}

func TestClient_SignForum(t *testing.T) {
	data, err := C.SignForum(GameIdMC)
	if IsCode(err, CodeHasSigned) {
		t.Log("forum has signed")
		return
	}
	testutil.LogNoErr(t, err, data)
}

func TestClient_GetSignForum(t *testing.T) {
	data, err := C.GetSignForum(GameIdMC)
	testutil.LogNoErr(t, err, data)
}
