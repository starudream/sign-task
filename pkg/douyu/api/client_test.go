package api

import (
	"fmt"
	"os"
	"testing"

	"github.com/starudream/go-lib/core/v2/utils/osutil"

	"github.com/starudream/sign-task/pkg/douyu/config"
)

var C *Client

func TestMain(m *testing.M) {
	accounts := config.C().Accounts
	if len(accounts) == 0 {
		osutil.ExitErr(fmt.Errorf("no account"))
	}
	C = NewClient(accounts[0])

	err := C.Refresh()
	osutil.ExitErr(err)

	os.Exit(m.Run())
}
