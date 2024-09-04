package main

import (
	"github.com/starudream/sign-task/pkg/douyu/cmd"
)

func init() {
	rootCmd.AddCommand(cmd.DouyuCmd)
}
