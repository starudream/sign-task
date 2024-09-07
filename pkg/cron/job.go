package cron

import (
	"fmt"
	"math/rand/v2"
	"time"

	"github.com/starudream/go-lib/core/v2/config"
	"github.com/starudream/go-lib/core/v2/gh"
	"github.com/starudream/go-lib/core/v2/slog"
	"github.com/starudream/go-lib/core/v2/task"
	"github.com/starudream/go-lib/core/v2/utils/maputil"
	"github.com/starudream/go-lib/core/v2/utils/osutil"
	"github.com/starudream/go-lib/cron/v2"
)

type Job interface {
	Name() string
	Do()
}

type Config struct {
	Disable bool   `json:"disable" yaml:"disable"`
	Spec    string `json:"spec"    yaml:"spec"`
	Startup bool   `json:"startup" yaml:"startup"`
	Jitter  int    `json:"jitter"  yaml:"jitter"`
}

var jobs = maputil.SyncMap[string, Job]{}

func Register(cron Job) {
	_, exist := jobs.LoadOrStore(cron.Name(), cron)
	if exist {
		osutil.ExitErr(fmt.Errorf("cron job %s already exists", cron.Name()))
	}
}

func Run() {
	jobs.Range(func(k string, v Job) bool {
		c := Config{}
		_ = config.Unmarshal(k+".cron", &c)
		if c.Disable {
			return true
		}
		if c.Startup {
			task.Go(func() {
				slog.Debug("cron job %s startup", k)
				v.Do()
				slog.Debug("cron job %s startup done", k)
			})
		}
		if c.Spec != "" {
			err := cron.AddJob(c.Spec, k, func() {
				jitter := int(gh.Ternary(c.Jitter > 0, float64(c.Jitter)*rand.Float64(), 0))
				if jitter > 0 {
					slog.Debug("cron job %s jitter %d seconds", k, jitter)
					time.Sleep(time.Duration(jitter) * time.Second)
				}
				v.Do()
			})
			if err != nil {
				osutil.ExitErr(fmt.Errorf("add cron job %s error: %w", k, err))
			}
			slog.Info("add cron job %s success", k)
		} else {
			slog.Info("%s is empty, skip", k+".cron.spec")
		}
		return true
	})
	cron.Run()
}
