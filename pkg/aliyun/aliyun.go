package aliyun

import (
	"github.com/starudream/sign-task/pkg/aliyun/config"
	"github.com/starudream/sign-task/pkg/cron"
)

func init() {
	cron.Register(aliyun{})
}

type aliyun struct{}

func (aliyun) Name() string {
	return "aliyun"
}

func (j aliyun) Do() {
	for _, account := range config.C().Accounts {
		j.do(account)
	}
}

func (j aliyun) do(a config.Account) {
}
