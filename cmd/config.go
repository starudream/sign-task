package main

import (
	"github.com/starudream/go-lib/cobra/v2"

	"github.com/starudream/sign-task/pkg/cfg"
)

var (
	configCmd = cobra.NewCommand(func(c *cobra.Command) {
		c.Use = "config"
		c.Short = "Manage config"
	})

	configInitCmd = cobra.NewCommand(func(c *cobra.Command) {
		c.Use = "init"
		c.Aliases = []string{"save"}
		c.Short = "Init config"
		c.RunE = func(cmd *cobra.Command, args []string) error {
			return cfg.Save()
		}
	})
)

func init() {
	configCmd.AddCommand(configInitCmd)

	rootCmd.AddCommand(configCmd)
}
