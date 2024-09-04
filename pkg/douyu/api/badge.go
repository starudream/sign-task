package api

import (
	"bytes"
	"cmp"
	"fmt"
	"net/http"
	"slices"
	"strconv"
	"strings"
	"time"

	"golang.org/x/net/html"

	"github.com/starudream/sign-task/util"
)

// Badge 徽章
type Badge struct {
	Room     int       // 房间号
	Anchor   string    // 主播名称
	Name     string    // 名称
	Level    int       // 等级
	Intimacy float64   // 亲密度
	Rank     int       // 排名
	AccessAt time.Time // 获得时间
}

func (c *Client) ListBadges() ([]*Badge, error) {
	resp, err := c.R().SetCookies(c.yCookie()).Get(Addr + "/member/cp/getFansBadgeList")
	if err != nil {
		return nil, fmt.Errorf("[douyu] list badges error: %w", err)
	}

	if code := resp.StatusCode(); code != http.StatusOK {
		return nil, fmt.Errorf("[douyu] http status code %d not 200", code)
	}

	root, err := html.Parse(bytes.NewReader(resp.Body()))
	if err != nil {
		return nil, fmt.Errorf("[douyu] parse html error: %w", err)
	}

	title := util.NodeTitle(root)
	if !strings.Contains(title, "我的头衔") {
		return nil, fmt.Errorf("[douyu] page title not match: %s", title)
	}

	table := util.NodeSearch(root, func(node *html.Node) bool {
		return node.Type == html.ElementNode && strings.TrimSpace(node.Data) == "table" && util.NodeAttrExists(node, func(attr html.Attribute) bool {
			return attr.Key == "class" && strings.Contains(attr.Val, "fans-badge-list")
		})
	})
	if table == nil {
		return nil, fmt.Errorf("[douyu] no table")
	}

	tbody := util.NodeSearch(table, func(node *html.Node) bool {
		return node.Type == html.ElementNode && strings.TrimSpace(node.Data) == "tbody"
	})
	if tbody == nil {
		return nil, fmt.Errorf("[douyu] no table body")
	}

	trs := util.NodeChildren(tbody, "tr")
	if len(trs) == 0 {
		return nil, fmt.Errorf("[douyu] no table rows")
	}

	badges := make([]*Badge, len(trs))

	for i, tr := range trs {
		badge := &Badge{}
		for _, attr := range tr.Attr {
			switch attr.Key {
			case "data-fans-room":
				badge.Room, _ = strconv.Atoi(attr.Val)
			case "data-fans-level":
				badge.Level, _ = strconv.Atoi(attr.Val)
			case "data-fans-intimacy":
				badge.Intimacy, _ = strconv.ParseFloat(attr.Val, 64)
			case "data-fans-rank":
				badge.Rank, _ = strconv.Atoi(attr.Val)
			case "data-fans-gbdgts":
				v, _ := strconv.Atoi(attr.Val)
				badge.AccessAt = time.Unix(int64(v), 0)
			}
		}
		badge.Anchor = util.NodeAttrSearch(tr, func(attr html.Attribute) bool { return attr.Key == "data-anchor_name" })
		badge.Name = util.NodeAttrSearch(tr, func(attr html.Attribute) bool { return attr.Key == "data-bn" })
		badges[i] = badge
	}

	slices.SortFunc(badges, func(a, b *Badge) int {
		return cmp.Compare(b.Intimacy, a.Intimacy)
	})

	return badges, nil
}
