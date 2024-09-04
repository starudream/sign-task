package cmd

import (
	"github.com/starudream/go-lib/cobra/v2"
)

var KuroCmd = cobra.NewCommand(func(c *cobra.Command) {
	c.Use = "kuro"
	c.Short = "Manage kuro"
})
