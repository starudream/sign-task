package geetest

import (
	"fmt"
	"net/url"
	"time"

	"github.com/starudream/go-lib/core/v2/config"
	"github.com/starudream/go-lib/core/v2/gh"
	"github.com/starudream/go-lib/resty/v2"
)

// https://www.kancloud.cn/ttorc/ttorc/3119237
type ttConfig struct {
	Key    string `json:"key"     yaml:"key"`
	ItemId string `json:"item_id" yaml:"item_id"`
	Proxy  string `json:"proxy"   yaml:"proxy"`
}

var (
	tt = ttConfig{}

	ttClient *resty.Client
)

func init() {
	_ = config.Unmarshal("geetest.tt", &tt)
}

func ttR() *resty.Request {
	if ttClient == nil {
		ttClient = resty.New().
			SetTimeout(30*time.Second).
			SetHeader("Accept-Encoding", "gzip").
			SetHeader("User-Agent", resty.UAWindowsChrome)
		if tt.Proxy != "" {
			_, err := url.Parse(tt.Proxy)
			if err == nil {
				ttClient.SetProxy(tt.Proxy)
			}
		}
	}
	return ttClient.R()
}

func TTKey() string {
	return tt.Key
}

func TTPoint(req *V3Param) (string, error) {
	form := gh.MS{
		"appkey": gh.Ternary(req.Key != "", req.Key, tt.Key),
	}
	res, err := resty.ParseResp[*ttResp, *ttResp](
		ttR().SetError(&ttResp{}).SetResult(&ttResp{}).SetFormData(form).Post("http://api.ttocr.com/api/points"),
	)
	if err != nil {
		return "-1", fmt.Errorf("[ttocr] %w", err)
	}
	return res.Points, nil
}

func TT(req *V3Param) (*V3Data, error) {
	resultId, err := ttRecognize(req)
	if err != nil {
		return nil, err
	}

	ch := make(chan *V3Data, 1)

	go func() {
		var v *V3Data
		for {
			time.Sleep(3 * time.Second)
			v, err = ttResult(req, resultId)
			if v1, ok1 := resty.AsRespErr(err); ok1 {
				if v2, ok2 := v1.Response.Result().(interface{ GetStatus() int }); ok2 && v2.GetStatus() == 2 {
					continue
				}
			}
			ch <- v
			return
		}
	}()

	select {
	case v := <-ch:
		return v, err
	case <-time.After(60 * time.Second):
		return nil, fmt.Errorf("[ttocr] recognize timeout")
	}
}

type ttResp struct {
	Status   int     `json:"status"`
	Msg      string  `json:"msg"`
	ResultId string  `json:"resultid,omitempty"`
	Data     *V3Data `json:"data,omitempty"`
	Points   string  `json:"points,omitempty"`
}

func (t *ttResp) GetStatus() int {
	if t == nil {
		return 0
	}
	return t.Status
}

func (t *ttResp) IsSuccess() bool {
	return t.Status == 1
}

func (t *ttResp) String() string {
	return fmt.Sprintf("status: %d, msg: %s", t.Status, t.Msg)
}

func ttRecognize(req *V3Param) (string, error) {
	form := gh.MS{
		"appkey":    gh.Ternary(req.Key != "", req.Key, tt.Key),
		"gt":        req.GT,
		"challenge": req.Challenge,
		"referer":   req.Referer,
		"itemid":    gh.Ternary(req.ItemId != "", req.ItemId, "388"),
	}
	res, err := resty.ParseResp[*ttResp, *ttResp](
		ttR().SetError(&ttResp{}).SetResult(&ttResp{}).SetFormData(form).Post("http://api.ttocr.com/api/recognize"),
	)
	if err != nil {
		return "", fmt.Errorf("[ttocr] %w", err)
	}
	return res.ResultId, nil
}

func ttResult(req *V3Param, resultId string) (*V3Data, error) {
	form := gh.MS{
		"appkey":   gh.Ternary(req.Key != "", req.Key, tt.Key),
		"resultid": resultId,
	}
	res, err := resty.ParseResp[*ttResp, *ttResp](
		ttR().SetError(&ttResp{}).SetResult(&ttResp{}).SetFormData(form).Post("http://api.ttocr.com/api/results"),
	)
	if err != nil {
		return nil, fmt.Errorf("[ttocr] %w", err)
	}
	return res.Data, nil
}
