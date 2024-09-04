package main

import (
	"github.com/starudream/sign-task/pkg/skland/cmd"
)

func init() {
	rootCmd.AddCommand(cmd.SklandCmd)
}
