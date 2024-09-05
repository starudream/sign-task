package tieba

import (
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

func (tieba) Name() string {
	return "tieba"
}

func (j tieba) Do() {
	for _, account := range config.C().Accounts {
		j.do(account)
	}
}

func (j tieba) do(a config.Account) {
	util.NtfyJob(j, a.GetKey(), job.SignForum(a).String())
}
