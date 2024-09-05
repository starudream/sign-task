package cmd

import (
	"fmt"

	"github.com/starudream/go-lib/cobra/v2"

	"github.com/starudream/sign-task/pkg/tieba/config"
	"github.com/starudream/sign-task/pkg/tieba/job"
)

var (
	signCmd = cobra.NewCommand(func(c *cobra.Command) {
		c.Use = "sign"
		c.Short = "Sign task manually"
	})

	signForumCmd = cobra.NewCommand(func(c *cobra.Command) {
		c.Use = "forum <account phone>"
		c.Short = "Sign forum"
		c.Run = func(cmd *cobra.Command, args []string) {
			fmt.Println(job.SignForum(config.GetAccountOrFirst(args...)).String())
		}
	})
)

func init() {
	signCmd.AddCommand(signForumCmd)

	TiebaCmd.AddCommand(signCmd)
}
