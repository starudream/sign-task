package douyu

import (
	"github.com/starudream/sign-task/pkg/cron"
	"github.com/starudream/sign-task/pkg/douyu/config"
	"github.com/starudream/sign-task/pkg/douyu/job"
)

func init() {
	cron.Register(douyu{})
}

type douyu struct {
}

func (douyu) Name() string {
	return "douyu"
}

func (j douyu) Do() {
	for _, account := range config.C().Accounts {
		j.do(account)
	}
}

func (j douyu) do(a config.Account) {
	cron.Ntfy(j, a.GetKey(), job.Refresh(a).String())

	cron.Ntfy(j, a.GetKey(), job.Renewal(a).String())
}
