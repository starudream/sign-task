package cmd

import (
	"fmt"

	"github.com/starudream/go-lib/cobra/v2"

	"github.com/starudream/sign-task/pkg/miyoushe/config"
	"github.com/starudream/sign-task/pkg/miyoushe/job"
)

var (
	signCmd = cobra.NewCommand(func(c *cobra.Command) {
		c.Use = "sign"
		c.Short = "Sign task manually"
	})

	signGameCmd = cobra.NewCommand(func(c *cobra.Command) {
		c.Use = "game <account phone>"
		c.Short = "Sign game"
		c.RunE = func(cmd *cobra.Command, args []string) (err error) {
			account := config.GetAccountOrFirst(args...)
			account, err = job.Refresh(account)
			if err != nil {
				return err
			}
			fmt.Println(job.SignGame(account).String())
			return
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

	MiyousheCmd.AddCommand(signCmd)
}
