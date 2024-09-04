package job

import (
	"bytes"
	"cmp"
	_ "embed"
	"fmt"
	"slices"
	"strings"
	"text/template"

	"github.com/starudream/sign-task/pkg/skland/api"
	"github.com/starudream/sign-task/pkg/skland/config"
	"github.com/starudream/sign-task/util"
)

type SignGameRecord struct {
	GameId        string
	GameName      string
	PlayerName    string
	PlayerUid     string
	PlayerChannel string
	HasSigned     bool
	IsSuccess     bool
	Award         string
	Error         error
}

type SignGameRecords struct {
	Records []SignGameRecord
	Error   error
}

//go:embed game.tmpl
var gameTplRaw string

var gameTpl = template.Must(template.New("skland-sign-game").Parse(gameTplRaw))

func (rs SignGameRecords) String() string {
	val := util.ToMap[any](rs)
	buf := &bytes.Buffer{}
	_ = gameTpl.Execute(buf, val)
	return strings.TrimSpace(buf.String())
}

func SignGame(account config.Account) (records SignGameRecords) {
	c := api.NewClient(account)

	players, err := c.ListPlayer()
	if err != nil {
		records.Error = fmt.Errorf("list player error: %w", err)
		return
	}

	for _, app := range players.List {
		for _, player := range app.BindingList {
			record := SignGamePlayer(c, app, player)
			records.Records = append(records.Records, record)
		}
	}

	slices.SortFunc(records.Records, func(a, b SignGameRecord) int {
		if sub := cmp.Compare(a.GameId, b.GameId); sub != 0 {
			return sub
		}
		return cmp.Compare(a.PlayerUid, b.PlayerUid)
	})

	return
}

func SignGamePlayer(c *api.Client, app *api.PlayersByApp, player *api.Player) (record SignGameRecord) {
	record.GameId = api.GameIdByCode[app.AppCode]
	record.GameName = app.AppName
	record.PlayerName = player.NickName
	record.PlayerUid = player.Uid
	record.PlayerChannel = player.ChannelName

	if record.GameId == "" {
		record.Error = fmt.Errorf("game code %s not supported", app.AppCode)
		return
	}

	list, err := c.ListSignGame(record.GameId, player.Uid)
	if err != nil {
		record.Error = fmt.Errorf("list sign game error: %w", err)
		return
	}

	today := list.Records.Today()
	if len(today) > 0 {
		record.HasSigned = true
		record.Award = today.ShortString(list.ResourceInfoMap)
		return
	}

	sgd, err := c.SignGame(record.GameId, player.Uid)
	if err != nil {
		if api.IsCode(err, api.CodeGameHasSigned) {
			record.HasSigned = true
		} else {
			record.Error = fmt.Errorf("sign game error: %w", err)
			return
		}
	} else {
		record.Award = sgd.Awards.ShortString()
	}

	return
}
