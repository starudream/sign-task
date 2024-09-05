package cfg

import (
	"github.com/goccy/go-yaml"
)

var options = []yaml.EncodeOption{
	yaml.Indent(2),
	yaml.IndentSequence(true),
	yaml.UseLiteralStyleIfMultiline(true),
	yaml.WithComment(yaml.CommentMap{
		"$":          []*yaml.Comment{yaml.HeadComment("", " https://github.com/starudream/sign-task", "")},
		"$.geetest":  []*yaml.Comment{yaml.HeadComment(" 打码")},
		"$.douyu":    []*yaml.Comment{yaml.HeadComment(" 斗鱼")},
		"$.kuro":     []*yaml.Comment{yaml.HeadComment(" 库街区")},
		"$.miyoushe": []*yaml.Comment{yaml.HeadComment(" 米游社")},
		"$.skland":   []*yaml.Comment{yaml.HeadComment(" 森空岛")},
		"$.tieba":    []*yaml.Comment{yaml.HeadComment(" 百度贴吧")},
		"$.log":      []*yaml.Comment{yaml.HeadComment(" 日志", "  https://pkg.go.dev/github.com/starudream/go-lib/core/v2/config/global#Config")},
		"$.ntfy":     []*yaml.Comment{yaml.HeadComment(" 通知", "  https://pkg.go.dev/github.com/starudream/go-lib/ntfy/v2#Config")},
	}),
}
