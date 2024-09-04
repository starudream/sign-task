package cmd

import (
	"github.com/starudream/go-lib/cobra/v2"
)

var XiaoheiheCmd = cobra.NewCommand(func(c *cobra.Command) {
	c.Use = "xiaoheihe"
	c.Short = "Manage xiaoheihe"
})
