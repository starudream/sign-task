package api

import (
	"github.com/starudream/go-lib/core/v2/gh"
)

type ListForumData struct {
	baseResp
	ForumList []*Forum `json:"forum_list"`
}

func (t *ListForumData) GetForums() []*Forum {
	if t == nil {
		return nil
	}
	return t.ForumList
}

type Forum struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	LevelId   int    `json:"level_id"`
	LevelName string `json:"level_name"`
	CurScore  int    `json:"cur_score"`
}

func (c *Client) ListForum() ([]*Forum, error) {
	r := c.R().SetCookies(c.cookie()).SetHeader("subapp-type", "hybrid").SetFormData(gh.MS{"subapp_type": "hybrid"})
	data, err := Exec[*ListForumData](r, "POST", AddrC+"/c/f/forum/like")
	return data.GetForums(), err
}

type SignForumData struct {
	baseResp
	UserInfo *SignForumInfo `json:"user_info"`
}

func (t *SignForumData) GetInfo() *SignForumInfo {
	if t == nil {
		return nil
	}
	return t.UserInfo
}

type SignForumInfo struct {
	UserId   int `json:"user_id"`
	IsSignIn int `json:"is_sign_in"`
}

func (c *Client) SignForum(name string) (*SignForumInfo, error) {
	tbs, err := c.TBS()
	if err != nil {
		return nil, err
	}
	r := c.R().SetCookies(c.cookie()).SetFormData(gh.MS{"tbs": tbs, "kw": name})
	data, err := Exec[*SignForumData](addSign(r), "POST", AddrC+"/c/c/forum/sign")
	return data.GetInfo(), err
}
