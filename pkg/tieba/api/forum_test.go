package api

import (
	"testing"

	"github.com/starudream/go-lib/core/v2/utils/testutil"
)

func TestClient_ListForum(t *testing.T) {
	data, err := C.ListForum()
	testutil.LogNoErr(t, err, data, len(data))
}

func TestClient_SignForum(t *testing.T) {
	data, err := C.SignForum("艾尔登法环")
	if IsCode(err, CodeHasSigned) {
		t.Log("forum has signed")
		return
	}
	testutil.LogNoErr(t, err, data)
}
