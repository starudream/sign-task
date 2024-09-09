package cmd

import (
	"github.com/starudream/go-lib/cobra/v2"
)

var AliyunCmd = cobra.NewCommand(func(c *cobra.Command) {
	c.Use = "aliyun"
	c.Short = "Manage aliyun"
})
