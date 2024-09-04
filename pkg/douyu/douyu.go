package douyu

import (
	"fmt"

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

var _ cron.Job = (*douyu)(nil)

func (douyu) Name() string {
	return "douyu"
}

func (j douyu) Do() {
	for _, account := range config.C().Accounts {
		j.do(account)
	}
}

func (douyu) do(account config.Account) {
	rf := job.Refresh(account)
	util.Ntfy(fmt.Sprintf("%s %s", account.Phone, rf))

	rn := job.Renewal(account)
	util.Ntfy(fmt.Sprintf("%s %s", account.Phone, rn))
}
