package main

import (
	"github.com/starudream/sign-task/pkg/volcengine/cmd"
)

func init() {
	rootCmd.AddCommand(cmd.VolcengineCmd)
}
