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
		util.NtfyJob(j, a.GetKey(), fmt.Sprintf("执行失败（%s）", err))
		return
	}

	util.NtfyJob(j, a.GetKey(), job.SignGame(a).String())

	util.NtfyJob(j, a.GetKey(), job.SignForum(a).String())
}
