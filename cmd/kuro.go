package main

import (
	"github.com/starudream/sign-task/pkg/kuro/cmd"
)

func init() {
	rootCmd.AddCommand(cmd.KuroCmd)
}
