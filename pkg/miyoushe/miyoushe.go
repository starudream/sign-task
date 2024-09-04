package miyoushe

import (
	"fmt"

	"github.com/starudream/sign-task/pkg/cron"
	"github.com/starudream/sign-task/pkg/miyoushe/config"
	"github.com/starudream/sign-task/pkg/miyoushe/job"
	"github.com/starudream/sign-task/util"
)

func init() {
	cron.Register(miyoushe{})
}

type miyoushe struct {
}

var _ cron.Job = (*miyoushe)(nil)

func (miyoushe) Name() string {
	return "miyoushe"
}

func (j miyoushe) Do() {
	for _, account := range config.C().Accounts {
		j.do(account)
	}
}

func (miyoushe) do(account config.Account) {
	account, err := job.Refresh(account)
	if err != nil {
		util.Ntfy(fmt.Sprintf("%s 执行失败（%s）", account.Phone, err))
		return
	}

	sg := job.SignGame(account)
	util.Ntfy(fmt.Sprintf("%s %s", account.Phone, sg))

	sf := job.SignForum(account)
	util.Ntfy(fmt.Sprintf("%s %s", account.Phone, sf))
}
