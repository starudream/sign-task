package cron

import (
	"fmt"

	"github.com/starudream/sign-task/util"
)

func Ntfy(j Job, k, s string) {
	util.Ntfy(fmt.Sprintf("[%s] %s\n%s", j.Name(), k, s))
}
