package api

import (
	"strconv"
	"strings"
	"time"

	"github.com/starudream/go-lib/core/v2/gh"

	"github.com/starudream/sign-task/pkg/geetest"
)

type ListGameRoleData struct {
	List []*GameRole `json:"list"`
}

type GameRole struct {
	GameBiz    string `json:"game_biz"`
	Region     string `json:"region"`
	GameUid    string `json:"game_uid"`
	Nickname   string `json:"nickname"`
	Level      int    `json:"level"`
	IsChosen   bool   `json:"is_chosen"`
	RegionName string `json:"region_name"`
	IsOfficial bool   `json:"is_official"`
}

func (c *Client) ListGameRole(gameBiz string) (*ListGameRoleData, error) {
	req := c.R().SetCookies(c.sToken()).SetQueryParam("game_biz", gameBiz)
	return Exec[*ListGameRoleData](req, "GET", AddrTakumi+"/binding/api/getUserGameRolesByStoken")
}

var signGameAddrByName = map[string]string{
	GameNameZZZ: AddrActNap,
}

func signGameAddr(gameName string) string {
	addr, ok := signGameAddrByName[gameName]
	if ok {
		return addr
	}
	return AddrTakumi
}

var signGameHeaderByName = map[string]string{
	GameNameZZZ: "zzz",
}

func signGameHeader(gameName string) string {
	header, ok := signGameHeaderByName[gameName]
	if ok {
		return header
	}
	return gameName
}

type SignGameData struct {
	Code      string `json:"code"`
	Success   int    `json:"success"`
	IsRisk    bool   `json:"is_risk"`
	RiskCode  int    `json:"risk_code"`
	Gt        string `json:"gt"`
	Challenge string `json:"challenge"`
}

func (t *SignGameData) IsRisky() bool {
	return t.IsRisk
}

func (c *Client) SignGame(gameName, actId, region, uid string, gt *geetest.V3Data) (*SignGameData, error) {
	body := gh.MS{"lang": "zh-cn", "act_id": actId, "region": region, "uid": uid}
	req := c.R().SetHeader(xRpcSignGame, signGameHeader(gameName)).SetCookies(c.sToken()).SetCookies(c.cToken()).SetBody(body)
	return Exec[*SignGameData](req, "POST", signGameAddr(gameName)+"/event/luna/sign", gt)
}

type GetSignGameData struct {
	TotalSignDay  int    `json:"total_sign_day"`
	Today         string `json:"today"`
	IsSign        bool   `json:"is_sign"`
	IsSub         bool   `json:"is_sub"`
	Region        string `json:"region"`
	SignCntMissed int    `json:"sign_cnt_missed"`
	ShortSignDay  int    `json:"short_sign_day"`
}

func (c *Client) GetSignGame(gameName, actId, region, uid string) (*GetSignGameData, error) {
	query := gh.MS{"lang": "zh-cn", "act_id": actId, "region": region, "uid": uid}
	req := c.R().SetHeader(xRpcSignGame, signGameHeader(gameName)).SetCookies(c.sToken()).SetCookies(c.cToken()).SetQueryParams(query)
	return Exec[*GetSignGameData](req, "GET", signGameAddr(gameName)+"/event/luna/info")
}

type ListSignGameData struct {
	Month      int                 `json:"month"`
	Biz        string              `json:"biz"`
	Resign     bool                `json:"resign"`
	Awards     []*SignGameAward    `json:"awards"`
	ExtraAward *SignGameExtraAward `json:"short_extra_award"`
}

type SignGameAward struct {
	Name      string `json:"name"`
	Cnt       int    `json:"cnt"`
	CreatedAt string `json:"created_at,omitempty"`
}

type SignGameExtraAward struct {
	HasExtraAward  bool   `json:"has_extra_award"`
	StartTime      string `json:"start_time"`
	EndTime        string `json:"end_time"`
	List           []any  `json:"list"`
	StartTimestamp string `json:"start_timestamp"`
	EndTimestamp   string `json:"end_timestamp"`
}

func (c *Client) ListSignGame(gameName, actId string) (*ListSignGameData, error) {
	query := gh.MS{"lang": "zh-cn", "act_id": actId}
	req := c.R().SetHeader(xRpcSignGame, signGameHeader(gameName)).SetCookies(c.sToken()).SetQueryParams(query)
	return Exec[*ListSignGameData](req, "GET", signGameAddr(gameName)+"/event/luna/home")
}

type ListSignGameAwardData struct {
	Total int            `json:"total"`
	List  SignGameAwards `json:"list"`
}

func (t *ListSignGameAwardData) GetList() SignGameAwards {
	if t == nil {
		return nil
	}
	return t.List
}

type SignGameAwards []*SignGameAward

func (v1 SignGameAwards) Today() (v2 SignGameAwards) {
	today := time.Now().Format(time.DateOnly)
	for i := range v1 {
		if strings.HasPrefix(v1[i].CreatedAt, today) {
			v2 = append(v2, v1[i])
		}
	}
	return
}

func (v1 SignGameAwards) ShortString() string {
	v2 := make([]string, len(v1))
	for i, v := range v1 {
		v2[i] = v.Name + "*" + strconv.Itoa(v.Cnt)
	}
	return strings.Join(v2, ", ")
}

func (c *Client) ListSignGameAward(gameName, actId, region, uid string) (list SignGameAwards, _ error) {
	for page, total := 1, -1; ; page++ {
		query := gh.MS{"lang": "zh-cn", "act_id": actId, "region": region, "uid": uid, "current_page": strconv.Itoa(page), "page_size": "10"}
		req := c.R().SetHeader(xRpcSignGame, signGameHeader(gameName)).SetCookies(c.sToken()).SetCookies(c.cToken()).SetQueryParams(query)
		data, err := Exec[*ListSignGameAwardData](req, "GET", signGameAddr(gameName)+"/event/luna/award", SignDS2)
		if err != nil {
			return nil, err
		}
		list = append(list, data.List...)
		if page == 1 && total == -1 {
			total = data.Total
		}
		total -= len(data.List)
		if total <= 0 {
			return
		}
	}
}

func (c *Client) ListSignGameAwardPage(gameName, actId, region, uid string, page, pageSize int) (list SignGameAwards, _ error) {
	query := gh.MS{"lang": "zh-cn", "act_id": actId, "region": region, "uid": uid, "current_page": strconv.Itoa(page), "page_size": strconv.Itoa(pageSize)}
	req := c.R().SetHeader(xRpcSignGame, signGameHeader(gameName)).SetCookies(c.sToken()).SetCookies(c.cToken()).SetQueryParams(query)
	data, err := Exec[*ListSignGameAwardData](req, "GET", signGameAddr(gameName)+"/event/luna/award", SignDS2)
	return data.GetList(), err
}
