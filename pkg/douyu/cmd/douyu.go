package cmd

import (
	"github.com/starudream/go-lib/cobra/v2"
)

var DouyuCmd = cobra.NewCommand(func(c *cobra.Command) {
	c.Use = "douyu"
	c.Short = "Manage douyu"
})
