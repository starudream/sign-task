package api

type QueryBalanceAcctData struct {
	AccountID        int    `json:"AccountID"`
	ArrearsBalance   string `json:"ArrearsBalance"`
	AvailableBalance string `json:"AvailableBalance"`
	CashBalance      string `json:"CashBalance"`
	CreditLimit      string `json:"CreditLimit"`
	FreezeAmount     string `json:"FreezeAmount"`
}

func (c *Client) QueryBalanceAcct() (*QueryBalanceAcctData, error) {
	return Exec[*QueryBalanceAcctData](c.R(), "GET", "https://open.volcengineapi.com", "/", "QueryBalanceAcct", "2022-01-01", Region, "billing", c.account)
}
