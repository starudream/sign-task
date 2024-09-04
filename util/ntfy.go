package util

import (
	"context"
	"errors"

	"github.com/starudream/go-lib/core/v2/slog"
	"github.com/starudream/go-lib/ntfy/v2"
)

func Ntfy(text string) {
	err := ntfy.Notify(context.Background(), text)
	if err != nil && !errors.Is(err, ntfy.ErrNoConfig) {
		slog.Error("notify error: %v", err)
	}
}
