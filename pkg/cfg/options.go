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
		"$.rr":       []*yaml.Comment{yaml.HeadComment(" 人人打码", "  http://www.rrocr.com")},
		"$.tt":       []*yaml.Comment{yaml.HeadComment(" 套套打码", "  http://www.ttocr.com")},
		"$.douyu":    []*yaml.Comment{yaml.HeadComment(" 斗鱼", "  https://www.douyu.com")},
		"$.kuro":     []*yaml.Comment{yaml.HeadComment(" 库街区", "  https://www.kurobbs.com")},
		"$.miyoushe": []*yaml.Comment{yaml.HeadComment(" 米游社", "  https://www.miyoushe.com")},
		"$.skland":   []*yaml.Comment{yaml.HeadComment(" 森空岛", "  https://www.skland.com")},
		"$.tieba":    []*yaml.Comment{yaml.HeadComment(" 百度贴吧", "  https://tieba.baidu.com")},
		"$.log":      []*yaml.Comment{yaml.HeadComment(" 日志", "  https://pkg.go.dev/github.com/starudream/go-lib/core/v2/config/global#Config")},
		"$.ntfy":     []*yaml.Comment{yaml.HeadComment(" 通知", "  https://pkg.go.dev/github.com/starudream/go-lib/ntfy/v2#Config")},
	}),
}
