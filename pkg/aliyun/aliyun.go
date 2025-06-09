package aliyun

import (
	"fmt"
	"strconv"

	"github.com/starudream/go-lib/core/v2/utils/poolutil"

	"github.com/starudream/sign-task/pkg/aliyun/api"
	"github.com/starudream/sign-task/pkg/aliyun/config"
	"github.com/starudream/sign-task/pkg/cron"
	"github.com/starudream/sign-task/util"
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
		for t := util.GetToday(); t.After(util.GetToday(-3)); t = t.AddDate(0, 0, -1) {
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
				_, _ = fmt.Fprintf(buf, "[%s]\n", bill.BillingDate)
				_, _ = fmt.Fprintf(buf, "%s(%s)：%.04f\n", bill.ProductName, bill.ProductCode, bill.PretaxAmount)
			}
		}
		if buf.Len() > 0 {
			cron.Ntfy(j, "阿里云", buf.String())
		}
	}

	{
		items, err := c.DescribeInstanceBill(&api.DescribeInstanceBillReq{
			BillingCycle:     util.GetToday().Format("2006-01"),
			ProductCode:      "cdt",
			Granularity:      "MONTHLY",
			IsBillingItem:    true,
			IsHideZeroCharge: false,
		})
		if err != nil {
			cron.Ntfy(j, "阿里云", fmt.Sprintf("执行失败（%s）", err))
		} else {
			usage, usageUnit := float64(0), ""
			for _, item := range items {
				_usage, _ := strconv.ParseFloat(item.Usage, 64)
				if _usage > 0 {
					usage += _usage
				}
				if item.UsageUnit != "" {
					usageUnit = item.UsageUnit
				}
			}
			cron.Ntfy(j, "阿里云", fmt.Sprintf("CDT 本月已使用：%.04f %s", usage, usageUnit))
		}
	}
}
