package main

import (
	"github.com/starudream/sign-task/pkg/qcloud/cmd"
)

func init() {
	rootCmd.AddCommand(cmd.QCloudCmd)
}
