package cmd

import (
	"github.com/starudream/go-lib/cobra/v2"
)

var TiebaCmd = cobra.NewCommand(func(c *cobra.Command) {
	c.Use = "tieba"
	c.Short = "Manage tieba"
})
