package cmd

import (
	"github.com/starudream/go-lib/cobra/v2"
)

var MiyousheCmd = cobra.NewCommand(func(c *cobra.Command) {
	c.Use = "miyoushe"
	c.Short = "Manage miyoushe"
})
