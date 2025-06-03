package util

import (
	"context"
	"errors"

	"github.com/starudream/go-lib/core/v2/config"
	"github.com/starudream/go-lib/core/v2/slog"
	"github.com/starudream/go-lib/ntfy/v2"
)

func Ntfy(topic, title, text string) {
	var options []ntfy.Option
	switch config.Get("ntfy.module").String() {
	case "ntfy":
		options = append(options,
			ntfy.WithHeader(map[string]string{"x-markdown": "true"}),
			ntfy.WithExtra(map[string]string{"topic": topic, "title": title}),
		)
	default:
		text = "[" + topic + "]\n" + title + "\n" + text
	}
	err := ntfy.Notify(context.Background(), text, options...)
	if err != nil && !errors.Is(err, ntfy.ErrNoConfig) {
		slog.Error("notify error: %v", err)
	}
}
