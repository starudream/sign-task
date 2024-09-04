package cmd

import (
	"fmt"

	"github.com/starudream/go-lib/cobra/v2"

	"github.com/starudream/sign-task/pkg/skland/config"
	"github.com/starudream/sign-task/pkg/skland/job"
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
			sg := job.SignGame(account)
			fmt.Println(sg.String())
			return
		}
	})
)

func init() {
	signCmd.AddCommand(signGameCmd)

	SklandCmd.AddCommand(signCmd)
}
