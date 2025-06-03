package api

import (
	"github.com/starudream/sign-task/util"
)

type QueryAccountBalanceData struct {
	AvailableCashAmount string `json:"AvailableCashAmount"`
	MybankCreditAmount  string `json:"MybankCreditAmount"`
	Currency            string `json:"Currency"`
	AvailableAmount     string `json:"AvailableAmount"`
	CreditAmount        string `json:"CreditAmount"`
	QuotaLimit          string `json:"QuotaLimit"`
}

func (c *Client) QueryAccountBalance() (*QueryAccountBalanceData, error) {
	return Exec[*QueryAccountBalanceData](c.R(), "POST", AddrBSS, "/", "QueryAccountBalance", "2017-12-14", c.account)
}

type QueryAccountBillReq struct {
	PageNum  int `json:"PageNum"`
	PageSize int `json:"PageSize"`

	BillingCycle     string `json:"BillingCycle"`
	IsGroupByProduct bool   `json:"IsGroupByProduct"`
	Granularity      string `json:"Granularity"` // MONTHLY or DAILY
	BillingDate      string `json:"BillingDate"`
}

type QueryAccountBillData struct {
	PageNum    int `json:"PageNum"`
	PageSize   int `json:"PageSize"`
	TotalCount int `json:"TotalCount"`

	AccountID    string `json:"AccountID"`
	AccountName  string `json:"AccountName"`
	BillingCycle string `json:"BillingCycle"`

	Items *QueryAccountBillItems `json:"Items"`
}

type QueryAccountBillItems struct {
	Item []*QueryAccountBillItem `json:"Item"`
}

type QueryAccountBillItem struct {
	BillingDate  string  `json:"BillingDate"`
	ProductCode  string  `json:"ProductCode"`
	ProductName  string  `json:"ProductName"`
	PretaxAmount float64 `json:"PretaxAmount"`
}

func (c *Client) QueryAccountBill(req *QueryAccountBillReq) ([]*QueryAccountBillItem, error) {
	var items []*QueryAccountBillItem
	for i := 1; i < 100; i++ {
		req.PageNum = i
		req.PageSize = 100
		data, err := Exec[*QueryAccountBillData](c.R().SetQueryParams(util.ToMapString(req)), "POST", AddrBSS, "/", "QueryAccountBill", "2017-12-14", c.account)
		if err != nil {
			return nil, err
		}
		items = append(items, data.Items.Item...)
		if len(items) >= data.TotalCount {
			break
		}
	}
	return items, nil
}
