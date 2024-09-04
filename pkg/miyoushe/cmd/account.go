package cmd

import (
	"fmt"

	"github.com/starudream/go-lib/cobra/v2"
	"github.com/starudream/go-lib/tablew/v2"

	"github.com/starudream/sign-task/pkg/miyoushe/config"
	"github.com/starudream/sign-task/pkg/miyoushe/job"
)

var (
	accountCmd = cobra.NewCommand(func(c *cobra.Command) {
		c.Use = "account"
		c.Short = "Manage account"
	})

	accountInitCmd = cobra.NewCommand(func(c *cobra.Command) {
		c.Use = "init <account phone>"
		c.Short = "Init account device information"
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
			config.AddAccount(config.Account{Phone: args[0], Device: config.NewDevice()})
			return config.Save()
		}
	})

	accountLoginCmd = cobra.NewCommand(func(c *cobra.Command) {
		c.Use = "login <account phone>"
		c.Short = "Login account"
		c.Args = func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return fmt.Errorf("requires account phone")
			}
			_, exist := config.GetAccount(args[0])
			if !exist {
				return fmt.Errorf("account %s not exists", args[0])
			}
			return nil
		}
		c.RunE = func(cmd *cobra.Command, args []string) error {
			return job.Login(config.GetAccountOrFirst(args[0]))
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
	accountCmd.AddCommand(accountLoginCmd)
	accountCmd.AddCommand(accountListCmd)

	MiyousheCmd.AddCommand(accountCmd)
}
