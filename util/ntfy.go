package util

import (
	"context"
	"errors"
	"fmt"

	"github.com/starudream/go-lib/core/v2/slog"
	"github.com/starudream/go-lib/ntfy/v2"

	"github.com/starudream/sign-task/pkg/cron"
)

func Ntfy(text string) {
	err := ntfy.Notify(context.Background(), text)
	if err != nil && !errors.Is(err, ntfy.ErrNoConfig) {
		slog.Error("notify error: %v", err)
	}
}

func NtfyJob(j cron.Job, k, s string) {
	Ntfy(fmt.Sprintf("[%s] %s\n%s", j.Name(), k, s))
}
