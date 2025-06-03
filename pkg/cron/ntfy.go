package cron

import (
	"github.com/starudream/sign-task/util"
)

func Ntfy(j Job, k, s string) {
	util.Ntfy(j.Name(), k, s)
}
