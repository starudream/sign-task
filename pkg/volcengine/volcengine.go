package volcengine

import (
	"fmt"
	"time"

	"github.com/starudream/go-lib/core/v2/utils/poolutil"

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

	{
		balance, err := c.QueryBalanceAcct()
		if err != nil {
			cron.Ntfy(j, "火山引擎", fmt.Sprintf("执行失败（%s）", err))
		} else {
			cron.Ntfy(j, "火山引擎", fmt.Sprintf("可用余额：%s，现金余额：%s", balance.AvailableBalance, balance.CashBalance))
		}
	}

	{
		buf := poolutil.BytesBuffer1024.Get()
		defer poolutil.BytesBuffer1024.Put(buf)
		now := time.Now()
		et := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
		be := now.AddDate(0, 0, -3)
		for t := et; t.After(be); t = t.AddDate(0, 0, -1) {
			bills, err := c.ListBillDetail(&api.ListBillDetailReq{
				BillPeriod:    t.Format("2006-01"),
				ExpenseDate:   t.Format("2006-01-02"),
				GroupPeriod:   1,
				GroupTerm:     0,
				IgnoreZero:    1,
				NeedRecordNum: 1,
			})
			if err != nil {
				cron.Ntfy(j, "火山引擎", fmt.Sprintf("执行失败（%s）", err))
				break
			}
			for _, bill := range bills {
				_, _ = fmt.Fprintf(buf, "[%s]\n", bill.ExpenseDate)
				_, _ = fmt.Fprintf(buf, "%s(%s)：%s\n", bill.ProductZh, bill.Product, bill.PayableAmount)
			}
		}
		if buf.Len() > 0 {
			cron.Ntfy(j, "火山引擎", buf.String())
		}
	}
}
