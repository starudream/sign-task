package api

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/starudream/go-lib/core/v2/gh"
)

type Role struct {
	UserId     int    `json:"userId"`
	GameId     int    `json:"gameId"`
	ServerId   string `json:"serverId"`
	ServerName string `json:"serverName"`
	RoleId     string `json:"roleId"`
	RoleName   string `json:"roleName"`
	IsDefault  bool   `json:"isDefault"`
}

func (c *Client) ListRole(gid int) ([]*Role, error) {
	req := c.R().SetFormData(gh.MS{"gameId": strconv.Itoa(gid)})
	return Exec[[]*Role](req, "POST", "/gamer/role/list")
}

type Good struct {
	SerialNum int    `json:"serialNum,omitempty"`
	Type      int    `json:"type,omitempty"`
	GoodsName string `json:"goodsName"`
	GoodsNum  int    `json:"goodsNum"`
	GoodsId   GoodId `json:"goodsId,string"`
	SigInDate string `json:"sigInDate,omitempty"`
}

type Goods []*Good

type GoodId int

func (v *GoodId) UnmarshalJSON(bs []byte) error {
	i, err := strconv.Atoi(strings.Trim(string(bs), "\""))
	if err == nil {
		*v = GoodId(i)
	}
	return err
}

func (v1 Goods) Today() (v2 Goods) {
	today := time.Now().Format(time.DateOnly)
	for i := range v1 {
		if strings.HasPrefix(v1[i].SigInDate, today) {
			v2 = append(v2, v1[i])
		}
	}
	return
}

func (v1 Goods) ShortString() string {
	v2 := make([]string, len(v1))
	for i, v := range v1 {
		v2[i] = v.GoodsName + "*" + strconv.Itoa(v.GoodsNum)
	}
	return strings.Join(v2, ", ")
}

func (v1 Goods) ShortStringByMap(m map[int]*Good) string {
	v2 := make([]string, len(v1))
	for i, v := range v1 {
		v2[i] = m[int(v.GoodsId)].GoodsName + "*" + strconv.Itoa(v.GoodsNum)
	}
	return strings.Join(v2, ", ")
}

type SignGameData struct {
	TodayList    Goods `json:"todayList"`
	TomorrowList Goods `json:"tomorrowList"`
}

func (c *Client) SignGame(gid int, sid, rid string, uid int) (*SignGameData, error) {
	month := fmt.Sprintf("%02d", time.Now().Month())
	req := c.R().SetFormData(gh.MS{"gameId": strconv.Itoa(gid), "serverId": sid, "roleId": rid, "userId": strconv.Itoa(uid), "reqMonth": month})
	return Exec[*SignGameData](req, "POST", "/encourage/signIn/v2")
}

type ListSignGameData struct {
	DisposableGoodsList Goods `json:"disposableGoodsList"`
	DisposableSignNum   int   `json:"disposableSignNum"`

	SignInGoodsConfigs Goods `json:"signInGoodsConfigs"`
	SignLoopGoodsList  Goods `json:"signLoopGoodsList"`
	SigInNum           int   `json:"sigInNum"`

	NowServerTimes  string `json:"nowServerTimes"`
	EventStartTimes string `json:"eventStartTimes"`
	EventEndTimes   string `json:"eventEndTimes"`
	ExpendGold      int    `json:"expendGold"`
	ExpendNum       int    `json:"expendNum"`
	IsSigIn         bool   `json:"isSigIn"`
	OmissionNnm     int    `json:"omissionNnm"`
}

func (v *ListSignGameData) GoodsMap() map[int]*Good {
	m := map[int]*Good{}
	for _, good := range v.DisposableGoodsList {
		m[int(good.GoodsId)] = good
	}
	for _, good := range v.SignInGoodsConfigs {
		m[int(good.GoodsId)] = good
	}
	for _, good := range v.SignLoopGoodsList {
		m[int(good.GoodsId)] = good
	}
	return m
}

func (c *Client) ListSignGame(gid int, sid, rid string, uid int) (*ListSignGameData, error) {
	req := c.R().SetFormData(gh.MS{"gameId": strconv.Itoa(gid), "serverId": sid, "roleId": rid, "userId": strconv.Itoa(uid)})
	return Exec[*ListSignGameData](req, "POST", "/encourage/signIn/initSignInV2")
}

func (c *Client) ListSignGameRecord(gid int, sid, rid string, uid int) (Goods, error) {
	req := c.R().SetFormData(gh.MS{"gameId": strconv.Itoa(gid), "serverId": sid, "roleId": rid, "userId": strconv.Itoa(uid)})
	return Exec[Goods](req, "POST", "/encourage/signIn/queryRecordV2")
}
