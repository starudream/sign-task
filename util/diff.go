package util

import (
	"strings"

	"github.com/kr/pretty"
)

func Diff(a, b any) string {
	return strings.Join(pretty.Diff(a, b), ", ")
}
