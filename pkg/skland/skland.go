package skland

import (
	"fmt"

	"github.com/starudream/sign-task/pkg/cron"
	"github.com/starudream/sign-task/pkg/skland/config"
	"github.com/starudream/sign-task/pkg/skland/job"
)

func init() {
	cron.Register(skland{})
}

type skland struct{}

func (skland) Name() string {
	return "skland"
}

func (j skland) Do() {
	for _, account := range config.C().Accounts {
		j.do(account)
	}
}

func (j skland) do(a config.Account) {
	a, err := job.Refresh(a)
	if err != nil {
		cron.Ntfy(j, a.GetKey(), fmt.Sprintf("执行失败（%s）", err))
		return
	}

	cron.Ntfy(j, a.GetKey(), job.SignGame(a).String())
}
