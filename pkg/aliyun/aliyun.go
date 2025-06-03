package aliyun

import (
	"fmt"
	"time"

	"github.com/starudream/go-lib/core/v2/utils/poolutil"

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

	{
		balance, err := c.QueryAccountBalance()
		if err != nil {
			cron.Ntfy(j, "阿里云", fmt.Sprintf("执行失败（%s）", err))
		} else {
			cron.Ntfy(j, "阿里云", fmt.Sprintf("可用额度：%s\n现金余额：%s", balance.AvailableAmount, balance.AvailableCashAmount))
		}
	}

	{
		buf := poolutil.BytesBuffer1024.Get()
		defer poolutil.BytesBuffer1024.Put(buf)
		now := time.Now()
		et := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
		be := now.AddDate(0, 0, -3)
		for t := et; t.After(be); t = t.AddDate(0, 0, -1) {
			bills, err := c.QueryAccountBill(&api.QueryAccountBillReq{
				BillingCycle:     t.Format("2006-01"),
				IsGroupByProduct: true,
				Granularity:      "DAILY",
				BillingDate:      t.Format("2006-01-02"),
			})
			if err != nil {
				cron.Ntfy(j, "阿里云", fmt.Sprintf("执行失败（%s）", err))
				break
			}
			for _, bill := range bills {
				buf.WriteString(fmt.Sprintf("[%s]\n", bill.BillingDate))
				buf.WriteString(fmt.Sprintf("%s(%s)：%.04f\n", bill.ProductName, bill.ProductCode, bill.PretaxAmount))
			}
		}
		if buf.Len() > 0 {
			cron.Ntfy(j, "阿里云", buf.String())
		}
	}
}
