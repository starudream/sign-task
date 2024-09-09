package miyoushe

import (
	"fmt"

	"github.com/starudream/sign-task/pkg/cron"
	"github.com/starudream/sign-task/pkg/miyoushe/config"
	"github.com/starudream/sign-task/pkg/miyoushe/job"
)

func init() {
	cron.Register(miyoushe{})
}

type miyoushe struct {
}

func (miyoushe) Name() string {
	return "miyoushe"
}

func (j miyoushe) Do() {
	for _, account := range config.C().Accounts {
		j.do(account)
	}
}

func (j miyoushe) do(a config.Account) {
	a, err := job.Refresh(a)
	if err != nil {
		cron.Ntfy(j, a.GetKey(), fmt.Sprintf("执行失败（%s）", err))
		return
	}

	cron.Ntfy(j, a.GetKey(), job.SignGame(a).String())

	cron.Ntfy(j, a.GetKey(), job.SignForum(a).String())
}
