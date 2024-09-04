package api

import (
	"strconv"
	"time"

	"github.com/starudream/go-lib/core/v2/gh"
)

type GiftList struct {
	List []*Gift `json:"list"`
}

func (gl *GiftList) Find(id int) *Gift {
	if gl == nil {
		return nil
	}
	for i := range gl.List {
		if gl.List[i].Id == id {
			return gl.List[i]
		}
	}
	return nil
}

func (gl *GiftList) FirstNotEmpty(ids ...int) int {
	for i := 0; i < len(ids); i++ {
		if gl.Find(ids[i]).GetCount() > 0 {
			return ids[i]
		}
	}
	return -1
}

type Gift struct {
	Id       int    `json:"id"`       // id
	Name     string `json:"name"`     // 名称
	Count    int    `json:"count"`    // 现有数量
	Exp      int    `json:"exp"`      // 经验
	Intimate int    `json:"intimate"` // 亲密度
	Met      int    `json:"met"`      // 过期时间

	Price     int `json:"price"     table:",ignored"` // 价值
	PriceType int `json:"priceType" table:",ignored"` // 价值类型（不确定）2-免费礼物
	PropType  int `json:"propType"  table:",ignored"` // 礼物类型（不确定）2-免费礼物 5-等级礼包 6-分区喇叭
}

func (g *Gift) GetCount() int {
	if g == nil {
		return -1
	}
	return g.Count
}

func (g *Gift) TodayExpired() bool {
	if g == nil {
		return false
	}
	y, m, d := time.Now().Date()
	t := time.Date(y, m, d+1, 0, 0, 0, 0, time.Local)
	return t.Unix() >= int64(g.Met)
}

func (c *Client) SendGift(rootId, giftId, count int) (*GiftList, error) {
	form := gh.MS{"roomId": strconv.Itoa(rootId), "propId": strconv.Itoa(giftId), "propCount": strconv.Itoa(count)}
	req := c.R().SetCookies(c.yCookie()).SetFormData(form)
	return Exec[*GiftList](req, "POST", "/japi/prop/donate/mainsite/v1")
}

func (c *Client) ListGift() (*GiftList, error) {
	query := gh.MS{"rid": strconv.Itoa(RoomYYF)}
	req := c.R().SetCookies(c.yCookie()).SetQueryParams(query)
	return Exec[*GiftList](req, "GET", "/japi/prop/backpack/web/v1")
}
