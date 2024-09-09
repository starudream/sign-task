package cmd

import (
	"fmt"

	"github.com/starudream/go-lib/cobra/v2"
	"github.com/starudream/go-lib/core/v2/utils/fmtutil"
	"github.com/starudream/go-lib/tablew/v2"

	"github.com/starudream/sign-task/pkg/aliyun/config"
)

var (
	accountCmd = cobra.NewCommand(func(c *cobra.Command) {
		c.Use = "account"
		c.Short = "Manage account"
	})

	accountInitCmd = cobra.NewCommand(func(c *cobra.Command) {
		c.Use = "init"
		c.Short = "Init account information"
		c.Long = c.Short + "\n" + "Aliyun RAM: https://ram.console.aliyun.com/users"
		c.RunE = func(cmd *cobra.Command, args []string) error {
			id := fmtutil.Scan("please enter AccessKey ID: ")
			if id == "" {
				return nil
			}
			secret := fmtutil.Scan("please enter AccessKey Secret: ")
			if secret == "" {
				return nil
			}
			config.AddAccount(config.Account{Id: id, Secret: secret})
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

	AliyunCmd.AddCommand(accountCmd)
}
