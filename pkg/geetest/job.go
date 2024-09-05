package geetest

import (
	"fmt"

	"github.com/starudream/sign-task/pkg/cron"
	"github.com/starudream/sign-task/util"
)

func init() {
	cron.Register(geetest{})
}

type geetest struct{}

func (geetest) Name() string {
	return "geetest"
}

func (j geetest) Do() {
	if TTKey() != "" {
		point, err := TTPoint(&V3Param{})
		if err != nil {
			util.NtfyJob(j, "套套", fmt.Sprintf("执行失败（%s）", err))
		} else {
			util.NtfyJob(j, "套套", fmt.Sprintf("剩余点数：%s", point))
		}
	}

	if RRKey() != "" {
		point, err := RRPoint(&V3Param{})
		if err != nil {
			util.NtfyJob(j, "人人", fmt.Sprintf("执行失败（%s）", err))
		} else {
			util.NtfyJob(j, "人人", fmt.Sprintf("剩余点数：%d", point))
		}
	}
}
