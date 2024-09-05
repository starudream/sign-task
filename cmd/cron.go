package main

import (
	"github.com/starudream/go-lib/cobra/v2"
	"github.com/starudream/go-lib/service/v2"

	_ "github.com/starudream/sign-task/pkg/douyu"
	_ "github.com/starudream/sign-task/pkg/kuro"
	_ "github.com/starudream/sign-task/pkg/miyoushe"
	_ "github.com/starudream/sign-task/pkg/skland"
	_ "github.com/starudream/sign-task/pkg/tieba"
)

var cronCmd = cobra.NewCommand(func(c *cobra.Command) {
	c.Use = "cron"
	c.Short = "Run as cron job"
	c.RunE = func(cmd *cobra.Command, args []string) error {
		return service.New(NAME, nil).Run()
	}
})

func init() {
	rootCmd.AddCommand(cronCmd)
}
