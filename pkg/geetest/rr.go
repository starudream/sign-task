package geetest

import (
	"context"
	"fmt"
	"time"

	"github.com/starudream/go-lib/core/v2/config"
	"github.com/starudream/go-lib/core/v2/gh"
	"github.com/starudream/go-lib/resty/v2"
)

// https://www.kancloud.cn/rrocr/rrocr/2294926
type rrConfig struct {
	Key string `json:"key" yaml:"key"`
}

var rr = rrConfig{}

func init() {
	_ = config.Unmarshal("geetest.rr", &rr)
}

func RRKey() string {
	return rr.Key
}

func RRPoint(req *V3Param) (int, error) {
	form := gh.MS{
		"appkey": gh.Ternary(req.Key != "", req.Key, rr.Key),
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	res, err := resty.ParseResp[*rrResp, *rrResp](
		resty.R().SetContext(ctx).SetError(&rrResp{}).SetResult(&rrResp{}).SetFormData(form).Post("http://api.rrocr.com/api/integral.html"),
	)
	if err != nil {
		return -1, fmt.Errorf("[rrocr] %w", err)
	}
	return res.Integral, nil
}

type rrResp struct {
	Status   int     `json:"status"`
	Msg      string  `json:"msg"`
	Code     int     `json:"code,omitempty"`
	Data     *V3Data `json:"data,omitempty"`
	Integral int     `json:"integral,omitempty"`
}

func (t *rrResp) IsSuccess() bool {
	return t.Status == 0
}

func (t *rrResp) String() string {
	return fmt.Sprintf("status: %d, msg: %s, code: %d", t.Status, t.Msg, t.Code)
}

func RR(req *V3Param) (*V3Data, error) {
	form := gh.MS{
		"appkey":    gh.Ternary(req.Key != "", req.Key, rr.Key),
		"gt":        req.GT,
		"challenge": req.Challenge,
		"referer":   req.Referer,
	}
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	res, err := resty.ParseResp[*rrResp, *rrResp](
		resty.R().SetContext(ctx).SetError(&rrResp{}).SetResult(&rrResp{}).SetFormData(form).Post("http://api.rrocr.com/api/recognize.html"),
	)
	if err != nil {
		return nil, fmt.Errorf("[rrocr] %w", err)
	}
	return res.Data, nil
}
