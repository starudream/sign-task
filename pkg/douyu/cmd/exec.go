package cmd

import (
	"fmt"

	"github.com/starudream/go-lib/cobra/v2"

	"github.com/starudream/sign-task/pkg/douyu/config"
	"github.com/starudream/sign-task/pkg/douyu/job"
)

var (
	execCmd = cobra.NewCommand(func(c *cobra.Command) {
		c.Use = "exec"
		c.Short = "Exec task manually"
	})

	execRefreshCmd = cobra.NewCommand(func(c *cobra.Command) {
		c.Use = "refresh <account phone>"
		c.Short = "Exec refresh"
		c.Run = func(cmd *cobra.Command, args []string) {
			fmt.Println(job.Refresh(config.GetAccountOrFirst(args...)).String())
		}
	})

	execRenewalCmd = cobra.NewCommand(func(c *cobra.Command) {
		c.Use = "renewal <account phone>"
		c.Short = "Exec renewal"
		c.Run = func(cmd *cobra.Command, args []string) {
			fmt.Println(job.Renewal(config.GetAccountOrFirst(args...)).String())
		}
	})
)

func init() {
	execCmd.AddCommand(execRefreshCmd)
	execCmd.AddCommand(execRenewalCmd)

	DouyuCmd.AddCommand(execCmd)
}
