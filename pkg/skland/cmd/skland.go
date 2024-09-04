package cmd

import (
	"github.com/starudream/go-lib/cobra/v2"
)

var SklandCmd = cobra.NewCommand(func(c *cobra.Command) {
	c.Use = "skland"
	c.Short = "Manage skland"
})
