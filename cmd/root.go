package main

import (
	"github.com/starudream/go-lib/cobra/v2"
)

const NAME = "sign-task"

var rootCmd = cobra.NewRootCommand(func(c *cobra.Command) {
	c.Use = NAME

	cobra.AddConfigFlag(c)
})
