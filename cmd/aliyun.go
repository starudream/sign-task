package main

import (
	"github.com/starudream/sign-task/pkg/aliyun/cmd"
)

func init() {
	rootCmd.AddCommand(cmd.AliyunCmd)
}
