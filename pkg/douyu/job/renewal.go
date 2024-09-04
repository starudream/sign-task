package job

import (
	"bytes"
	_ "embed"
	"fmt"
	"strings"
	"text/template"

	"github.com/starudream/go-lib/core/v2/slog"

	"github.com/starudream/sign-task/pkg/douyu/api"
	"github.com/starudream/sign-task/pkg/douyu/config"
	"github.com/starudream/sign-task/util"
)

type RenewalRecord struct {
	Badges1 []*api.Badge
	Gifts1  []*api.Gift
	Badges2 []*api.Badge
	Gifts2  []*api.Gift
	Skip    bool
	Error   error
}

//go:embed renewal.tmpl
var renewalTplRaw string

var renewalTpl = template.Must(template.New("douyu-renewal").Parse(renewalTplRaw))

func (r RenewalRecord) String() string {
	val := util.ToMap[any](r)
	buf := &bytes.Buffer{}
	_ = renewalTpl.Execute(buf, val)
	return strings.TrimSpace(buf.String())
}

func Renewal(account config.Account) (record RenewalRecord) {
	c := api.NewClient(account)

	err := c.Refresh()
	if err != nil {
		record.Error = fmt.Errorf("refresh error: %w", err)
		return
	}

	badges1, err := c.ListBadges()
	if err != nil {
		record.Error = fmt.Errorf("list badges error: %w", err)
		return
	}
	record.Badges1 = badges1

	gifts1, err := c.ListGift()
	if err != nil {
		record.Error = fmt.Errorf("list gifts error: %w", err)
		return
	}
	record.Gifts1 = gifts1.List

	id := gifts1.FirstNotEmpty(api.GiftFansGlowSticks, api.GiftGlowSticks)
	if id == -1 {
		record.Error = fmt.Errorf("no free gifts")
		return
	}

	gift := gifts1.Find(id)

	if account.IgnoreExpiredCheck {
		slog.Info("ignore expired check")
	} else if !gift.TodayExpired() {
		record.Skip = true
		return
	}

	count := gift.GetCount()

	for i := 0; i < len(account.Assigns); i++ {
		a := account.Assigns[i]

		if count <= 0 {
			break
		}

		if a.All {
			count, err = SendGift(c, a.Room, id, count)
			if err != nil {
				record.Error = err
				return
			}
			continue
		}

		if a.Room <= 0 {
			for j := 0; j < len(badges1); j++ {
				count, err = SendGift(c, badges1[j].Room, id, a.Count)
				if err != nil {
					record.Error = err
					return
				}
			}
			continue
		}

		count, err = SendGift(c, a.Room, id, a.Count)
		if err != nil {
			record.Error = err
			return
		}
	}

	gifts2, err := c.ListGift()
	if err != nil {
		record.Error = fmt.Errorf("list gifts error: %w", err)
		return
	}
	record.Gifts2 = gifts2.List

	badges2, err := c.ListBadges()
	if err != nil {
		record.Error = fmt.Errorf("list badges error: %w", err)
		return
	}
	record.Badges2 = badges2

	return
}

func SendGift(c *api.Client, roomId, giftId, count int) (int, error) {
	slog.Info("send gift(%d) count(%d) to room(%d)", giftId, count, roomId)

	gifts, err := c.SendGift(roomId, giftId, count)
	if err != nil {
		return -1, fmt.Errorf("send gift error: %w", err)
	}

	return gifts.Find(giftId).GetCount(), nil
}
