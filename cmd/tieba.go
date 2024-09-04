package main

import (
	"github.com/starudream/sign-task/pkg/tieba/cmd"
)

func init() {
	rootCmd.AddCommand(cmd.TiebaCmd)
}
