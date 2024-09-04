package cmd

import (
	"fmt"

	"github.com/starudream/go-lib/cobra/v2"
	"github.com/starudream/go-lib/core/v2/utils/fmtutil"
	"github.com/starudream/go-lib/tablew/v2"

	"github.com/starudream/sign-task/pkg/tieba/config"
)

var (
	accountCmd = cobra.NewCommand(func(c *cobra.Command) {
		c.Use = "account"
		c.Short = "Manage account"
	})

	accountInitCmd = cobra.NewCommand(func(c *cobra.Command) {
		c.Use = "init <account phone>"
		c.Short = "Init account information"
		c.Args = func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return fmt.Errorf("requires account phone")
			}
			_, exist := config.GetAccount(args[0])
			if exist {
				return fmt.Errorf("account %s already exists", args[0])
			}
			return nil
		}
		c.RunE = func(cmd *cobra.Command, args []string) error {
			bduss := fmtutil.Scan("please enter cookie BDUSS: ")
			if bduss == "" {
				return nil
			}
			config.AddAccount(config.Account{Phone: args[0], BDUSS: bduss})
			return config.Save()
		}
	})

	accountListCmd = cobra.NewCommand(func(c *cobra.Command) {
		c.Use = "list"
		c.Aliases = []string{"ls"}
		c.Short = "List account"
		c.Run = func(cmd *cobra.Command, args []string) {
			fmt.Println(tablew.Structs(config.C().Accounts))
		}
	})
)

func init() {
	accountCmd.AddCommand(accountInitCmd)
	accountCmd.AddCommand(accountListCmd)

	TiebaCmd.AddCommand(accountCmd)
}
