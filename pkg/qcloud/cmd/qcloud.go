package cmd

import (
	"github.com/starudream/go-lib/cobra/v2"
)

var QCloudCmd = cobra.NewCommand(func(c *cobra.Command) {
	c.Use = "qcloud"
	c.Short = "Manage qcloud"
})
