package util

import (
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func FormatInt(i int) string {
	p := message.NewPrinter(language.English)
	return p.Sprintf("%d", i)
}
