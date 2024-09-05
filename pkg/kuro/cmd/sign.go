package cmd

import (
	"fmt"

	"github.com/starudream/go-lib/cobra/v2"

	"github.com/starudream/sign-task/pkg/kuro/config"
	"github.com/starudream/sign-task/pkg/kuro/job"
)

var (
	signCmd = cobra.NewCommand(func(c *cobra.Command) {
		c.Use = "sign"
		c.Short = "Sign task manually"
	})

	signGameCmd = cobra.NewCommand(func(c *cobra.Command) {
		c.Use = "game <account phone>"
		c.Short = "Sign game"
		c.Run = func(cmd *cobra.Command, args []string) {
			fmt.Println(job.SignGame(config.GetAccountOrFirst(args...)).String())
		}
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
	signCmd.AddCommand(signGameCmd)
	signCmd.AddCommand(signForumCmd)

	KuroCmd.AddCommand(signCmd)
}
