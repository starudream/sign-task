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

type ListBillDetailReq struct {
	Offset        int    `json:"Offset"`
	Limit         int    `json:"Limit"`
	BillPeriod    string `json:"BillPeriod"`
	ExpenseDate   string `json:"ExpenseDate"`
	GroupPeriod   int    `json:"GroupPeriod"`
	GroupTerm     int    `json:"GroupTerm"`
	IgnoreZero    int    `json:"IgnoreZero"`
	NeedRecordNum int    `json:"NeedRecordNum"`
}

type ListBillDetailData struct {
	Offset int `json:"Offset"`
	Limit  int `json:"Limit"`
	Total  int `json:"Total"`

	List []*ListBillDetailItem `json:"List"`
}

type ListBillDetailItem struct {
	ExpenseDate   string `json:"ExpenseDate"`
	Product       string `json:"Product"`
	ProductZh     string `json:"ProductZh"`
	PayableAmount string `json:"PayableAmount"`
}

func (c *Client) ListBillDetail(req *ListBillDetailReq) ([]*ListBillDetailItem, error) {
	var items []*ListBillDetailItem
	for i := 0; i < 100; i++ {
		req.Offset = i * 100
		req.Limit = 100
		data, err := Exec[*ListBillDetailData](c.R().SetBody(req), "POST", "https://open.volcengineapi.com", "/", "ListBillDetail", "2022-01-01", Region, "billing", c.account)
		if err != nil {
			return nil, err
		}
		items = append(items, data.List...)
		if len(items) >= data.Total {
			break
		}
	}
	return items, nil
}
