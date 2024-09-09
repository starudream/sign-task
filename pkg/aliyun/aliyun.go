package aliyun

import (
	"fmt"

	"github.com/starudream/sign-task/pkg/aliyun/api"
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
	c := api.NewClient(a)

	balance, err := c.QueryAccountBalance()
	if err != nil {
		cron.Ntfy(j, "阿里云", fmt.Sprintf("执行失败（%s）", err))
	} else {
		cron.Ntfy(j, "阿里云", fmt.Sprintf("可用额度：%s，现金余额：%s", balance.AvailableAmount, balance.AvailableCashAmount))
	}
}
