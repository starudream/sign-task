package job

import (
	"bytes"
	"cmp"
	_ "embed"
	"fmt"
	"slices"
	"strings"
	"text/template"

	"github.com/starudream/sign-task/pkg/kuro/api"
	"github.com/starudream/sign-task/pkg/kuro/config"
	"github.com/starudream/sign-task/util"
)

type SignGameRecord struct {
	GameId     int
	GameName   string
	ServerName string
	RoleId     string
	RoleName   string
	HasSigned  bool
	IsSuccess  bool
	Award      string
	Error      error
}

type SignGameRecords struct {
	Records []SignGameRecord
	Error   error
}

//go:embed game.tmpl
var gameTplRaw string

var gameTpl = template.Must(template.New("kuro-sign-game").Parse(gameTplRaw))

func (rs SignGameRecords) String() string {
	val := util.ToMap[any](rs)
	buf := &bytes.Buffer{}
	_ = gameTpl.Execute(buf, val)
	return strings.TrimSpace(buf.String())
}

func SignGame(account config.Account) (records SignGameRecords) {
	c := api.NewClient(account)

	roles := make([]*api.Role, 0)
	for gid, name := range api.GameNameById {
		rs, err := c.ListRole(gid)
		if err != nil {
			records.Error = fmt.Errorf("list game [%s] role error: %w", name, err)
			continue
		}
		roles = append(roles, rs...)
	}

	for _, role := range roles {
		record := SignGameRole(c, role)
		records.Records = append(records.Records, record)
	}

	slices.SortFunc(records.Records, func(a, b SignGameRecord) int {
		if sub := cmp.Compare(a.GameId, b.GameId); sub != 0 {
			return sub
		}
		return cmp.Compare(a.RoleId, b.RoleId)
	})

	return
}

func SignGameRole(c *api.Client, role *api.Role) (record SignGameRecord) {
	record.GameId = role.GameId
	record.GameName = api.GameNameById[record.GameId]
	record.ServerName = role.ServerName
	record.RoleId = role.RoleId
	record.RoleName = role.RoleName

	if record.GameName == "" {
		record.Error = fmt.Errorf("game id %d not supported", record.GameId)
		return
	}

	records, err := c.ListSignGameRecord(role.GameId, role.ServerId, role.RoleId, role.UserId)
	if err != nil {
		record.Error = fmt.Errorf("list sign game record error: %w", err)
		return
	}

	today := records.Today()
	if len(today) > 0 {
		record.HasSigned = true
		record.Award = today.ShortString()
		return
	}

	sgd, err := c.SignGame(role.GameId, role.ServerId, role.RoleId, role.UserId)
	if err != nil {
		if api.IsCode(err, api.CodeHasSigned) {
			record.HasSigned = true
		} else {
			record.Error = fmt.Errorf("sign game error: %w", err)
			return
		}
	} else {
		list, err2 := c.ListSignGame(role.GameId, role.ServerId, role.RoleId, role.UserId)
		if err2 != nil {
			record.Error = fmt.Errorf("list sign game error: %w", err2)
			return
		}
		record.Award = sgd.TodayList.ShortStringByMap(list.GoodsMap())
	}

	return
}
