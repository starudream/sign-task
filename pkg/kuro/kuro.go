package kuro

import (
	"fmt"

	"github.com/starudream/sign-task/pkg/cron"
	"github.com/starudream/sign-task/pkg/kuro/config"
	"github.com/starudream/sign-task/pkg/kuro/job"
	"github.com/starudream/sign-task/util"
)

func init() {
	cron.Register(kuro{})
}

type kuro struct{}

var _ cron.Job = (*kuro)(nil)

func (kuro) Name() string {
	return "kuro"
}

func (j kuro) Do() {
	for _, account := range config.C().Accounts {
		j.do(account)
	}
}

func (kuro) do(account config.Account) {
	sg := job.SignGame(account)
	util.Ntfy(fmt.Sprintf("%s %s", account.Phone, sg))

	sf := job.SignForum(account)
	util.Ntfy(fmt.Sprintf("%s %s", account.Phone, sf))
}
