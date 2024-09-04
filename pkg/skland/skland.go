package skland

import (
	"fmt"

	"github.com/starudream/sign-task/pkg/cron"
	"github.com/starudream/sign-task/pkg/skland/config"
	"github.com/starudream/sign-task/pkg/skland/job"
	"github.com/starudream/sign-task/util"
)

func init() {
	cron.Register(skland{})
}

type skland struct{}

var _ cron.Job = (*skland)(nil)

func (skland) Name() string {
	return "skland"
}

func (j skland) Do() {
	for _, account := range config.C().Accounts {
		j.do(account)
	}
}

func (skland) do(account config.Account) {
	account, err := job.Refresh(account)
	if err != nil {
		util.Ntfy(fmt.Sprintf("%s 执行失败（%s）", account.Phone, err))
		return
	}

	sg := job.SignGame(account)
	util.Ntfy(fmt.Sprintf("%s %s", account.Phone, sg))
}
