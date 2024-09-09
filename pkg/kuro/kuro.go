package kuro

import (
	"github.com/starudream/sign-task/pkg/cron"
	"github.com/starudream/sign-task/pkg/kuro/config"
	"github.com/starudream/sign-task/pkg/kuro/job"
)

func init() {
	cron.Register(kuro{})
}

type kuro struct{}

func (kuro) Name() string {
	return "kuro"
}

func (j kuro) Do() {
	for _, account := range config.C().Accounts {
		j.do(account)
	}
}

func (j kuro) do(a config.Account) {
	cron.Ntfy(j, a.GetKey(), job.SignGame(a).String())

	cron.Ntfy(j, a.GetKey(), job.SignForum(a).String())
}
