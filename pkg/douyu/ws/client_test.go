package ws

import (
	"fmt"
	"testing"

	"github.com/starudream/go-lib/core/v2/utils/osutil"
	"github.com/starudream/go-lib/core/v2/utils/testutil"

	"github.com/starudream/sign-task/pkg/douyu/api"
	"github.com/starudream/sign-task/pkg/douyu/config"
)

func TestLogin(t *testing.T) {
	accounts := config.C().Accounts
	if len(accounts) == 0 {
		osutil.ExitErr(fmt.Errorf("no account"))
	}
	c := api.NewClient(accounts[0])

	err := c.Refresh()
	testutil.LogNoErr(t, err)

	err = Login(LoginParams{Room: api.RoomYYF, Stk: c.Stk, Ltkid: c.Ltkid, Username: c.Username})
	testutil.LogNoErr(t, err)
}
