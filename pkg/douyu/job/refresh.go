package job

import (
	"bytes"
	_ "embed"
	"fmt"
	"strings"
	"text/template"
	"time"

	"github.com/starudream/sign-task/pkg/douyu/api"
	"github.com/starudream/sign-task/pkg/douyu/config"
	"github.com/starudream/sign-task/pkg/douyu/ws"
	"github.com/starudream/sign-task/util"
)

const (
	refreshLoop = 5
)

type RefreshRecord struct {
	Gifts1 []*api.Gift
	Gifts2 []*api.Gift
	Error  error
}

//go:embed refresh.tmpl
var refreshTplRaw string

var refreshTpl = template.Must(template.New("douyu-refresh").Parse(refreshTplRaw))

func (r RefreshRecord) String() string {
	val := util.ToMap[any](r)
	buf := &bytes.Buffer{}
	_ = refreshTpl.Execute(buf, val)
	return strings.TrimSpace(buf.String())
}

func Refresh(account config.Account) (record RefreshRecord) {
	c := api.NewClient(account)

	err := c.Refresh()
	if err != nil {
		record.Error = fmt.Errorf("refresh error: %w", err)
		return
	}

	gifts1, err := c.ListGift()
	if err != nil {
		record.Error = fmt.Errorf("list gift error: %w", err)
		return
	}
	record.Gifts1 = gifts1.List

	loopCount := 0

login:

	err = ws.Login(ws.LoginParams{
		Room:     account.Room,
		Stk:      c.Stk,
		Ltkid:    c.Ltkid,
		Username: c.Username,
	})
	if err != nil {
		if loopCount >= refreshLoop {
			record.Error = fmt.Errorf("login error: %w", err)
			return
		}
		loopCount++
		time.Sleep(30 * time.Second)
		goto login
	}

	gifts2, err := c.ListGift()
	if err != nil {
		record.Error = fmt.Errorf("list gift error: %w", err)
		return
	}
	record.Gifts2 = gifts2.List

	return
}
