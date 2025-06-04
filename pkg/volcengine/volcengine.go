package volcengine

import (
	"fmt"
	"maps"
	"slices"
	"time"

	"github.com/starudream/go-lib/core/v2/utils/poolutil"

	"github.com/starudream/sign-task/pkg/cron"
	"github.com/starudream/sign-task/pkg/volcengine/api"
	"github.com/starudream/sign-task/pkg/volcengine/config"
	"github.com/starudream/sign-task/util"
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

	j.doBill(c)

	j.doArk(c)
}

func (j volcengine) doBill(c *api.Client) {
	buf := poolutil.BytesBuffer1024.Get()
	defer poolutil.BytesBuffer1024.Put(buf)

	for t := util.GetToday(); t.After(util.GetToday(-3)); t = t.AddDate(0, 0, -1) {
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

func (j volcengine) doArk(c *api.Client) {
	endpoints, err := c.ListArkEndpoints(&api.ListArkEndpointsReq{})
	if err != nil {
		cron.Ntfy(j, "火山方舟", fmt.Sprintf("执行失败（%s）", err))
		return
	}
	for _, ep := range endpoints {
		j.doArkEp(c, ep)
	}
}

func (j volcengine) doArkEp(c *api.Client, ep *api.ListArkEndpointsItem) {
	usages, err := c.GetArkUsage(&api.GetArkUsageReq{
		StartTime:   int(util.GetToday(-3).Unix()),
		EndTime:     int(util.GetToday(1).Unix() - 1),
		Interval:    86400,
		EndpointIds: []string{ep.Id},
	})
	if err != nil {
		cron.Ntfy(j, "火山方舟", fmt.Sprintf("执行失败（%s）", err))
		return
	}

	data := map[string]map[string]int{}
	for _, usage := range usages {
		for _, item := range usage.MetricItems {
			for _, v := range item.Values {
				k := time.Unix(int64(v.Timestamp), 0).Format("2006-01-02 15:04:05")
				if _, ok := data[k]; !ok {
					data[k] = map[string]int{}
				}
				data[k][usage.Name] += v.Value
			}
		}
	}
	if len(data) == 0 {
		return
	}

	buf := poolutil.BytesBuffer1024.Get()
	defer poolutil.BytesBuffer1024.Put(buf)

	_, _ = fmt.Fprintf(buf, "%s\n", ep.Name)

	keys := slices.Collect(maps.Keys(data))
	slices.Sort(keys)
	slices.Reverse(keys)
	for _, key := range keys {
		_, _ = fmt.Fprintf(buf, "[%s]\n", key)
		names := slices.Collect(maps.Keys(data[key]))
		slices.Sort(names)
		for _, name := range names {
			cname := name
			switch name {
			case "PromptTokens":
				cname = "输入"
			case "CompletionTokens":
				cname = "输出"
			}
			_, _ = fmt.Fprintf(buf, "%s：%s\n", cname, util.FormatInt(data[key][name]))
		}
	}
	if buf.Len() > 0 {
		cron.Ntfy(j, "火山方舟", buf.String())
	}
}
