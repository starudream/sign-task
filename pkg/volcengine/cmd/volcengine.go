package cmd

import (
	"github.com/starudream/go-lib/cobra/v2"
)

var VolcengineCmd = cobra.NewCommand(func(c *cobra.Command) {
	c.Use = "volcengine"
	c.Short = "Manage volcengine"
})
