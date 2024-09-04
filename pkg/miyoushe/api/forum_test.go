package api

import (
	"testing"

	"github.com/starudream/go-lib/core/v2/utils/testutil"

	"github.com/starudream/sign-task/pkg/geetest"
)

func TestClient_ListPost(t *testing.T) {
	data, err := C.ListPost(ForumIdSR, "")
	testutil.LogNoErr(t, err, data)
}

func TestClient_ListFeedPost(t *testing.T) {
	data, err := C.ListFeedPost(GameIdSR)
	testutil.LogNoErr(t, err, data)
}

const PostId = "57050295"

func TestClient_GetPost(t *testing.T) {
	data, err := C.GetPost(PostId)
	testutil.LogNoErr(t, err, data)
}

func TestClient_UpvotePost(t *testing.T) {
	err := C.UpvotePost(PostId, false)
	testutil.LogNoErr(t, err)
}

func TestClient_CollectPost(t *testing.T) {
	err := C.CollectPost(PostId, false)
	testutil.LogNoErr(t, err)
}

func TestClient_SharePost(t *testing.T) {
	data, err := C.SharePost(PostId)
	testutil.LogNoErr(t, err, data)
}

func TestClient_SignForum(t *testing.T) {
	gt := &geetest.V3Data{
		Challenge: "",
		Validate:  "",
		Seccode:   "",
	}

	data, err := C.SignForum(GameIdSR, gt)
	if IsRetCode(err, RetCodeForumNeedGeetest) {
		t.Log("bbs need geetest")
		return
	}
	testutil.LogNoErr(t, err, data)
}

func TestClient_GetSignForum(t *testing.T) {
	data, err := C.GetSignForum(GameIdSR)
	testutil.LogNoErr(t, err, data)
}
