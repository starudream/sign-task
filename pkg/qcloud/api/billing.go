package api

type DescribeAccountBalanceData struct {
	respError

	Uin                      int    `json:"Uin"`
	RealBalance              int    `json:"RealBalance"`
	CashAccountBalance       int    `json:"CashAccountBalance"`
	IncomeIntoAccountBalance int    `json:"IncomeIntoAccountBalance"`
	PresentAccountBalance    int    `json:"PresentAccountBalance"`
	FreezeAmount             int    `json:"FreezeAmount"`
	OweAmount                int    `json:"OweAmount"`
	IsAllowArrears           bool   `json:"IsAllowArrears"`
	IsCreditLimited          bool   `json:"IsCreditLimited"`
	Balance                  int    `json:"Balance"`
	CreditAmount             int    `json:"CreditAmount"`
	CreditBalance            int    `json:"CreditBalance"`
	RealCreditBalance        int    `json:"RealCreditBalance"`
	RequestId                string `json:"RequestId"`
}

func (c *Client) DescribeAccountBalance() (*DescribeAccountBalanceData, error) {
	return Exec[*DescribeAccountBalanceData](c.R(), "POST", "https://billing.tencentcloudapi.com", "/", "DescribeAccountBalance", "2018-07-09", "", "billing", c.account)
}
