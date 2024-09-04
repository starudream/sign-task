package job

import (
	"bytes"
	"cmp"
	_ "embed"
	"fmt"
	"slices"
	"strings"
	"text/template"

	"github.com/starudream/go-lib/core/v2/slog"

	"github.com/starudream/sign-task/pkg/geetest"
	"github.com/starudream/sign-task/pkg/miyoushe/api"
	"github.com/starudream/sign-task/pkg/miyoushe/config"
	"github.com/starudream/sign-task/util"
)

type SignGameRecord struct {
	GameId    string
	GameName  string
	RoleName  string
	RoleUid   string
	HasSigned bool
	IsRisky   bool
	IsSuccess bool
	Verify    int
	Award     string
	Skip      bool
	Error     error
}

type SignGameRecords struct {
	Records []SignGameRecord
	Error   error
}

//go:embed game.tmpl
var gameTplRaw string

var gameTpl = template.Must(template.New("miyoushe-sign-game").Parse(gameTplRaw))

func (rs SignGameRecords) String() string {
	val := util.ToMap[any](rs)
	buf := &bytes.Buffer{}
	_ = gameTpl.Execute(buf, val)
	return strings.TrimSpace(buf.String())
}

func SignGame(account config.Account) (records SignGameRecords) {
	c := api.NewClient(account)

	roles, err := c.ListGameRole("")
	if err != nil {
		records.Error = fmt.Errorf("list game role error: %w", err)
		return
	}

	for _, role := range roles.List {
		record := SignGameRole(c, role, account)
		if record.Skip {
			continue
		}
		records.Records = append(records.Records, record)
	}

	slices.SortFunc(records.Records, func(a, b SignGameRecord) int {
		if sub := cmp.Compare(a.GameId, b.GameId); sub != 0 {
			return sub
		}
		return cmp.Compare(a.RoleUid, b.RoleUid)
	})

	return
}

func SignGameRole(c *api.Client, role *api.GameRole, account config.Account) (record SignGameRecord) {
	record.RoleName = role.Nickname
	record.RoleUid = role.GameUid

	record.GameId = api.GameIdByBiz[role.GameBiz]
	record.GameName = api.GameCNNameById[record.GameId]

	if record.GameId == "" {
		record.Error = fmt.Errorf("game biz %s not supported", role.GameBiz)
		return
	}

	if len(account.SignGameIds) > 0 && !slices.Contains(account.SignGameIds, record.GameId) {
		record.Skip = true
		return
	}

	gameName := strings.Split(role.GameBiz, "_")[0]

	home, err := c.GetHome(record.GameId)
	if err != nil {
		record.Error = fmt.Errorf("get home error: %w", err)
		return
	}

	actId := home.GetSignActId()
	if actId == "" {
		record.Error = fmt.Errorf("get sign act id error: %w", err)
		return
	}

	today, err := c.GetSignGame(gameName, actId, role.Region, role.GameUid)
	if err != nil {
		record.Error = fmt.Errorf("get sign game error: %w", err)
		return
	}

	var (
		gt  *geetest.V3Data
		sgd *api.SignGameData
	)

	if today.IsSign {
		record.HasSigned = true
		goto award
	}

sign:

	sgd, err = c.SignGame(gameName, actId, role.Region, role.GameUid, gt)
	if err != nil {
		if api.IsRetCode(err, api.RetCodeGameHasSigned) {
			record.HasSigned = true
		} else {
			record.Error = fmt.Errorf("sign game error: %w", err)
			return
		}
	} else if sgd.IsRisky() {
		record.IsRisky = true
		record.Verify++
		gt, err = dm(&geetest.V3Param{GT: sgd.Gt, Challenge: sgd.Challenge})
		if err == nil {
			goto sign
		} else {
			slog.Error("dm error: %v", err)
			if record.Verify < verifyRetry {
				slog.Info("retry sign, count: %d", record.Verify)
				goto sign
			} else {
				record.Error = fmt.Errorf("dm max retry and give up")
				return
			}
		}
	} else {
		record.IsSuccess = true
	}

award:

	award, err := c.ListSignGameAwardPage(gameName, actId, role.Region, role.GameUid, 1, 3)
	if err != nil {
		record.Error = fmt.Errorf("list sign game award error: %w", err)
		return
	}
	record.Award = award.Today().ShortString()

	return
}
