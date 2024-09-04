package xiaoheihe

import (
	"github.com/starudream/sign-task/pkg/cron"
)

func init() {
	cron.Register(xiaoheihe{})
}

type xiaoheihe struct {
}

var _ cron.Job = (*xiaoheihe)(nil)

func (xiaoheihe) Name() string {
	return "xiaoheihe"
}

func (xiaoheihe) Do() {
}
