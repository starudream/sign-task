package volcengine

import (
	"fmt"

	"github.com/starudream/sign-task/pkg/cron"
	"github.com/starudream/sign-task/pkg/volcengine/api"
	"github.com/starudream/sign-task/pkg/volcengine/config"
)

func init() {
	cron.Register(volcengine{})
}

type volcengine struct{}

func (volcengine) Name() string {
	return "volcengine"
}

func (j volcengine) Do() {
	for _, account := range config.C().Accounts {
		j.do(account)
	}
}

func (j volcengine) do(a config.Account) {
	c := api.NewClient(a)

	balance, err := c.QueryBalanceAcct()
	if err != nil {
		cron.Ntfy(j, "火山引擎", fmt.Sprintf("执行失败（%s）", err))
	} else {
		cron.Ntfy(j, "火山引擎", fmt.Sprintf("可用余额：%s，现金余额：%s", balance.AvailableBalance, balance.CashBalance))
	}
}
