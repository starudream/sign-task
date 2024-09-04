package main

import (
	"github.com/starudream/sign-task/pkg/miyoushe/cmd"
)

func init() {
	rootCmd.AddCommand(cmd.MiyousheCmd)
}
