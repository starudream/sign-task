package douyu

import (
	"github.com/starudream/sign-task/pkg/cron"
	"github.com/starudream/sign-task/pkg/douyu/config"
	"github.com/starudream/sign-task/pkg/douyu/job"
	"github.com/starudream/sign-task/util"
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
	util.NtfyJob(j, a.GetKey(), job.Refresh(a).String())

	util.NtfyJob(j, a.GetKey(), job.Renewal(a).String())
}
