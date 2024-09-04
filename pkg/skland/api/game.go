package api

import (
	"strconv"
	"strings"
	"time"

	"github.com/starudream/go-lib/core/v2/gh"
)

type ListPlayerData struct {
	List []*PlayersByApp `json:"list"`
}

type PlayersByApp struct {
	AppCode     string    `json:"appCode"`
	AppName     string    `json:"appName"`
	DefaultUid  string    `json:"defaultUid"`
	BindingList []*Player `json:"bindingList"`
}

type Player struct {
	Uid             string `json:"uid"`
	ChannelName     string `json:"channelName"`
	ChannelMasterId string `json:"channelMasterId"`
	NickName        string `json:"nickName"`
	IsOfficial      bool   `json:"isOfficial"`
	IsDefault       bool   `json:"isDefault"`
	IsDelete        bool   `json:"isDelete"`
}

func (c *Client) ListPlayer() (*ListPlayerData, error) {
	return Exec[*ListPlayerData](c.R(), "GET", AddrZonai+"/api/v1/game/player/binding", c.account)
}

type SignGameData struct {
	Ts     string         `json:"ts"`
	Awards SignGameAwards `json:"awards"`
}

type SignGameAwards []*SignGameAward

func (t SignGameAwards) ShortString() string {
	v := make([]string, len(t))
	for i, a := range t {
		v[i] = a.Resource.Name + "*" + strconv.Itoa(a.Count)
	}
	return strings.Join(v, ", ")
}

type SignGameAward struct {
	Type     string       `json:"type"`
	Count    int          `json:"count"`
	Resource *SignGameRes `json:"resource"`
}

type SignGameRes struct {
	Id     string `json:"id"`
	Type   string `json:"type"`
	Name   string `json:"name"`
	Rarity int    `json:"rarity"`
}

func (c *Client) SignGame(gid, uid string) (*SignGameData, error) {
	req := c.R().SetBody(gh.M{"gameId": gid, "uid": uid})
	return Exec[*SignGameData](req, "POST", AddrZonai+"/api/v1/game/attendance", c.account)
}

type ListAttendanceData struct {
	CurrentTs       string                  `json:"currentTs"`
	Calendar        []*Calendar             `json:"calendar"`
	Records         CalendarRecords         `json:"records"`
	ResourceInfoMap map[string]*SignGameRes `json:"resourceInfoMap"`
}

type CalendarRecords []*CalendarRecord

func (v1 CalendarRecords) Today() (v2 CalendarRecords) {
	now := time.Now()
	zero := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	zeroTs := strconv.FormatInt(zero.Unix(), 10)
	for _, r := range v1 {
		if r.Ts == zeroTs {
			v2 = append(v2, r)
		}
	}
	return
}

func (v1 CalendarRecords) ShortString(m map[string]*SignGameRes) string {
	v2 := make([]string, len(v1))
	for i, v := range v1 {
		v2[i] = m[v.ResourceId].Name + "*" + strconv.Itoa(v.Count)
	}
	return strings.Join(v2, ", ")
}

type Calendar struct {
	ResourceId string `json:"resourceId"`
	Type       string `json:"type"`
	Count      int    `json:"count"`
	Available  bool   `json:"available"`
	Done       bool   `json:"done"`
}

type CalendarRecord struct {
	Ts         string `json:"ts"`
	ResourceId string `json:"resourceId"`
	Type       string `json:"type"`
	Count      int    `json:"count"`
}

func (c *Client) ListSignGame(gid, uid string) (*ListAttendanceData, error) {
	req := c.R().SetQueryParams(gh.MS{"gameId": gid, "uid": uid})
	return Exec[*ListAttendanceData](req, "GET", AddrZonai+"/api/v1/game/attendance", c.account)
}
