package qcloud

import (
	"fmt"

	"github.com/starudream/sign-task/pkg/cron"
	"github.com/starudream/sign-task/pkg/qcloud/api"
	"github.com/starudream/sign-task/pkg/qcloud/config"
)

func init() {
	cron.Register(qcloud{})
}

type qcloud struct{}

func (j qcloud) Name() string {
	return "qcloud"
}

func (j qcloud) Do() {
	for _, account := range config.C().Accounts {
		j.do(account)
	}
}

func (j qcloud) do(a config.Account) {
	c := api.NewClient(a)

	{
		balance, err := c.DescribeAccountBalance()
		if err != nil {
			cron.Ntfy(j, "腾讯云", fmt.Sprintf("执行失败（%s）", err))
		} else {
			cron.Ntfy(j, "腾讯云", fmt.Sprintf("可用额度：%.02f\n现金余额：%.02f", float64(balance.Balance)/100, float64(balance.CashAccountBalance)/100))
		}
	}
}
