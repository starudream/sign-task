package api

import (
	"fmt"
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

	DefaultRole *PlayerRole `json:"defaultRole"`
}

func (p *Player) GetDefaultRole() *PlayerRole {
	if p == nil || p.DefaultRole == nil {
		return &PlayerRole{}
	}
	return p.DefaultRole
}

type PlayerRole struct {
	ServerId   string `json:"serverId"`
	ServerType string `json:"serverType"`
	ServerName string `json:"serverName"`
	RoleId     string `json:"roleId"`
	Nickname   string `json:"nickname"`
	Level      int    `json:"level"`
	IsDefault  bool   `json:"isDefault"`
	IsBanned   bool   `json:"isBanned"`
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
	Rarity int    `json:"rarity,omitempty"` // 明日方舟
	Count  int    `json:"count,omitempty"`  // 终末地
}

func (c *Client) SignGame(gid, uid, rid, sid string) (*SignGameData, error) {
	switch gid {
	case GameIdArknights:
		req := c.R().SetBody(gh.M{"gameId": gid, "uid": uid})
		return Exec[*SignGameData](req, "POST", AddrZonai+"/api/v1/game/attendance", c.account)
	case GameIdEndfield:
		req := c.R().SetHeader("sk-game-role", fmt.Sprintf("%s_%s_%s", gid, rid, sid)).
			SetBody(gh.MS{"gameId": gid, "roleId": rid, "serverId": sid})
		return Exec[*SignGameData](req, "POST", AddrZonai+"/api/v1/game/endfield/attendance", c.account)
	}
	return nil, fmt.Errorf("unsupported game: %s", gid)
}

type ListAttendanceData struct {
	CurrentTs       string                  `json:"currentTs"`
	Calendar        []*Calendar             `json:"calendar"`
	ResourceInfoMap map[string]*SignGameRes `json:"resourceInfoMap"`
	Records         CalendarRecords         `json:"records"`  // 明日方舟
	HasToday        bool                    `json:"hasToday"` // 终末地
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
	ResourceId string `json:"resourceId"` // 明日方舟
	AwardId    string `json:"awardId"`    // 终末地
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

func (c *Client) ListSignGame(gid, uid, rid, sid string) (*ListAttendanceData, error) {
	switch gid {
	case GameIdArknights:
		req := c.R().SetQueryParams(gh.MS{"gameId": gid, "uid": uid})
		return Exec[*ListAttendanceData](req, "GET", AddrZonai+"/api/v1/game/attendance", c.account)
	case GameIdEndfield:
		req := c.R().SetHeader("sk-game-role", fmt.Sprintf("%s_%s_%s", gid, rid, sid)).
			SetQueryParams(gh.MS{"gameId": gid, "roleId": rid, "serverId": sid})
		return Exec[*ListAttendanceData](req, "GET", AddrZonai+"/api/v1/game/endfield/attendance", c.account)
	}
	return nil, fmt.Errorf("unsupported game: %s", gid)
}
