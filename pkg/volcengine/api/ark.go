package api

type ListArkEndpointsReq struct {
	PageNumber int `json:"PageNumber"`
	PageSize   int `json:"PageSize"`
}

type ListArkEndpointsData struct {
	PageNumber int `json:"PageNumber"`
	PageSize   int `json:"PageSize"`
	TotalCount int `json:"TotalCount"`

	Items []*ListArkEndpointsItem `json:"Items"`
}

type ListArkEndpointsItem struct {
	Id          string `json:"Id"`
	Name        string `json:"Name"`
	Description string `json:"Description"`
	ProjectName string `json:"ProjectName"`
	Status      string `json:"Status"`
}

func (c *Client) ListArkEndpoints(req *ListArkEndpointsReq) ([]*ListArkEndpointsItem, error) {
	var items []*ListArkEndpointsItem
	for i := 0; i < 100; i++ {
		req.PageNumber = i + 1
		req.PageSize = 100
		data, err := Exec[*ListArkEndpointsData](c.R().SetBody(req), "POST", "https://open.volcengineapi.com", "/", "ListEndpoints", "2024-01-01", Region, "ark", c.account)
		if err != nil {
			return nil, err
		}
		items = append(items, data.Items...)
		if len(items) >= data.TotalCount {
			break
		}
	}
	return items, nil
}

type GetArkUsageReq struct {
	StartTime   int      `json:"StartTime"`
	EndTime     int      `json:"EndTime"`
	Interval    int      `json:"Interval"`
	EndpointIds []string `json:"EndpointIds,omitempty"`
}

type GetArkUsageData struct {
	UsageResults []*GetArkUsageResult `json:"UsageResults"`
}

type GetArkUsageResult struct {
	Name        string                   `json:"Name"`
	MetricItems []*GetArkUsageMetricItem `json:"MetricItems"`
}

type GetArkUsageMetricItem struct {
	Values []*GetArkUsageMetricValue `json:"Values"`
}

type GetArkUsageMetricValue struct {
	Timestamp int `json:"Timestamp"`
	Value     int `json:"Value"`
}

func (c *Client) GetArkUsage(req *GetArkUsageReq) ([]*GetArkUsageResult, error) {
	data, err := Exec[*GetArkUsageData](c.R().SetBody(req), "POST", "https://open.volcengineapi.com", "/", "GetUsage", "2024-01-01", Region, "ark", c.account)
	if err != nil {
		return nil, err
	}
	return data.UsageResults, nil
}
