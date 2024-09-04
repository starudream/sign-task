package main

import (
	"context"

	"github.com/starudream/go-lib/core/v2/config"
	"github.com/starudream/go-lib/service/v2"

	"github.com/starudream/sign-task/pkg/cron"
)

func init() {
	args := []string{"cron"}
	if c := config.LoadedFile(); c != "" {
		args = append(args, "-c", c)
	}
	service.AddCommand(rootCmd, service.New(NAME, serviceRun, service.WithArguments(args...)))
}

func serviceRun(context.Context) {
	cron.Run()
}
