package xiaoheihe

import (
	"github.com/starudream/sign-task/pkg/cron"
)

func init() {
	cron.Register(xiaoheihe{})
}

type xiaoheihe struct {
}

func (xiaoheihe) Name() string {
	return "xiaoheihe"
}

func (xiaoheihe) Do() {
}
