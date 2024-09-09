package api

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
