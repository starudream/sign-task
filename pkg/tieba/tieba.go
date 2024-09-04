package tieba

import (
	"fmt"

	"github.com/starudream/sign-task/pkg/cron"
	"github.com/starudream/sign-task/pkg/tieba/config"
	"github.com/starudream/sign-task/pkg/tieba/job"
	"github.com/starudream/sign-task/util"
)

func init() {
	cron.Register(tieba{})
}

type tieba struct {
}

var _ cron.Job = (*tieba)(nil)

func (tieba) Name() string {
	return "tieba"
}

func (j tieba) Do() {
	for _, account := range config.C().Accounts {
		j.do(account)
	}
}

func (tieba) do(account config.Account) {
	sf := job.SignForum(account)
	util.Ntfy(fmt.Sprintf("%s %s", account.Phone, sf))
}
